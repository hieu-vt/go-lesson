# Learn Golang from 200Lab team 

## Build golang for linux
```ecma script level 3
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app
```

## Docker build
```dockerfile
docker run -d --name fd-restaurant --network my-net -e GINPORT=8080 -e JWT_PROVIDER_SECRET="i_love_you_3000" -e MYSQL_GORM_DB_TYPE="mysql" -e MYSQL_GORM_DB_URI="root:ead8686ba57479778a76e@tcp(mysql:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local" -e USER_API_SECRET="http://fd-user:8082" -p 8080:8080  food-delivery-service:1.0
```
## Docker save
```dockerfile
docker save food-delivery-service -o food-delivery-service.tar
```

## Docker load
```dockerfile
docker load -i food-delivery-service.tar
```

## Check nginx config
```dockerfile
cat ./nginx-conf/default.conf
```

## Install nats pubsub
```dockerfile
docker run -d --name nats --network my-network -p 4222:4222 -p 8222:8222 nats
```


## Install nats redis
```dockerfile
docker run --name redis -e ALLOW_EMPTY_PASSWORD=yes bitnami/redis:latest
```

## Install BufCli
```dockerfile
brew install protobuf
brew install bufbuild/buf/buf
```
[Link download and buf cli](https://docs.buf.build/installation/)

## Go plugins for the protocol compiler
```dockerfile
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
### Update your PATH so that the protoc compiler can find the plugins
```
export PATH="$PATH:$(go env GOPATH)/bin"
```
[Link download and install protobuf](https://grpc.io/docs/languages/go/quickstart/)

## Note about GRPC
* Bình thường thì khi 2 service kết nối với nhau bằng grpc thì thường sẽ đi qua con proxy chứ không kết nối trực tiếp
