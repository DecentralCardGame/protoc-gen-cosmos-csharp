# protoc-gen-cosmosCsharp
A protoc plugin for generating CosmCs compatible tx routes from protobuf.

## Example usage
bug.gen.yaml
```yaml
version: v1
managed:
enabled: true
plugins:
- plugin: cosmosCsharp
out: out
```

## License and Development
This is licensed under the GPL-v3 License, part of the [DecentralCardgame](https://github.com/DecentralCardGame) project and developed by [lxgr-linux](https://github.com/lxgr-linux).