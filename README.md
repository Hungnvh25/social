2 - Project Architecture

3 - Building a server from TCP to HTTP

![image](https://github.com/user-attachments/assets/f29fc312-fd94-4be7-a8e4-aec9fbf14e82)

![image](https://github.com/user-attachments/assets/4c4bf9bb-e3f5-4c99-a2d0-6a72394273d3)

Go mod init …..
go get -u github.com/go-chi/chi/v5  
go install github.com/air-verse/air@latest

-tạo database postgres qua docker
docker compose up --build

-Tạo table
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate create -seq -ext sql -dir ./cmd/migrate/migrations/ create_users

migrate -path=cmd/migrate/migrations/ -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" up


![alt text](image.png)
