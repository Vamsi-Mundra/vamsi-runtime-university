module github.com/heroku/vamsi-runtime-university

go 1.17

// +heroku goVersion go1.14
// +heroku install ./...

require (
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220114011407-0dd24b26b47d // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220118154757-00ab72f36ad5 // indirect
)
