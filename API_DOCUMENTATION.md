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
| `username`    | `string` | Yes      | Desired username   |
| `email`       | `string` | Yes      | User email address         |
| `password`    | `string` | Yes      | Desired password            |
