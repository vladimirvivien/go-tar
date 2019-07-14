package main

import (
	"archive/tar"
	"log"
	"os"
)

func main() {
	tarPath := "out.tar"
	tarFile, err := os.Create(tarPath)
	if err != nil {
		log.Fatal(err)
	}
	defer tarFile.Close()

	tw := tar.NewWriter(tarFile)
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
}
