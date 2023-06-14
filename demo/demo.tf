terraform {
  required_providers {
    setnested = {
      source  = "chrismarget/setnested"
    }
  }
}

provider "setnested" {}

resource "setnested_problem" "x" {
  sna = [
    { required_string = "foo" },
    { required_string = "bar" },
  ]
}


