# regend

Generate AWS Redshift CREATE TABLE DDL for an existing table.

## Requirements

- Go 1.21+ (module declares `go 1.21`).

## Installation

Using Go tooling:

```bash
go install github.com/shuymn/regend/cmd/regend@latest
```

From source:

```bash
make build
./bin/regend <table_name>
```

## Usage

```bash
regend <table_name>
```

The command connects to Redshift using your configured credentials, runs an embedded query against the system catalogs, and prints the assembled DDL to stdout.

Example (with environment variables):

```bash
REGEND_REDSHIFT_HOST=example.abc123.us-west-2.redshift.amazonaws.com \
REGEND_REDSHIFT_PORT=5439 \
REGEND_REDSHIFT_USER=myuser \
REGEND_REDSHIFT_PASSWORD=mypassword \
REGEND_REDSHIFT_DATABASE=dev \
regend my_table
```

## Configuration

Provide Redshift connection settings via a TOML file and/or environment variables. Environment variables override values loaded from the file.

### Configuration file

By default, the tool looks for the following files:

1. `$XDG_CONFIG_HOME/regend/regend.toml`
2. `$HOME/.regend.toml`

Example contents:

```toml
[redshift]
host = "foo.bar.com"
port = 5439
user = "foobar"
password = "password1234"
database = "nyan"
```

### Environment variables

```text
REGEND_REDSHIFT_HOST
REGEND_REDSHIFT_PORT
REGEND_REDSHIFT_USER
REGEND_REDSHIFT_PASSWORD
REGEND_REDSHIFT_DATABASE
```

An example `.envrc.example` is provided if you use `direnv`.

## Output

The output is a Redshift DDL statement reconstructed from system catalogs (columns, encodings, defaults, constraints excluding foreign keys, distribution style/key, and sort keys), for example:

```sql
CREATE TABLE IF NOT EXISTS "public"."my_table"
(
  -- columns...
)
DISTSTYLE EVEN
SORTKEY (
  -- sort keys...
);
```

## License

- `cli/generate_tbl_ddl.sql` — Licensed under the Amazon Software License (see awslabs/amazon-redshift-utils)
- All other code — MIT (c) shuymn
