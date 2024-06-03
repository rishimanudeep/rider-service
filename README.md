# Rider-Service

This is a Go application for managing rider service.
The Rider Service manages rider-related operations, providing functionalities such as rider registration, availability tracking, and order assignment. It ensures efficient handling of rider data and real-time updates on rider availability, facilitating seamless integration with the order processing system for optimal delivery management.
## Features

The Rider Service offers comprehensive features tailored for efficient rider management and delivery operations. Key functionalities include:

-Rider Registration: Seamlessly onboard new riders onto the platform, capturing essential details such as name, contact information, and vehicle details.

-Availability Tracking: Real-time monitoring of rider availability, enabling the system to efficiently assign orders based on rider proximity and availability status.

-Order Assignment: Facilitates the allocation of orders to available riders.

-Location Tracking: Continuous tracking of rider locations, providing accurate updates on rider movements and enabling customers to track their orders in real-time.
## Getting Started with Rider-Service

### Requirements

- A working Go environment - [https://go.dev/dl/](https://go.dev/dl/)
- Check the go version with command: go version.
- One should also be familiar with the Golang syntax. [Golang Tour](https://tour.golang.org/) has an excellent guided tour and highly recommended.

### Installation

## GOFR as dependency used for migrations

- To get the GOFR as a dependency, use the command:
  `go get gofr.dev`

- Then use the command `go mod tidy`, to download the necessary packages.


### To Run Server

Run `go run main.go` command in CLI.

## Usage

The application provides the following RESTful endpoints:

- `POST /rider`: Register a new rider.
- `PUT /rider/{riderid}/location`: Update a riders location
- `GET /riders/nearby"`: Retrieve all nearby riders to the location.
- `PUT /rider/{riderid}/availability`: updates riders availability status.
- `PUT /rider/{id}`: Update Rider ID.
- `GET /rider/{id}`: Get Rider by ID.

## Dependencies

The application uses the following dependencies:

- `gofr.dev/pkg/gofr`: A Go web framework used for handling HTTP requests.
- `Rider-service/handlers`: Handlers package for handling HTTP requests related to riders.
- `Rider-service/services`: Services package for business logic related to riders.
- `Rider-service/stores`:Store package for handling db operations related to riders.

For any information please reach out to me via rishimanudeepg@gmail.com