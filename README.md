# terraform-provider-graylog

This repo is a fork of [terraform-provider-graylog](https://github.com/terraform-provider-graylog/terraform-provider-graylog)
with a couple of fixes to make it work with Graylog 4.1 and higher, by replacing the
`full_name` field with with `first_name` and `last_name`.

Linking to the upstream repos docs on registry.terraform.io below.

It can be used from Terraform with the following configuration.
```hcl
    graylog = {
      source  = "phonero.github.io/phonero/graylog"
      version = "1.1.0"
    }
```

As this repo exists primarily for our own usage, PR's might be ignored.

Terraform Provider for [Graylog](https://docs.graylog.org/)

- [Document](https://registry.terraform.io/providers/terraform-provider-graylog/graylog/latest/docs)
- [Terraform Registry](https://registry.terraform.io/providers/terraform-provider-graylog/graylog/latest/docs)
 
## License

[MIT](LICENSE)
