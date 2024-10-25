# Nome do serviço Docker Compose
DOCKER_COMPOSE = docker-compose

# Nome do binário
BINARY_NAME = api-reciclagem

# Comando para rodar o Docker Compose e construir as imagens
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

# Para acessar o shell do container da aplicação Go
.PHONY: shell
shell:
	$(DOCKER_COMPOSE) exec api sh

# Comando para rodar o linting usando golangci-lint
.PHONY: lint
lint:
	$(DOCKER_COMPOSE) run api golangci-lint run

# Para fazer o build da aplicação Go dentro do container
.PHONY: build
build:
	$(DOCKER_COMPOSE) run api go build -o $(BINARY_NAME) ./cmd/main.go

# Para rodar a aplicação Go (se necessário)
.PHONY: run-app
run-app:
	$(DOCKER_COMPOSE) exec api ./$(BINARY_NAME)

# Para verificar o código antes do commit
.PHONY: check
check: lint test

# Para rodar tudo (build, lint, test)
.PHONY: all
all: clean up lint test build
