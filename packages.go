// +build packages

//go:generate go build
//go:generate upx dude
//go:generate holo-build --force --format=rpm holo.toml
//go:generate holo-build --force --format=pacman holo.toml
//go:generate rm dude
package main
