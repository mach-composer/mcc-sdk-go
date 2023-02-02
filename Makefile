export GO_POST_PROCESS_FILE=gofmt -w

clean:
	rm -f *.go
	rm -rf docs
	rm -rf api
	rm -rf tmp

generate:
	rm -rf mccsdk
	openapi-generator generate \
		-i ../mcc-api-documentation/openapi-public.yaml \
		-g go \
		--enable-post-process-file \
		--model-package mccsdk/mccmodels \
		--http-user-agent mcc-sdk-go \
		--additional-properties=packageName=mccsdk,generateInterfaces=true
	mkdir mccsdk
	mv *.go mccsdk
	go mod tidy
