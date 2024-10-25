package keycloak

import (
	coreError "api/internal/core/error"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type KeycloakAuthenticator struct {
	Verifier *oidc.IDTokenVerifier
	Config   oauth2.Config
}

func NewKeycloakAuthenticator() (*KeycloakAuthenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, os.Getenv("KEYCLOAK_URL")+"/realms/"+os.Getenv("KEYCLOAK_REALM"))
	if err != nil {
		return nil, coreError.WrapError(err, "falha ao inicializar o Keycloak Provider")
	}

	config := oauth2.Config{
		ClientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		Endpoint:     provider.Endpoint(),
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: config.ClientID})

	return &KeycloakAuthenticator{
		Verifier: verifier,
		Config:   config,
	}, nil
}

func (a *KeycloakAuthenticator) ValidateToken(r *http.Request) (bool, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		err := coreError.NewErrorf("token ausente")
		coreError.LogError(err)
		return false, err
	}

	token = token[7:]

	_, err := a.Verifier.Verify(context.Background(), token)
	if err != nil {
		return false, coreError.WrapError(err, "falha na validação do token")
	}
	return true, nil
}

func (a *KeycloakAuthenticator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			err := coreError.NewErrorf("token ausente")
			coreError.LogError(err)
			http.Error(w, "Token ausente", http.StatusUnauthorized)
			return
		}

		token = token[7:] 

		_, err := a.Verifier.Verify(context.Background(), token)
		if err != nil {
			wrappedErr := coreError.WrapError(err, "token inválido")
			coreError.LogError(wrappedErr)
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *KeycloakAuthenticator) GetUserRoles(r *http.Request) ([]string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		err := coreError.NewErrorf("token ausente")
		coreError.LogError(err)
		return nil, err
	}
	token = token[7:] 

	idToken, err := a.Verifier.Verify(context.Background(), token)
	if err != nil {
		return nil, coreError.WrapError(err, "falha na verificação do token")
	}

	var claims struct {
		RealmAccess struct {
			Roles []string `json:"roles"`
		} `json:"realm_access"`
	}
	if err := idToken.Claims(&claims); err != nil {
		return nil, coreError.WrapError(err, "erro ao extrair claims do token")
	}

	return claims.RealmAccess.Roles, nil
}

func (a *KeycloakAuthenticator) getAccessToken() (string, error) {
	tokenURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", os.Getenv("KEYCLOAK_URL"), os.Getenv("KEYCLOAK_REALM"))

	data := url.Values{}
	data.Set("client_id", os.Getenv("KEYCLOAK_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("KEYCLOAK_CLIENT_SECRET"))
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", coreError.WrapError(err, "erro ao criar requisição para obter access token")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", coreError.WrapError(err, "erro ao enviar requisição para obter access token")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", coreError.NewErrorf("falha ao obter access token: status %d", resp.StatusCode)
	}

	var responseData struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return "", coreError.WrapError(err, "erro ao decodificar resposta de access token")
	}

	return responseData.AccessToken, nil
}

func (a *KeycloakAuthenticator) CreateUser(username, email, password string) error {
	accessToken, err := a.getAccessToken()
	if err != nil {
		return coreError.WrapError(err, "falha ao obter token de acesso para criação de usuário")
	}

	url := fmt.Sprintf("%s/admin/realms/%s/users", os.Getenv("KEYCLOAK_URL"), os.Getenv("KEYCLOAK_REALM"))

	user := map[string]interface{}{
		"username": username,
		"email":    email,
		"enabled":  true,
		"credentials": []map[string]string{
			{"type": "password", "value": password, "temporary": "false"},
		},
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		return coreError.WrapError(err, "erro ao criar JSON para novo usuário")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return coreError.WrapError(err, "erro ao criar requisição para o Keycloak")
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return coreError.WrapError(err, "erro ao enviar requisição para o Keycloak")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return coreError.NewErrorf("falha ao criar usuário no Keycloak: status %d", resp.StatusCode)
	}

	return nil
}
