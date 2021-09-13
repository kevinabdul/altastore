# Alta Store

## Endpoint

| Method | Endpoint | Description| Authentication | Authorization
|:-----:|:--------|:----------| :----------:| :----------:|
| GET    |   /users|Get list of all users| Yes | No
| GET   |   /users/:id| Get User By User Id. "id" paramater must be equal with "userId" in jwt | Yes | Yes
| POST  |   /users| Register a new user | No | No
| PUT   |   /users/:id | Update User by User Id. "id" paramater must be equal with "userId" in jwt | Yes | Yes
| DELETE|   /users/:id | Delete User by User Id. "id" paramater must be equal with "userId" in jwt | Yes | Yes
| | | | | |
| | | | | |
| | | | | |
| | | | | |
POST | /login | Login existing user| No | No
| | | | | |
| | | | | |
| | | | | |
| | | | | |
|GET   |  /products| Get all product. Can be filtered by category query | No | No
| | | | | |
| | | | | |
| | | | | |
| | | | | |
GET    | /carts/:id | Get Cart By User Id. "id" paramater must be equal with "userId" in jwt | Yes | Yes
PUT    | /carts/:id | Update Cart By User Id. "id" paramater must be equal with "userId" in jwt | Yes | Yes
DELETE | /carts/:id | Delete Cart By User Id. "id" paramater must be equal with "userId" in jwt | Yes | Yes
| | | | | |
| | | | | |
| | | | | |
| | | | | |
|GET   |  /checkout | Get checkout information. Return value depends on "userId" claims in jwt | Yes | No
|POST  |  /checkout | Checked out User's Cart based on jwt's userId claims. | Yes | No
| | | | | |
| | | | | |
| | | | | |
| | | | | |
|GET   |  /payments | Get all pending payment. Return value depends on "userId" claims in jwt | Yes | No
|POST  |  /payments | Resolves one pending payment | Yes | No

