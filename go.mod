module github.com/gnolang/gnomobile

go 1.20

require golang.org/x/mobile v0.0.0-20230531173138-3c911d8e3eda

require (
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/tools v0.12.1-0.20230818130535-1517d1a3ba60 // indirect
)

replace golang.org/x/mobile => github.com/berty/mobile v0.0.10 // temporary, see https://github.com/golang/mobile/pull/58 and https://github.com/golang/mobile/pull/82
