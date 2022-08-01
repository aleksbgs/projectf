**Run**
* Please run MONGODB first so services can have ready connection
```shell
docker-compose up 
```
* After  Database, run make proto
```shell
make proto
```

* After  proto and create documentation, run make faceit to create bin files in bin folder
```shell
make faceit
```

* start grpc/http server and grpc client 
```shell
./bin/server
./bin/client
```


* GRPC HEALTH PROBE
* go install github.com/grpc-ecosystem/grpc-health-probe@v0.4.8
```shell
grpc-health-probe -addr="0.0.0.0:50051" -service=""
```



* **grpc and http Gateway,** exposes api for accounts management
* **User,** account management grpc service, responsible for CRUD operations on user accounts

   Swagger Documentation
* http://localhost:8080/swagger/#/  
  
  Mongo Express to check it's grpc and http works in database
* http://localhost:8081/
  
   Postman Collection
* https://www.getpostman.com/collections/766b10b63ff5c0a74db6

