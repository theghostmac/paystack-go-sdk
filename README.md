# paystack-go-sdk
Unofficial Go SDK for Paystack.com

## Overview
The `paystack-go-sdk` provides a convenient and straightforward way to interact with the Paystack API using Go. This SDK covers various Paystack API endpoints, allowing you to manage transactions, customers, dedicated virtual accounts, and transfers programmatically.

## Features
- Manage transactions: initialize, verify, export, partial debit
- Manage customers: create, list, fetch, update, validate, whitelist/blacklist, deactivate authorization
- Manage dedicated virtual accounts: create, assign, list, fetch, requery, deactivate, split transactions, remove split
- Manage transfers: initiate, finalize, bulk transfer, list, fetch, verify

Check the [Changelog](CHANGELOG.md) for upcoming features and TODOs.

## Installation
To install the `paystack-go-sdk` package, use the following command:
```sh
go get github.com/theghostmc/paystack-go-sdk
```

## Usage

You can learn about the usage for every function currently supported by visiting [the examples folder](examples)

## Running Tests
You can run the tests using the provided `Makefile`:
```shell
make test
```

Note that there are functions that do not return any data from Paystack under the `tests` folder. This is
probably because the local tests and CI/CD test job uses a test secret key (`sk_test`).
