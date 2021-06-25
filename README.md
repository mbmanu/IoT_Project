# [IOT PROJECT - GODS EYE](https://github.com/mbmanu/IoT_Project)
Doorstep monitoring system built using Raspberry Pi, Python and OpenCV. 

## Requirements
- Python 3.8
- [opencv-python](https://pypi.org/project/opencv-python/)
- [opencv-contrib-python](https://pypi.org/project/opencv-contrib-python/)

## Training the model
```
    $ python3 face_training.py
```
## Run - recognizes face and updates the database
```
    $ python3 face_recog.py
```

# [IOT PROJECT API](https://github.com/shujanpannag/iot_project_api)
This is a REST API built for IOT project using Golang, Gin and Gorm.
## Requirements
- Golang
- [Go Gin](https://pkg.go.dev/github.com/gin-gonic/gin)
- [Gorm](https://pkg.go.dev/gorm.io/gorm)
## Environment variables
- HOST : Hostname
- DBPORT : PostgreSQl Server Port
- DBUSER : Database User Name
- DBNAME : Database Name
- DBPASSWORD : Database Password
- EMAIL_ADDR : From Email
- EMAIL_PSWD : From Email Password
- EMAIL_HOST : SMTP Host
- EMAIL_PORT : SMTP Port
## Install
```
    $ go build
```
## Run the app
```
    $ go run cmd/main.go
```
## REST API
The REST API Methods to the above app is described below.
### Node Log
`GET /v0/nodelog/:date`
``` 
    curl -i -H 'Content-Type: application/json' http://localhost:8080/v0/nodelog/2000-06-28
```
`POST /v0/nodelog/`
``` 
    curl -i -H 'Content-Type: application/json' -d '{"relname": " " , "ipaddr": " ", "datetime": " "}' http://localhost:8080/v0/nodelog/
```

### User Log
`GET /v0/userrel/:rel`

``` 
    curl -i -H 'Content-Type: application/json' http://localhost:8080/v0/nodelog/name
 ```

`POST /v0/userrel/`

``` 
    curl -i -H 'Content-Type: application/json' -d '{"name": " ", "rel": " ", "email": " "}' http://localhost:8080/v0/userrel
```

`PUT /v0/userrel/:name`

 ``` 
    curl -i -H 'Content-Type: application/json' -d '{"name": " ", "rel": " ", "email": " "}' http://localhost:8080/v0/name
```
    
`POST /v0/userrel/:name`

``` 
    curl -i -H 'Content-Type: application/json' http://localhost:8080/v0/name
```

### Server Log
`GET /v0/getlog`

``` 
    curl http://localhost:8080/v0/getlog
```
## Contributions
Feel free to contribute. You know the drill fork and request a PR ;)
