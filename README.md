# logIn

login/register in Golang with GIN framework (API as Backend)

## Installation

Need to have GoLang version 1.13 or above.

To check
```bash
go version
```

To install all packages, go inside repo and use: 
```bash
go get
```

## Usage

To start this API, use:
```bash
go run main.go
```

## Setup database:
Sql query for creating database is provided in db folder, And also, if needed db dump is provided in dump folder inside db.
(psql)

## End points of API

For register:
http://localhost:8000/

Body:
{
	"username":"tanmay",
	"email":"tanmay@piy.com",
	"password":"abc"
}

For login:
http://localhost:8000/login
Body:
{
	"email":"tanmay@piy.com",
	"password":"abc"
}

