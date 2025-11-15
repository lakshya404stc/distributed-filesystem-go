package main

import (
	"io"
	"log"
	"os"
	"path"
)

func CASPathTransformFunc(key string) (string, error) {
 return key, nil
}

type PathTransformFunc func(string) (string, error)

func DefaultPathTransformFunc(key string) (string, error){
	return key, nil
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

	filename := "testFileName.txt"
	pathAndFilename := path.Join(pathName, filename)
	f, err := os.Create(pathAndFilename)
	if err != nil {
		return err
	}
	n, err := io.Copy(f, r) 
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Written (%d) bytes to disk: %s", n, pathAndFilename)
	return nil
}
