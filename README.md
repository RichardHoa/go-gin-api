# go-gin-api

`main.go`: 
1. Create database
2. Create API server and put the database, decide the port the API will run
3. Run the database

`api.go`:
1. Create API struct
2. Create Router and subrouter (if any)
3. Each service will have it own store, and handler, the input of handler would be the subrouter

A mux router match the request url to the corresponding handler

The handler contains an interface that can do certain task, these tasks are associated with datbase, but the handler does not contains the database. 

Because it's an interface. the store.go will satisfy all the methods. the store.go will contains the database with all them methods to interact with the database
