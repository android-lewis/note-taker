run:
	@go run ./cmd/note-taker/
build:
	@go build ./cmd/note-taker/
test:
	@go test -v ./pkg/...
benchmark:
	@go test ./pkg/... -bench=. -benchtime=10s
install:
	@go install ./cmd/note-taker/