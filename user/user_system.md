# User System

## How it works
* On every request server checks if token has been given inside the cookie
* If token is given, server checks if token is valid, if it is, it will give add the user token to the request object as context
* Token is generated through Login/Register API call

## TODO
* [ ] Add a way to register
* [ ] Add tests for API calls

### Notes
    * Frontend has to send a request to renew the token every n minutes