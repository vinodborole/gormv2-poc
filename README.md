# GoRM V2 POC
This project showcases GoRM v2 features

GoRM features include
1. Logging [Two Options]

    a. Into a separate file    
    b  Using existing logrus     
2. Connection Pooling
3. Transaction and Nested transaction
4. Transactions with save points
5. Bulk Insert 

Apart from GoRM features, it also make use of https://github.com/smallnest/gen library to auto generate gorm models

Core following features

1. Openapi v3 integration
2. contextual, structured logging
3. GoRM V2 - [example with mysql dialect]
4. graceful shutdown
5. Cron jobs
6. Configuration through simple Yaml
7. Build, Lint, Fmt through a common shell script


# Usage

### Step 1

Clone the project set the GOPATH to git clone directory and resolve dependency using dep ensure library management

```
#> git clone https://github.com/vinodborole/gormv2-poc.git

#> cd gormv2-poc/src/app
#gormv2-poc/src/app> dep ensure

```
### Step 2
Modify /gormv2-poc/yaml/config.yaml as per your requirement

### Step 3
Create Empty Database

```
#>mysql -u root -p
Enter password: 
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MySQL connection id is 84882
Server version: 8.0.21 MySQL Community Server - GPL

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MySQL [(none)]>CREATE DATABASE myapp;
Query OK, 1 row affected (0.024 sec)
```
### Step 4
Run the schema script from etc/db.sql to generate tables

### Step 5
Build the Project [This will also autogenerate the DB models based on the database schema]
```
#> cd gormv2-poc/scripts
#gormv2-poc/scripts> sh build.sh

```
### Step 6

Execute App
```
#gormv2-poc/bin> ./app

```

### Step 7

Execute REST API
```
#>curl localhost:8080/v1/about
{"name":"gormv2-poc","description":"Gorm version 2 POC","url":"http://localhost","port":"8080"
```


# Structure

This structure is inspired from clean code architecture

```
├── src // put all handlers here.
│   ├── app
│   │   └── config
│   │       └── yaml 
│   │       │   └── config.yaml //define custom config for the app
│   │       └── config.go  //config read
│   ├── domain
│       └── domain.go //domain objects if any can be defined here
│   ├── gateway
        └── appcontext 
            └── appContext.go // context management across app
        └── appRepository.go  // implementation for database opeartion on model object
        └── databaseRepository.go // implementation for common database operations commit/rollback/getconnection
│   ├── infra
        └── constants //define constants used in app
        └── database
            └── model
                └──*.go //define database models/tables struct
            └── operations.go //database connection/transaction management framework
        └── logging
             └── auditlog.go // framework for audit log on every request
        └── rest
            └── converter 
                └── modelConverter.go //convert DB data to rest api response domain object
            └── generated //swagger gen code 
            └── handler
                └── logger.go // logger initializer
                └── myApp.go // rest api handler
                └── responseHandler.go   //marshal and respone of type success/error
            └── api.yaml //openapi v3 yaml - restapi definition
            └── apiserver.go // rest api server listener
            └── routers.go  //define navigation for rest apis
│   ├── usecase //business login for usecases/features supported by application 
        └── interactorinterface
            └── databaseRepository.go  //interface to interact with DB handlers
│   ├── utils
        └── httpUtils.go //json response generator 
    ├── main.go //entry point 

```

# Libraries
Few of the libraries that I am using as part of this project are
1. gorm.io/driver/mysql - Mysql dialect to communicate with mysql DB
2. github.com/google/uuid - Unique ID library used in audit log framework
3. github.com/gorilla/mux - Used as http server listener
4. gorm.io/gorm - Used as an orm library to talk to database
5. github.com/sirupsen/logrus - logging library
6. gopkg.in/yaml.v2 - Yaml config reader
