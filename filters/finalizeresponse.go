package filters

import "github.com/zalando/skipper/filters"

// This filter is the default filter for every configuration of skrop
// It
type finalizeResponse struct{}

const (
	FinalizeResponseName = "finalizeResponse"
)

func NewFinalizeResponse() filters.Spec {
	return &finalizeResponse{}
}

func (s *finalizeResponse) Name() string {
	return FinalizeResponseName
}

func (s *finalizeResponse) CreateFilter(args []interface{}) (filters.Filter, error) {
	if len(args) != 0 {
		return nil, filters.ErrInvalidFilterParameters
	}

	return &finalizeResponse{}, nil
}

func (s *finalizeResponse) Request(ctx filters.FilterContext) {}

//finalize the response calling the transformer for the image one last time before returning the image to the client
func (s *finalizeResponse) Response(ctx filters.FilterContext) {
	FinalizeResponse(ctx)
}
