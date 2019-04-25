## Usage
### Kutt Account
You need to have [Kutt][1] account before use this application.

### Setup Environment Variable
To using this application you need to set environment variable `KUTT_TOKEN` for [Kutt][1] API Token that you can found [here](https://kutt.it/settings). By default, `shorty` setup for 5 second timeout, if you willing to change less or more you can set environment variable `KUTT_TIMEOUT`.
In case you already set custom domain on your [Kutt][1] account, you need to set environment variable `KUTT_CUSTOM_DOMAIN` for your custom domain.
All possible environment variable that can be set showing in `.env.sample`, you can use that by copying with this command `cp .env.sample .env`, then shorty will automatically load up on current directory on file `.env`.

[1]: https://kutt.it

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