package main

import (
	"context"
	"flag"
	"log"

	"github.com/heroku/vamsi-runtime-university/client"
	__ "github.com/heroku/vamsi-runtime-university/spec"
	"github.com/joeshaw/envdecode"
	"google.golang.org/grpc"
)

type Config struct {
	GrpcHostName string `env:"SERVER_URL,default=grpc-server.herokuapp.com:80"`
}

func createConn(hostName string) *grpc.ClientConn {
	conn, err := grpc.Dial(hostName, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

func createGrpcClient(conn *grpc.ClientConn) (*client.RouteGuide, context.Context) {
	rgc := __.NewRouteGuideClient(conn)
	rg := client.NewRouteGuide(rgc)
	ctx := context.Background()
	return rg, ctx
}

func main() {

	flag.String("help", "", "Enter a latitude & longitude to receive the feature for that point.")
	var latitude = flag.Int("latitude", 0, "Enter a latitude value.")
	var longitude = flag.Int("longitude", 0, "Enter a longitude value.")
	flag.Parse()
	// flag.String("help", "", "")
	var cfg Config
	err := envdecode.Decode(&cfg)
	if err != nil {
		log.Fatalln(err)
	}
	connection := createConn(cfg.GrpcHostName)
	defer connection.Close()
	routeGuide, ctx := createGrpcClient(connection)
	point := []__.Point{{Latitude: int32(*latitude), Longitude: int32(*longitude)}}
	results, err := routeGuide.GetFeatures(ctx, point)

	if err != nil {
		log.Fatalln(err)
	}

	for _, result := range results {
		log.Println(" Feature Name : %v", result.Name)
		log.Println(result.Name)
	}
	log.Println("Results : ")
	log.Println(results)
}
