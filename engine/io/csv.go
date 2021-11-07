package io

import (
	"fmt"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	"log"
	"os"
)

type CsvHandler struct {
	Filename string
}

func (c *CsvHandler) Write(pairs <-chan *internal.Pair) {
	f, err := os.Create(c.Filename)
	if err != nil {
		log.Fatalf("cannot open file for writing: %s", err)
	}
	defer f.Close()
	defer f.Sync()

	for pair := range pairs {
		_, err := f.WriteString(fmt.Sprintf("%s,%s\n", pair.GetKey().GetKey(), pair.GetValue().GetValue()))
		if err != nil {
			log.Fatalf("cannot write to file: %s", err)
		}
	}
}

func (c *CsvHandler) Read(split *internal.Split) <-chan *internal.Pair {
	panic("not implemented yet")
}

func (c *CsvHandler) Split(spec *external.InputSpec) []internal.Split {
	panic("not implemented yet")
}
