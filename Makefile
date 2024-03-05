SHELL := /bin/bash

check-program = $(foreach exec,$(1),$(if $(shell PATH="$(PATH)" which $(exec)),,$(error "Missing deps: no '$(exec)' in PATH")))

OS := $(shell uname)

# Get the temporary directory of the system
TEMPDIR := $(shell dirname $(shell mktemp -u))

APP_NAME ?= gnoboard

# Define the directory that contains the current Makefile
make_dir := $(realpath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
cache_dir := $(make_dir)/.cache
gnoboard_dir := $(make_dir)/examples/react-native/gnoboard

# Argument Defaults
APP_OUTPUT_DIR ?= $(make_dir)/examples/react-native/$(APP_NAME)
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
ifeq ($(OS),Darwin)
	@echo "generate iOS framework"
	cd $(APP_OUTPUT_DIR); $(MAKE) node_modules
	cd $(APP_OUTPUT_DIR); $(MAKE) ios/$(APP_NAME).xcworkspace TEMPLATE_PROJECT=$(APP_NAME)
endif

# Build Android aar & jar
build.android: generate $(gnocore_aar) $(gnocore_jar)
	cd $(APP_OUTPUT_DIR); $(MAKE) node_modules

# Generate API from protofiles
generate: api.generate

# Clean and generate
regenerate: api.clean api.generate

# Clean all generated files
clean: bind.clean api.clean

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
	$(call check-program, yarn)
	cd api; yarn

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

yarn_basic_dependencies := @bufbuild/protobuf @connectrpc/connect @connectrpc/connect-web react-native-polyfill-globals react-native-url-polyfill web-streams-polyfill@3.2.1 react-native-get-random-values text-encoding base-64 react-native-fetch-api
yarn_basic_dev_dependencies = @tsconfig/react-native babel-plugin-module-resolver
OUTPUT_DIR := $(make_dir)/examples/react-native

new-app:
ifndef APP_NAME
	$(error APP_NAME is undefined. Please set APP_NAME to the name of your app)
endif
	$(MAKE) new-react-native-app OUTPUT_DIR=$(make_dir)/examples/react-native
	$(MAKE) add-app-json-entry APP_NAME=$(APP_NAME) APP_OUTPUT_DIR=$(make_dir)/examples/react-native
	$(MAKE) copy-js-files APP_NAME=$(APP_NAME) APP_OUTPUT_DIR=$(make_dir)/examples/react-native
	$(MAKE) new-app-build-android APP_NAME=$(APP_NAME) APP_OUTPUT_DIR=$(make_dir)/examples/react-native
ifeq ($(OS),Darwin)
	$(MAKE) new-app-build-ios APP_NAME=$(APP_NAME) APP_OUTPUT_DIR=$(make_dir)/examples/react-native
	$(MAKE) copy-ios-project-pbxproj
endif

# creates a new react native app using Expo script. Also creates ios and android folders
new-react-native-app:
	$(call check-program, yarn)
	@mkdir -p $(OUTPUT_DIR)
	@echo "creating a new gno awesome project at: $(OUTPUT_DIR)"
	cd $(OUTPUT_DIR) && yarn create expo $(APP_NAME) --template expo-template-blank-typescript
	@echo "Creating ios and android folders"
	cd $(OUTPUT_DIR)/$(APP_NAME) && yarn expo prebuild
	@echo "Installing yarn dependencies"
	cd $(OUTPUT_DIR)/$(APP_NAME) && yarn add ${yarn_basic_dependencies} && yarn add -D ${yarn_basic_dev_dependencies}
	@echo "Building GnoCore.xcframework for the new app"
	$(MAKE) build.ios APP_NAME=$(APP_NAME) APP_OUTPUT_DIR=$(OUTPUT_DIR)/$(APP_NAME)

# copy js files from gnoboard to the new app
copy-js-files:
	$(call check-program, jq)
	@echo "Copying js files"
	@mkdir -p $(OUTPUT_DIR)/$(APP_NAME)/src/grpc
	@mkdir -p $(OUTPUT_DIR)/$(APP_NAME)/src/hooks
	@cp -r $(gnoboard_dir)/src/grpc $(OUTPUT_DIR)/$(APP_NAME)/src
	@cp -r $(gnoboard_dir)/src/hooks $(OUTPUT_DIR)/$(APP_NAME)/src
	@cp -r $(gnoboard_dir)/src/native_modules $(OUTPUT_DIR)/$(APP_NAME)/src
	@cp -r $(gnoboard_dir)/Makefile $(OUTPUT_DIR)/$(APP_NAME)
	@cp -r $(gnoboard_dir)/android/.gitignore $(OUTPUT_DIR)/$(APP_NAME)/android
	@cp -r $(gnoboard_dir)/ios/.gitignore $(OUTPUT_DIR)/$(APP_NAME)/ios
	@cp $(make_dir)/templates/tsconfig.json $(OUTPUT_DIR)/$(APP_NAME)/tsconfig.json
	@cp $(make_dir)/templates/babel.config.js $(OUTPUT_DIR)/$(APP_NAME)/babel.config.js
	@cp $(make_dir)/templates/metro.config.js $(OUTPUT_DIR)/$(APP_NAME)/metro.config.js
	@cp $(make_dir)/templates/App.tsx $(OUTPUT_DIR)/$(APP_NAME)/App.tsx

# build GnoCore.xcframework for the new app
new-app-build-ios:
	@echo "Copying ios files"
	@mkdir -p $(OUTPUT_DIR)/$(APP_NAME)/ios/$(APP_NAME)/Sources
	@cp -r $(gnoboard_dir)/ios/gnoboard/Sources $(OUTPUT_DIR)/$(APP_NAME)/ios/$(APP_NAME)/
	@cp $(gnoboard_dir)/ios/gnoboard/gnoboard-Bridging-Header.h $(OUTPUT_DIR)/$(APP_NAME)/ios/$(APP_NAME)/$(APP_NAME)-Bridging-Header.h
	@cp -r $(gnoboard_dir)/ios/Sources $(OUTPUT_DIR)/$(APP_NAME)/ios/
	@cd $(OUTPUT_DIR)/$(APP_NAME) && $(MAKE) ios/$(APP_NAME).xcworkspace TEMPLATE_PROJECT=$(APP_NAME)

# build the Android project files for the new app
new-app-build-android:
	@echo "Copying Android project files"
	@$(eval MAIN_APP_PATH := $(shell find $(OUTPUT_DIR)/$(APP_NAME)/android/app/src/main/java -name 'MainApplication.kt')) # ./android/app/src/main/java/com/anonymous/<APP_NAME>/MainApplication.kt
	@$(eval MAIN_APP_FILE := $(shell basename $(MAIN_APP_PATH))) # MainApplication.kt
	@$(eval MAIN_APP_DIR := $(shell dirname $(MAIN_APP_PATH))) # ./android/app/src/main/java/com/anonymous/<APP_NAME>
	@$(eval PACKAGE_PREFIX := $(shell sed -n 's/package \(.*\)\.$(APP_NAME)/\1/'p $(MAIN_APP_DIR)/$(MAIN_APP_FILE))) # e.g. com.anonymous
	@cp -r $(gnoboard_dir)/android/app/src/main/java/land/gno/gobridge $(MAIN_APP_DIR)/.. # copy gobridge directory
	@cp -r $(gnoboard_dir)/android/app/src/main/java/land/gno/rootdir $(MAIN_APP_DIR)/.. # copy rootdir directory
	@perl -pi -e "s/land\.gno/$(PACKAGE_PREFIX)/" $(MAIN_APP_DIR)/../gobridge/* # in gobridge, replace land.gno by PACKAGE_PREFIX (e.g. com.anonymous)
	@perl -pi -e "s/land\.gno/$(PACKAGE_PREFIX)/" $(MAIN_APP_DIR)/../rootdir/* # in rootdir, replace land.gno by PACKAGE_PREFIX (e.g. com.anonymous)
	@perl -pi -e '/^package ./ and $$_.="\nimport '"$(PACKAGE_PREFIX)"'.gobridge.GoBridgePackage"' $(MAIN_APP_DIR)/$(MAIN_APP_FILE) # add the right import path for gobridge (e.g. import com.anonymous.gobridge.GoBridgePackage)
	@perl -pi -e '/^package ./ and $$_.="\nimport '"$(PACKAGE_PREFIX)"'.rootdir.RootDirPackage"' $(MAIN_APP_DIR)/$(MAIN_APP_FILE) # add the right import path for rootdir (e.g. import com.anonymous.rootdir.RootDirPackage)
	@perl -pi -e 's/return PackageList\(this\)\.packages/return PackageList\(this\)\.packages.apply \{\n\t\t\t\tadd(RootDirPackage())\n\t\t\t\tadd(GoBridgePackage())\n\t\t\t\}/' $(MAIN_APP_DIR)/$(MAIN_APP_FILE) # replace the default package list by one adding gobridge and rootdir
	@perl -pi -e '/^def projectRoot/ and $$_.="def frameworkDir = \"\$$\{rootDir\.getAbsoluteFile\(\)\.getParentFile\(\)\.getParentFile\(\)\.getParentFile\(\)\.getParentFile\(\)\.getAbsolutePath\(\)\}/framework\"\n"' $(OUTPUT_DIR)/$(APP_NAME)/android/app/build.gradle # add the projectRoot variable in build.gradle
	@perl -pi -e '/^dependencies/ and $$_.="\timplementation fileTree(dir: \"\$$\{frameworkDir\}/android\", include: \[\"\*\.jar\", \"\*\.aar\"\]\)\n"' $(OUTPUT_DIR)/$(APP_NAME)/android/app/build.gradle # add the framework dependency in build.gradle
	@cd $(OUTPUT_DIR)/$(APP_NAME) && $(MAKE) node_modules TEMPLATE_PROJECT=$(APP_NAME)

JSON_FILE := $(OUTPUT_DIR)/$(APP_NAME)/app.json
# add tsconfigPaths entry to app.json
add-app-json-entry:
	@echo "Adding tsconfigPaths entry to app.json"
	jq '.expo.experiments = {"tsconfigPaths": true}'  $(JSON_FILE) > $(JSON_FILE).tmp && mv $(JSON_FILE).tmp  $(JSON_FILE)

# copy ios project.pbxproj from gnoboard to the new app and replace 'gnoboard' with the new app name
copy-ios-project-pbxproj:
	@echo "Copying ios project.pbxproj"
	@cp $(gnoboard_dir)/ios/gnoboard.xcodeproj/project.pbxproj $(OUTPUT_DIR)/$(APP_NAME)/ios/$(APP_NAME).xcodeproj/project.pbxproj
	@cd $(OUTPUT_DIR)/$(APP_NAME)/ios/$(APP_NAME).xcodeproj
	@sed -i.pbxproj 's/gnoboard/$(APP_NAME)/g' $(OUTPUT_DIR)/$(APP_NAME)/ios/$(APP_NAME).xcodeproj/project.pbxproj

.PHONY: new-app new-react-native-app copy-js-files new-app-build-ios add-app-json-entry copy-ios-project-pbxproj
