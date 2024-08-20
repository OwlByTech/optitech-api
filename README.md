<h1 align="center">
    optitech API
</h1>

## Getting Started

To start the project in development mode, you must copy the `.env.example` to `.env`:
and then you should run the following command:

```bash
docker compose -f compose-dev.yml up
```

Run DocuStream

```bash
docker run -it --rm -p 5000:3000 docu-stream:latest
```

Then you must run the migration and seeder inside the app container.

```bash
docker compose -f compose-dev.yml exec app bash
```

And you must first run the migrations using the cli

```bash
 go run cmd/cli/main.go migrate up
```

the seeder

```bash
 go run cmd/cli/main.go seed up
```

and then install mjml in the container

```bash
apt-get update
```

now install npm and mjml

```bash
apt-get install -y nodejs npm
```

```bash
npm install -g mjml
```

run this comant to execute convert-mjml

```bash
go run cmd/cli/main.go convert-mjml
```

## Production

To run the project in production mode, you should add the environment variable SETUP=first-time to automatically run migrations, seeders, and conversions. Do not use this variable after the first time.

For the first run:

```bash
SETUP=first-time docker compose up
```

For subsequent runs:

```bash
docker compose up
```
