# Go RPC Demo

This project is an illustration of what a project leveraging gRPC looks like.

For insight into the structure of this folder, check out [this link](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html)

## Cmd

This folder contains the two build artifacts for this project, `greeter` and `greeterd`, a CLI tool and a daemon.

The CLI tool makes gRPC calls to the daemon.

## Greeter

Contains the application code implementing the greeter service, which is leveraged by `greeterd` and queried by `greeter`.
