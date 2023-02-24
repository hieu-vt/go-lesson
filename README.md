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

