.PHONY: assets
assets: $(GOPATH)/bin/go-assets-builder
		@echo run go-assets-builder...
		@$(GOPATH)/bin/go-assets-builder --output=config/cyml.go -p=config env/*.yml
		@echo complete [go-assets-builder]