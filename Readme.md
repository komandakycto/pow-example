# Proof of Work (PoW) TCP Server [![Build Status](https://github.com/komandakycto/pow-example/workflows/build/badge.svg)](https://github.com/komandakycto/pow-example/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/komandakycto/pow-example)](https://goreportcard.com/report/github.com/komandakycto/pow-example) [![Coverage Status](https://coveralls.io/repos/github/komandakycto/pow-example/badge.svg?branch=master)](https://coveralls.io/github/komandakycto/pow-example?branch=master)

## Overview

This project implements a "Word of Wisdom" TCP server protected from DDOS attacks using Proof of Work (PoW). The server is designed to challenge clients with a
PoW problem that must be solved before a quote from a predefined list can be retrieved. This mechanism ensures that clients perform a certain amount of
computational work before receiving the desired information, thereby mitigating the risk of DDOS attacks.

## Project Implementation

### Hashcash Algorithm

For this project, we chose the Hashcash algorithm for the PoW implementation. Hashcash is a well-known and widely used PoW algorithm that requires finding a
nonce such that when hashed, the resulting hash has a certain number of leading zeros. This method is simple to implement and provides a sufficient level of
security for our use case.

### Why Hashcash?

* Simplicity: The Hashcash algorithm is straightforward to implement using standard cryptographic hash functions.
* Security: The algorithm is designed to be computationally expensive to solve but easy to verify, making it effective at mitigating DDOS attacks.
* Proven: Hashcash is a proven algorithm, used in various applications, including email spam prevention and cryptocurrency mining.

### Running the Project

#### Prerequisites

- Docker Engine (version 20.10 or later)
- Docker Compose (version 1.29 or later)
- Go (version 1.18 or later) for local development

```bash
make run
```