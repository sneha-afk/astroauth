# astroauth api

## Register user

```json
POST /v1/register
```

- **Headers**:  
  - `Content-Type`: `application/json` (required)  
- **Request Body**:  
  Provide user details in JSON format.  

#### Request Fields:  

| Field         | Type     | Required | Description                       |
|---------------|----------|----------|-----------------------------------|
| `username`    | `string` | Yes      | Desired username (up to 25 characters) |
| `email`       | `string` | Yes      | User email address         |
| `password`    | `string` | Yes      | Desired password (8 to 72 characters) |

## Login

```json
POST /v1/login
```

- **Headers**:  
  - `Content-Type`: `application/json` (required)  
- **Request Body**:  
  Provide login attempt in JSON format.

Usernames are forced to be unique, so login with either username or email.

| Field         | Type     | Required | Description                       |
|---------------|----------|----------|-----------------------------------|
| `username`    | `string` | If not with email      | Attempt to login with username |
| `email`       | `string` | If not with username   | Attempt to login with email    |
| `password`    | `string` | Yes      | Attempt with this password |
