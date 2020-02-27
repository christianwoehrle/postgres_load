docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
docker run -it --rm --link some-postgres:postgres postgres psql -h postgres -U postgres
Password for user postgres:
psql (12.2 (Debian 12.2-1.pgdg100+1))
Type "help" for help.

postgres=# create database store
postgres-# \connect store
FATAL:  database "store" does not exist
Previous connection kept
postgres-# create table id(id bigint GENERATED ALWAYS AS IDENTITY, created_at timestamp, updated_at timestamp)
