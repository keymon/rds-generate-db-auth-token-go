# rds-generate-db-auth-token-go

Generates the AWS RDS IAM token with a golang binary.

In a nutshell, a golang binary to run:
```
aws rds generate-db-auth-token \
	--hostname "${PGHOST}" \
	--port "${PGPORT}" \
	--region "${REGION}" \
	--username "${PGUSER}"
```

This is a very light binary, and handy when you do not want to
download all the AWS cli in python to just generate these creds in
some script or tool.

Based on: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.Connecting.Go.html

# Install

Build it:

```
go install github.com/keymon/rds-generate-db-auth-token-go/cmd/rds-auth-token@latest
```

Or download it from the releases section.

# Usage

```
# Providing AWS creds are loaded

export PGHOST=mysqlcluster.cluster-123456789012.us-east-1.rds.amazonaws.com
export PGPORT=3306
export PGUSER=myuser

export PGPASSWORD="$(./rds-auth-token --host "${PGHOST}" --port "${PGPORT}" --user "${PGUSER}")"

# Additional params
export PGSSLMODE=require
export PGDATABASE=mydb

# Test it with psql
psql
```


