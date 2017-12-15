package stats_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/goadesign/goa"
)

// StatsController implements the stats resource.
type StatsController struct {
	*goa.Controller
}

var newLineTerm = false
var prettyPrint = false

// NewStatsController creates a stats controller.
func NewStatsController(service *goa.Service) *StatsController {
	return &StatsController{Controller: service.NewController("StatsController")}
}

// ShowStats runs the ShowStats action.
func (c *StatsController) ShowStats(ctx *ShowStatsContext) error {
	values := ctx.URL.Query()
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
		return ctx.InternalServerError(body)
	}

	return ctx.OK(body)
}

// ShowStatsContext providers the stats ShowStats action context.
type ShowStatsContext struct {
	context.Context
	*goa.RequestData
	*goa.ResponseData
}

// NewShowStatsContext parses the incoming request URL and body, performs validations and creates the
// context used by the acquisition controller ShowStats action.
func NewShowStatsContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowStatsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowStatsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowStatsContext) OK(r string) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowStatsContext) InternalServerError(r string) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	return ctx.REsponseData.Service.Send(ctx.Context, 500, r)
}

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// StatsControllerInterface is the controller interface for the Stats actions.
type StatsControllerInterface interface {
	goa.Muxer
	ShowStats(*ShowStatsContext) error
}

// MountStatsController "mounts" a Stats resource controller on the given service.
func MountStatsController(service *goa.Service, ctrl StatsControllerInterface) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request.
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowStatsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ShowStats(rctx)
	}
	service.Mux.Handle("GET", "/stats", ctrl.MuxHandler("ShowStats", h, nil))
	service.LogInfo("mount", "ctrl", "Stats", "action", "ShowStats", "route", "GET /stats")
}
