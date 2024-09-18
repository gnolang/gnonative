ANDROID_DEVICE ?=

check-program = $(foreach exec,$(1),$(if $(shell PATH="$(PATH)" which $(exec)),,$(error "Missing deps: no '$(exec)' in PATH")))

TEMPLATE_PROJECT := gnoboard

# - Node: Handle node_modules

build.android: node_modules
	$(call check-program, npx)
	npx expo run:android

build.ios: node_modules
	$(call check-program, npx)
	npx expo run:ios

node_modules: package.json package-lock.json
	$(call check-program, npm)
	(npm install && touch $@) || true
.PHONY: node_modules

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
