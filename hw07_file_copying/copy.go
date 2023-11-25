package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	file, err := os.OpenFile(fromPath, os.O_RDWR, 0666) //nolint
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnsupportedFile, err)
	}
	defer func() { _ = file.Close() }()

	inf, err := file.Stat()
	if err != nil {
		return fmt.Errorf("getting stat: %w", err)
	}

	buf := bufio.NewReaderSize(file, int(inf.Size()))
	if offset > inf.Size() {
		return errors.New("offset exceeds file size")
	}
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrOffsetExceedsFileSize, err)
	}
	newFile, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("failed to trying create file: %w", err)
	}
	defer func() { _ = newFile.Close() }()

	if limit == 0 || limit > int64(buf.Size()) {
		limit = int64(buf.Size())
	}

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(buf)
	_, err = io.CopyN(newFile, barReader, limit)
	bar.Finish()
	if err != nil {
		log.Println(err)
	}
	return nil
}
