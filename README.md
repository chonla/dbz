DBZ
=====

Initialize a database from a YML file.

## Supported Drivers

* Sqlite

## Execute

```
go run main.go -conf example.yml
```

## Dump sql command (no execution)

```
go run main.go -conf example.yml -dump
```

## Overwrite output

```
go run main.go -conf example.yml -overwrite
```

## License

MIT: http://chonla.mit-license.org/