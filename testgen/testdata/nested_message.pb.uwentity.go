// Code generated by protoc-gen-uwentity. DO NOT EDIT.
// source: testgen/proto/testdata/nested_message.proto
package testdata

// GetEntityIdentifier returns the value from the field marked as the identifier
func (m *NestedMessage) GetEntityIdentifier() string {
	return m.GetSimpleMessage().GetEntityIdentifier()
}
