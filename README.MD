# Gin User Management

An implementation for Restful APIs for user managements 
## Structure


```
api/v1        
bootstrap/
lib/
migrations/
models/
public/
repository/
services/
static/
```


## Prerequisites

Make sure you have installed all the following prerequisites on your machine:
* [GO](https://golang.org/doc/install)
* [Mysql](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)

``` bash
$ composer require :vendor/:package_name
```
## Install and Run
 
To run the application you'll need to create a database and copy `.env.example` file into `.env` and update the DB credentials,
also you need to add the DB credentials into the connection string in `dbconfig.yml`.
Once changed you'll need to run the `migrate.sh` file with the following command to create the users table in the database

```shell
bash migrate.sh up
```


## Usage
To start the application run the command
``` shell
go run server.go
```