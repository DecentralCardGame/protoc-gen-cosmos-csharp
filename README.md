# protoc-gen-cosmos-csharp

A protoc plugin for generating CosmCs compatible tx routes from protobuf.

## Installation

```shell
go install github.com/DecentralCardGame/protoc-gen-cosmos-csharp@latest
```

## Example usage

bug.gen.yaml

```yaml
version: v1
managed:
    enabled: true
    go_package_prefix:
        default: github.com/yourModule
plugins:
    - plugin: cosmos-csharp
          out: out
          opt: suffix=suffix.pb.cs
    - plugin: buf.build/protocolbuffers/csharp
          out: out
          opt: file_extension=.pb.cs,base_namespace=
    - plugin: buf.build/grpc/csharp
          out: out
          opt: no_server,file_suffix=Grpc.pb.cs,base_namespace=
```

## License and Development

This is licensed under the GPL-v3 License, part of the [DecentralCardgame](https://github.com/DecentralCardGame) project
and developed by [lxgr-linux](https://github.com/lxgr-linux).