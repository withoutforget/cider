DATABASE_URL := env('DATABASE_URL')

[doc("Run Go app outside Docker")]
@run:
    go run ./cmd/app/main.go
[doc("Show lines of code")]
@lines:
    cloc -vcs=git .

tests:
    cd ./tests/ && uv run main.py
[doc("Apply all migrations")]
migrate-up:
    migrate -database {{DATABASE_URL}} -path ./migrations up

[doc("Rollback last migration")]
migrate-down:
    migrate -database {{DATABASE_URL}} -path ./migrations down 1

[doc("Create new migration")]
migrate-create name:
    migrate create -ext sql -dir ./migrations -seq {{name}}
