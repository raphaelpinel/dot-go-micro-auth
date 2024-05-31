# dot go micro broker
Microservices with Go at the dot conf

## Branches
Start with `step-1-auth-start`, then `step1-auth-solution`

## Steps
- initialize the module with `go mod init authentication`
Note that you may need to manually change the version in the `go.mod` file from the format `1.22.0` to `1.22`

- import the router packages:
    - `go get github.com/go-chi/chi/v5`
    - `go get github.com/go-chi/chi/v5/middleware`
    - `go get github.com/go-chi/cors`
- import the best recommended driver for Postgres:
    - `go get github.com/jackc/pgconn`
    - `go get github.com/jackc/pgx/v4`
    - `go get github.com/jackc/pgx/v4/stdlib`

## How to run the docker image
- Start Docker on your computer
- Clone the project `https://github.com/raphaelpinel/dot-go-micro-project`, be sure to checkout a branch that includes the authentication service, like `step-2-docker-compose-for-authentication-solution` inside a folder named `project` folder and run the command `docker-compose up -d`
