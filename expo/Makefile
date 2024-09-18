PROJECT_DIR := ..

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
#
# go files and dependencies
go_framework_files := $(shell find $(PROJECT_DIR)/framework -iname '*.go')
go_service_files := $(shell find $(PROJECT_DIR)/service -iname '*.go')
go_files := $(go_framework_files) $(go_service_files)
go_deps := $(PROJECT_DIR)/go.mod $(PROJECT_DIR)/go.sum $(go_files)

build: build.android build.ios

clean: clean.android clean.ios
	cd $(PROJECT_DIR); $(MAKE) clean
.PHONY: clean

######### ANDROID #########

build.android: $(ANDROID_FRAMEWORK_DST)
.PHONY: build.android

$(ANDROID_FRAMEWORK_SRC): $(go_deps)
	cd $(PROJECT_DIR); $(MAKE) framework.android

$(ANDROID_FRAMEWORK_DST): $(ANDROID_FRAMEWORK_SRC)
	mkdir -p $(ANDROID_FRAMEWORK_OUTPUT_DIR)
	cp $^ $(ANDROID_FRAMEWORK_OUTPUT_DIR)

clean.android:
	rm -fr $(ANDROID_FRAMEWORK_OUTPUT_DIR)
.PHONY: clean.android

######### IOS #########

build.ios: $(IOS_FRAMEWORK_DST)
.PHONY: build.ios

$(IOS_FRAMEWORK_SRC): $(go_deps)
	cd $(PROJECT_DIR); $(MAKE) framework.ios

$(IOS_FRAMEWORK_DST): $(IOS_FRAMEWORK_SRC)
	mkdir -p $(IOS_FRAMEWORK_OUTPUT_DIR)
	cp -r $^ $(IOS_FRAMEWORK_OUTPUT_DIR)
	touch $(IOS_FRAMEWORK_DST)

clean.ios:
	rm -fr $(IOS_FRAMEWORK_OUTPUT_DIR)
.PHONY: clean.ios

######### NPM #########
npm.pack: build
	npm install
	npm pack
.PHONY: npm.pack

npm.publish: build
	npm install
	npm publish
.PHONY: npm.publish
