# API E-Commerce Basic

This repo serves as basic skeleton to develop production grade REST API using Go + Gin framework, in this case for an e-commerce.

## Features

- [x] Database migrations using `golang-migrate`
- [x] Password hashing with Argon2
- [x] Automatic Swagger API documentation with `github.com/swaggo/swag`
- [x] Authentication using JWT (access + refresh token)
- [x] Reproducible build using Docker and Docker Compose
- [x] Load config from environment variables
- [x] Request input validation
- [x] Error handling
- [x] API health check
- [x] API timeout

## TODO

- [ ] Unit test
- [ ] Integration test
- [ ] RBAC (Role Based Access Control)
- [ ] CI/CD using Github Actions
- [ ] Structured logging
- [ ] Rate limiter

## Build
```bash
$ git clone https://github.com/akhmadzaki/api-ecommerce-basic.git
$ docker compose up --build -d
```

To stop container

```bash
$ docker compose down
```

## Contribute

Please submit pull request or issues to improve this repo.