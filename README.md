# go-mvc-api

Welcome to the Go MVC API Project! This repository presents a simple API built in Go, following the Model-View-Controller (MVC) architecture. The project encompasses key components such as controllers, services, repositories, models, and views, providing a clear separation of concerns.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The Go MVC API Project demonstrates a structured approach to building APIs in Go, adhering to the MVC design pattern. The goal is to create maintainable, modular code by segregating responsibilities into distinct layers: controllers for handling requests, services for business logic, repositories for data access, models for data representation, and views for rendering output.

## Features

- Implements a simple API following the MVC architecture.
- Enforces separation of concerns for better code organization.
- Utilizes MongoDB for data storage.
- Secures sensitive data using `.env` configuration.

## Technologies Used

- Go: Core programming language.
- MongoDB: Database for data storage.
- `.env` file: Secure configuration storage.

## Getting Started

1. Clone the repository:

   ```bash
     git clone https://github.com/suegoiee/go-mvc-api.git
     cd go-mvc-api
   ```
2. Install required dependencies:

  ```bash
    go mod download
  ```
3. Configure your .env file with required settings. See the [Configuration](#configuration) section for guidance.

4. Build and run the project:

  ```bash
    go run main.go
  ```
## Configuration
Copy the .env.example file to .env and replace values with your own:

  ```makefile
    USERNAME=your-mongodb-user-name
    PASSWORD=your-mongodb-password
    HOST=your-mongodb-uri
    PORT=your-mongodb-port
  ```
## Usage
This project is using main.go as its route file, before you execute, please make sure the port is not occupied.

## Contributing
Feel free to contribute by forking the repository and submitting pull requests. Your contributions are valued!

## License
This project is licensed under the MIT License.
