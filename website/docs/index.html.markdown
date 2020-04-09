---
layout: "ciao"
page_title: "Provider: CIAO"
sidebar_current: "docs-ciao-index"
description: |-
  The ciao provider provides utilities for working with a ciao server.
---

# ciao Provider

The cio provider provides utilities for working with a *ciao* server.
It provides a resource to manage checks.


## Example Usage

```hcl
  resource "ciao_check" "autonubil" {
    name ="autonubil"
    cron ="*/4 * * * *"
    url = "https://www.autonubil.de"
    active = true
  }

```
