User Service Project
Overview
This User Service is a Golang-based project designed to manage user data. It includes functionalities for creating users and storing them in a database, all implemented without using the repository pattern for simplicity.

Features
User creation with email and password.
Data storage in a PostgreSQL database.
Basic validation for email and password.
HTTP server for handling user creation requests.
Getting Started
Prerequisites
Golang (version 1.x or later)
PostgreSQL database
Installation
Clone the repository:
bash
Copy code
git clone https://github.com/your-username/user-service.git
Navigate to the project directory:
bash
Copy code
cd user-service
Configuration
Update the PostgreSQL database connection settings in pkg/user/model.go.
Running the Service
Execute the following command in the root directory of the project:

arduino
Copy code
go run ./cmd
The service will start and listen on localhost:8080.

Usage
Send a POST request to localhost:8080/create with a JSON body containing email and password to create a new user. Example using curl:

bash
Copy code
curl -X POST http://localhost:8080/create -d '{"Email":"test@example.com","Password":"password123"}'
Testing
Run the tests with the following command:

bash
Copy code
go test ./...
Contributing
Contributions to this project are welcome. Please fork the repository and submit a pull request.

