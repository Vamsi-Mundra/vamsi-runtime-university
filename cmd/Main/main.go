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
	ServerURL string `env:"SERVER_URL,default=grpc-server.herokuapp.com:80"`
}

func main() {
	var lat = flag.Int("lat", 407838351, "latitude of the point")
	var long = flag.Int("long", -746143763, "longitude of the point")
	flag.Parse()

	var cfg Config
	err := envdecode.Decode(&cfg)
	if err != nil {
		log.Fatal("Could not decode env variables: ", err)
	}

	points := []__.Point{
		{Latitude: int32(*lat), Longitude: int32(*long)},
	}

	conn, err := grpc.Dial(cfg.ServerURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not create client connection: ", err)
	}
	defer conn.Close()
	log.Println("connection created")

	rg := client.NewRouteGuide(__.NewRouteGuideClient(conn))
	features, err := rg.GetFeatures(context.Background(), points)
	if err != nil {
		log.Fatal("couldn't not get features: ", err)
	}

	log.Println("Features:", features)
}
