# Usa uma imagem base compatível com Go
FROM golang:1.23-bullseye AS builder

WORKDIR /app

# Copia os arquivos de dependências do Go
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código-fonte
COPY . .

# Compila o binário principal a partir do arquivo main.go
RUN go build -o api-reciclagem ./cmd/main.go

# Usa uma imagem para o runtime
FROM golang:1.23-bullseye

WORKDIR /app

# Copia o binário gerado na etapa anterior
COPY --from=builder /app/api-reciclagem .

# Expõe a porta 8080
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./api-reciclagem"]
