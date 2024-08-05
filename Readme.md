# Proof of Work (PoW) TCP Server [![Build Status](https://github.com/komandakycto/pow-example/workflows/build/badge.svg)](https://github.com/komandakycto/pow-example/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/komandakycto/pow-example)](https://goreportcard.com/report/github.com/komandakycto/pow-example) [![Coverage Status](https://coveralls.io/repos/github/komandakycto/pow-example/badge.svg?branch=master)](https://coveralls.io/github/komandakycto/pow-example?branch=master)

## Overview

This project implements a "Word of Wisdom" TCP server protected from DDOS attacks using [Proof of Work (PoW)](https://en.wikipedia.org/wiki/Proof_of_work#:~:text=Proof%20of%20work%20(PoW)%20is,computational%20effort%20has%20been%20expended). The server is designed to challenge clients with a
PoW problem that must be solved before a quote from a predefined list can be retrieved. This mechanism ensures that clients perform a certain amount of
computational work before receiving the desired information, thereby mitigating the risk of DDOS attacks.

## Project Implementation

### Hashcash Algorithm

For this project, we chose the Hashcash algorithm for the PoW implementation. [Hashcash](https://en.wikipedia.org/wiki/Hashcash#:~:text=Hashcash%20is%20a%20cryptographic%20hash,proof%20can%20be%20verified%20efficiently) is a well-known and widely used PoW algorithm that requires finding a
nonce such that when hashed, the resulting hash has a certain number of leading zeros. This method is simple to implement and provides a sufficient level of
security for our use case.

### Why Hashcash?

* **Simplicity**: Hash-based PoW algorithms are straightforward to implement and understand. They rely on well-established cryptographic hash functions like SHA-256, which are widely supported in many programming languages, including Go.

* **Security**: Cryptographic hash functions are designed to be computationally difficult to reverse or predict, making it challenging for an attacker to bypass the PoW without performing the necessary work.

* **Configurability**: The difficulty of a hash-based PoW can be easily adjusted by changing the number of leading zeroes required in the hash. This allows for fine-tuning the level of effort required from the client.

* **Resource** Usage: Hash-based PoW primarily consumes CPU resources, which are readily available on most devices. This makes it accessible and practical for a wide range of applications.

* **Proven** Effectiveness: Hashcash and similar hash-based PoW algorithms have been proven effective in real-world applications, such as Bitcoin mining and email spam prevention.

### Running the Project

#### Prerequisites

- Docker Engine
- Docker Compose
- Go (version 1.18 or later) for local development

```bash
make run
```