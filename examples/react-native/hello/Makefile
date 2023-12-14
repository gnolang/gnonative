ANDROID_DEVICE ?=

check-program = $(foreach exec,$(1),$(if $(shell PATH="$(PATH)" which $(exec)),,$(error "Missing deps: no '$(exec)' in PATH")))

TEMPLATE_PROJECT := gnoboard

# - Node: Handle node_modules

node_modules: package.json yarn.lock
	$(call check-program, yarn)
	(yarn && touch $@) || true
.PHONY: node_modules

ios/${TEMPLATE_PROJECT}.xcworkspace: ios/${TEMPLATE_PROJECT}.xcodeproj ios/Podfile node_modules
	$(call check-program, bundle)
	cd ios && pod install
	touch $@
.PHONY: ios/${TEMPLATE_PROJECT}.xcworkspace

# - Helpers

android.reverse_tcp:
	$(call check-program, adb)
	$(if $(ANDROID_DEVICE),,$(eval ANDROID_DEVICE = $(shell adb devices | tail +2 | head -1 | cut -f 1)))
	@if [ -z "$(ANDROID_DEVICE)" ]; then \
	  >&2 echo "ERROR: no Android device found"; exit 1; \
	fi
	adb -s $(ANDROID_DEVICE) reverse tcp:8081 tcp:8081 # metro
.PHONY: android.reverse_tcp
