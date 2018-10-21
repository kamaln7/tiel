## 🐦

```go
// error handling omitted //

// Create the tiel instance
t, err := tiel.New(nil, nil)

// Configure routes
t.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("chirp chirp"))
})

// Start the server
err = t.Sing()
```

↓

```
~ $ curl localhost:3000
chirp chirp
```
