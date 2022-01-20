package client

import (
	"context"

	__ "github.com/heroku/vamsi-runtime-university/spec"
)

type RouteGuide struct {
	client __.RouteGuideClient
}

func (rg *RouteGuide) GetFeatures(ctx context.Context, points []__.Point) ([]__.Feature, error) {
	//TODO Add your implementation here
	features := make([]__.Feature, len(points))
	for i := 0; i < len(points); i++ {
		feat, e := rg.client.GetFeature(ctx, &points[i])
		if e != nil {
			return nil, e
		}
		features[i] = *feat
	}
	return features, nil
}
