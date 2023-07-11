package main

import (
	"github.com/stretchr/testify/require"
	"github.com/utilitywarehouse/protoc-gen-uwentity/testgen/testgen/data"
	"testing"
)

func TestSimpleGetEntityIdentifier(t *testing.T) {
	msg := &data.SimpleMessage{
		Id: "test",
	}
	require.Equal(t, "test", msg.GetEntityIdentifier())
}

func TestNestedGetEntityIdentifier(t *testing.T) {
	simpleMessage := &data.SimpleMessage{
		Id: "test",
	}

	msg := &data.NestedMessage{
		SimpleMessage: simpleMessage,
	}

	require.Equal(t, "test", msg.GetEntityIdentifier())
}
