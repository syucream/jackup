# jackup

Jack up your DDL and translate between MySQL and Spanner.

## How to use

- Pass your DDL via stdin or `-f` option

```
$ jackup -f path/to/ddl.sql
```

- With any options you hope

```
$ ./jackup -h
Usage of ./jackup:
  -allow-convert-string
        Convert between long string (default true)
  -f string
        path to DDL .sql file
  -remove-index-name
        Remove long index name (default true)
  -strict
        Strict check
  -t string
        convert target (spanner2mysql, ...) (default "spanner2mysql")
```
