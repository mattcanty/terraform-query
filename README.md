# Terraform Query

Quickly query a Terraform provider's data type.

Such as:

```shell
➜ ~ tfq github_repository full_name hashicorp/terraform | jq .git_clone_url
"git://github.com/hashicorp/terraform.git"
```
