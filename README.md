
# Redis Integration in Golang for multiple Data structure


# start redis docker

$ docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest

$ cd redis-go-integration && go mod download && go run main.go