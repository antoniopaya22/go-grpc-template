
# Proto parameters
PROTO_FILES_PATH=proto
PROTO_OUT=proto

gen-proto:
	protoc -I $(PROTO_FILES_PATH) --go_out=plugins=grpc:$(PROTO_OUT) $(PROTO_FILES_PATH)/*.proto

clean-proto:
	rm -f $(PROTO_OUT)/*.pb.go