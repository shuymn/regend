# regend

Generate AWS Redshift table DDL.

# Installation

```bash
go get github.com/shuymn/regend/cmd/regend
```

# Usage

```bash
regend table
```

# Configuration

Following configuration file or environment variables are required to execute.

## Configuration File

By default check the following files.

1. `$XDG_CONFIG_HOME/regend/regend.toml`
2. `$HOME/.regend.toml`

The contents of the file are as follows.

```toml
[redshift]
host = "foo.bar.com"
port = 1234
user = "foobar"
password = "password1234"
database = "nyan"
```

## Envitonment variable

```
REGEND_REDSHIFT_HOST
REGEND_REDSHIFT_PORT
REGEND_REDSHIFT_USER
REGEND_REDSHIFT_PASSWORD
REGEND_REDSHIFT_DATABASE
```


# License

- static/generate_tbl_ddl.sql - Licensed under the [Amazon Software License](http://aws.amazon.com/asl/)
  - ref: https://github.com/awslabs/amazon-redshift-utils
- the others - MIT (c) shuymn
