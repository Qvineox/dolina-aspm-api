PROTO_DIR=.
GO_OUT=../gen/go
TS_OUT=../gen/ts
GW_OPENAPI_OUT=../gen

all: protoc_gen_go protoc_gen_go_gw

dump_variables:
	echo "Using protoc include path: ${PROTOC_INC}" && \
	echo "Using GRPC gateway include path: ${PROTO_GRPC_GW_INC}"

protoc_gen_go:
	cd ../api/proto && protoc -I . -I $(PROTOC_INC) -I $(PROTO_GRPC_GW_INC) \
		--go_out=$(GO_OUT) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT) \
		--go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		applications/v1/*.proto \
		components/v1/*.proto \
		exploits/v1/*.proto \
		projects/v1/*.proto \
		reports/v1/*.proto \
		defects/v1/*.proto \
		workers/v1/*.proto \
		common/v1/*.proto \
		cvss/v2/*.proto \
		cvss/v3/*.proto \
		cvss/v4/*.proto \
		cve/v1/*.proto \
		openapi.proto

# https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_openapi_output/

protoc_gen_go_gw:
	cd ../api/proto && protoc -I . -I $(PROTOC_INC) -I $(PROTO_GRPC_GW_INC) \
		--grpc-gateway_out=$(GO_OUT) \
		--grpc-gateway_opt=paths=source_relative \
		--grpc-gateway_opt use_opaque_api=true \
		--openapiv2_out=$(GW_OPENAPI_OUT) \
		--openapiv2_opt=allow_merge=true \
		applications/v1/*.proto \
		components/v1/*.proto \
		exploits/v1/*.proto \
		projects/v1/*.proto \
		reports/v1/*.proto \
		defects/v1/*.proto \
		workers/v1/*.proto \
		common/v1/*.proto \
		cvss/v2/*.proto \
		cvss/v3/*.proto \
		cvss/v4/*.proto \
		cve/v1/*.proto \
		openapi.proto

protoc_gen_ts:
	cd ../api/proto && protoc -I . -I $(PROTO_DIR) -I $(PROTOC_INC) -I $(PROTO_GRPC_GW_INC) \
		--ts_proto_out=$(TS_OUT) \
		--ts_proto_opt=esModuleInterop=true,forceLong=string,outputServices=grpc-js \
		applications/v1/*.proto \
		cyclonedx/v17/*.proto \
		cyclonedx/v16/*.proto \
		components/v1/*.proto \
		exploits/v1/*.proto \
		projects/v1/*.proto \
		reports/v1/*.proto \
		defects/v1/*.proto \
		workers/v1/*.proto \
		upload/v1/*.proto \
		common/v1/*.proto \
		sbom/v1/*.proto \
		cve/v1/*.proto \
		openapi.proto

clean:
	del /S /Q $(GO_OUT)\* || echo ""
	del /S /Q $(TS_OUT)\* || echo ""

.PHONY: all go gateway ts clean
