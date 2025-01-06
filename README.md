# go-gin-api

## How to Run

### 1. Configure Database Info
You will need to fill in these fields in the `.env` file:

- **PublicHost**:            
- **Port**:                  
- **DBUser**:                 
- **DBPassword**:             
- **DBAddress**:              
- **DBName**:                 
- **JWTExpirationInSeconds**: 
- **JWTSecret**:              

To know the default values in case `.env` does not have values, refer to `env.go`.

---

### 2. Get All Dependencies
Run the following command to fetch all necessary dependencies:

```bash
make get
```

---

### 3. Run the app

```bash
air
```


## Project Overview

### `main.go`
1. **Create Database**: Initializes the database connection.
2. **Create API Server**:
   - Integrates the database.
   - Specifies the port the API will run on.
3. **Run the Database**: Starts the database service.

### `api.go`
1. **API Struct**: Defines the structure of the API.
2. **Router and Subrouter**:
   - Creates a main router.
   - Configures subrouters (if needed) for different endpoints.
3. **Handler Input**:
   - Each service has its own store and handler.
   - The handler receives the subrouter as input.

### Routing
- A **Mux Router** matches the request URL to the corresponding handler.

### Handlers
- The handler contains an interface that performs specific tasks.
- These tasks are associated with the database.
- **Note**: The handler does not directly contain the database; instead, it uses an interface.

### `store.go`
- Implements all methods defined in the handler interface.
- Contains the database and provides methods to interact with it.


## Project Structure

```bash
.
├── Makefile
├── README.md
├── cmd
│   ├── api
│   │   └── api.go
│   ├── config
│   │   └── env.go
│   ├── db
│   │   └── db.go
│   ├── main.go
│   ├── migrate
│   │   └── migrations
│   │       ├── 20250102041039_add-user-table.down.sql
│   │       ├── 20250102041039_add-user-table.up.sql
│   │       ├── 20250102042144_add-product-table.down.sql
│   │       ├── 20250102042144_add-product-table.up.sql
│   │       ├── 20250102042211_add-orders-table.down.sql
│   │       ├── 20250102042211_add-orders-table.up.sql
│   │       ├── 20250102042253_add-orders-items-table.down.sql
│   │       ├── 20250102042253_add-orders-items-table.up.sql
│   │       ├── 20250102072343_add-products-to-products-table.down.sql
│   │       ├── 20250102072343_add-products-to-products-table.up.sql
│   │       └── main.go
│   ├── services
│   │   ├── auth
│   │   │   ├── jwt.go
│   │   │   ├── password.go
│   │   │   └── password_test.go
│   │   ├── cart
│   │   │   ├── routes.go
│   │   │   └── service.go
│   │   ├── health
│   │   │   └── routes.go
│   │   ├── order
│   │   │   └── store.go
│   │   ├── product
│   │   │   ├── routes.go
│   │   │   ├── routes_test.go
│   │   │   └── store.go
│   │   └── user
│   │       ├── routes.go
│   │       ├── routes_test.go
│   │       └── store.go
│   ├── types
│   │   └── type.go
│   └── utils
│       └── util.go
├── go.mod
├── go.sum
```