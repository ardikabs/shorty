# Shorty

## Description
Shorty is an application client to shortener your long url on shortener URL provider. Currently only support [kutt](https://kutt.it) provider.

## Diagram Flow
```
<user> - <shorty> - <shortenerProvider>
```
Basically it just client interact with shortener provider.

## Installation
1. Go installation. Follow this guide [here](https://golang.org/doc/install). Also other setup like GOPATH/GOROOT on your $PATH.
2. Run this command `go get -u github.com/ardikabs/shorty/app/cli`
3. Make sure you already follow step #1.

## Usage
### Setup Environment Variable
To using this application you need to set environment variable `KUTT_TOKEN`, that variable are API TOKEN from [Kutt](https://kutt.it), you can see on the settings page.
For other setup, like maybe you already set custom domain on your Kutt account, you need to set environment variable `KUTT_CUSTOM_DOMAIN`.

### `shorty --help`
This command will show you all the option and guide to follow this application.

### `shorty list`
This command will show you list of available urls on the Kutt provider.

### `shorty submit [url]`
This command will submit your selected URL to be shorten on Kutt provider.
__Available flags__: (you can see the explaination [here](https://github.com/thedevs-network/kutt#api))
1. `-c` or `--customurl`
1. `-p` or `--password`
1. `-r` or `--reuse`

### `shorty delete [url]`
This command will delete your selected short URL from Kutt provider