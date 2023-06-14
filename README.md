# terraform-provider-setnested

This is a skeleton of a terraform provider hastily thrown together to demonstrate a problem which cropped up in Terraform 1.5.0 but which did not appear in 1.4.6.

Running `terraform plan` in the `demo/` directory with the new version of Terraform produces the following error:

```text
Planning failed. Terraform encountered an error while generating this plan.

╷
│ Error: Provider produced invalid plan
│ 
│ Provider "registry.terraform.io/chrismarget/setnested" planned an invalid value for setnested_problem.x.sna: count in plan (cty.UnknownVal(cty.Number))
│ disagrees with count in config (cty.NumberIntVal(2)).
│ 
│ This is a bug in the provider, which should be reported in the provider's own issue tracker.
╵
```

The error seems to revolve around the fact that member objects of the `sna` attribute (a set) have a `Computed` attribute within.

I understand that these elements will be `Unknown` until a value is computed, and until such time as their value is completely known, members cannot be dereferenced from within the set.

It's less clear that this is the problem though. The error seems to suggest that the set members can't be *counted* while some values are unknown.

Also, it's not clear that plan generation should require dereferencing set members. Terraform 1.4.6 doesn't seem to have this problem. It generates a plan just fine:

```text
  # setnested_problem.x will be created
  + resource "setnested_problem" "x" {
      + sna = [
          + {
              + computed_int    = (known after apply)
              + required_string = "bar"
            },
          + {
              + computed_int    = (known after apply)
              + required_string = "foo"
            },
        ]
    }
```