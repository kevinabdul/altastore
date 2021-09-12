# Alta Store Task 

## Endpoint
## Users
| Method | Endpoint | Description|
|:-----|:--------|:----------|
| GET    |   /users||
| GET   |   /users/:id|
| GET   |   /users/:id/carts| Get cart by user's id. id paramater must be equal with user_id in jwt
| POST  |   /users|
| PUT   |   /users/:id| Update User by id. id paramater must be equal with user_id in jwt
| PUT   |   /users/:id/carts| Update cart by user's id. id paramater must be equal with user_id in jwt
| DELETE|   /users/:id

### Products
| Method | Endpoint | Description|
|:-----|:--------|:----------|
|GET   |  /products| Get all product. Can be filtered by category query



### Checkout
| Method | Endpoint | Description|
|:-----|:--------|:----------|
|GET   |     /checkout
|POST  |  /checkout

