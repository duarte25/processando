# Usa a imagem oficial do Golang
FROM golang:1.23.5-alpine

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o código fonte para o diretório de trabalho
COPY . .

# Baixa as dependências
RUN go mod download

# Compila a aplicação
RUN go build -o main .

# Expõe a porta 8080
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]