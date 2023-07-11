package main

import (
	"github.com/stretchr/testify/require"
	"github.com/utilitywarehouse/protoc-gen-uwentity/testdata/go/github.com/utilitywarehouse/protoc-gen-uwentity/testdata"
	"testing"
)

func TestSimpleGetEntityIdentifier(t *testing.T) {
	msg := &testdata.SimpleMessage{
		Id: "test",
	}

	require.Equal(t, "test", msg.GetEntityIdentifier())
}

func TestNestedGetEntityIdentifier(t *testing.T) {
	simpleMessage := &testdata.SimpleMessage{
		Id: "test",
	}

	msg := &testdata.NestedMessage{
		SimpleMessage: simpleMessage,
	}

	require.Equal(t, "test", msg.GetEntityIdentifier())
}
