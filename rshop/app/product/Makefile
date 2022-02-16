APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$a)
API_PROTO_FILES=$(shell cd ../../api/$(APP_RELATIVE_PATH) && find . -name *.proto)

INTERNAL_PROTO_FILES=$(shell find . -name '*.proto')

grpc:
	#echo $(APP_RELATIVE_PATH)
	cd ../../api/$(APP_RELATIVE_PATH) && protoc --proto_path=. \
	--proto_path=../../third_party \
	--go_out=paths=source_relative:. \
    --go-grpc_out=paths=source_relative:. \
   	$(API_PROTO_FILES)

.PHONY: proto
proto:
	protoc --proto_path=. \
           --proto_path=../../third_party \
           --go_out=paths=source_relative:. \
           $(INTERNAL_PROTO_FILES)

run:
	go run service/cmd/server/main.go service/cmd/server/wire_gen.go -conf service/configs/config.yaml