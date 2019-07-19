package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	contents := map[string]string{
		"index.html": `<head><title>Hello</title></head><body>Hello!</body>`,
		"lang.json":  `[{"code":"eng","name":"English"},{"code":"fr","name":"French"}]`,
		"songs.txt":  `Claire de la lune, The Valkyrie, Swan Lake`,
	}

	for name, content := range contents {
		hdr := &tar.Header{
			Name: name,
			Mode: 0600,
			Size: int64(len(content)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(content)); err != nil {
			log.Fatal(err)
		}
	}

	// close the tar writer before reading from buf
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// Open and iterate through the files in the archive.
	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
