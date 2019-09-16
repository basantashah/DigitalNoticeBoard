
# DigitalNoticeBoardLogin
    # Under Development
    > This is the api used for digital notice board

###### Register new user
```
{
    "email": "Basanta.Shah@test.net",
    "password": "test@123"
}
```

###### Postnotice
```
{
    "title": "this is the title",
    "expiry": "2018-09-22T12:42:31+07:00",
    "subject": "New routing published",
    "content": "Please follow the new routing published in website",
    "department": "RTE department",
    "urgent": "TRUE",
    "status": "TRUE"
}
```

###### The tree diagram of the directory is
```
.
├── README.md
├── app
│   ├── auth.go
│   ├── errors.go
│   └── middleware.go
├── controllers
│   ├── authControllers.go
│   └── noticesControllers.go
├── go.mod
├── go.sum
├── main.go
├── models
│   ├── accounts.go
│   ├── base.go
│   └── notice.go
└── utils
    └── util.go

4 directories, 13 files
```