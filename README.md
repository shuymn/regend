# regen

Generate AWS Redshift table DDL.

# Installation

```bash
go get github.com/shuymn/regen/cmd/regen
```

# Configuration

Following configuration file or environment variables are required to execute.

## Configuration File

By default check the following files.

1. `$XDG_CONFIG_HOME/regen/regen.toml`
2. `$HOME/.regen.toml`

The contents of the file are as follows.

```toml
[regen]
host = "foo.bar.com"
port = 1234
user = "foobar"
password = "password1234"
database = "nyan"
```

## Envitonment variable

```
REDSHIFT_HOST
REDSHIFT_PORT
REDSHIFT_USER
REDSHIFT_PASSWORD
REDSHIFT_DATABASE
```


# License

- static/generate_tbl_ddl.sql - Licensed under the [Amazon Software License](http://aws.amazon.com/asl/)
  - ref: https://github.com/awslabs/amazon-redshift-utils
- the others - MIT (c) shuymn
