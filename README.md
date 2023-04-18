# Gitti - Git HTTP Runner

Manage `Git` repositories with ease via Webhooks.

## Usage

    gitti [command] [flags]

### Commands

    config      Show the current config
    serve       Start the Gitti server
    help        Help about any command

### Flags

    -h, --help      Show help
    -V, --version   Show version

## Setup

`gitti config` will create a config file in the current directory. Under the section `[repos]` you can add your repositories. 
The key will be used as name for the webhook and the value is the local file path to the repository. The `port` flag is optional and defaults to `4000`.

    gitti config
    nano config.toml
    gitti serve --port 4000

### Example `config.toml`

    [repos]
    myrepo = "/path/to/myrepo"

## API

Description of the API endpoints. You can use `curl` to test the endpoints. The `:repo` must match the name of the repository as defined in the config file.

### Dashboard

Provides a simple overview of all repositories.

    GET /

```
curl http://localhost:4000
```

### Status

Does a `git status` on the repository.

    GET /status/:repo

```
curl http://localhost:4000/status/myrepo
```

### Log

Does a `git log` on the repository.

    GET /log/:repo

```
curl http://localhost:4000/log/myrepo
```

### Pull

Does a `git pull` on the repository.

    GET /pull/:repo

```    
curl http://localhost:4000/pull/myrepo
```    

### Checkout

Does a `git checkout` on the repository.

    GET /checkout/:repo
    GET /checkout/:repo?branch=:branch
    GET /checkout/:repo?branch=:branch&stash

```
curl http://localhost:4000/checkout/myrepo?branch=master&stash
```

### Prune

Deletes all temporary branches created by `gitti`.

    GET /prune/:repo

```    
curl http://localhost:4000/prune/myrepo
```

## Authentication

If `[auth]` is `enabled = true` in the config file, you need to provide the `token` in the `Authorization` header. 
The `token` is a random string that is generated on the first run of `gitti serve`. If you like, you can change it in the config file.

```
curl -H "Authorization: <token>" http://localhost:4000/checkout/myrepo?branch=test
```

## Systemd

You can use the provided `gitti.service` file to run `gitti` as a service. Change the `User` and `Group` to your needs. 
Keep in mind that the user needs write access to the repositories!

    sudo cp gitti.service /etc/systemd/system/gitti.service
    sudo systemctl daemon-reload
    sudo systemctl enable --now gitti
    sudo systemctl status gitti

## Logging

The log is written to `gitti.log`.

---
Copyright © 2023 Typomedia Foundation. All rights reserved.