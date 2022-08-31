## General info
Go-Api adalah simple project CRUD menggunakan GoLang


## Technologies
Project ini menggunakan beberapa teknologi yaitu :
* Golang Versi 1.18
* Echo Labstack 
* GORM
* Viper
* Postgre SQL

## Setup
Sebelum menjalankan Api, pastikan PostgreSQL sudah running dan buat sebuah database dengan nama **go_api** dan atur credentials database yang ada di **config/development.json**

Postman collection: [link](https://www.getpostman.com/collections/f181256659ef69f180e3)

Berikut langkah-langkah untuk running repo ini

```
$ git clone https://github.com/tony-hr/go-api.git
$ go mod download
$ go run main.go
```

## Migration

Untuk menjalankan migration dengan seeder

```DBEVENT=migrate go run main.go```
