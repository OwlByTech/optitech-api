<h1 align="center">
    optitech API
</h1>

## Getting Started

To start the project in development mode, you must copy the `.env.example` to `.env`:
and then you should run the following command:

```bash
docker compose -f compose-dev.yml up
```

Then you must run the migration and seeder inside the app container.

```bash
docker compose -f compose-dev.yml exec app  bash
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
