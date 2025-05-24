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

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```

## Environment Variables

The solution uses Environment variables to control the init state fo the app.
You can setup the PORT, INIT_PACKS and API_KEY.
On development environment just use .env file to setup those.

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
docker run -e API_KEY=secret -e INIT_PACKS=10,100,1000 -p 7000:7000 packr-api
```