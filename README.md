# Goalboard

An application dedicated towards helping users create, track and an emphasis on Goal sharing.

## Design specs

https://www.figma.com/file/WlrtJXxCnsjjGGoJaYPqCk/Goalboard?type=design&node-id=0-1&mode=design&t=aQd7KCJ9X6oOg2HJ-0

MVP

https://www.figma.com/file/WlrtJXxCnsjjGGoJaYPqCk/Goalboard?type=design&node-id=122-1570&mode=design&t=aQd7KCJ9X6oOg2HJ-0

## Tech stack

- GoLang
- React (Vite) with TypeScript
- Docker

## Development

### Setup

Copy over your .env.example files and rename them as .env.

Ex:

```bash
cp ./backend/.env.example ./backend/.env
```

Make sure to update the variables in there if needed!

Install docker and docker compose if you haven't already:

https://www.docker.com/get-started/

After that we can create the containers:

```bash
docker compose build
```

And then run them in the backround:

```bash
docker compose up -d
```

### Migrations

Before you get started, you'll want to migrate the database. Right now we're doing this manually. The process will look something like this:

Go into the backend container:

```bash
docker exec -it goalboard-backend-1 sh
```

Run each migration. The command per file looks something like this:

```bash
psql postgres://postgres:<password>@db/postgres -a -f /code/database/migrations/<my-migration-file>.up.sql
```

We'll need to figure out a out to manage our migrations to make this a bit easier.

PS: Make sure to replace and use the correct settings for your current project setup!
