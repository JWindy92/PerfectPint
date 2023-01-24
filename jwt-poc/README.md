## JWT Proof of Concept

### App flow

1. User Signs up via json POST
```
{
    "name": "John",
    "email": "john@pp.com",
    "password": "1234",
    "role": "admin"
}
```
2. User logs in via email/password
```
{
    "email": "john@pp.com",
    "password": "1234"
}
```
3. If sign in is successful, a JWT is returned
```
{
    "role": "admin",
    "email": "john@pp.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImpvaG5AcHAuY29tIiwiZXhwIjoxNjc0NTg5OTQ4LCJyb2xlIjoiYWRtaW4ifQ.rCQ984wDMbSWdJNyxpT5nhSzSl5PdFCvdlDVr0xiXqw"
}
```
4. Token is stored locally, to be passed to further requests in Token header

5. Further requests are passed through authentication middleare which checks whether or not the user is authorized to view the route
