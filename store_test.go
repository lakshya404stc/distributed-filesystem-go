package main

import (
	"bytes"
	"fmt"
	"testing"
)


func TestStorePathTransformFunc(t *testing.T) {
	var key string = "testpathnumber1"
	pathName, err := CASPathTransformFunc(key)
	if err != nil {
		t.Error(err) 
	}
	fmt.Println(pathName)
	if pathName.Pathname != "2ad85/b1c05/d0930/01012/86512/80a47/010b7/48f1d" {
		t.Error("[ERROR] Path name mismatch")
	}
}

func TestStore(t *testing.T) {
	var storeOpts StoreOpts = StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(storeOpts)
	err := s.writeStream("PathForTesting", bytes.NewReader([]byte("Some Bytes for testing")))
	if err != nil {
		t.Error(err)
	}
}
