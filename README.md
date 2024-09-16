# In Memory CRUD App

Esse projeto é um cadastro básico de usuários, usando a memória como persistência.

## Requisitos

- [x] Nome e sobrenome do usuário deve ter no mínimo 2 e no máximo 20 caracteres.
- [x] Biografia do usuário deve ter no mínimo 20 e no máximo 450 caracteres.

## Uso

Com [Go](https://go.dev/) instalado, inicie o servidor na raiz do projeto com `go run ./cmd`.

Rotas do servidor:

### Criação de usuário

Exemplo de payload:

```json
{
    "first_name": "hello",
    "last_name": "world",
    "biography": "this is a hello world biography"
}
``` 

POST `http://localhost:8080/api/users`

```bash
curl -v -X \
-d '{"first_name": "hello", "last_name": "world", "biography": "this is a hello world biography"}' \
http://localhost:8080/api/users
```

### Listagem de usuários cadastrados

GET `http://localhost:8080/api/users`

```bash
curl -v http://localhost:8080/api/users
```

### Listagem de usuário específico

GET `http://localhost:8080/api/users/id`

```bash
curl -v http://localhost:8080/api/users/id
```

### Atualização de usuário específico

PUT `http://localhost:8080/api/users/id`

```bash
curl -v -X PUT \
-d '{"first_name": "updated", "last_name": "record", "biography": "new biography for the updated record"}' \
http://localhost:8080/api/users/id
```

### Remoção de usuário específico

DELETE `http://localhost:8080/api/users/id`

```bash
curl -v -X DELETE http://localhost:8080/api/users/id
```