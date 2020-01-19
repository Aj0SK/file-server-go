package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
)

func hashMD5(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func listFiles(dirname string) string {
	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var list string = ""

	for _, file := range files {
		list += fmt.Sprintf("%s\t%d\t%s\n", file.Name(), file.Size(), file.ModTime().String())
	}

	return list
}
