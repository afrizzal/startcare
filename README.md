# Startcare - Fundraiser App

Startcare is a Go application designed for fundraising purposes. It provides a platform for users to create and manage campaigns, as well as support others' initiatives. This README provides instructions on setting up and running the application.

## Prerequisites

Make sure you have Go installed on your machine. You can download and install it from [here](https://golang.org/dl/).

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/afrizzal/startcare.git
    ```

2. Change into the project directory:

    ```bash
    cd startcare
    ```

3. Set up the database:

    - Ensure that MySQL is running on your machine.
    - Modify the `dsn` variable in `main.go` with your MySQL database connection details.

4. Install dependencies:

    ```bash
    go mod tidy
    ```

5. Run the application:

    ```bash
    go run main.go
    ```

   This will start the Startcare application, accessible at [http://localhost:8080](http://localhost:8080).

## API Endpoints

- **Register User:**
    ```http
    POST /api/v1/users
    ```

- **Login:**
    ```http
    POST /api/v1/sessions
    ```

- **Check Email Availability:**
    ```http
    POST /api/v1/email_checkers
    ```

- **Upload Avatar (Authenticated):**
    ```http
    POST /api/v1/avatars
    ```

- **Get Campaigns:**
    ```http
    GET /api/v1/campaigns
    ```

- **Get Campaign by ID:**
    ```http
    GET /api/v1/campaigns/:id
    ```

- **Create Campaign (Authenticated):**
    ```http
    POST /api/v1/campaigns
    ```

## Authentication

The application uses JWT (JSON Web Token) for authentication. To make authenticated requests, include the generated JWT token in the `Authorization` header with the "Bearer" prefix.

Example:
   ```http
   Authorization: Bearer <your-generated-token>