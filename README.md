# Axxonsoft

Run package by cmd/app/*.go <br>
or<br>
by cmd/app/main.go

### Env
Env files example
```bash
DB_HOST=localhost
DB_USER=axxonsoft
DB_PASS=axxonsoft
DB_NAME=axxonsoft

AMQP_USER=lazyuser
AMQP_PASS=qwerty
AMQP_HOST=localhost
AMQP_PORT=5672
```

### Curls
```bash
curl --location 'http://localhost:8080/api/v1/proxy/task' \
--header 'Content-Type: application/json' \
--data '{
    "method": "GET",
    "url": "https://google.kz"
}'
```

```bash
curl --location 'http://localhost:8080/api/v1/proxy/task/90bfb1a4-d4d2-44ca-9925-e453ca514909'
```


### Dockerfiles

Dockerfile for local Postgres:
```yaml
version: "3.9"
services:
  postgres:
    image: postgres:13.3
    environment:
      POSTGRES_DB: "axxonsoft"
      POSTGRES_USER: "axxonsoft"
      POSTGRES_PASSWORD: "axxonsoft"
    volumes:
      - <YOUR_LOCAL_PATH>/postgres_axxonsoft/data:/var/lib/postgresql/data
    ports:
      - "5433:5432"
```

Dockerfile for local RMQ:
```yaml
version: "2.1"
services:
  rabbitmq:
    image: rabbitmq:3.10.7-management
    hostname: rabbitmq
    restart: always
    environment:
      - RABBITMQ_DEFAULT_USER=lazyuser
      - RABBITMQ_DEFAULT_PASS=qwerty
      - RABBITMQ_SERVER_ADDITIONAL_ERL_ARGS=-rabbit disk_free_limit 2147483648
      - /Users/lazy_owl/Desktop/docker/rabbitmq/data:/var/lib/rabbitmq
    ports:
      - 15672:15672
      - 5672:5672
```