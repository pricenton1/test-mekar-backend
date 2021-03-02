## How to run this project 

1. configure environment in Makefile or config.env
2. import db_user.sql
3. run the project

# build project
$ go build -o main 

# Run app 
go run main.go
or
make run

# FrontEnd Application
https://github.com/pricenton1/test-mekar-frontend
And setting package.json
"proxy": "http://localhost:8000"

## API Spec

GET USERS

    URL : http://localhost:8000/users?keyword=&page=0&limit=5
    HEADERS :
        key : token
        value :eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21LZXkiOiJSYWhhc2lhIiwiZXhwaXJlZEF0IjoiMjAyMS0wMy0wMSAwNzoyNjozNyIsIm5hbWUiOiJhbnRvbkBnbWFpbC5jb20ifQ.TBvcI3LUlChXi1zAPfFnjxSpyQj-d6Wr9mHEoI2rL5c
    RESPONSE : {
        "statusCode": 200,
        "message": "OK",
        "payload": [
            {
                "userID": "db179baa-559b-4967-b75c-9e316d6564af",
                "userNIK": "1234567891123",
                "userName": "UJI COBA 2",
                "userBirth": "2000-03-09",
                "userJob": {
                    "jobID": "8219c7cc-30bb-11eb-b405-c85b766bafe8",
                    "jobLabel": "PNS"
                },
                "userEducation": {
                    "educationID": "5df1331d-30bb-11eb-b405-c85b766bafe8",
                    "educationLabel": "Sarjana"
                },
                "userStatus": "1",
                "userCreated": "2021-02-28 21:55:05"
            },
            {
                "userID": "8a95efea-e6d6-47c3-88e7-d57b03f9712e",
                "userNIK": "3209120503980005",
                "userName": "UJI COBA",
                "userBirth": "1998-03-05",
                "userJob": {
                    "jobID": "82225b07-30bb-11eb-b405-c85b766bafe8",
                    "jobLabel": "Wiraswasta"
                },
                "userEducation": {
                    "educationID": "5dedbec5-30bb-11eb-b405-c85b766bafe8",
                    "educationLabel": "Diploma"
                },
                "userStatus": "1",
                "userCreated": "2020-12-01 09:41:42"
            }
        ]
    }

GET USER BY ID

    URL : http://localhost:8000/users/8a95efea-e6d6-47c3-88e7-d57b03f9712e
    HEADERS :
        key : token
        value :eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21LZXkiOiJSYWhhc2lhIiwiZXhwaXJlZEF0IjoiMjAyMS0wMy0wMSAwNzoyNjozNyIsIm5hbWUiOiJhbnRvbkBnbWFpbC5jb20ifQ.TBvcI3LUlChXi1zAPfFnjxSpyQj-d6Wr9mHEoI2rL5c
    RESPONSE :{
        "statusCode": 200,
        "message": "OK",
        "payload": {
            "userID": "8a95efea-e6d6-47c3-88e7-d57b03f9712e",
            "userNIK": "3209120503980005",
            "userName": "UJI COBA",
            "userBirth": "1998-03-05",
            "userJob": {
                "jobID": "82225b07-30bb-11eb-b405-c85b766bafe8",
                "jobLabel": "Wiraswasta"
            },
            "userEducation": {
                "educationID": "5dedbec5-30bb-11eb-b405-c85b766bafe8",
                "educationLabel": "Diploma"
            },
            "userStatus": "1",
            "userCreated": "2020-12-01 09:41:42"
        }
    }

CREATE NEW USER

    URL : http://localhost:8000/users/register
    HEADERS :
        key : token
        value :eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21LZXkiOiJSYWhhc2lhIiwiZXhwaXJlZEF0IjoiMjAyMS0wMy0wMSAwNzoyNjozNyIsIm5hbWUiOiJhbnRvbkBnbWFpbC5jb20ifQ.TBvcI3LUlChXi1zAPfFnjxSpyQj-d6Wr9mHEoI2rL5c
    BODY RAW :
            {
                "userNIK": "1234567891123",
                "userName": "UJI COBA 2",
                "userBirth": "2000-03-09",
                "userJob": {
                    "jobID": "8219c7cc-30bb-11eb-b405-c85b766bafe8"
                },
                "userEducation": {
                    "educationID": "5df1331d-30bb-11eb-b405-c85b766bafe8"
                }
            }
    RESPONSE :
        {
            "statusCode": 201,
            "message": "Accepted"
        }

UPDATE USER 

    URL : http://localhost:8000/users/update/dd3542db-2666-405b-9b99-291b99e637b1
    HEADERS :
        key : token
        value :eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21LZXkiOiJSYWhhc2lhIiwiZXhwaXJlZEF0IjoiMjAyMS0wMy0wMSAwNzoyNjozNyIsIm5hbWUiOiJhbnRvbkBnbWFpbC5jb20ifQ.TBvcI3LUlChXi1zAPfFnjxSpyQj-d6Wr9mHEoI2rL5c
    BODY RAW :
        {
            "userNIK": "12345678910",
            "userName": "UJI COBA 3",
            "userBirth": "2000-03-08",
            "userJob": {
                "jobID": "8219c7cc-30bb-11eb-b405-c85b766bafe8"
            },
            "userEducation": {
                "educationID": "5df1331d-30bb-11eb-b405-c85b766bafe8"
            }
        }
    RESPONSE : 
        {
            "statusCode": 200,
            "message": "Ok"
        }

