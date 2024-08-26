# Simsples API para aprender GO lang

aprendendo golang e criando uma api com autenticação JWT e blocklist

## Tecnologias

- Go
- Gorm
- Postgres
- Redis
- Testify

## Como utilizar?

### Rodar aplicação

Setar as variáveis de ambiente, copiar o conteúdo do `.env.example`

lembrando que é necessário do bancos `postgres` e `redis` rodando.

```bash
go run cmd/server/main.go
```

### Rodar aplicação através do dockerfile

```bash
docker compose up
```
