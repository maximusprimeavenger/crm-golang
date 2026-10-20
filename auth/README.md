# User authentication service
User authentication and authorization microservice for the e-commerce platform. Handles user registration, login, JWT-based authentication, and role management.

## Requirements:
- Golang 1.23+
- Docker & Docker Compose
- MongoDB (preferably)

## Setup Instructions  

Clone the repository and set up the environment:  

```sh
git clone https://github.com/NIC-Golang/user-auth-service.git
cd user-auth-service
cp .env.example .env
```

Also, you need to configure your admin account in conf.yaml(user-auth-service/config/conf.yaml). Add the parameters to set up your admin account.
Here's the example:
```
password: admin
phone: 111111111
email: admin@mail.ru
name: admin
```
## Available Makefile Commands:
```sh
make help       # Show available commands  
make install    # Install dependencies (go mod tidy)  
make run        # Run the server  
make stop       # Stop the server  
make restart    # Restart the server  
make compile    # Compile the application  
make clean      # Clean the build cache  
make test       # Run the test container (docker compose -f docker-compose.test.yml up -d)  
make build      # Build the docker container  
make up         # Run the docker container  
make down       # Stop the docker container
```

## Running the service:
Make sure Docker Engine is running, then execute:
```
docker compose build
docker compose up
```
To stop the service:
```
docker compose down
```

If you are using an older Docker version, use docker-compose instead of docker compose.