DELETE USER 

    URL : http://localhost:8000/users/delete/dd3542db-2666-405b-9b99-291b99e637b1
    HEADERS :
        key : token
        value :eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21LZXkiOiJSYWhhc2lhIiwiZXhwaXJlZEF0IjoiMjAyMS0wMy0wMSAwNzoyNjozNyIsIm5hbWUiOiJhbnRvbkBnbWFpbC5jb20ifQ.TBvcI3LUlChXi1zAPfFnjxSpyQj-d6Wr9mHEoI2rL5c
    RESPONSE :
        {
            "statusCode": 200,
            "message": "Ok"
        }

GET JOBS 

     URL : http://localhost:8000/jobs
    HEADERS :
        key : token
        value :eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21LZXkiOiJSYWhhc2lhIiwiZXhwaXJlZEF0IjoiMjAyMS0wMy0wMSAwNzoyNjozNyIsIm5hbWUiOiJhbnRvbkBnbWFpbC5jb20ifQ.TBvcI3LUlChXi1zAPfFnjxSpyQj-d6Wr9mHEoI2rL5c
    RESPONSE :
        {
            "statusCode": 200,
            "message": "OK",
            "payload": [
                {
                    "jobID": "8219c7cc-30bb-11eb-b405-c85b766bafe8",
                    "jobLabel": "PNS"
                },
                {
                    "jobID": "821e309f-30bb-11eb-b405-c85b766bafe8",
                    "jobLabel": "Wirausaha"
                },
                {
                    "jobID": "82225b07-30bb-11eb-b405-c85b766bafe8",
                    "jobLabel": "Wiraswasta"
                }
            ]
        }

GET EDUCATIONS

    URL : http://localhost:8000/educations
    HEADERS :
        key : token
        value :eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21LZXkiOiJSYWhhc2lhIiwiZXhwaXJlZEF0IjoiMjAyMS0wMy0wMSAwNzoyNjozNyIsIm5hbWUiOiJhbnRvbkBnbWFpbC5jb20ifQ.TBvcI3LUlChXi1zAPfFnjxSpyQj-d6Wr9mHEoI2rL5c
    RESPONSE :
        {
            "statusCode": 200,
            "message": "OK",
            "payload": [
                {
                    "educationID": "5ddbb91e-30bb-11eb-b405-c85b766bafe8",
                    "educationLabel": "SD"
                },
                {
                    "educationID": "5de17471-30bb-11eb-b405-c85b766bafe8",
                    "educationLabel": "SMP"
                },
                {
                    "educationID": "5de57c7c-30bb-11eb-b405-c85b766bafe8",
                    "educationLabel": "SMA"
                },
                {
                    "educationID": "5dedbec5-30bb-11eb-b405-c85b766bafe8",
                    "educationLabel": "Diploma"
                },
                {
                    "educationID": "5df1331d-30bb-11eb-b405-c85b766bafe8",
                    "educationLabel": "Sarjana"
                },
                {
                    "educationID": "5df4cbfb-30bb-11eb-b405-c85b766bafe8",
                    "educationLabel": "Magister"
                },
                {
                    "educationID": "5df8d140-30bb-11eb-b405-c85b766bafe8",
                    "educationLabel": "Doktor"
                }
            ]
        }

LOGIN ADMIN 

    URL : http://localhost:8000/account/login
    BODY RAW :
        {
            "email": "anton@gmail.com",
            "password":"12345"
        }
    RESPONSE :
        {
            "statusCode": 200,
            "message": "OK",
            "payload": {
                "accountID": "85d6d016-7cc8-4ccf-b6a8-d78714d8cef4",
                "email": "anton@gmail.com",
                "password": "$2a$04$wwTjDKqA6.0OOzI2MQmDO./nx.VyqpA.1jWx6RRblUEHJZEdnvwr.",
                "token": {
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21LZXkiOiJSYWhhc2lhIiwiZXhwaXJlZEF0IjoiMjAyMS0wMy0wMSAxNzoxMjowOCIsIm5hbWUiOiJhbnRvbkBnbWFpbC5jb20ifQ.ytR-Ry4ukeWFhk8PBUuRn_FKsgizmGxuUGBPtsnmL88"
                }
            }
        }

REGISTER ADMIN 

    URL : http://localhost:8000/account/register
    BODY RAW : 
        {
            "email": "anton@gmail.com",
            "password":"12345"
        }
    RESPONSE : 
        {
            "statusCode": 201,
            "message": "Accepted"
        }
