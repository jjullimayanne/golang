#!/bin/sh

# Define the URL to check
KEYCLOAK_URL="http://keycloak:8080/realms/recycling"

# Wait for the Keycloak service to be available
until $(curl --output /dev/null --silent --head --fail "$KEYCLOAK_URL/.well-known/openid-configuration"); do
    printf '.'
    sleep 5
done

echo "\nKeycloak is up and running!"
