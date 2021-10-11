# WasmX

This project represents the [Juno](https://github.com/desmos-labs/juno) implementation for
the [x/wasm module](https://github.com/cosmwasm/wasmd).

It extends the custom Juno behavior with custom message handlers for all the wasm messages. This allows to store
the needed data inside a [PostgreSQL](https://www.postgresql.org/) database on top of
which [GraphQL](https://graphql.org/) APIs can then be created using [Hasura](https://hasura.io/)
