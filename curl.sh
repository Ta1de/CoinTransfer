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

curl -X POST http://localhost:8080/api/sendCoin \
     -H "Authorization: Bearer $token1" \
     -H "Content-Type: application/json" \
     -d '{"ToUser": "pers2", "Amount": 100}'

echo \

curl -X POST http://localhost:8080/api/sendCoin \
     -H "Authorization: Bearer $token2" \
     -H "Content-Type: application/json" \
     -d '{"ToUser": "pers1", "Amount": 100}'

echo \

curl -X GET http://localhost:8080/api/buy/cup \
     -H "Authorization: Bearer $token1" \
     -H "Content-Type: application/json"

echo \

curl -X GET http://localhost:8080/api/buy/cup \
     -H "Authorization: Bearer $token1" \
     -H "Content-Type: application/json"

echo \

curl -X GET http://localhost:8080/api/buy/book \
     -H "Authorization: Bearer $token1" \
     -H "Content-Type: application/json"

echo \

curl -X GET http://localhost:8080/api/info \
     -H "Authorization: Bearer $token1" \
     -H "Content-Type: application/json"