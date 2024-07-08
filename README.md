# Coffee App

The Coffee App is a web application that allows users to register as sellers or buyers. Sellers can manage their products, and buyers can place orders for products.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [User Management](#user-management)
  - [Product Management](#product-management)
  - [Order Management](#order-management)
- [Usage](#usage)
- [License](#license)

## Features

- User registration (Buyer/Seller)
- Product management (Create, Read)
- Order management (Create, Read)

## Prerequisites

- Go 1.19 or later
- Docker
- Docker Compose

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/manimovassagh/coffee-app.git
    cd coffee-app
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Set up the database with Docker:

    ```sh
    docker-compose up -d
    ```

4. Ensure Air is installed for hot reloading:

    ```sh
    make ensure-air
    ```

## Running the Application

To run the application with hot reloading:

```sh
make hot
