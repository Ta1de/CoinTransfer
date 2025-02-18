migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

token1=$(curl -X POST http://localhost:8080/api/auth \
  -H 'Content-Type: application/json' \
  -d '{"username": "pers1", "password": "password"}' | jq -r '.token')

token2=$(curl -X POST http://localhost:8080/api/auth \
  -H 'Content-Type: application/json' \
  -d '{"username": "pers2", "password": "password"}' | jq -r '.token')

token3=$(curl -X POST http://localhost:8080/api/auth \
  -H 'Content-Type: application/json' \
  -d '{"username": "pers3", "password": "password"}' | jq -r '.token')

token4=$(curl -X POST http://localhost:8080/api/auth \
  -H 'Content-Type: application/json' \
  -d '{"username": "pers4", "password": "password"}' | jq -r '.token')

docker exec -it cointransfer-db psql -U postgres -c "UPDATE users SET coins = 5000 WHERE username = 'pers1';"

curl -X POST http://localhost:8080/api/sendCoin \
     -H "Authorization: Bearer $token1" \
     -H "Content-Type: application/json" \
     -d '{"ToUser": "pers2", "Amount": 100}'

curl -X POST http://localhost:8080/api/sendCoin \
     -H "Authorization: Bearer $token2" \
     -H "Content-Type: application/json" \
     -d '{"ToUser": "pers1", "Amount": 100}'

curl -X POST http://localhost:8080/api/sendCoin \
     -H "Authorization: Bearer $token1" \
     -H "Content-Type: application/json" \
     -d '{"ToUser": "pers2", "Amount": 1000}'

curl -X POST http://localhost:8080/api/sendCoin \
     -H "Authorization: Bearer $token1" \
     -H "Content-Type: application/json" \
     -d '{"ToUser": "pers3", "Amount": 1000}'

curl -X POST http://localhost:8080/api/sendCoin \
     -H "Authorization: Bearer $token1" \
     -H "Content-Type: application/json" \
     -d '{"ToUser": "pers4", "Amount": 1000}'