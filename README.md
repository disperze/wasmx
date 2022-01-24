# WasmX
[![Go Report Card](https://goreportcard.com/badge/github.com/disperze/wasmx)](https://goreportcard.com/report/github.com/disperze/wasmx)

This project represents the [Juno](https://github.com/forbole/juno) implementation for
the [x/wasm module](https://github.com/cosmwasm/wasmd).

It extends the custom Juno behavior with custom message handlers for wasm message. This allows to store
the needed data inside a [PostgreSQL](https://www.postgresql.org/) database on top of
which [GraphQL](https://graphql.org/) APIs can then be created using [Hasura](https://hasura.io/)

## Usage
To know how to setup and run wasmx, please refer to the [docs folder](.docs).

## TODO

- [] Save code and store contract_version, checksum
- [] Save tokens info
- [] Handle IBC Msg to get contract calls.