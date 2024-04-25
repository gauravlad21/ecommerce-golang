# ecommerce-golang

tech used:
- gin framework
- sqlc database library with postgres
- goroutines for async call
- go-migration for migrate database
- docker and docker-compose


below is individual micro-service's structure
.
├── Dockerfile
├── common          ====> contains common functionalities
│   ├── common.go
│   ├── enum.go
│   ├── response.go
│   ├── struct.go
│   └── viper.go    ====> config reader
├── config.json     ====> configuration file. update env variable to overwrite it
├── controller      ====> controller/handler
│   ├── comm.go
│   └── handlers.go
├── service         ====> business logic
│   ├── interface.go
│   └── service.go
├── dbhelper        ====> database layer
│   ├── connection_setup.go
│   ├── dbUtils.go
│   ├── dpOperations.go
│   └── sqlc         
│       ├── dbsqlc   ====> auto generated folder
│       │   ├── db.go
│       │   ├── models.go
│       │   └── query.sql.go
│       ├── query.sql.   ====> actual queries
│       ├── schema.sql
│       └── sqlc.yaml
├── go.mod
├── go.sum
├── server.go
└── urls_mappings
    └── urls.go      ====> entrypoint


- migration folder and Makefile are parallel to docker-compose.yaml.
- once you build docker images using 'docker-compose up' command, run 'make migrate-up' to create tables in docker images.
- [or simlpy go to individual microservice and start server]


================================================


- used optimistic locking for updating product count using version column.
- used goroutines for async call
- used jwt tokens, and added to 'Authorization' header
- for authentication, function/middleware is added in urls.go
- for authorization, it should be authenticated and used for user's user-type (check used in product/service/service.go -> AddProduct). same way you can use for perticular user's email/id to check that same user is allowed for GetOrder. [casbin can be used for authorization, i already implemented in my past company]


deployment:
- make migrate-up
- docker-compose up

testing apis:
- curl --location 'http://localhost:8001/user-auth/hello'
- curl --location 'http://localhost:8002/product/hello'
- curl --location 'http://localhost:8003/order-management/hello'

you can check other urls in url_mappings/urls.go in each microservice. (do not forget to add url-prefix. like 'user-auth', 'product', 'order-management')