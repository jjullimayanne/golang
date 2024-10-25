# Use uma imagem base de build e runtime mais compatível
FROM golang:1.23-bullseye AS builder

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário a partir do arquivo main.go na pasta cmd
RUN go build -o api-reciclagem ./cmd/main.go

# Usa a mesma imagem para runtime
FROM golang:1.23-bullseye

WORKDIR /app

# Copia o binário gerado da etapa anterior
COPY --from=builder /app/api-reciclagem .

# Exponha a porta 8080
EXPOSE 8080

# Executa o binário
CMD ["./api-reciclagem"]
