package discogs
import "errors"

const URI = "/database/search"

type SearchParameters struct {
	// See; https://www.discogs.com/developers/#page:database,header:database-search
	Query		string    `json:"q,omitempty"`
	// Your search query
	// string (optional) Example: nirvana
	Type			string	`choices:"release,master,artist,label"`
	// string (optional) Example: release
	// String. One of release, master, artist, label

	Title			string
	// Search by combined “Artist Name - Release Title” title field.
	// string (optional) Example: nirvana - nevermind

	ReleaseTitle	string
	// string (optional) Example: nevermind
	// Search release titles.

	Credit			string
	// string (optional) Example: kurt
	// Search release credits.

	Artist			string
	// string (optional) Example: nirvana
	// Search artist names.

	Anv				string
	// string (optional) Example: nirvana
	// Search artist ANV.

	Label			string
	// string (optional) Example: dgc
	// Search label names.

	Genre			string
	// string (optional) Example: rock
	// Search genres.

	Style			string
	// string (optional) Example: grunge
	// Search styles.

	Country			string
	// string (optional) Example: canada
	// Search release country.

	Year			string
	// string (optional) Example: 1991
	// Search release year.

	Format			string
	// string (optional) Example: album
	// Search formats.

	Catno			string
	// string (optional) Example: DGCD-24425
	// Search catalog number.

	Barcode			string
	// string (optional) Example: 7 2064-24425-2 4
	// Search barcodes.

	Track			string
	// string (optional) Example: smells like teen spirit
	// Search track titles.

	Submitter		string
	// string (optional) Example: milKt
	// Search submitter username.

	Contributor		string
	// string (optional) Example: jerome99
	// Search contributor usernames.
}

func (SearchParameters) Validate() (error) {
	return errors.New("invalid search parameters")
}

func (SearchParameters) Run() (results interface{}, err error) {
	return nil, errors.New("search failed")
}

func Search(query string, parameters SearchParameters)  {
	search := SearchParameters{
		Query: query,
	}
	search.Run()
}
