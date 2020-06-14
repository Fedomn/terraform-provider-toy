package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type fileManifestEntry struct {
	Path  string
	IsDir bool
}

func main() {
	filePath := os.Args[1]
	fd, _ := os.Open(filePath)
	dec := gob.NewDecoder(fd)

	entry := fileManifestEntry{}
	_ = dec.Decode(&entry)
	fmt.Printf("IsDir: %t\n", entry.IsDir)
	fmt.Println("Path:")
	fmt.Println(entry.Path)
}
