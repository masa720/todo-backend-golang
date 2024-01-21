format:
	@find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/masa720/todo-backend-golang"
verify:
	@staticcheck ./... && go vet ./...
unit-test:
	@go test ./... -coverprofile=./test_results/cover.out && go tool cover -html=./test_results/cover.out -o ./test_results/cover.html
serve:
	@docker-compose -f build/docker-compose.yaml up