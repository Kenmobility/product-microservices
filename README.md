# Product microservice API

This service fetches data from GitHub's public APIs to retrieve repository commits, saves the data in a persistent store, and continuously monitors the repository for changes at a set interval.

## Requirements- Golang 1.20+
- PostgreSQL

## Setup1. Clone the repository and download go dependencies
```bash
git clone https://git@github.com:Kenmobility/product-microservices.git
cd product-microservices
go mod tidy
```
## 2. Set up environment variables:
- use the .env.example format to set up a .env file with environmental variables or run the below
```bash
    export DATABASE_HOST="localhost"
    export DATABASE_PORT=5438
    export DATABASE_USER=root
    export DATABASE_PASSWORD=secret
    export DATABASE_NAME=product_microservices_db
```
## 3. Edit makefile
- edit makefile to match your database environmental variables on setup2:

## 4. Run makefile commands 
- run 'make postgres' to pull and run PostgreSQL instance as docker container
```bash
make postgres
```
- run 'make createdb' to create a database
```bash
make createdb
```
## Testing (optional)

Run 'make test' to run the unit tests:
```bash
make test
```
## 5. Start web server
- run 'make server' to start the service
```bash
make server
```
  
