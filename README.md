## General info
Go-Api adalah simple project CRUD menggunakan GoLang


## Technologies
Project ini menggunakan beberapa teknologi yaitu :
* Golang Versi 1.15
* Echo Labstack 
* GORM
* Viper
* MySQL

## Setup
Sebelum menjalankan Api, pastikan MySQL sudah running dan buat sebuah database dengan nama **go_api** dan atur credentials database yang ada di **config/development.json**

Postman collection: [link](https://www.getpostman.com/collections/2edf41dc3d3959190125)

Berikut langkah-langkah untuk running repo ini

```
$ git clone https://github.com/tony-hr/go-api.git
$ go mod download
$ go run main.go
```

## Migration

Untuk menjalankan migration dengan seeder

```DBEVENT=migrate go run main.go```
