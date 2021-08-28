protoc  \
-I. \
-I $GOPATH/src/ \
--proto_path=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 \
--proto_path=${GOPATH}/src/github.com/google/protobuf/src \
--go_out="plugins=grpc:."  \
--validate_out="lang=go:." \
--govalidators_out=. \
./interfaces/test_server/*.proto


or

protoc  \
-I ./interfaces/test_server \
-I $GOPATH/src/ \
--proto_path=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 \
--proto_path=${GOPATH}/src/github.com/google/protobuf/src \
--go_out="plugins=grpc:."  \
--validate_out="lang=go:." \
--govalidators_out=. \
./interfaces/test_server/test.proto


protoc  \
-I ./interfaces/test_server \
-I $GOPATH/src/ \
--go_out="plugins=grpc:."  \
--validate_out="lang=go:." \
--govalidators_out=. \
./interfaces/test_server/test.proto