package discogs
import (
	"testing"
)

func TestSearchFunction(*testing.T) {
	api := API{}
	params := map[string]string{
		"type": "release",
	}
	results, err := api.Search("blackroc", params)
	if err != nil {
		panic(err)
	}
	if results == nil {
		panic("empty result set")
	}
}

func TestSearchInputValidation(*testing.T) {
	api := API{}
	badParam := map[string]string{
		"typeee": "release",
	}
	results, err := api.Search("blackroc", badParam)
	if err == nil {
		panic("invalid search parameter given, but was not caught")
	}
	if results != nil {
		panic("result set given for bad search")
	}
	invalidChoice := map[string]string{
		"type": "album",
	}
	results, err = api.Search("blackroc", invalidChoice)
	if err == nil {
		panic("invalid choice given for parameter, but was not caught")
	}
	if results != nil {
		panic("result set given for bad search")
	}
}
