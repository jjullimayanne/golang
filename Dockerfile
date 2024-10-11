# Etapa de build
FROM golang:1.19 AS builder

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar o go.mod e go.sum para dentro do container
COPY go.mod go.sum ./

# Baixar as dependências Go
RUN go mod download

# Copiar o código-fonte
COPY . .

# Compilar a aplicação
RUN go build -o api-reciclagem ./cmd/main.go

# Etapa final - Container leve para execução da aplicação
FROM debian:buster

# Definir o diretório de trabalho
WORKDIR /app

# Copiar o binário compilado da etapa anterior
COPY --from=builder /app/api-reciclagem .

# Expor a porta 8080
EXPOSE 8080

# Executar o binário
CMD ["./api-reciclagem"]
