package io

import (
	"bufio"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	"log"
	"os"
	"strconv"
)

type TextHandler struct {
}

func (t *TextHandler) Split(spec *external.InputSpec) []internal.Split {
	splits := make([]internal.Split, 0)
	for _, file := range spec.GetInputFiles() {
		splits = addSplitsForFile(spec, file, splits)
	}
	return splits
}

func addSplitsForFile(spec *external.InputSpec, file *external.DFSFile, splits []internal.Split) []internal.Split {
	//loc := strings.ReplaceAll(file.GetLocation(), "/mnt/d", "D:") //TODO
	loc := file.GetLocation()
	f, err := os.Open(loc)
	if err != nil {
		log.Fatalf("cannot split files: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	buf := make([]byte, spec.GetInputSplitSizeBytes())
	scanner.Buffer(buf, int(spec.GetInputSplitSizeBytes()))

	var offset, prevOffset int64
	for scanner.Scan() {
		length := int64(len(scanner.Bytes()))
		if offset+length > spec.GetInputSplitSizeBytes() {
			splits = append(splits, internal.Split{Source: file, Offset: prevOffset, Limit: offset})
			prevOffset = offset
		}
		offset += length + 1 // 1 byte for stripped newline
	}
	splits = append(splits, internal.Split{Source: file, Offset: prevOffset, Limit: offset})
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return splits
}

func (t *TextHandler) Read(split *internal.Split) <-chan *external.Pair {
	c := make(chan *external.Pair)
	go func() {
		defer close(c)

		f, err := os.Open(split.Source.GetLocation())
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		_, err = f.Seek(split.Offset, 0)
		if err != nil {
			log.Fatalln(err)
		}
		scanner := bufio.NewScanner(f)
		buf := make([]byte, split.Limit-split.Offset)
		scanner.Buffer(buf, int(split.Limit-split.Offset))

		offset := split.Offset
		for scanner.Scan() {
			c <- &external.Pair{Key: &external.Key{Key: strconv.FormatInt(offset, 10)}, Value: &external.Value{Value: scanner.Text()}}
			offset += int64(len(scanner.Bytes())) + 1
			if offset >= split.Limit {
				break
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	return c
}

func (t *TextHandler) Write(pairs <-chan *external.Pair) {
	panic("not implemented yet")
}
