ANDROID_DEVICE ?=

check-program = $(foreach exec,$(1),$(if $(shell PATH="$(PATH)" which $(exec)),,$(error "Missing deps: no '$(exec)' in PATH")))

PROJECT_DIR := $(shell go list -m -f '{{.Dir}}')
# This example links the local @gnolang/gnonative module. The gomobile frameworks and the built JS
# both come from there, so we build them in place rather than copying into this app's dirs.
EXPO_MODULE_DIR := $(PROJECT_DIR)/expo

# gomobile framework outputs, copied into the expo module by its own build.ios/build.android targets.
IOS_FRAMEWORK_DST := $(EXPO_MODULE_DIR)/ios/Frameworks/GnoCore.xcframework
ANDROID_FRAMEWORK_DST := $(EXPO_MODULE_DIR)/android/libs/gnocore.aar

# go files and dependencies
go_framework_files := $(shell find $(PROJECT_DIR)/framework -iname '*.go')
go_service_files := $(shell find $(PROJECT_DIR)/service -iname '*.go')
go_deps := $(PROJECT_DIR)/go.mod $(PROJECT_DIR)/go.sum $(go_framework_files) $(go_service_files)

ts_check: node_modules
	npm run ts:check
.PHONY: ts_check

# - Node: Handle node_modules

node_modules: package.json
	$(call check-program, npm)
	(npm install && touch $@) || true
.PHONY: node_modules

# Build the local @gnolang/gnonative module (JS) so autolinking resolves it.
expo_module:
	$(call check-program, npm)
	cd $(EXPO_MODULE_DIR) && npm install && npm run prepare
.PHONY: expo_module

######### IOS #########

ios: node_modules ts_check $(IOS_FRAMEWORK_DST) expo_module
	$(call check-program, npx)
	npx expo run:ios
.PHONY: ios

# Builds the gomobile xcframework and copies it into the expo module (build.ios also runs build.api).
$(IOS_FRAMEWORK_DST): $(go_deps)
	cd $(EXPO_MODULE_DIR) && $(MAKE) build.ios

######### ANDROID #########

android: node_modules ts_check $(ANDROID_FRAMEWORK_DST) expo_module
	$(call check-program, npx)
	npx expo run:android
.PHONY: android

# Builds the gomobile aar and copies it into the expo module (build.android also runs build.api).
$(ANDROID_FRAMEWORK_DST): $(go_deps)
	cd $(EXPO_MODULE_DIR) && $(MAKE) build.android

# - Helpers

android.reverse_tcp:
	$(call check-program, adb)
	$(if $(ANDROID_DEVICE),,$(eval ANDROID_DEVICE = $(shell adb devices | tail +2 | head -1 | cut -f 1)))
	@if [ -z "$(ANDROID_DEVICE)" ]; then \
	  >&2 echo "ERROR: no Android device found"; exit 1; \
	fi
	adb -s $(ANDROID_DEVICE) reverse tcp:8081 tcp:8081 # metro
	adb -s $(ANDROID_DEVICE) reverse tcp:36657 tcp:36657 # gnodev
.PHONY: android.reverse_tcp
