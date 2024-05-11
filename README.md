![Header](https://github.com/GwartneyDev/goalboard-collab/blob/main/Capture.png)

> Goalboard is under current development, and we are working hard to bring our vison to reality!

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

Before you get started, you'll want to migrate your database.

Go into the backend container:

```bash
docker exec -it goalboard-backend-1 bash
```

Run the migration scripts:

```bash
/bin/bash migrate.sh up
```
