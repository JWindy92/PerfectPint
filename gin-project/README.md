### Run Service
`go run cmd/gin-project/main.go`

### API

`GET http://localhost:8000/albums`

```
[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    }
]
```

`GET http://localhost:8000/albums/[id]`

```
{
    "id": "3",
    "title": "Sarah Vaughan and Clifford Brown",
    "artist": "Sarah Vaughan",
    "price": 39.99
}
```


`POST http://localhost:8000/albums`

BODY

```
{
    "id": "4",
    "title": "The New Album for the API",
    "artist": "John Wid",
    "price": 39.99
}
```

RESPONSE

```
{
    "id": "4",
    "title": "The New Album for the API",
    "artist": "John Wid",
    "price": 39.99
}
```