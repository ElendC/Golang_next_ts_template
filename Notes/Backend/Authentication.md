# JSON WEB TOKEN (JWT)
https://jwt.io/#debugger
## What is it?
- Digitally signed information using HMAC, RSA or ECDSA.
- Used for Authorization and information exchange.
- Consist of: 
  - Header: Contains meta-data about the token.
  - Payload: Claims we define in the code (statements about the entity: username, roles, expiration etc.)
  - Signature: Secret key defined to validate token (From webProg: "Some secret"). Should be long and random, and maybe changed once in a while.

## How it works
- When user successfully logs in, a JTW will be returned to the user.
- When user access a protected route/resource, the user sends the JWT with the request.
- Typically, in the **Authorization** header using the **Bearer** schema
- Content of header should look like: `Authorization: Bearer <token>`

## Implementing Logic
### Creating the token
- **Location**: `backend/utils/jwt.go` -> `func GenerateJWT()`


- **Header** is created by the Signing method specified from jwt.NewWithClaims(...)


- **Payload** added with `type Claims`.
- Custom claims can be added as new fields.
- `jwt.RegisteredClaims` should contain the standard JWT claims (Ex: `exp`, `iat`, `nbf`)
- Benefits of using `jwt.RegisteredClaims`** includes: 
  - Helper methods (`jwt.NewNumericDate` etc.).
  - Handles the claims automatically, error reduction and more!


- **Signature** is added at the return, with the secret key as parameter.
- Requests from the frontend will contain this key, authenticating the user.
- When using in the frontend:
  - Least safe method.: `localStorage.setItem("token", "your-jwt-token");`
  - Only stored during session: `sessionStorage.setItem("token", "your-jwt-token");` 
  - Most secure: Send the token as HTTP-only cookie. Browser will include it automatically with every request.


## getSecretKey
- Gets the secret key stored inside a variable (env or somewhere safe), if and only if the 
signing method is the expected signing method.
- Enforcing algorithm type this way prevents clients to manipulate the 'alg' field in the JWT header
- This was for example set to 'none' in older versions, which would bypass the verification logic.
    
### Verifying the token
- Checks if the JWT returned from the frontend is valid.
- Uses the `getSecretKey` to check if client key is the same as server key




# Login route

# Register route