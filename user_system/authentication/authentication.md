# Authentication

## Description
This is the authentication package of the system. It is responsible for authenticating users and providing them with a token that they can use to access the system.

Based losely on JWT tokens. Inspiration taken from: https://medium.com/swlh/building-a-user-auth-system-with-jwt-using-golang-30892659cc0

## TOKEN
The token is encrypted using private key. On every request the token is regenerated.
The token contains the following information:
* expiration time (unix nano timestamp)
* Username for identification
* Signature for verification

