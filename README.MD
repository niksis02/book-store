# Book Store application

Book Store is a simple Go back-end application to order/list books.

## Getting Started

### Run the application with docker:

Copy .env file

```
cp .env.sample .env

```

Run the application

```
make up
```

### Run the application without docker

1. Build the app

```
go build
```

2. Create a postgres database instance.
3. Configure environment variables in .env file.
4. Run the application

```
./book-store
```
