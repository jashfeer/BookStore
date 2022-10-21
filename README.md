# BookStore

* Simple RESTful API to create, read, update and deleted books
* Postgres database implementad 
* User authentication implemented so only loggined users can access the create, read, update and deleted operations

## Endpoints
### User login
```
POST /login
 Now the application have only single user
 {
    "username":"user"
   "password":"123"
 }
```
### User logout
```
GET /logout
```

### Get All Books
```
GET /books
```
### Get Single Books
```
GET /books{id}
```
### Create Book
```
POST /books
 request sample
 {
    "title":"RED"
   "price":"990"
    "author" :"jashfeer"
 }
```
### Update Book
```
PUT /books/{id}
 request sample
 {
    "title":"RED"
   "price":"990"
    "author" :"jashfeer"
 }
```
### Delete Books
```
DELETE /books/{id}
```


