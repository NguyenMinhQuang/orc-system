DB_URL=mysql://root:root@tcp(127.0.0.1:3306)/orc-sys?sslmode=disable

# Go migrate Mysql
force:
	migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/orc-sys?multiStatements=true" -path migrations force 1
install:
	# scoop install migrate  ==> information: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
create:
	migrate create -ext sql -dir migrations -seq user_info  # create file sql migrate
up:
	migrate -path migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/orc-sys?multiStatements=true" -verbose up
down:
	migrate -path migration -database "mysql://root:root@tcp(127.0.0.1:3306)/orc-sys?multiStatements=true" -verbose down
##############################
migrate-up:
	go run ./cmd/migrate/main.go
install-linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
run:
	go run ./cmd/app/main.go

.PHONY: run force install create up down migrate-up install-linter
