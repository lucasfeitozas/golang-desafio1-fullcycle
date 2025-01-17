# Desafio 01 Full Cycle - Cotação do Dólar
## Descrição do Projeto

Este projeto consiste em dois componentes principais: `client.go` e `server.go`. O objetivo é obter a cotação do dólar em tempo real.

## server.go

O `server.go` é responsável por fornecer a cotação do dólar. Ele faz uma requisição a uma API externa para obter a cotação atual e expõe essa informação através de um servidor HTTP.

### Funcionamento

1. Inicializa um servidor HTTP na porta especificada.
2. Faz uma requisição a uma API externa para obter a cotação do dólar.
3. Retorna a cotação do dólar em formato JSON quando acessado via HTTP.

## client.go

O `client.go` é responsável por consumir o serviço fornecido pelo `server.go` e exibir a cotação do dólar.

### Funcionamento

1. Faz uma requisição HTTP ao servidor `server.go`.
2. Recebe a cotação do dólar em formato JSON.
3. Exibe a cotação do dólar no console.

## Como Executar

### Executar o Servidor

Para iniciar o servidor, execute o seguinte comando:

```sh
go run server.go
```

### Executar o Cliente

Para iniciar o cliente, execute o seguinte comando:

```sh
go run client.go
```

Certifique-se de que o servidor esteja em execução antes de iniciar o cliente.

## Requisitos

- Go 1.16 ou superior
- Conexão com a internet para acessar a API externa

## Estrutura do Projeto

```
golang-desafio1-fullcycle/
├── client.go
├── server.go
└── README.md
```