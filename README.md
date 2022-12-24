# Polyamanita Server

Repository concerned with all things related to the Polyamanita server side

## Setup

To develop the server on a local machine, set up the following:

### AWS CLI

Follow the instructions provided here to [install](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html) and [configure](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html) an aws account with profile name "polyamanita"

### Environment

Create a `.env` file with the following contents

```env
AWS_PROFILE=polyamanita
JWT_KEY=some-key
ENVIRONMENT=dev
```

### Swagger

To install swag-go:

`go install github.com/swaggo/swag/cmd/swag@latest`

To update swagger docs

`swag init -g main.go && swag fmt`
