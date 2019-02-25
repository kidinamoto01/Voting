# Voting

## Design Goal

The goal for this application is to build a `voting system` for people to reach offlice consensus.

## Application Architecture
```
./votingsystem
├── Gopkg.toml
├── Makefile
├── app.go
├── cmd
│   ├── votecli
│   │   └── main.go
│   └── voted
│       └── main.go
└── x
    └── voting
        ├── client
        │   ├── cli
        │   │   ├── query.go
        │   │   └── tx.go
        │   ├── rest
        │   │   └── rest.go
        │   └── module_client.go
        ├── codec.go
        ├── handler.go
        ├── keeper.go
        ├── msgs.go
        └── querier.go
```

## State

The state represents your application at a given moment. It tells how much token each account possesses, what are the owners and price of each name, and to what value each name resolves to.

The state of tokens and accounts is defined by the auth and bank modules, which means you don't have to concern yourself with it for now. What you need to do is define the part of the state that relates specifically to your nameservice module.

In the SDK, everything is stored in one store called the multistore. Any number of key/value stores (called KVStores in the Cosmos SDK) can be created in this multistore. For your application, you need to store:

A mapping of id to name. Create a nameStore in the multistore to hold this data.
A mapping of id to voters. Create a voterStore in the multistore to hold this data.
A mapping of id to results. Create a resultStore in the multistore to hold this data.

## Messages

Messages are contained in transactions. They trigger state transitions. Each module defines a list of messages and how to handle them. Here are the messages you need to implement the desired functionality for your nameservice application:
