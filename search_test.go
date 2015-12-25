package discogs
import (
	"testing"
	"reflect"
	"fmt"
)

func TestParameters(*testing.T) {
	p := searchParameters{}
	typ := reflect.TypeOf(p)
	for i := 0; i < typ.NumField(); i++ {
    f := typ.Field(i)
		choices := f.Tag.Get("choices")
		if len(choices) != 0 {
			fmt.Println(f.Name, "choices:", choices)
		}
	}
}

func TestSearchFunction(*testing.T) {
	params := searchParameters{
		Query: "Blakroc",
		Type: "album",
	}
	results, err := params.Run()
	if err != nil {
		panic("search error")
	}
	if results == nil {
		panic("empty result set")
	}
}
