# How to use

Generate a `keys.yaml` with the keys needed
```
keys:
  - foo
  - bar
  - top.middle.bottom
```

Run it against the source yaml
```
go run main.go -k keys.yaml -i source.yaml -o output.yaml
```