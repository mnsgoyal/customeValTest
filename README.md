protoc  \
-I ./interfaces/ \
-I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 \
--proto_path=${GOPATH}/src/github.com/google/protobuf/src \
--proto_path=${GOPATH}/src/github.com/maanasasubrahmanyam-sd/customValidation \
--proto_path=. \
--go_out="plugins=grpc:./generated"  \
--validate_out="lang=go:./generated" \
./interfaces/test_server/*.proto