package main

import (
	"bytes"
	"testing"
)

func TestStore(t *testing.T) {
	var storeOpts StoreOpts = StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}

	s := NewStore(storeOpts)
	err := s.writeStream("PathForTesting", bytes.NewReader([]byte("Some Bytes for testing")))
	if err != nil {
		t.Error(err)
	}
}
