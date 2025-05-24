# Project packr

A project for solving example Unbound Knapsack Problem

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Run the test suite:
```bash
make test
```

## Running FrontEnd

The Front end is a simple Svelte web app.
To run it you will need node installed.

```bash
cd client
npm install
npm run dev
```

To build the app:
```bash
npx vite build
```
The build is done in client/dist folder as a target.

## Environment Variables

The solution uses Environment variables to control the init state fo the app.
You can setup the PORT, INIT_PACKS and API_KEY for Backend.
You can setup VITE_API_KEY and VITE_API_BASE_URL for Frontend

On development environment just use .env file to setup those, theres .env.template file in repository with example values, just rename it to .env
Frontend and backend has separate .env files.


INIT_PACKS are just packs delimited with comma

```
INIT_PACKS=10,100,1000
```

## Docker

To run the app as a docker container simply build it:
```bash
docker build -t packr-api .
```

And run it:
```bash
docker run -p 7000:7000 packr-api
```

You can supply environment variables:
```bash
docker run -e API_KEY=secret -e INIT_PACKS=10,100,1000 -e PORT=7000 -p 7000:7000 packr-api
```

## API

Simple REST API to add and remove packs, view them and get solution.
The state is held on server and is only one for whole app, so multiple users are not allowed.
The initial packs state can be defined by INIT_PACKS enviroment variable.

The API is restricted with API_KEY if header X-Api-Key is not provided, 401 is expected.

### GET /api/packs

Returns a list of current packs

```
curl http://localhost:7000/api/packs
```

example response:
```
{"packs":[5000,2000,1000,500,250]}
```

### POST /api/packs

Adds a new pack to the state
Required query parameter pack

```
curl -X POST http://localhost:7000/api/packs?pack=500
```

For valid request is HTTP status 201

For NAN, 0 and negative numbers status 400 with reason in body

### DELETE /api/packs

Removed a pack from the state if it existed
Required query parameter pack

```
curl -X DELETE http://localhost:7000/api/packs?pack=500
```

For valid request is HTTP status 200

For NAN, 0 and negative numbers status 400 with reason in body

### GET /api/solution

Runs the solution algorithm and returns the result
Required query parameter amount

```
curl http://localhost:7000/api/solution?amount=12001
```

example response
```
{"solution":{"5000":2,"2000":1,"250":1}}
```

For valid request is HTTP status 200

For NAN, 0 and negative numbers status 400 with reason in body