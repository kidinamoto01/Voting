# Voting

## Application Goals

This application is designed to let users create ballot and 
others could vote on them.

Here are the modules you will need for the ballot application:

* `auth`: This module defines accounts and fees and gives access to these functionalities to the rest of your application.
* `bank`: This module enables the application to create and manage tokens and token balances.
* `ballot`: This module does not exist yet! It will handle the core logic for the ballot application.

## State

The state represents your application at a given moment. It tells how much token each account possesses, 
what are the owners and price of each name, and to what value each name resolves to.

The state of tokens and accounts is defined by the `auth` and `bank` modules, 
which means you don't have to concern yourself with it for now. 
What you need to do is define the part of the state that relates specifically to your `ballot` module.

In the SDK, everything is stored in one store called the multistore. Any number of key/value stores (called KVStores in the Cosmos SDK) 
can be created in this multistore. For your application, you need to store:

A mapping of name to ballot contents. Create a `ballotStore` in the multistore to hold this data.
A mapping of name to owner. Create a `ownerStore` in the multistore to hold this data.
A mapping of address to ballot. Create a `historyStore` in the multistore to hold this data.


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


## Messages

Messages are contained in transactions. They trigger state transitions. Each module defines a list of messages and how to handle them. Here are the messages you need to implement the desired functionality for your nameservice application:
