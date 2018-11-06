package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	chunkSize := flag.Int("size", 15, "chunk size in MB")
	fileName := flag.String("file", "", "file name to check")
	flag.Parse()

	if *fileName == "" {
		flag.Usage()
		os.Exit(1)
	}

	reader, err := os.Open(*fileName)
	defer reader.Close()
	if err != nil {
		panic(err.Error())
	}

	chunkSizeInBytes := 1024 * 1024 * *chunkSize
	buffer := make([]byte, chunkSizeInBytes)
	hasher := md5.New()

	scanner := bufio.NewScanner(reader)
	scanner.Buffer(buffer, chunkSizeInBytes)
	scanner.Split(splitByBufferSize)

	totalChunks := 0
	var md5bytes []byte

	for scanner.Scan() {
		_, err := hasher.Write(scanner.Bytes())
		if err != nil {
			panic(err.Error())
		}
		md5bytes = append(md5bytes, hasher.Sum(nil)...)
		totalChunks++
		hasher.Reset()
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	hasher.Write(md5bytes)
	filemd5 := hex.EncodeToString(hasher.Sum(nil))

	fmt.Printf("%s-%d\n", filemd5, totalChunks)
}

func splitByBufferSize(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	return len(data), data[0:], nil
}
