ALTER USER postgres WITH PASSWORD 'password';

https://github.com/golang-migrate/migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

docker run --name my-postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -e POSTGRES_USER=postgres -d postgres

migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path ./internal/database/migrations version
migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path ./internal/database/migrations up 1
migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path ./internal/database/migrations down 1

migrate create -ext="sql" -dir="internal/database/migrations" -seq -digits=3 partition_users

go get -u gorm.io/gorm

https://postgrespro.ru/docs/postgresql/10/ddl-partitioning#DDL-PARTITIONING-DECLARATIVE