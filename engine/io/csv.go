package io

import (
	"bufio"
	"fmt"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	"log"
	"os"
)

type CsvHandler struct {
	Filename string
}

func (c *CsvHandler) Write(pairs <-chan *external.Pair) {
	f, err := os.OpenFile(c.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("cannot open file for writing: %s", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()

	for pair := range pairs {
		_, err := w.WriteString(fmt.Sprintf("%s,%s\n", pair.GetKey().GetKey(), pair.GetValue().GetValue()))
		if err != nil {
			log.Fatalf("cannot write to file: %s", err)
		}
	}
}

func (c *CsvHandler) Read(split *internal.Split) <-chan *external.Pair {
	panic("not implemented yet")
}

func (c *CsvHandler) Split(spec *external.InputSpec) []internal.Split {
	panic("not implemented yet")
}
