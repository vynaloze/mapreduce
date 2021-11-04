package master

import (
	"fmt"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	"github.com/vynaloze/mapreduce/engine/io"
)

func (s *scheduler) split(spec *external.InputSpec) ([]internal.Split, error) {
	textHandler := &io.TextHandler{}

	splits := make([]internal.Split, 0)
	for _, file := range spec.GetInputFiles() {
		switch format := file.GetFormat().GetFormat(); format {
		case external.FileFormat_TEXT:
			splits = append(splits, textHandler.Split(spec)...)
		default:
			return nil, fmt.Errorf("unsupported format: %+v", format)
		}
	}
	return splits, nil
}
