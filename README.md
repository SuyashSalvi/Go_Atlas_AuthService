# Go Atlas - JWT Auth Service

This is a simple authentication service written in Go, integrated with MongoDB Atlas for data storage and JWT (JSON Web Tokens) for authentication.

## Features

- User signup: Register new users with a username and password.
- User login: Authenticate users with their username and password and issue JWT tokens.
- Access token refresh: Provide an endpoint for refreshing access tokens.
- Session control: Manage user sessions using JWT tokens.

## Prerequisites

- Go 1.16 or later installed on your system.
- MongoDB Atlas account with a cluster set up.

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/SuyashSalvi/Go_Atlas_AuthService.git
   ```

2. Navigate to the project directory:

   ```bash
   cd Go_Atlas_AuthService
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## Configuration

1. Set up MongoDB Atlas:
   - Create a cluster on MongoDB Atlas.
   - Get the connection URI for your cluster.

2. Configure environment variables:
   - Create a `.env` file in the root directory.
   - Add the following variables to the `.env` file:

     ```dotenv
     ATLAS_URI=mongodb+srv://<username>:<password>@<cluster-name>/<database-name>
     JWT_SECRET=your-secret-key
     ```

   Replace `<username>`, `<password>`, `<cluster-name>`, and `<database-name>` with your MongoDB Atlas credentials.

## Usage

1. Run the application:

   ```bash
   go run main.go
   ```

2. The service will start running on `http://localhost:8080`.

## Endpoints

- `POST /signup`: Register a new user.
- `POST /login`: Authenticate user and generate JWT token.
- `POST /token`: Refresh access token.

## Dependencies

- [gorilla/mux](https://github.com/gorilla/mux): HTTP router for building RESTful APIs in Go.
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver): Official MongoDB driver for Go.
- [jwt-go](https://github.com/dgrijalva/jwt-go): JWT implementation for Go.

## Contributing

Contributions are welcome! Feel free to open issues and pull requests.
