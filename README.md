# Shorty

## Description
Shorty is an application client to shortener your long url on shortener URL provider. Currently only support [Kutt](https://kutt.it) provider.

## Diagram Flow
```
<user> - <shorty> - <shortenerProvider>
```
Basically it just client interact with shortener provider.

## Development Guide
### Prerequisites
* Go 1.11 or later
### Setup
* Install Git. [Here](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
* Install Go 1.11 or later. [Here](https://golang.org/doc/install)
* Clone this repo on your `$GOPATH`
    Run this command `git clone git:github.com/ardikabs/shorty.git`

### Build and run binary file
To build binary file:
`make build`

### Unit Test
To check the unit test:
`make test`

## Installation Guide
Check the [release page](https://github.com/ardikabs/shorty/releases).
Please read [user guide](USAGE.md) for further use.