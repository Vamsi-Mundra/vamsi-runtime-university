package client

import (
	"context"
	"errors"
	"reflect"
	"testing"

	__ "github.com/heroku/vamsi-runtime-university/spec"
	"google.golang.org/grpc"
)

func TestGetFeatures(t *testing.T) {

	expectedFeature, testPoints := setup()
	rgClient := &fake_RouteGuideCLient{features: expectedFeature}
	routeGuide := NewRouteGuide(rgClient)
	ctx := new(context.Context)
	receivedFeature, err := routeGuide.GetFeatures(*ctx, testPoints)

	if err != nil {
		t.Fatalf("We didn't receive any features")
		return
	}

	if !reflect.DeepEqual(expectedFeature, receivedFeature) {
		t.Fatalf("We didn't receive Correct features")
		return
	}

	points := []__.Point{{Latitude: -1, Longitude: 0}}

	_, error_1 := routeGuide.GetFeatures(*ctx, points)
	if error_1 == nil {
		t.Fatalf("Expected an Error")
	}
}

func setup() ([]__.Feature, []__.Point) {
	//Add a location with a valid name
	point1 := &__.Point{Latitude: 90, Longitude: 0}
	feature1 := __.Feature{Name: "North Pole", Location: point1}

	point2 := &__.Point{Latitude: -90, Longitude: 0}
	feature2 := __.Feature{Name: "South Pole", Location: point2}

	points := []__.Point{*point1, *point2}
	features := []__.Feature{feature1, feature2}

	return features, points
}

type fake_RouteGuideCLient struct {
	features []__.Feature
}

func (fr *fake_RouteGuideCLient) GetFeature(ctx context.Context, in *__.Point, opts ...grpc.CallOption) (*__.Feature, error) {
	for _, feature := range fr.features {
		if reflect.DeepEqual(feature.Location, in) {
			return &feature, nil
		}
	}
	return nil, errors.New("Didn't find the location")
}

func (fr *fake_RouteGuideCLient) ListFeatures(ctx context.Context, in *__.Rectangle, opts ...grpc.CallOption) (__.RouteGuide_ListFeaturesClient, error) {
	return nil, nil
}

func (fr *fake_RouteGuideCLient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (__.RouteGuide_RecordRouteClient, error) {
	return nil, nil
}

func (fr *fake_RouteGuideCLient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (__.RouteGuide_RouteChatClient, error) {
	return nil, nil
}
