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

* start grpc/https server and grpc client 
```shell
./bin/server
./bin/client
```



* **grpc and http Gateway,** exposes api for accounts management
* **User,** account management grpc service, responsible for CRUD operations on user accounts
   Swagger Documentation
* **http://localhost:8080/swagger/#/  
  Mongo Express to check it's grpc and http works in database
* **http://localhost:8081/
