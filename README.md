# Go-To-Do-List-Backend
Backend REST API for To-Do List in Go and MySQL 

Database Stuff -  
```
CREATE TABLE `todo_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(45),
  `description` longtext,
  PRIMARY KEY (`id`)
) 
```

APIs developed -  

GET `/getAllItems` - To get all TODOs in the Database.   
Request - 
```
{
}
```

Response - 
```
{
    "data": [
        {
            "id": 5,
            "title": "Watch Sacred Games S01",
            "description": "I'm already late!"
        },
        {
            "id": 6,
            "title": "Watch Sacred Games S02 ",
            "description": "Can't be late for this!"
        }
    ],
    "message": "Success",
    "responseCode": 200
}
```

POST `/deleteItem`-   
Request - 
```
{
    "id": 123
}
```
Response - 
```
{
    "message": "DeleteItem",
    "responseCode": 200,
    "rowsAffected": 1
}
```

POST `/addItem`
Request - 
```
{
    "name": "Title of the TODO ",
    "description": "Some description on that"
}
```
Response - 
```
{
    "insertedId": 7,
    "message": "AddItem",
    "responseCode": 200
}
```

### Todos

- Add Created Date and Updated Date in the Database.
- Organize Code into different packages. 

### Go-Libraries Used

- database/sql
- go-gin
- go-sql-driver/mysql


### Development

Want to contribute? Great!
