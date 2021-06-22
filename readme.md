# protoc-gen-uwentity

Custom protoc plugin that generates a `GetEntityIdentifier` getter against the generated protobuf message.

We use this to retrieve the partition key when writing the event to the message broker. The benefit is the protobuf
is your source of truth.

**1) Pull in the extended field options**

[buf](https://buf.build/) makes that super easy!

```yaml
version: v1beta1
name: buf.build/utilitywarehouse/some-project
deps:
  - buf.build/utilitywarehouse/protoc-gen-uwentity
```

**2) Annotate the field**

```protobuf
import "uw/entity/v1/identifier.proto";

message UserCreatedEvent {
  string user_id = 1 [(uw.entity.v1.identifier) = true];
}
```

