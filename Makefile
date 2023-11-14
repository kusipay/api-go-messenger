build:
	go mod download
	@for item in cmd/*; do \
		item_name=$$(basename $$item); \
		GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -tags lambda.norpc -o bin/$$item_name/bootstrap $$item/main.go; \
		zip -j bin/$$item_name.zip bin/$$item_name/bootstrap; \
	done

clean:
	rm -rf ./bin ./vendor

deploy:
	npx serverless deploy

remove:
	npx serverless remove
