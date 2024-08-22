# Usa uma imagem base oficial do Golang
FROM golang:1.23 AS build

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Configura a arquitetura para x86_64
ENV GOARCH=amd64

# Copia os arquivos da aplicação para o container
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Compila a aplicação
RUN go build -o urlzipper ./cmd/main.go

# Usa uma imagem base menor para a execução da aplicação
FROM alpine:latest AS run

# Instala dependências para a execução do binário
RUN apk --no-cache add ca-certificates

# Copia o binário da etapa de construção
COPY --from=build /app/urlzipper /urlzipper

# Define o diretório de trabalho
WORKDIR /

RUN chmod +x /urlzipper

# Expõe a porta 8080
EXPOSE 8080

# Define o comando padrão a ser executado quando o container iniciar
CMD ["./urlzipper"]
