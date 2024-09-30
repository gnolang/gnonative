ANDROID_DEVICE ?=

TEMPLATE_PROJECT := gnoboard

check-program = $(foreach exec,$(1),$(if $(shell PATH="$(PATH)" which $(exec)),,$(error "Missing deps: no '$(exec)' in PATH")))

PROJECT_DIR := $(shell go list -m -f '{{.Dir}}')

ANDROID_FRAMEWORK_SRC_DIR := $(PROJECT_DIR)/framework/android
ANDROID_FRAMEWORK_OUTPUT_DIR := android/libs
ANDROID_FRAMEWORK_FILES := gnocore.aar gnocore-sources.jar
ANDROID_FRAMEWORK_SRC := $(addprefix $(ANDROID_FRAMEWORK_SRC_DIR)/, $(ANDROID_FRAMEWORK_FILES))
ANDROID_FRAMEWORK_DST := $(addprefix $(ANDROID_FRAMEWORK_OUTPUT_DIR)/, $(ANDROID_FRAMEWORK_FILES))

IOS_FRAMEWORK_SRC_DIR := $(PROJECT_DIR)/framework/ios
IOS_FRAMEWORK_OUTPUT_DIR := ios/Frameworks
IOS_FRAMEWORK_FILES := GnoCore.xcframework
IOS_FRAMEWORK_SRC := $(addprefix $(IOS_FRAMEWORK_SRC_DIR)/, $(IOS_FRAMEWORK_FILES))
IOS_FRAMEWORK_DST := $(addprefix $(IOS_FRAMEWORK_OUTPUT_DIR)/, $(IOS_FRAMEWORK_FILES))

# go files and dependencies
go_framework_files := $(shell find $(PROJECT_DIR)/framework -iname '*.go')
go_service_files := $(shell find $(PROJECT_DIR)/service -iname '*.go')
go_files := $(go_framework_files) $(go_service_files)
go_deps := $(PROJECT_DIR)/go.mod $(PROJECT_DIR)/go.sum $(go_files)

ts_check:
	npm run ts:check

# - Node: Handle node_modules

node_modules: ts_check package.json package-lock.json
	$(call check-program, npm)
	(npm install && touch $@) || true
.PHONY: node_modules

######### ANDROID #########

android: node_modules $(ANDROID_FRAMEWORK_SRC)
	$(call check-program, npx)
	npx expo run:android
.PHONY: build.android

$(ANDROID_FRAMEWORK_SRC): $(go_deps)
	cd $(PROJECT_DIR); $(MAKE) framework.android

######### IOS #########

ios: node_modules $(IOS_FRAMEWORK_SRC)
	$(call check-program, npx)
	npx expo run:ios
.PHONY: build.ios

$(IOS_FRAMEWORK_SRC): $(go_deps)
	cd $(PROJECT_DIR); $(MAKE) framework.ios

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
