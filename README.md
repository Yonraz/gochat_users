# User Service

This service is responsible for managing user information, including storing usernames and their statuses (online/offline). It supports CRUD operations triggered by events from the authentication service and provides a single endpoint to retrieve user information based on sorting methods.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Installation](#installation)
- [API Endpoints](#api-endpoints)
  - [Get Users](#get-users)
- [Event Consumers](#event-consumers)
  - [User Registered Handler](#user-registered-handler)
  - [User Logged-in Handler](#user-logged-in-handler)
  - [User Signed-out Handler](#user-signed-out-handler)
- [Models](#models)
  - [User](#user)
- [Services](#services)
  - [Create User](#create-user)
  - [Login User](#login-user)
  - [Signout User](#signout-user)
- [Consumers](#consumers)
  - [User Registered Handler](#user-registered-handler-1)
  - [User Logged-in Handler](#user-logged-in-handler-1)
  - [User Signed-out Handler](#user-signed-out-handler-1)
- [License](#license)

## Features

- Stores user information (username and status).
- Supports CRUD operations when events occur (user registration, login, logout).
- Consumes messages from RabbitMQ.
- Uses PostgreSQL as the database.
- Provides an endpoint to retrieve users with optional sorting parameters.

## Technologies

- Go
- Gin Framework
- GORM
- RabbitMQ
- PostgreSQL

## Installation

_*This project is part of a super-module (gochat-app), it's intended to run as a microservice on a k8s cluster. these instructions are for local installation._

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/gochat_users.git
   cd gochat_users
   ```

2. **Set up the database:**

   Ensure you have PostgreSQL installed and create a database for the service.

3. **Set up RabbitMQ:**

   Ensure you have RabbitMQ installed and running.

4. **Set up environment variables:**

   Create a `.env` file with the necessary configuration for your PostgreSQL and RabbitMQ instances.

   ```env
   DB_HOST=localhost
   DB_USER=youruser
   DB_PASSWORD=yourpassword
   DB_NAME=yourdb
   DB_PORT=5432

   RABBITMQ_URL=amqp://guest:guest@localhost:5672/
   ```

5. **Run the service:**

   ```bash
   go run main.go
   ```

## API Endpoints

### Get Users

- **URL:** `/api/users`
- **Method:** `GET`
- **Query Parameters:**
  - `sort` (optional): Field to sort by (`status`, `username`). Default is `status`.
  - `direction` (optional): Sort direction (`asc`, `desc`). Default is `desc`.
- **Response:**

  - `200 OK`: Returns a list of users.

    ```json
    [
      {
        "ID": 1,
        "Username": "user1",
        "Status": "online"
      },
      {
        "ID": 2,
        "Username": "user2",
        "Status": "offline"
      }
    ]
    ```

## Event Consumers

The service consumes events from RabbitMQ to handle user registration, login, and logout events.

### User Registered Handler

- **Queue:** `UserRegistrationQueue`
- **Key:** `UserRegisteredKey`
- **Exchange:** `UserEventsExchange`

This handler processes user registration events. When a new user registers, the Auth Service emits a `UserRegisteredKey` event to the `UserEventsExchange`. The User Service listens to this event on the `UserRegistrationQueue` and creates a new user in the database with an initial status of `offline`.

### User Logged-in Handler

- **Queue:** `UserLoginQueue`
- **Key:** `UserLoggedInKey`
- **Exchange:** `UserEventsExchange`

This handler processes user login events. When a user logs in, the Auth Service emits a `UserLoggedInKey` event to the `UserEventsExchange`. The User Service listens to this event on the `UserLoginQueue` and updates the user's status to `online` in the database.

### User Signed-out Handler

- **Queue:** `UserSignoutQueue`
- **Key:** `UserSignedoutKey`
- **Exchange:** `UserEventsExchange`

This handler processes user logout events. When a user logs out, the Auth Service emits a `UserSignedoutKey` event to the `UserEventsExchange`. The User Service listens to this event on the `UserSignoutQueue` and updates the user's status to `offline` in the database.

## Models

### User

```go
type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Status   constants.UserStatus
}
```
