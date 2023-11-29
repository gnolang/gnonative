SHELL := /bin/bash

check-program = $(foreach exec,$(1),$(if $(shell PATH="$(PATH)" which $(exec)),,$(error "Missing deps: no '$(exec)' in PATH")))

# Get the temporary directory of the system
TEMPDIR := $(shell dirname $(shell mktemp -u))

TEMPLATE_PROJECT := gnoboard

# Define the directory that contains the current Makefile
make_dir := $(realpath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
cache_dir := $(make_dir)/.cache
react_native_dir := $(make_dir)/$(TEMPLATE_PROJECT)

# Argument Defaults
IOS_OUTPUT_FRAMEWORK_DIR ?= $(react_native_dir)/ios/Frameworks
ANDROID_OUTPUT_LIBS_DIR ?= $(react_native_dir)/android/libs
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
go_files := $(shell find . -iname '*.go')
go_deps := go.mod go.sum $(go_files)

# rewrite shell path
# this is mostly for gomobile to have the correct gobind in his path
PATH := $(GO_BIND_BIN_DIR):$(PATH)

# * Main commands

# `all` and `build` command builds everything (generate, build.ios, build.android)
all build: generate build.ios build.android

# Build iOS framework
build.ios: generate $(gnocore_xcframework)
	cd $(react_native_dir); $(MAKE) node_modules
	cd $(react_native_dir); $(MAKE) ios/$(TEMPLATE_PROJECT).xcworkspace TEMPLATE_PROJECT=$(TEMPLATE_PROJECT)

# Build Android aar & jar
build.android: generate $(gnocore_aar) $(gnocore_jar)
	cd $(react_native_dir); $(MAKE) node_modules

# Generate API from protofiles
generate: api.generate

# Clean and generate
regenerate: api.clean api.generate

# Clean all generated files
clean: bind.clean api.clean

# Force clean (clean and remove node_modules)
fclean: clean
	rm -rf node_modules
	rm -rf $(cache_dir)

.PHONY: generate regenerate build.ios build.android clean fclean

# - API : Handle API generation and cleaning

api.generate: _api.generate.protocol
api.clean: _api.clean.protocol

# - API - rpc

protos_src := $(wildcard service/rpc/*.proto)
gen_src := $(protos_src) Makefile buf.gen.yaml $(wildcard service/gnomobiletypes/*.go)
gen_sum := gen.sum

_api.generate.protocol: $(gen_sum)
_api.clean.protocol:
	rm -f service/rpc/*.pb.go
	rm -f service/rpc/rpcconnect/*.connect.go
	rm -f $(TEMPLATE_PROJECT)/src/api/*.{ts,js}
	rm -f $(gen_sum)

$(gen_sum): $(gen_src)
	$(call check-program, shasum buf)
	@shasum $(gen_src) | sort -k 2 > $(gen_sum).tmp
	@diff -q $(gen_sum).tmp $(gen_sum) || ( \
		cd misc/genproto && go run . && cd ../.. ; \
		buf generate service/rpc; \
		shasum $(gen_src) | sort -k 2 > $(gen_sum).tmp; \
		mv $(gen_sum).tmp $(gen_sum); \
		go mod tidy \
	)

.PHONY: api.generate _api.generate.protocol _api.clean.protocol

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

$(gnocore_xcframework): $(bind_init_files) $(go_deps)
	@mkdir -p $(dir $@)
	# need to use `nowatchdog` tags, see https://github.com/libp2p/go-libp2p-connmgr/issues/98
	$(gomobile) bind -v \
		-cache $(cache_dir)/ios-gmomobile \
		-tags 'nowatchdog' -prefix=Gno \
		-o $@ -target ios ./framework
_bind.clean.ios:
	rm -rf $(gnocore_xcframework)

# - Bind - android aar and jar

$(gnocore_aar): $(bind_init_files) $(go_deps)
	@mkdir -p $(dir $@) .cache/bind/android
	$(gomobile) bind -v \
		-cache $(cache_dir)/android-gmomobile \
		-javapkg=gnolang.gno \
		-o $@ -target android -androidapi 21 ./framework
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

# Script to create a new app

YARN_BASIC_DEPENDENCIES := @bufbuild/protobuf @connectrpc/connect @connectrpc/connect-web react-native-polyfill-globals react-native-url-polyfill web-streams-polyfill react-native-get-random-values text-encoding base-64 react-native-fetch-api

new-app:
ifndef APP_NAME
	$(error APP_NAME is undefined. Please set APP_NAME to the name of your app)
endif
	@mkdir -p ./examples/
	@echo "creating a new gno awesome project"
	cd examples && yarn create expo $(APP_NAME) --template expo-template-blank-typescript
	@echo "Creating ios and android folders"
	cd examples/$(APP_NAME) && yarn expo prebuild
	@echo "Installing yarn dependencies"
	cd examples/$(APP_NAME) && yarn add ${YARN_BASIC_DEPENDENCIES}

# create folders (api, grpc, hooks)
	@mkdir -p ./examples/$(APP_NAME)/src/api
	@mkdir -p ./examples/$(APP_NAME)/src/grpc
	@mkdir -p ./examples/$(APP_NAME)/src/hooks
# copy the essential files
	@cp -r $(react_native_dir)/src/api ./examples/$(APP_NAME)/src
	@cp -r ./$(TEMPLATE_PROJECT)/src/grpc ./examples/$(APP_NAME)/src
	@cp -r ./$(TEMPLATE_PROJECT)/src/hooks ./examples/$(APP_NAME)/src
	@cp -r ./$(TEMPLATE_PROJECT)/src/native_modules ./examples/$(APP_NAME)/src
	@cp -r ./$(TEMPLATE_PROJECT)/Makefile ./examples/$(APP_NAME)
	@cp -r ./$(TEMPLATE_PROJECT)/android/.gitignore ./examples/$(APP_NAME)/android
	@cp -r ./$(TEMPLATE_PROJECT)/ios/.gitignore ./examples/$(APP_NAME)/ios
	@cp ./templates/tsconfig.json ./examples/$(APP_NAME)/tsconfig.json
	@cp ./templates/App.tsx ./examples/$(APP_NAME)/App.tsx
	$(MAKE) add-app-json-entry
# copy ios Sources
	@mkdir -p ./examples/$(APP_NAME)/ios/$(APP_NAME)/Sources
	@cp -r $(react_native_dir)/ios/$(TEMPLATE_PROJECT)/Sources ./examples/$(APP_NAME)/ios/$(APP_NAME)/
	@cp $(react_native_dir)/ios/$(TEMPLATE_PROJECT)/$(TEMPLATE_PROJECT)-Bridging-Header.h ./examples/$(APP_NAME)/ios/$(APP_NAME)/$(APP_NAME)-Bridging-Header.h
	@cp -r $(react_native_dir)/ios/Sources ./examples/$(APP_NAME)/ios/

# build GnoCore.xcframework for the new app
	@echo "Building GnoCore.xcframework for the new app"
	$(MAKE) build.ios TEMPLATE_PROJECT=$(APP_NAME) react_native_dir=./examples/$(APP_NAME)

# copy ios project.pbxproj from gnoboard
	$(MAKE) copy-ios-project-pbxproj


JSON_FILE := $(make_dir)/examples/$(APP_NAME)/app.json
# add tsconfigPaths entry to app.json
add-app-json-entry:
	@echo "Adding tsconfigPaths entry to app.json"
	jq '.expo.experiments = {"tsconfigPaths": true}'  $(JSON_FILE) > $(JSON_FILE).tmp && mv $(JSON_FILE).tmp  $(JSON_FILE)

# copy ios project.pbxproj from gnoboard to the new app and replace 'gnoboard' with the new app name
copy-ios-project-pbxproj:
	@echo "Copying ios project.pbxproj"
	@cp $(react_native_dir)/ios/$(TEMPLATE_PROJECT).xcodeproj/project.pbxproj ./examples/$(APP_NAME)/ios/$(APP_NAME).xcodeproj/project.pbxproj
	@cd examples/$(APP_NAME)/ios/$(APP_NAME).xcodeproj
	@sed -i.pbxproj 's/gnoboard/$(APP_NAME)/g' examples/$(APP_NAME)/ios/$(APP_NAME).xcodeproj/project.pbxproj

.PHONY: new-app copy-ios-project-pbxproj add-app-json-entry

