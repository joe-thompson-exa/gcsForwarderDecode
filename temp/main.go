package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var gzips []string
	var directory string

	// parse the flags
	flag.StringVar(&directory, "directory", "./discarded-events", "Please provide a directory to read the files from")
	flag.Parse()

	// walk the directory structure and create slice of paths to gzips
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if strings.HasSuffix(path, "proto.gz") {
			gzips = append(gzips, path)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	// unzip each gzip in place
	for _, filename := range gzips {
		gzipfile, err := os.Open(filename)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		reader, err := gzip.NewReader(gzipfile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer reader.Close()

		newfilename := strings.TrimSuffix(filename, ".gz")

		writer, err := os.Create(newfilename)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer writer.Close()

		if _, err = io.Copy(writer, reader); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
