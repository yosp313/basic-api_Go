# Basic API

This is a basic API built with Go. It uses the `gorm` package for ORM and SQLite for the database.

## Structure

The project is structured as follows:

- `configs/`: Contains configuration files.
- `controllers/`: Contains controller files.
- `middleware/`: Contains middleware files.
- `models/`: Contains model files.
- `routes/`: Contains route files.
- `services/`: Contains service files.
- `utils/`: Contains utility files.
- `server.go`: The main entry point of the application.

## Setup

1. Install Go.
2. Clone this repository.
3. Run `go tidy` to download the necessary dependencies.
4. Run `go run server.go` to start the server.

## Usage

The API has the following endpoints:

- `GET /`: Shows the name of the url path
- `GET /ip`: Shows info about your ip address
- `POST /product`: Create a new product.
- `GET /product`: List all products.
- `GET /product/{id}`: Find a product by ID.
- `PUT /product/{id}`: Update a product by ID.
- `DELETE /product/{id}`: Delete a product by ID.

## To Do

- [ ] Make a routes file to manage all the routes for the products route.
- [ ] Handle the request errors to prevent unexpected behaviours.
