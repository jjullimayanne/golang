package auth

import (
    "net/http"
    "github.com/coreos/go-oidc"
    "golang.org/x/net/context"
)

type KeycloakAuthenticator struct {
    Verifier *oidc.IDTokenVerifier
}

func NewKeycloakAuthenticator(issuerURL string, clientID string) (*KeycloakAuthenticator, error) {
    provider, err := oidc.NewProvider(context.Background(), issuerURL)
    if err != nil {
        return nil, err
    }

    config := &oidc.Config{
        ClientID: clientID,
    }

    verifier := provider.Verifier(config)

    return &KeycloakAuthenticator{
        Verifier: verifier,
    }, nil
}

func (a *KeycloakAuthenticator) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        _, err := a.Verifier.Verify(context.Background(), token)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
