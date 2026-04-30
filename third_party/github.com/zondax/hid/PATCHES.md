# Local patches

This is a vendored copy of `github.com/zondax/hid` v0.9.2.

Local changes:

- Android is excluded from `hid_enabled.go`.
- Android is included in `hid_disabled.go`.
- Android is excluded from `wchar.go`.

See `android-no-cgo-hid.patch` for the exact patch applied to the upstream
v0.9.2 files.

This makes Android gomobile builds use the no-op HID implementation instead of
compiling the bundled hidapi/libusb C code. Newer Android NDKs define
`pthread_barrier_t`, which conflicts with the old hidapi Android compatibility
shim bundled in v0.9.2.
