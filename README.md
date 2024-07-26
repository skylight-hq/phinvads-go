# phinvads-go

PHIN VADS written in Go. Load `phinvads.dump` into a PostgreSQL database using `pg_restore`, run the app with [air](https://github.com/air-verse/air), and go!

### Dev setup

1. Clone this repo:

    ```bash
    git clone https://github.com/skylight-hq/phinvads-go.git
    cd phinvads-go
    ```

1. Install and run [PostgreSQL](https://www.postgresql.org/download/), or just run `docker compose up`
1. Create an empty database:

    ```bash
    psql -c 'CREATE DATABASE phinvads'
    ```

1. Load the database dump file:

    ```bash
    pg_restore -d phinvads --no-owner --role=$(whoami) phinvads.dump
    ```

1. Install [Go](https://go.dev/doc/install)
1. Install [air](https://github.com/air-verse/air):

    ```bash
    go install github.com/air-verse/air@latest
    ```

1. Install [mkcert](https://github.com/FiloSottile/mkcert)
1. Create your own self-signed certs for local development:  

    ```bash
    cd tls
    mkcert -install
    mkcert localhost
    cd ..
    ```

1. Run the app!

    ```bash
    air
    ```
