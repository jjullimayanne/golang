# API de Gestão de Reciclagem e Recompensas

Este projeto é uma API desenvolvida em Go para a gestão de atividades de reciclagem e recompensas, utilizando **PostgreSQL** como banco de dados e **Keycloak** para autenticação.

## Pré-requisitos

Certifique-se de ter o seguinte instalado no seu sistema:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

## Variáveis de Ambiente

Antes de rodar o projeto, crie um arquivo `.env` com as seguintes configurações:

~~~
# Configuração do Banco de Dados PostgreSQL
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=reciclagem
DB_HOST=db
DB_PORT=5432
DB_SSLMODE=disable

# Configuração de Autenticação Keycloak
KEYCLOAK_URL=http://localhost:8080
KEYCLOAK_REALM=myrealm
KEYCLOAK_CLIENT_ID=myclientid
KEYCLOAK_CLIENT_SECRET=myclientsecret

~~~

Como rodar o projeto
1. Subir os containers (API + PostgreSQL)
Você pode subir a aplicação Go e o banco de dados PostgreSQL usando o Docker e o Docker Compose com o seguinte comando:

~~~
make up

~~~
Isso vai:

Construir a imagem Docker da aplicação Go.

Iniciar o container PostgreSQL.
Subir a aplicação no container e conectar ao banco de dados.


2. Rodar a aplicação com logs
Se quiser ver os logs da aplicação e do banco de dados em tempo real, rode:
~~~
make run

~~~

3. Parar os containers
Para parar a aplicação e o banco de dados, use o comando:

~~~
make stop

~~~

Keycloak config Realm

Um Realm no Keycloak é um ambiente isolado que gerencia de forma independente usuários, credenciais, grupos, roles e clientes (aplicativos) de autenticação e autorização. Pense no realm como uma "área de controle" onde você define todos os recursos de segurança para um conjunto específico de usuários e aplicações.
~~~JSON
{
  "resources": [
    {
      "name": "Recicláveis",
      "type": "endpoint",
      "uris": ["/recycle/add", "/recycle/history"],
      "scopes": ["view", "add"]
    },
    {
      "name": "Wallet",
      "type": "endpoint",
      "uris": ["/wallet/balance", "/wallet/history"],
      "scopes": ["view"]
    },
    {
      "name": "Cupons",
      "type": "endpoint",
      "uris": ["/coupons/available", "/coupons/redeem"],
      "scopes": ["view", "redeem"]
    },
    {
      "name": "Administração de Cupons",
      "type": "admin-endpoint",
      "uris": ["/admin/coupons/add", "/admin/coupons/update", "/admin/coupons/delete"],
      "scopes": ["manage"]
    }
  ],
  "policies": [
    {
      "name": "User Policy",
      "type": "role",
      "roles": ["user"]
    },
    {
      "name": "Admin Policy",
      "type": "role",
      "roles": ["admin"]
    }
  ],
  "permissions": [
    {
      "name": "Access Recicláveis",
      "type": "scope",
      "scopes": ["view", "add"],
      "resources": ["Recicláveis"],
      "policies": ["User Policy"]
    },
    {
      "name": "Access Wallet",
      "type": "scope",
      "scopes": ["view"],
      "resources": ["Wallet"],
      "policies": ["User Policy"]
    },
    {
      "name": "Access Cupons",
      "type": "scope",
      "scopes": ["view", "redeem"],
      "resources": ["Cupons"],
      "policies": ["User Policy"]
    },
    {
      "name": "Manage Cupons Admin",
      "type": "scope",
      "scopes": ["manage"],
      "resources": ["Administração de Cupons"],
      "policies": ["Admin Policy"]
    }
  ]
}


~~~

# Como configurar e acessar o `keycloack`

No arquivo `docker-compose.yml`, o `keycloack` foi configurado como um serviço Docker separado, com algumas variáveis de ambiente e configurações específicas.

~~~yaml
  keycloak:
    image: quay.io/keycloak/keycloak:latest
    environment:
      KEYCLOAK_ADMIN: admin
      KEYCLOAK_ADMIN_PASSWORD: admin
    ports:
      - "8081:8080"
    command: ["start-dev"]
    networks:
      - api-net

~~~

### Acesse o Painel de Administração do Keycloak
O contêiner Keycloak está configurado para escutar na porta 8080 internamente, mas está mapeado para a porta 8081 no sistema host.
Assim, você acessa o Keycloak via `http://localhost:8081` no navegador, mas o contêiner comunica-se internamente com a porta 8080.

Faça login com as credenciais de administrador:
No  `docker-compose.yml`, esses valores foram definidos com as variáveis de ambiente:
~~~
Usuário: admin
Senha: admin
~~~
