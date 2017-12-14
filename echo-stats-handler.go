package stats_handler

import (
	"encoding/json"
	"net/http"

	"github.com/fukata/golang-stats-api-handler"
	"github.com/labstack/echo"
)

var newLineTerm = false
var prettyPrint = false

// NewLineTermEnabled enable termination with newline for response body.
func NewLineTermEnabled() {
	newLineTerm = true
}

// NewLineTermDisabled disable termination with newline for response body.
func NewLineTermDisabled() {
	newLineTerm = false
}

// PrettyPrintEnabled enable pretty-print for response body.
func PrettyPrintEnabled() {
	prettyPrint = true
}

// PrettyPrintDisabled disable pritty-print for response body.
func PrettyPrintDisabled() {
	prettyPrint = false
}

// EchoStatsHandler show activity stats for echo
func EchoStatsHandler(c echo.Context) error {
	values := c.Request().URL.Query()
	for _, p := range []string{"1", "true"} {
		if values.Get("pp") == p {
			prettyPrint = true
		}
	}

	var jsonBytes []byte
	var jsonErr error
	if prettyPrint {
		jsonBytes, jsonErr = json.MarshalIndent(stats_api.GetStats(), "", "  ")
	} else {
		jsonBytes, jsonErr = json.Marshal(stats_api.GetStats())
	}
	var body string
	if jsonErr != nil {
		body = jsonErr.Error()
	} else {
		body = string(jsonBytes)
	}

	if newLineTerm {
		body += "\n"
	}

	if jsonErr != nil {
		return c.String(http.StatusInternalServerError, body)
	}
	return c.String(http.StatusOK, body)
}
