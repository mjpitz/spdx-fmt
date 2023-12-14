define HELP_TEXT
Welcome!

Targets:
  help		provides help text
  legal		prepends legal header to source code
  deps		install dependencies
  spdx		regenerate the spdx files

endef
export HELP_TEXT

help:
	@echo "$$HELP_TEXT"

legal:
	@addlicense -f ./legal/header.txt -skip tmpl .

deps:
	@echo "installing addlicense" && go install github.com/google/addlicense@latest
	@echo "installing bom" && go install sigs.k8s.io/bom/cmd/bom@latest

spdx:
	@bom generate --config .sbom.yaml --format json --output spdx.json .
	@go run main.go -input spdx.json -output SPDX.md
