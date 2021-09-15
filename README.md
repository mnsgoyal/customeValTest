


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


working command
$ protoc   --proto_path=./interfaces/   --proto_path=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0   --proto_path=${GOPATH}/pkg/mod/github.com/maanasasubrahmanyam-sd/go-proto-validators@v0.3.2   --go_out="plugins=grpc:./generated"   --validate_out="lang=go:./generated"   ./interfaces/test_server/*.proto

This is gem
Rohit@Rohit MINGW64 ~/go/src/github.com/maanasasubrahmanyam-sd/customeValTest (master)
$  protoc   --proto_path=./interfaces/   --proto_path=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0   --proto_path=${GOPATH}/pkg/mod/github.com/maanasasubrahmanyam-sd/go-proto-validators@v0.3.7   --go_out="plugins=grpc:./generated"   --validate_out="lang=go:./generated" --go_opt=Mvalidator.proto=github.com/maanasasubrahmanyam-sd/go-proto-validators    ./interfaces/test_server/*.proto

for miwitkow
$ protoc \
> -I. \
> -I $GOPATH/src/ \
> -I $GOPATH/src/github.com/google/protobuf/src/ \
> --proto_path=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 \
> --go_out=plugins=grpc:./generated \
> --govalidators_out=./generated \
> ./interfaces/test_server/*.proto


go get github.com/maanasasubrahmanyam-sd/go-proto-validators/protoc-gen-govalidators

go get github.com/maanasasubrahmanyam-sd/test/protoc-gen-govalidators

For my
protoc \
-I. \
-I $GOPATH/src/ \
-I $GOPATH/src/github.com/google/protobuf/src/ \
--proto_path=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 \
--proto_path=${GOPATH}/pkg/mod/github.com/maanasasubrahmanyam-sd/go-proto-validators@v1.4.1 \
--go_out=plugins=grpc:./ \
--govalidators_out=./ \
./interfaces/test_server/*.proto

protoc \
-I. \
-I $GOPATH/src/ \
-I $GOPATH/src/github.com/google/protobuf/src/ \
--proto_path=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 \
--proto_path=${GOPATH}/pkg/mod/github.com/maanasasubrahmanyam-sd/test@v0.0.9 \
--go_out=plugins=grpc:./ \
--govalidators_out=./ \
./interfaces/test_server/*.proto

----------------------------------------------------------------------------------------
go get github.com/maanasasubrahmanyam-sd/go-proto-validators/protoc-gen-govalidators1

protoc \
-I. \
-I $GOPATH/src/ \
-I $GOPATH/src/github.com/google/protobuf/src/ \
--proto_path=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 \
--proto_path=${GOPATH}/src/github.com/maanasasubrahmanyam-sd/go-proto-validators1 \
--go_out=plugins=grpc:./ \
--govalidators_out=./ \
./interfaces/test_server/*.proto

with latest version
protoc \
-I. \
-I $GOPATH/src/ \
-I $GOPATH/pkg/mod/github.com/protocolbuffers/protobuf@v3.17.3+incompatible/src/ \
--proto_path=$GOPATH/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 \
--proto_path=$GOPATH/pkg/mod/github.com/maanasasubrahmanyam-sd/go-proto-valiadators1@v0.0.0-20210914164940-4b15a3dbf816 \
--go-grpc_out=./ \
--govalidators_out=./ \
./interfaces/test_server/*.proto
