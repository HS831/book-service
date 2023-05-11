# book-service
A RESTful web-api service which contains all the basic CRUD APIs for any book-service application. This project uses Go and Gin as it's web framework.

### API Endpoints
1. GET /books :- Get all books
2. GET /books/:id :- Get only single book with given ID
3. POST /books :- Create/Post new Books
4. PATCH /books/:id :- Updates the given book with given ID
5. DELETE /books/:id :- Delete the book with given ID


### Technologies Used :- 
1. Go 
2. Gin Web Framework

### Steps to run this application on local machine :- 
```
STEP 1: Open git bash or git CLI in your machine

STEP 2: Clone this repository using command - 
``` git clone https://github.com/HS831/book-service.git```
        
STEP 3: Open the project folder or directory in VS Code or any other editor.

STEP 3: Open the VS code terminal and install all the dependencies using command-
``` go get . ```
        
STEP 4: Finally, run the command -
``` go run main.go ```
```
Your Application will be up and running on the http://localhost:3000 server.


### API Documentation :- 
https://documenter.getpostman.com/view/17373422/2s93eYUBmF

### Postman API Collection :- 
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/17373422-be58a9e3-0cfe-44fa-8a5a-246668075b5d?action=collection%2Ffork&collection-url=entityId%3D17373422-be58a9e3-0cfe-44fa-8a5a-246668075b5d%26entityType%3Dcollection%26workspaceId%3D8b7ad081-e86e-4420-a3e2-fa4ced2b1851)


