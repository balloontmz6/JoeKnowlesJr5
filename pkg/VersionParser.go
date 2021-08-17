package pkg

import (
	"net/http"

	"github.com/pierrre/imageserver"
	imageserver_http "github.com/pierrre/imageserver/http"
)

// VersionParser VersionParser
type VersionParser struct{}

// Parse implements imageserver/http.Parser.
func (parser *VersionParser) Parse(req *http.Request, params imageserver.Params) error {
	imageserver_http.ParseQueryString("ver", req, params)
	if !params.Has("ver") {
		return nil
	}
	ver, err := params.GetString("ver")
	if err != nil {
		return err
	}
	params.Set("ver", ver)
	return nil
}

// Resolve implements imageserver/http.Parser.
func (parser *VersionParser) Resolve(param string) string {
	if param == "ver" {
		return "ver"
	}
	return ""
}
