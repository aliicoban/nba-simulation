# Go Nba Simulation

Go Nba Simulation

## Installation

Run:

+ Standard
   ```
   go run cmd/main.go
  ```
+ Docker Container
    ``` 
    docker build . -t nba-simulation
    ```
    
    ```
    docker run -p 4444:4444 nba-simulation
    ```

    # Endpoint
+ http://localhost:4444/start
- The request payload of POST endpoint will nba simulation starts

