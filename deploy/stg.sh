APP_NAME=food-delivery

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME}

docker run -d --name ${APP_NAME} \
      --network my-net
      -e VIRTUAL_HOST=""
      -e DBConnectionStr="root:ead8686ba57479778a76e@tcp(mysql:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
      -p 8080:8080 \
      ${APP_NAME}