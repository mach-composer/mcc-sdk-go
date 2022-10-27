export GO_POST_PROCESS_FILE=gofmt -w

clean:
	rm -f *.go
	rm -rf docs
	rm -rf api
	rm -rf tmp

generate:
	openapi-generator generate \
		-i ../mcc-api-documentation/openapi.yaml \
		-g go \
		--enable-post-process-file \
		--model-package mccsdk/mccmodels \
		--http-user-agent mcc-sdk-go \
		--additional-properties=packageName=mccsdk
	mkdir mccsdk
	mv *.go mccsdk
	go mod tidy
