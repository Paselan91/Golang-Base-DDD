# golang ddd sample

## env

- Golang 1.18
  - Echo 3.3.10
- MySQL 8.0.23
- Docker 20.10.5
- docker-compose 1.29.2

## Installation
### Build & up Contaier
  - build
```zsh
docker compose build
```

- up container
```zsh
docker compose up
```

### migrate & sees

#### migrate
- request
```zsh
curl http://localhost:8080/api/v1/migrate
```
- resopnse
```
success migrate
```
#### seed
- request
```zsh
curl http://localhost:8080/api/v1/seed
```
- resopnse
```
success seed
```

### Test

####  Get user
- request
```zsh
curl http://localhost:8080/test_user/1 
```
- response
```json
{"id":1,"user_sub_id":"test_user_0","mail_address":"test_user_0@test.com"}
```

####  Create user
- request
```zsh
curl -X POST -H "Content-Type: application/json" -d '{"user_sub_id":"testUser123", "mail_address":"hoge-123@hoge.com"}' localhost:8080/test_user/create
```
- response
```json
{"id":124,"user_sub_id":"testUser123","mail_address":"hoge-123@hoge.com"}
```
