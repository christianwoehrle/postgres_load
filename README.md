# postgres_load
[![GoDoc](https://godoc.org/github.com/christianwoehrle/postgres_load?status.svg)](https://godoc.org/github.com/christianwoehrle/postgres_load)
[![CircleCI](https://img.shields.io/circleci/project/github/christianwoehrle/postgres_load.png)](https://circleci.com/gh/christianwoehrle/postgres_load)
[![Go Report Card](https://goreportcard.com/badge/github.com/christianwoehrle/postgres_load)](https://goreportcard.com/report/github.com/christianwoehrle/postgres_load)

Continuously inserts rows into a table, I use that for a failover test in AWS
## setup database
```
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
docker run -it --rm --link some-postgres:postgres postgres psql -h postgres -U postgres
Password for user postgres:
psql (12.2 (Debian 12.2-1.pgdg100+1))
Type "help" for help.

postgres=# create database postgresql
postgres-# \connect postgresql
FATAL:  database "store" does not exist
Previous connection kept
postgres-# create table id(id bigint GENERATED ALWAYS AS IDENTITY, created_at timestamp, updated_at timestamp)
```


## run it
```
go run postgres-load.go -user postgres -password password -url localhost -database postgresql
```
