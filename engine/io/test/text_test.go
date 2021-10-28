package test

import (
	"github.com/stretchr/testify/assert"
	external "github.com/vynaloze/mapreduce/api"
	internal "github.com/vynaloze/mapreduce/engine/api"
	"github.com/vynaloze/mapreduce/engine/io"
	"testing"
)

func TestTextSplit(t *testing.T) {
	// given
	format := &external.FileFormat{Format: external.FileFormat_TEXT}
	spec := &external.InputSpec{
		InputFiles: []*external.DFSFile{
			{Location: "input/01.txt", Format: format, SizeBytes: 4 * 1024},
			{Location: "input/02.txt", Format: format, SizeBytes: 4 * 1024},
		},
		InputSplitSizeBytes: 1.5 * 1024,
	}
	expected := []io.Split{
		{Source: spec.GetInputFiles()[0].GetLocation(), Offset: 0, Limit: 1304},
		{Source: spec.GetInputFiles()[0].GetLocation(), Offset: 1304, Limit: 2039},
		{Source: spec.GetInputFiles()[0].GetLocation(), Offset: 2039, Limit: 2683},
		{Source: spec.GetInputFiles()[0].GetLocation(), Offset: 2683, Limit: 3395},
		{Source: spec.GetInputFiles()[0].GetLocation(), Offset: 3395, Limit: 4097},
		{Source: spec.GetInputFiles()[1].GetLocation(), Offset: 0, Limit: 958},
		{Source: spec.GetInputFiles()[1].GetLocation(), Offset: 958, Limit: 1890},
		{Source: spec.GetInputFiles()[1].GetLocation(), Offset: 1890, Limit: 2445},
		{Source: spec.GetInputFiles()[1].GetLocation(), Offset: 2445, Limit: 2988},
		{Source: spec.GetInputFiles()[1].GetLocation(), Offset: 2988, Limit: 3520},
		{Source: spec.GetInputFiles()[1].GetLocation(), Offset: 3520, Limit: 4097},
	}
	th := io.TextHandler{}
	// when
	splits := th.Split(spec)
	// then
	assert.ElementsMatch(t, splits, expected)
}

func TestTextRead(t *testing.T) {
	// given
	split := io.Split{Source: "input/01.txt", Offset: 0, Limit: 1304}
	expected := []io.Pair{
		{Key: &internal.Key{Key: "0"}, Value: &internal.Value{Value: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed blandit mollis varius. Duis pretium sem eget tortor maximus porttitor. Aenean in libero metus. Maecenas sagittis, sem nec suscipit aliquet, ipsum augue eleifend sem, nec laoreet leo ligula id erat. Morbi leo purus, cursus et orci quis, posuere rhoncus ligula. Vivamus et lacusneque. In lacinia ultrices massa nec eleifend. Praesent porttitor orci et felis finibus, sed porta eros blandit. Sed quis interdum sem, ut fermentum turpis. In hac habitasse platea dictumst. Aenean tincidunt tellus est, ut hendrerit lectus ullamcorper sit amet."}},
		{Key: &internal.Key{Key: "599"}, Value: &internal.Value{Value: "Etiam finibus mi sed interdum molestie. Pellentesque vel nisi ullamcorper, sodales mi at, malesuada tellus. Cras eleifend lacus id ipsum rhoncus consequat. Donec aliquet auctor blandit. Nulla suscipit malesuada turpis sed mollis. Aenean tincidunt dictum est, a scelerisque magna varius vel. In nunc diam, faucibus sed mattis vitae, dapibus sed sem. Aenean in pretium augue. Quisque aliquam orci eget arcu molestie, eget mattis dolor vestibulum. Vivamus eu ultricies quam. Mauris aliquet elit velit, in tincidunt elit vestibulum a. Integer congue venenatis erat, et imperdiet ex dignissim id. Curabitur pretium erat ullamcorper sapien ullamcorper ornare. Phasellus vitae mauris quis dui elementum lacinia."}},
	}
	th := io.TextHandler{}
	// when
	pairs := make([]io.Pair, 0)
	for p := range th.Read(split) {
		pairs = append(pairs, p)
	}
	// then
	assert.ElementsMatch(t, pairs, expected)
}
