![alt tag](https://upload.wikimedia.org/wikipedia/commons/2/23/Golang.png)

[![License](https://img.shields.io/github/license/Massad/gin-boilerplate)](https://github.com/Geekinn/go-micro/blob/master/LICENSE) [![GitHub release (latest by date)](https://img.shields.io/github/v/release/Massad/gin-boilerplate)](https://github.com/Geekinn/go-micro/releases) [![Go Version](https://img.shields.io/github/go-mod/go-version/Massad/gin-boilerplate)](https://github.com/Geekinn/go-micro/blob/master/go.mod) [![DB Version](https://img.shields.io/badge/DB-PostgreSQL--latest-blue)](https://github.com/Geekinn/go-micro/blob/master/go.mod) [![DB Version](https://img.shields.io/badge/DB-Redis--latest-blue)](https://github.com/Geekinn/go-micro/blob/master/go.mod)

[![Build Status](https://travis-ci.org/Massad/gin-boilerplate.svg?branch=master)](https://travis-ci.org/Massad/gin-boilerplate) [![Go Report Card](https://goreportcard.com/badge/github.com/Geekinn/go-micro)](https://goreportcard.com/report/github.com/Geekinn/go-micro)


Welcome to **Geekinn Golang Gin boilerplate**

The fastest way to deploy a restful api's with [Gin Framework](https://github.com/gin-gonic/gin/) with a structured project that defaults to **PostgreSQL** database and **JWT** authentication middleware stored in **Redis** following Laravel structural guidelines and addition of wellknown golang packages.

## Configured with

- [gorm](https://gorm.io/): Go Relational Persistence
- [jwt-go](https://github.com/golang-jwt/jwt): JSON Web Tokens (JWT) as middleware
- [go-redis](https://github.com/go-redis/redis): Redis support for Go
- Go Modules
- Built-in **Custom Validators**
- Built-in **CORS Middleware**
- Built-in **RequestID Middleware**
- SSL Support
- Enviroment support
- Unit test
- Friendly Structure
- **Cobra** CMD added
- **Air** added for autoreload
- **Dlv** added for debugging
- Docker with Docker-compose added
- HTTP Request Sample Attached with Circuit Breaker
- Pagination Added
- **Prometheus** Added
- **Logrus** Added for access logging
- And few other important utilties to kickstart any project

### Installation

```
$ git clone https://github.com/ahsansalaldaha/Geekinn-Golang-Gin-Boilerplate.git
```

## Running Your Application

Rename .env_rename_me to .env and place your credentials

```
$ mv .env_rename_me .env
```

Generate SSL certificates (Optional)

> If you don't SSL now, change `SSL=TRUE` to `SSL=FALSE` in the `.env` file

```
$ mkdir cert/
```

```
$ sh generate-certificate.sh
```

> Make sure to change the values in .env for your databases

## Running your application

```
$ docker-compose up -d
```
this will auto build and run the application for **DEVELOPMENT MODE**. 

## Testing Your Application

```
$ go test -v ./tests/*
```

## Import Postman Collection (API's)

Download [Postman](https://www.getpostman.com/) -> Import 

Inside path **postman**

Includes the following:

- User
  - Login
  - Register
  - Refresh Token
  - Logout
- Article
  - Create
  - Update
  - Get Article
  - Get Articles
  - Paginate Articles
  - Delete
- API
  - TODO
  - Google (Http request)
- Prometheus

> In Login request in Tests tab:

```
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);

    var jsonData = JSON.parse(responseBody);
    pm.globals.set("token", jsonData.token.access_token);
    pm.globals.set("refresh_token", jsonData.token.refresh_token);

});
```

It captures the `access_token` from the success login in the **global variable** for later use in other requests.

Also, you will find in each request that needs to be authenticated you will have the following:

    Authorization -> Bearer Token with value of {{token}}

It's very useful when you want to test the APIs in Postman without copying and pasting the tokens.

## On You

You will need to implement the `refresh_token` mechanism in your application (Frontend).

> We have the `/v1/token/refresh` API here to use it.

_For example:_

If the API sends `401` Status Unauthorized, then you can send the `refresh_token` that you stored it before from the Login API in POST `/v1/token/refresh` to receive the new `access_token` & `refresh_token` and store them again. Now, if you receive an error in refreshing the token, that means the user will have to Login again as something went wrong.

That's just an example, of course you can implement your own way.

## Contribution

You are welcome to contribute to keep it up to date and always improving!

## Credit
HUGE SHOT OUT to **Massad** for inspiration.

The implementation of boilerplate stems with
[Massad/gin-boilerplate](https://github.com/Massad/gin-boilerplate) boilerplate with addition of coding structure from [Laravel](https://laravel.com/) with addition of wellknown packages for golang.

The implemented JWT inspired from this article: [Using JWT for Authentication in a Golang Application](https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr) worth reading it, thanks [Victor Steven](https://medium.com/@victorsteven)

---

## License

(The MIT License)

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


## SETUP Database
```
psql -U postgres
create database golang_gin_db;
\q

psql -U postgres -d golang_gin_db;
\dt
\q
```

## ROAD MAP
- [x] Add Validator https://github.com/go-ozzo/ozzo-validation 
- [x] Add Live Reload
- [x] Add GORM https://github.com/go-gorm/gorm
- [x] Add Delve Debugger
- [x] Add GORM Logger - Already implemented
- [x] Upgrade Golang & Gin to the latest
- [x] Add Prometeus Logger
- [x] Add Logrus Logger for file based logs
- [x] Add Article Pagination
- [x] Add Cmd
- [x] Add HTTP Request Package
- [ ] Proper Unit Testing of components


## Remove Unsed dependencies
First add your upgraded dependency and then do this which will remove old dependency if the previous dependency is not being used anymore
```
go mod tidy -v
```

## Upgrade Go
Run following commands to existing container
```
go mod edit -go=1.18
go mod tidy -v
```