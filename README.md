# go-htmx-template

This is a skeleton app using Go/Echo/Templ/sqlx in the backend and HTMX/AlpineJS/Tailwind/Vite in the frontend.

## Getting Started

* Copy the project to your local machine.
* Run `go mod tidy` to download the dependencies.
* In the client dir run `pnpm install` to download the client dependencies.
* Run `pnpm run build` to build the client bundle.
* Run `go run cmd/server/main.go` to start the server.
* The [gomon](https://github.com/jdudmesh/gomon) client is included so you can run `gomon` to start the server and have it automatically restart when you make changes.

## Features

* There is a separate `client` folder containing all of the client dependencies. The client bundle is built using Vite and copied to the `static/dist` folder along with any other dependencies (e.g. fonts). The contents of this folder are server statically by the Go server using the `/assets` path. To build the client run `pnpm run build` in the client folder. To watch for changes run `pnpm run build:watch`.
* Config object loaded from env vars.
* A simple Sqlite3/sqlx datastore (nothing is implemented beyond opening/closing the connection)
* The home page is rendered at startup and served from memory thereafter. There is an HTMX tag in the navbar which gets the auth buttons (logout or login/signup) from a separate endpoint.

## License

MIT. Have at it.