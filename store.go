package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func CASPathTransformFunc(key string) (PathKey, error) {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLen := len(hashStr) / blockSize
	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		var from, to int = i * blockSize, (i * blockSize) + blockSize
		path := hashStr[from:to]
		paths[i] = path
	}
	fullPath := path.Join(paths...)

	return PathKey{Pathname: fullPath, Original: hashStr}, nil
}

type PathKey struct {
	Pathname string
	Original string
}

type PathTransformFunc func(string) (PathKey, error)

func DefaultPathTransformFunc(key string) (PathKey, error) {
	return PathKey{key,key}, nil
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathName, err := s.PathTransformFunc(key)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	io.Copy(buf, r)

	filenameBuf := md5.Sum(buf.Bytes())
	filenameString := hex.EncodeToString(filenameBuf[:])

	pathAndFilename := path.Join(pathName, filenameString)
	f, err := os.Create(pathAndFilename)
	if err != nil {
		return err
	}
	n, err := io.Copy(f, buf)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Written (%d) bytes to disk: %s", n, pathAndFilename)
	return nil
}
