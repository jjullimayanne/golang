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