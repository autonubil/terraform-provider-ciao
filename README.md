# Terraform Privider for cioa

[ciao](https://brotandgames.com/ciao/) is an excellent HTTP(S) status checker

ciao checks HTTP(S) URL endpoints for a HTTP status code (or errors on the lower TCP stack) and sends a notification on status change via E-Mail or Webhooks.

It uses Cron syntax to schedule the checks and comes along with a Web UI and a RESTfull JSON API.

This provider wraps the ciao API into a simple terraform provider.

Maintainers
-----------

This provider plugin is maintained by the autonubil team at [autonubil](https://www.autonubil.de/).

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)


Usage
---------------------
```hcl

provider "ciao" {
    version = "~> 0.1"
    base_url = "http://127.0.0.1:8090"
    user     = "admin"
    password = "password"
    insecure = true
}

resource "ciao_check" "autonubil" {
    name ="autonubil"
    cron ="*/4 * * * *"
    url = "https://www.autonubil.de"
    active = true
}

```



Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/autonubil/terraform-provider-ciao`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:autonubil/terraform-provider-ciao
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/autonubil/terraform-provider-ciao
$ go build
```


## Configuration

The provider has three settings. The can be set in the provider or using the environment.

### Using HCL

```hcl
provider "ciao" {
    "base_url" = "http://127.0.0.1:8090"
    "user"     = "admin"
    "password" = "password"
    "insecure" = true
}

```

### Using the Environment

```bash
export CIAO_URL="http://127.0.0.1:8090"
export CIAO_USER="admin"
export CIAO_PASSWORD="password"
export CIAO_INSECURE=true
```


## Test setup

To run the acceptence test, you can start a local ciao server.

```bash
docker run --name ciao -p 8090:3000 -e BASIC_AUTH_USERNAME="admin" -e BASIC_AUTH_PASSWORD="password"   brotandgames/ciao
```