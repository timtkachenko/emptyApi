# emptyApi
project is based on open-api `swagger-go` and uses go modules

### api spec
located in `api-schema.yaml`

### db

add password and username to db dsn`./config/config.go`

simple structure 
```
CREATE TABLE public.tests (
    id integer NOT NULL,
    name character varying
);
```

### test 

test are trough multiple packages as there is not much of functionality and they will be too granular.

`go test ./...`


## Run

build `go build -o empty /srv/gospace/src/emptyApi/cmd-server/main.go`

optionally `./empty --port 9999`
 