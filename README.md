# Timeout Test Server

A simple Go server designed to test client-side timeout handling by deliberately not responding to requests.

## Overview

This server provides an endpoint that never sends a response back to the client, which forces the client to handle a timeout scenario. This is useful for:

- Testing how frontend applications handle API timeouts
- Verifying that client-side timeout mechanisms work correctly
- Simulating slow or unresponsive backend services

## Installation

### Prerequisites

- Go 1.23.4 or higher installed

### Setup & Usage

1. Clone this repository:

   ```
   git clone https://github.com/V4N1LLA-1CE/timeout-endpoint.git
   cd timeout-endpoint
   ```

2. Start the server:
   ```
   go run main.go
   ```

You should see a message: `Timeout test server running at http://localhost:8080`

2. Test the timeout endpoint:
   ```
   curl http://localhost:8080/api/timeout
   ```
   This request will hang indefinitely, as the server never responds.

## Configuration

The server runs on port 8080 by default. To change this, modify the `port` variable in `main.go`.

## License

[MIT License](LICENSE)
