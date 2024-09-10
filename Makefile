SHELL := /bin/bash

check-program = $(foreach exec,$(1),$(if $(shell PATH="$(PATH)" which $(exec)),,$(error "Missing deps: no '$(exec)' in PATH")))

OS := $(shell uname)

# Get the temporary directory of the system
TEMPDIR := $(shell dirname $(shell mktemp -u))

APP_NAME ?= gnoboard

# Define the directory that contains the current Makefile
make_dir := $(realpath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
cache_dir := $(make_dir)/.cache
gnoboard_dir := $(make_dir)/examples/js/react-native/gnoboard
expo_dir := $(make_dir)/expo

# Argument Defaults
APP_OUTPUT_DIR ?= $(make_dir)/examples/js/react-native/$(APP_NAME)
IOS_OUTPUT_FRAMEWORK_DIR ?= framework/ios
ANDROID_OUTPUT_LIBS_DIR ?= framework/android
GO_BIND_BIN_DIR ?= $(cache_dir)/bind

# IOS definitions
gnocore_xcframework := $(IOS_OUTPUT_FRAMEWORK_DIR)/GnoCore.xcframework

# Android definitions
gnocore_aar := $(ANDROID_OUTPUT_LIBS_DIR)/gnocore.aar
gnocore_jar := $(ANDROID_OUTPUT_LIBS_DIR)/gnocore-sources.jar

# Utility definitions
gomobile := $(GO_BIND_BIN_DIR)/gomobile
gobind := $(GO_BIND_BIN_DIR)/gobind

# go files and dependencies
go_framework_files := $(shell find framework -iname '*.go')
go_service_files := $(shell find service -iname '*.go')
go_files := $(go_framework_files) $(go_service_files)
go_deps := go.mod go.sum $(go_files)

# rewrite shell path
# this is mostly for gomobile to have the correct gobind in his path
PATH := $(GO_BIND_BIN_DIR):$(PATH)

# * Main commands

# `all` and `build` command builds everything (generate, build.ios, build.android)
all build: generate build.ios build.android

# Build iOS framework
build.ios: generate framework.ios
ifeq ($(OS),Darwin)
	@echo "generate iOS framework"
	cd $(APP_OUTPUT_DIR); $(MAKE) node_modules
	cd $(APP_OUTPUT_DIR); $(MAKE) ios/$(APP_NAME).xcworkspace TEMPLATE_PROJECT=$(APP_NAME)
endif

# Build Android aar & jar
build.android: generate framework.android
	cd $(APP_OUTPUT_DIR); $(MAKE) node_modules

# Generate API from protofiles
generate: api.generate

# Clean and generate
regenerate: api.clean api.generate

# Clean all generated files
clean: bind.clean

# Force clean (clean and remove node_modules)
fclean:
	cd $(APP_OUTPUT_DIR); ls; rm -rf node_modules
	rm -rf $(cache_dir)

.PHONY: generate regenerate build.ios build.android clean fclean

# - API : Handle API generation and cleaning

api.generate: _api.generate.protocol _api.generate.modules
api.clean: _api.clean.protocol _api.clean.modules

# - API - rpc

protos_src := $(wildcard service/rpc/*.proto)
gen_src := $(protos_src) Makefile buf.gen.yaml $(wildcard service/gnonativetypes/*.go)
gen_sum := gen.sum

_api.generate.protocol: $(gen_sum)
_api.clean.protocol:
	rm -f api/gen/go/*.pb.go
	rm -f api/gen/go/_goconnect/*.connect.go
	rm -f api/gen/es/*.{ts,js}
	rm -f $(gen_sum)

$(gen_sum): $(gen_src)
	$(call check-program, shasum buf)
	@shasum $(gen_src) | sort -k 2 > $(gen_sum).tmp
	@diff -q $(gen_sum).tmp $(gen_sum) || ( \
		cd misc/genproto && go run . && cd ../.. ; \
		buf generate api; \
		shasum $(gen_src) | sort -k 2 > $(gen_sum).tmp; \
		mv $(gen_sum).tmp $(gen_sum); \
		go mod tidy \
	)

_api.generate.modules:
	$(call check-program, npm)
	cd api; npm install

_api.clean.modules:
	cd api; rm -fr node_modules

.PHONY: api.generate _api.generate.protocol _api.generate.modules _api.clean.protocol _api.clean.modules

# - Bind : Handle gomobile bind

# - Bind - initialization
bind_init_files := $(TEMPDIR)/.tool-versions $(gobind) $(gomobile)

$(gobind): go.sum go.mod
	@mkdir -p $(dir $@)
	go build -o $@ golang.org/x/mobile/cmd/gobind && chmod +x $@

$(gomobile): $(gobind) go.sum go.mod
	@mkdir -p $(dir $@)
	go build -o $@ golang.org/x/mobile/cmd/gomobile && chmod +x $@
	$(gomobile) init || (rm -rf $@ && exit 1) # in case of failure, remove gomobile so we can init again

$(TEMPDIR)/.tool-versions: .tool-versions
	@echo "> copying current '.tool-versions' in '$(TEMPDIR)' folder in order to make asdf works"
	@echo "> this hack is needed in order for gomobile (who is building from '$(TEMPDIR)') bind to use the correct javac and go version"
	@cp -v $< $@

.PHONY: bind.init

# - Bind - ios framework

framework.ios: $(gnocore_xcframework)
.PHONY: framework.ios

$(gnocore_xcframework): $(bind_init_files) $(go_deps)
ifeq ($(OS),Darwin)
	@mkdir -p $(dir $@)
	# need to use `nowatchdog` tags, see https://github.com/libp2p/go-libp2p-connmgr/issues/98
	$(gomobile) bind -v \
		-cache $(cache_dir)/ios-gnonative \
		-tags 'nowatchdog' -prefix=Gno \
		-o $@ -target ios ./framework/service
endif
_bind.clean.ios:
	rm -rf $(gnocore_xcframework)

# - Bind - android aar and jar

framework.android: $(gnocore_aar) $(gnocore_jar)
.PHONY: framework.android

$(gnocore_aar): $(bind_init_files) $(go_deps)
	@mkdir -p $(dir $@) .cache/bind/android
	$(gomobile) bind -v \
		-cache $(cache_dir)/android-gnonative \
		-javapkg=gnolang.gno \
		-o $@ -target android -androidapi 21 ./framework/service
_bind.clean.android:
	rm -rf $(gnocore_jar) $(gnocore_aar)


# - Bind - cleaning

bind.clean: _bind.clean.ios _bind.clean.android
	rm -rf $(bind_init_files)

# - asdf

asdf.add_plugins:
	$(call check-program, asdf)
	@echo "Installing asdf plugins..."
	@set -e; \
	for PLUGIN in $$(cut -d' ' -f1 .tool-versions | grep "^[^\#]"); do \
		asdf plugin add $$PLUGIN || [ $$?==2 ] || exit 1; \
	done

asdf.install_tools: asdf.add_plugins
	$(call check-program, asdf)
	@echo "Installing asdf tools..."
	@asdf install

########################################
# Script to create a new app

npm_basic_dependencies := @gnolang/gnonative
OUTPUT_DIR := $(make_dir)/examples/js/react-native

new-app:
ifndef APP_NAME
	$(error APP_NAME is undefined. Please set APP_NAME to the name of your app)
endif
	$(call check-program, npm)
	npm config set @buf:registry https://buf.build/gen/npm/v1/
	$(MAKE) new-react-native-app OUTPUT_DIR=$(make_dir)/examples/js/react-native
	$(MAKE) copy-js-files APP_NAME=$(APP_NAME) APP_OUTPUT_DIR=$(make_dir)/examples/js/react-native

# creates a new react native app using Expo script. Also creates ios and android folders
new-react-native-app:
	$(call check-program, npm)
	$(call check-program, npx)
	@mkdir -p $(OUTPUT_DIR)
	@echo "creating a new gno awesome project at: $(OUTPUT_DIR)"
	cd $(OUTPUT_DIR) && npx create-expo-app@latest $(APP_NAME) --template expo-template-blank-typescript
	@echo "Creating ios and android folders"
	cd $(OUTPUT_DIR)/$(APP_NAME) && npx expo prebuild
	@echo "Installing npm dependencies"
	cd $(OUTPUT_DIR)/$(APP_NAME) && npm install ${npm_basic_dependencies}
	@echo "Building GnoCore.xcframework for the new app"
	$(MAKE) build.ios APP_NAME=$(APP_NAME) APP_OUTPUT_DIR=$(OUTPUT_DIR)/$(APP_NAME)

# copy js files from gnoboard to the new app
copy-js-files:
	@echo "Copying js files"
	@cp $(expo_dir)/example/App.tsx $(OUTPUT_DIR)/$(APP_NAME)/App.tsx

.PHONY: new-app configure-npm new-react-native-app copy-js-files
