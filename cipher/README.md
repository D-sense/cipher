# Cipher Service

This service exposes a RESTFUL, with two endpoints; one that enciphers a message, and one that deciphers it. 
Each endpoint accept two parameters, a string s to be enciphered or deciphered and an integer k that specifies the alphabet rotation factor. 
The service returns the enciphered or deciphered text, provided the input was valid, or an appropriate error. It also provides a simple CLI client that can be used to call the service to demonstrate its functionality.

## Running locally

### Without Container
1. Download the dependencies by running `make tidy`.
2. :
   - (a).  If you want to test the service via RESFull API, run `make api-run` from the root director, namely `cipher`.
   If everything goes well (logs are displayed in the terminal), you may start calling the endpoints.
    Example (1): sending a request for encrypting a string:
    `curl -X POST "http://localhost:8080/api/v1/encrypt" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"text\":\"Let's carve him as a dish fit for the gods.\",\"key\":14}"`
    Example (2): sending a request for decrypting a string:
    `curl -X POST "http://localhost:8080/api/v1/decrypt" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"text\":\"Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg.\",\"key\":14}"`
   - (b). If you want to test the service via CLI, run `make cli-run` from the root director, namely `cipher`.
    If everything goes well (logs are displayed in the terminal), you may start querying the service.
    Example (1): sending a request for encrypting a string:
    `./cmd/bin/cipher --action="encrypt" --text="Let's carve him as a dish fit for the gods." --key=14`
    Example (2): sending a request for decrypting a string:
    `./cmd/bin/cipher --action="decrypt" --text="Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg." --key=14`


### With Container
You can test the RESTFull API in a container. Ensure docker desktop application is on, then navigate to the root director, namely `cipher` and run:
`make docker-build`
- Once, the command is done running, run the command below to start the API service:
`make docker-build`
- Start querying the endpoints:
Example (1): sending a request for encrypting a string:
`curl -X POST "http://localhost:8080/api/v1/encrypt" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"text\":\"Let's carve him as a dish fit for the gods.\",\"key\":14}"`
Example (2): sending a request for decrypting a string:
`curl -X POST "http://localhost:8080/api/v1/decrypt" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"text\":\"Zsh'g qofjs vwa og o rwgv twh tcf hvs ucrg.\",\"key\":14}"`

## Running test coverage 
Navigate to the root director, namely `cipher` and run:
`make test`
