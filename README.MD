# Online Store Backend API

This project is a backend application for an online store, developed using Go and PostgreSQL. The backend focuses on providing RESTful APIs for the online store's functionality. The MVP (minimum viable product) criteria include the following features:

- Customer can view the product list by product category.
- Customer can add products to the shopping cart.
- Customers can see a list of products that have been added to the shopping cart.
- Customer can delete products from the shopping cart.
- Customers can checkout and make payment transactions.
- Customer login and registration.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/akmaliatrias08/home-test-shop.git
   ```
2. Navigate to the project directory:
    ```sh
   cd home-test-shop
    ```
3. Install dependencies:
    ```sh
   go mod tidy
    ```
4. Set up ```.env``` file by copy ```.env.example``` file that exist in repository
  ```
  APP_PORT=5000
  JWT_SECRET=93f3f72d-c80c-4de8-ac26-7e3eaa9fc9df //uuid

  DB_HOST=127.0.0.1
  DB_USER=postgres
  DB_PASSWORD=***
  DB_NAME=online_shop_home_test
  DB_PORT=5432
  ```
## Usage
### Accessing Swagger Documentation
This project includes Swagger for API documentation. To access the Swagger UI, follow these steps:
  1. Make sure the application is running:
  ```sh
  go run main.go
  ```
  2. Open your web browser and navigate to:
   ```sh
  http://localhost:8080/docs/index.html 
  ```
  The port is ```8080``` actually the ```APP_PORT``` in the ```.env``` file.


