# protoc-gen-uwentity

Custom protoc plugin that generates a `GetEntityIdentifier` getter against the generated protobuf message.

We use this to retrieve the partition key when writing the event to the message broker. The benefit is the protobuf
is your source of truth.

**1) Pull in the extended field options**

[Buf](https://buf.build/) makes that super easy! 

Add `protoc-gen-uwentity` to your project dependencies in your `buf.yaml` file:
```yaml
version: v1beta1
name: buf.build/utilitywarehouse/some-project
deps:
  - buf.build/utilitywarehouse/protoc-gen-uwentity
```

Install locally the plugin and add it to `buf.gen.yaml`:
```yaml
version: v1
plugins:
  # ... your other plugins
  - name: uwentity
    out: gen/go
    opt:
      - paths=source_relative
```

**2) Annotate the field**

```protobuf
import "uw/entity/v1/identifier.proto";

message UserCreatedEvent {
  string user_id = 1 [(uw.entity.v1.identifier) = true];
}
```

You can also choose to annotate a field of the type that has its own entity identifier already.
In this case, the UserCreationVerifiedEvent will have the same identifier as the UserCreatedEvent.

```protobuf
import "uw/entity/v1/identifier.proto";

message UserCreatedEvent {
  string user_id = 1 [(uw.entity.v1.identifier) = true];
}

message UserCreationVerifiedEvent {
  UserCreatedEvent user_created_event = 1 [(uw.entity.v1.identifier) = true];
}
```

## Enforcement

Any protobuf message that has the `Event` suffix will be checked that an identifier has been set.
You can skip this check by setting `option (uw.entity.v1.ignore) = true;` on the message.

If you want to restrict which directories are enforced, you can pass `enforce-dir=<path>` param to
protoc-gen-uwentity and only protobuf messages that live under that directory will be checked.
