# Goalboard

An application dedicated towards helping users create, track and an emphasis on Goal sharing.

## Design specs

https://www.figma.com/file/WlrtJXxCnsjjGGoJaYPqCk/Goalboard?type=design&node-id=0-1&mode=design&t=aQd7KCJ9X6oOg2HJ-0

MVP

https://www.figma.com/file/WlrtJXxCnsjjGGoJaYPqCk/Goalboard?type=design&node-id=122-1570&mode=design&t=aQd7KCJ9X6oOg2HJ-0

## Tech stack

- GoLang
- React (Vite) with TypeScript

## Development

### Steps for setup

- Install [Go](https://go.dev/doc/install)
- Install Nodejs (I like to use [NVM](https://github.com/nvm-sh/nvm))
- Install Go binaries `cd backend && go install`
- Install NodeJS packages (after going back to root) `cd frontend && npm i`

### Migrations

Before you get started, you'll want to migrate the database. Right now we're doing this manually. The command looks something like this:

```bash
psql -U postgres -d postgres -a -f ~/<path-to-project>/backend/database/migrations/<my-migration-file>.up.sql
```

Make sure to replace and use the correct settings for your current project setup!

### Steps for running

- While in /frontend `npm run dev`
- While in /backend `go run main.go`
