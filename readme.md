# component

# postgres with docker
- create postgres db with docker
```
docker pull postgres:alpine
```

- start container postgres
```
docker run --name kawaii_db_test -e POSTGRES_USER=kawaii -e POSTGRES_PASSWORD=123456 -p 4444:5432
```

- test postgres by connecting to docker bash
```
docker exec -it kawaii_db_test bash
```

- connect postgres with user kawaii 
```
psql -U kawaii
```

- list for all database 
```
\l
``` 

- create database
```
CREATE DATABASE kawaii_db_test;
```

- drop database
```
DROP DATABASE kawaii_db_test;
```

# migrate 
`https://github.com/golang-migrate/migrate`

install migrate
```go
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

test
```
migrate -version
```

## using migrate

1. create migration
```
cd pkg/databases/migrations
migrate create -ext sql -seq kawaii_db
```
2. migrate database

```
migrate -source file://E:/Repository/go/kawaii-shop/pkg/databases/migrations -database 'postgres://kawaii:123456@localhost:4444/kawaii_db_test?sslmode=disable' -verbose up
```

-verbose = logging 

# godotenv
`https://github.com/joho/godotenv`

install
```go
go install github.com/joho/godotenv/cmd/godotenv@latest
```

# go hot reload using air by cosmtrek
`https://github.com/cosmtrek/air`

install
```go
go install github.com/cosmtrek/air@latest
```

initial air
```
air init
```

settings .air.toml for clean temp file when exit application
```
clean_on_exit = true
```

start air
```
air -c .air.dev.toml
```

# go sqlx
`https://github.com/jmoiron/sqlx`

install
```go
go get github.com/jmoiron/sqlx
```

# go fiber
install 
```go
go get github.com/gofiber/fiber/v2
```

---
# .env example
```
APP_HOST=127.0.0.1
APP_PORT=3000
APP_NAME=kawaii-shop
APP_VERSION=v0.1.0
APP_BODY_LIMIT=10490000
APP_READ_TIMEOUT=60
APP_WRTIE_TIMEOUT=60
APP_FILE_LIMIT=2097000
APP_GCP_BUCKET=kawaii-shop-dev-bucket

JWT_SECRET_KEY=pehWw0SsewWgg6lb
JWT_ACCESS_EXPIRES=86400
JWT_REFRESH_EXPIRES=604800
JWT_ADMIN_KEY=LFuN8Tq6hXEW1ozR
JWT_API_KEY=nxh463PBcV7IPvff

DB_HOST=127.0.0.1
DB_PORT=4444
DB_PROTOCOL=tcp
DB_USERNAME=kawaii
DB_PASSWORD=123456
DB_DATABASE=kawaii_db_test
DB_SSL_MODE=disable
DB_MAX_CONNECTIONS=25
```