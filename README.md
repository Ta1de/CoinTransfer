# CoinTransfer

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' drop -f          
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
go run cmd/main.go 
docker exec -it bin/bash
docker ps
migrate create -ext sql -dir ./schema -seq init
docker pull postgres
docker run --name=cointransfer-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
psql -U postgres
UPDATE users SET coins = 1000 WHERE username = 'pers1';
brew install jq
brew install golang-migrate