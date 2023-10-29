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
		--additional-properties=packageName=mccsdk,generateInterfaces=true \
	 	--git-repo-id mcc-sdk-go \
		--git-user-id mach-composer \
		--api-name-suffix Api
	mkdir mccsdk
	mv *.go mccsdk
	rm -rf test
	go mod tidy
