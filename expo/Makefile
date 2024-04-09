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

build: build.android build.ios

clean: clean.android clean.ios
	cd $(PROJECT_DIR); $(MAKE) clean
.PHONY: clean

######### ANDROID #########

build.android: $(ANDROID_FRAMEWORK_DST)
.PHONY: build.android

$(ANDROID_FRAMEWORK_SRC):
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

$(IOS_FRAMEWORK_SRC):
	cd $(PROJECT_DIR); $(MAKE) framework.ios

$(IOS_FRAMEWORK_DST): $(IOS_FRAMEWORK_SRC)
	mkdir -p $(IOS_FRAMEWORK_OUTPUT_DIR)
	cp -r $^ $(IOS_FRAMEWORK_OUTPUT_DIR)

clean.ios:
	rm -fr $(IOS_FRAMEWORK_OUTPUT_DIR)
.PHONY: clean.ios
