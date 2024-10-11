# Nome do serviço no Docker Compose
DOCKER_COMPOSE=docker-compose

# Nome do binário
BINARY_NAME=api-reciclagem

# Comando para rodar o Docker Compose
.PHONY: up
up:
	$(DOCKER_COMPOSE) up -d --build

# Para rodar o Docker Compose no modo interativo (log)
.PHONY: run
run:
	$(DOCKER_COMPOSE) up

# Para parar os containers
.PHONY: stop
stop:
	$(DOCKER_COMPOSE) down

# Para rodar os testes usando Docker
.PHONY: test
test:
	$(DOCKER_COMPOSE) run api go test ./...

# Para limpar os containers, volumes, e imagens
.PHONY: clean
clean:
	$(DOCKER_COMPOSE) down --volumes --rmi all

# Para acessar o container da aplicação Go
.PHONY: shell
shell:
	$(DOCKER_COMPOSE) exec api sh
