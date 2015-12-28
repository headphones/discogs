package discogs
import "errors"

type API struct {}

func New(key string, secret string) (api *API, err error) {
	return api, errors.New("empty method")
}
