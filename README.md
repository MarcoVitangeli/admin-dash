# admin-dash

Simple SPA to showcase how we can achieve common webapp functionalities using [htmx](https://htmx.org)

## Running
create a local `products.db` empty file

now create your db schema using the `schema.sql` file

```shell
sqlite3 products.db < schema.sql
```

now run it (requires go 1.21 or upper)

```shell
go run ./main.go
```

go to localhost:8080