module github.com/isucon/isucon12-portal/bench-tool.go/benchrun

go 1.18

replace github.com/isucon/isucon12-portal/proto.go/isuxportal/resources => ../../proto.go/isuxportal/resources

require (
	github.com/golang/protobuf v1.5.2
	github.com/isucon/isucon12-portal/proto.go/isuxportal/resources v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.28.0
)
