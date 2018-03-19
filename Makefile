APP=hello-go-gl
VERSION=1.0

GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get

PROJECT_ROOT=github.com/$(APP)

build: build_osx

build_osx:
		env GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/osx/$(APP) $(PROJECT_ROOT)/cmd

clean:
		$(GOCLEAN)
		rm -fr bin/

deps:
		$(GOGET) github.com/go-gl/gl/v4.1-core/gl
		$(GOGET) github.com/go-gl/glfw/v3.2/glfw
