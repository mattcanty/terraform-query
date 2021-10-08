# Terraform Query

Quickly query a Terraform provider's data type.

Such as a [GitHub repository](https://registry.terraform.io/providers/integrations/github/latest/docs/data-sources/repository):

```shell
➜ ~ tfq github_repository full_name hashicorp/terraform | jq .git_clone_url
"git://github.com/hashicorp/terraform.git"
```

## Prerequisites

1. Terraform
