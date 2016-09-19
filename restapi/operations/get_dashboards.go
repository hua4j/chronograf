package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDashboardsHandlerFunc turns a function with the right signature into a get dashboards handler
type GetDashboardsHandlerFunc func(context.Context, GetDashboardsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDashboardsHandlerFunc) Handle(ctx context.Context, params GetDashboardsParams) middleware.Responder {
	return fn(ctx, params)
}

// GetDashboardsHandler interface for that can handle valid get dashboards params
type GetDashboardsHandler interface {
	Handle(context.Context, GetDashboardsParams) middleware.Responder
}

// NewGetDashboards creates a new http.Handler for the get dashboards operation
func NewGetDashboards(ctx *middleware.Context, handler GetDashboardsHandler) *GetDashboards {
	return &GetDashboards{Context: ctx, Handler: handler}
}

/*GetDashboards swagger:route GET /dashboards getDashboards

Pre-configured dashboards

Dashboards are a collection of `Cells` that visualize time-series data.


*/
type GetDashboards struct {
	Context *middleware.Context
	Handler GetDashboardsHandler
}

func (o *GetDashboards) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetDashboardsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}