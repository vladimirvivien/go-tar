package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	tarPath := "out.tar"
	files := map[string]string{
		"index.html": `<head><title>Hello</title></head><body>Hello!</body>`,
		"lang.json":  `[{"code":"eng","name":"English"},{"code":"fr","name":"French"}]`,
		"songs.txt":  `Claire de la lune, The Valkyrie, Swan Lake`,
	}

	tarWrite := func(data map[string]string) error {
		tarFile, err := os.Create(tarPath)
		if err != nil {
			return err
		}
		defer tarFile.Close()

		tw := tar.NewWriter(tarFile)
		defer tw.Close()

		for name, content := range data {
			hdr := &tar.Header{
				Name: name,
				Mode: 0600,
				Size: int64(len(content)),
			}
			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}
			if _, err := tw.Write([]byte(content)); err != nil {
				return err
			}
		}
		return nil
	}

	tarUnwrite := func() error {
		tarFile, err := os.Open(tarPath)
		if err != nil {
			return err
		}
		defer tarFile.Close()

		tr := tar.NewReader(tarFile)
		for {
			hdr, err := tr.Next()
			if err == io.EOF {
				break // End of archive
			}
			if err != nil {
				return err
			}
			fmt.Printf("Contents of %s: ", hdr.Name)
			if _, err := io.Copy(os.Stdout, tr); err != nil {
				return err
			}
			fmt.Println()
		}
		return nil
	}

	if err := tarWrite(files); err != nil {
		log.Fatal(err)
	}

	if err := tarUnwrite(); err != nil {
		log.Fatal(err)
	}
}
