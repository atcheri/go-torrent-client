package main

import (
	"fmt"
	"log"
	"os"

	"github.com/atcheri/go-torrent-client/internal/core/domain/torrentfile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nFinished downloading file %s\n", outPath)
}
