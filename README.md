# golang-auth0-example
An example backend with GoLang that uses auth0 for authentication

## server setup

0. Ensure you have air and task (https://youtu.be/Q7uh85_i0-M)
1. Create an auth0 account (https://auth0.com/)
2. Create a new api (https://auth0.com/docs/quickstart/backend/golang/01-authorization#create-a-middleware-to-validate-access-tokens)
3. Setup ```app.env``` file with your credentials, follow ```app.env.example```
4. Run ```task dev```

## client setup

1. Run ```cd web```, ```npm install```
2. Setup a new client application in auth0 (https://auth0.com/docs/quickstart/spa/react/01-login#create-a-new-application)
3. Fill in the ```web/.env``` file with your credentials, follow ```web/.env.example```
4. Run ```npm run dev```