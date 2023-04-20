.PHONY: test mock
flyway-migrate:
	./script/flyway.sh
refresh-locale:
	curl -X GET "https://api.localizely.com/v1/projects/2f76b4f0-4091-49d0-a1b0-c126b38510b0/files/download?export_empty_as=empty&include_tags=jojoy-backend&lang_codes=en&type=json" -H "accept: */*" -H "X-Api-Token: 64c35f96553c4a4b9e78c96a8d4a38b0cdfc6507207d4aea82dbddbf0498042d" -o config/resource/i18n/en.json
	curl -X GET "https://api.localizely.com/v1/projects/2f76b4f0-4091-49d0-a1b0-c126b38510b0/files/download?export_empty_as=empty&include_tags=jojoy-backend&lang_codes=es&type=json" -H "accept: */*" -H "X-Api-Token: 64c35f96553c4a4b9e78c96a8d4a38b0cdfc6507207d4aea82dbddbf0498042d" -o config/resource/i18n/es.json
	curl -X GET "https://api.localizely.com/v1/projects/2f76b4f0-4091-49d0-a1b0-c126b38510b0/files/download?export_empty_as=empty&include_tags=jojoy-backend&lang_codes=id&type=json" -H "accept: */*" -H "X-Api-Token: 64c35f96553c4a4b9e78c96a8d4a38b0cdfc6507207d4aea82dbddbf0498042d" -o config/resource/i18n/id.json
	curl -X GET "https://api.localizely.com/v1/projects/2f76b4f0-4091-49d0-a1b0-c126b38510b0/files/download?export_empty_as=empty&include_tags=jojoy-backend&lang_codes=pt&type=json" -H "accept: */*" -H "X-Api-Token: 64c35f96553c4a4b9e78c96a8d4a38b0cdfc6507207d4aea82dbddbf0498042d" -o config/resource/i18n/pt.json
	curl -X GET "https://api.localizely.com/v1/projects/2f76b4f0-4091-49d0-a1b0-c126b38510b0/files/download?export_empty_as=empty&include_tags=jojoy-backend&lang_codes=zh&type=json" -H "accept: */*" -H "X-Api-Token: 64c35f96553c4a4b9e78c96a8d4a38b0cdfc6507207d4aea82dbddbf0498042d" -o config/resource/i18n/zh.json
	go-localize -input config/resource/i18n -output internal/generated/localizations
generate_local:
	export GOPRIVATE=gitlab.wuren.com
	export MYSQL='root:@tcp(127.0.0.1:3306)/monster_base_backend' \
	&& ./script/generate.sh
generate:
	export GOPRIVATE=gitlab.wuren.com
	./script/generate.sh
lint:
	golint $$(go list ./... | grep -v /test | grep -v /internal/generated/ | grep -v pojo)
	# /opt/homebrew/Cellar/go/1.17/bin/golint $$(go list ./... | grep -v /test | grep -v /internal/generated/)
mock:
	./script/mock.sh
test: mock
	go test $$(sh -c "go list ./... | grep -v /test | grep -v /internal/generated/")
serve:
	go run -tags dynamic ./cmd/jojoy-cloud-game/main.go
clean:
	rm -rf test/generated