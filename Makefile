V := @

OUT_DIR := ./build
MAIN_PKG := .
COVERPROFILE := covprofile

.PHONY: vendor
vendor:
	$(V)go mod tidy
	$(V)go mod vendor
	$(V)git add vendor

default: build

.PHONY: build
build: build_prover build_verifier

.PHONY: build_prover
build_prover:
	@echo BUILDING $(OUT_DIR)/prover
	$(V)go build -mod=vendor -ldflags "-s -w" -o $(OUT_DIR)/prover $(MAIN_PKG)/cmd/prover
	@echo DONE

.PHONY: build_verifier
build_verifier:
	@echo BUILDING $(OUT_DIR)/verifier
	$(V)go build -mod=vendor -ldflags "-s -w" -o $(OUT_DIR)/verifier $(MAIN_PKG)/cmd/verifier
	@echo DONE

.PHONY: test
test:
	$(V)go test -race -covermode atomic -coverprofile=covprofile ./...

.PHONY: coverage
coverage: test
	go tool cover -html=$(COVERPROFILE)

.PHONY: run
run:
	docker compose up --build

.PHONY: docker-lint
docker-lint:
	docker run --rm -v `pwd`:/app -w /app golangci/golangci-lint:v1.59.1 golangci-lint run -v
