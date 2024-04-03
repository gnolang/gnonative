PROJECT_DIR := ..

build: android ios

android:
	cd $(PROJECT_DIR); $(MAKE) build.android
	cp $(PROJECT_DIR)/framework/android/*.{aar,jar} android/libs
.PHONY: android

ios:
	cd $(PROJECT_DIR); $(MAKE) build.ios
	cp -r $(PROJECT_DIR)/framework/ios/*.xcframework ios/Frameworks
.PHONY: ios

clean: clean.android clean.ios
.PHONY: clean

clean.android:
	rm -f android/lib/*.{aar,jar}
.PHONY: clean.android

clean.ios:
	rm -fr ios/Frameworks/*.xcframework
.PHONY: clean.ios
