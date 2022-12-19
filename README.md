# Polyamanita Server

Repository concerned with all things related to the Polyamanita server side

## Setup

To develop the server on a local machine, set up the following:

### Environment

Create a `.env` file with the following contents

```env
JWT_KEY="some-key"
```

### Swagger

To install swag-go:

`go install github.com/swaggo/swag/cmd/swag@latest`

To update swagger docs

`swag init -g main.go && swag fmt`
