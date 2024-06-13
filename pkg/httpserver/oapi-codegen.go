//go:generate go run github.com/telkomindonesia/openapi-utils/cmd/bundle ../../api/openapi-spec/src/profile.yml ../../api/openapi-spec/profile.yml
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config oapi-codegen.yml ../../api/openapi-spec/profile.yml
package httpserver

import (
	"context"

	"github.com/telkomindonesia/go-boilerplate/pkg/httpserver/internal/oapi"
)

var _ oapi.StrictServerInterface = oapiServerImplementation{}

type oapiServerImplementation struct {
	h *HTTPServer
}

// FindNinProfile implements oapi.StrictServerInterface.
func (s oapiServerImplementation) FindNinProfile(ctx context.Context, request oapi.FindNinRequestObject) (oapi.FindNinResponseObject, error) {
	panic("unimplemented")
}
