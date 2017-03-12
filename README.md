# SQL Reflect

_THIS IS HIGHLY EXPERIMENTAL AND UNDER ACTIVE DEVELOPMENT_

This library provides access to the structure of a SQL database,
allowing developers to "reflect" on the database itself. Why would you
need to do this?

- Find out all of the tables available in a database
- For a table, find out what columns it has
- Find out about the relationships between tables
  - Discover foreign keys, and automatically relate tables
  - Discover primary keys

## How It Works

This uses the `information_schema` database defined in the SQL standard.
The library is developed against PostgreSQL, but it should work on any
database that provides an approximately compliant implementation.

The `information_schema` tables (or views) provide information about the
structure of a database, and are designed to enable this sort of
reflection.

As the name of the library suggests, the code is designed to feel
roughly similar to Go's own reflection package. However, there is not a
one-to-one mapping between a concept like a table or column and Go's
concepts like type and value.

Under the hood, it uses
[squirrel](https://github.com/Masterminds/squirrel)
and [structable](https://github.com/Masterminds/structable) 4.

Tested on:

- Postgres 9.4, 9.6

## Terminology

For the most part, this library follows the terminology of the SQL
standard (e.g. `catalog` usually means `database`). Because I am most
familiar with PostgreSQL, it is likely that some Postgres terminology
slipped in here too.

Most of my understanding of the information schema comes from the
PostgreSQL documentation.

## TODO

- [ ] Enable querying CHECK constraints to actually return the
  constraint. (See `information_schame.check_constraints`)
