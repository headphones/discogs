package discogs
import (
	"errors"
	"fmt"
)

const searchURI = "/database/search"

type searchParameter struct {
	// choices is a map so that we might query string choices membership
	// without iterating over it (boolean is not currently used)
	Choices	map[string]bool
}

// See; https://www.discogs.com/developers/#page:database,header:database-search
var searchParameters = map[string]searchParameter {
	"query": searchParameter{
		// Your search query
		// string (optional) Example: nirvana
	},
	"type": searchParameter{
		// string (optional) Example: release
		// String. One of release, master, artist, label
		Choices: map[string]bool{
			"release": true,
			"master": true,
			"artist": true,
			"label": true,
		},
	},
	"title": searchParameter{
		// Search by combined “Artist Name - Release Title” title field.
		// string (optional) Example: nirvana - nevermind
	},
	"release_title": searchParameter{
		// string (optional) Example: nevermind
		// Search release titles.
	},
	"credit": searchParameter{
		// string (optional) Example: kurt
		// Search release credits.
	},
	"artist": searchParameter{
		// string (optional) Example: nirvana
		// Search artist names.
	},
	"anv": searchParameter{
		// string (optional) Example: nirvana
		// Search artist ANV.
	},
	"label": searchParameter{
		// string (optional) Example: dgc
		// Search label names.
	},
	"genre": searchParameter{
		// string (optional) Example: rock
		// Search genres.
	},
	"style": searchParameter{
		// string (optional) Example: grunge
		// Search styles.
	},
	"country": searchParameter{
		// string (optional) Example: canada
		// Search release country.
	},
	"year": searchParameter{
		// string (optional) Example: 1991
		// Search release year.
	},
	"format": searchParameter{
		// string (optional) Example: album
		// Search formats.
	},
	"catno": searchParameter{
		// string (optional) Example: DGCD-24425
		// Search catalog number.
	},
	"barcode": searchParameter{
		// string (optional) Example: 7 2064-24425-2 4
		// Search barcodes.
	},
	"track": searchParameter{
		// string (optional) Example: smells like teen spirit
		// Search track titles.
	},
	"submitter": searchParameter{
		// string (optional) Example: milKt
		// Search submitter username.
	},
	"contributor": searchParameter{
		// string (optional) Example: jerome99
		// Search contributor usernames.
	},
}

func validateQuery(params map[string]string) (error) {
	for p, v := range(params) {
		// check parameter name
		param, ok := searchParameters[p]
		if !ok {
			return errors.New(
				fmt.Sprintf("invalid parameter: %v (value was %v)", p, v))
		}
		// check parameter value
		if param.Choices != nil {
			if _, ok := param.Choices[v]; !ok {
				return errors.New(
					fmt.Sprintf("invalid choice \"%v\" for parameter \"%v\"", v, p))
			}
		}
	}
	// check parameter values
	for k, v := range(params) {
		if _, ok := searchParameters[k]; !ok {
			return errors.New(
				fmt.Sprintf("invalid parameter: %v (value was %v)", k, v))
		}
	}
	return nil
}

func (API) Search(query string, params map[string]string) (interface{}, error) {
	// params should be a set of two-member arrays (key, value)
	if err := validateQuery(params); err != nil {
		return nil, err
	}

	return params, nil
}
