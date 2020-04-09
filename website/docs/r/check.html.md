---
layout: "check"
page_title: "check: ciao_check"
sidebar_current: "docs-ciao-resource-check"
description: |-
  Creates a check definition.
---

# ciao_check


## Example Usage

```hcl
  resource "ciao_check" "autonubil" {
    name ="autonubil"
    cron ="*/4 * * * *"
    url = "https://www.autonubil.de"
    active = true
  }

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the check

* `url` - (Required) Target of the check

* `cron` - When to run the check (defaults to every 5 minutes)

* `active` - Active state of the check (defaults to true)


## Attributes Reference

No attributes are exported