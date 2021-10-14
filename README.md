# WasmX

This project represents the [Juno](https://github.com/desmos-labs/juno) implementation for
the [x/wasm module](https://github.com/cosmwasm/wasmd).

It extends the custom Juno behavior with custom message handlers for wasm message. This allows to store
the needed data inside a [PostgreSQL](https://www.postgresql.org/) database on top of
which [GraphQL](https://graphql.org/) APIs can then be created using [Hasura](https://hasura.io/)

> Currently indexes the instantiated contracts

## Usage
To know how to setup and run wasmx, please refer to the [docs folder](.docs).