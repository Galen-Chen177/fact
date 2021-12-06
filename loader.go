package fact

import (
	lua "github.com/yuin/gopher-lua"
)

// Preload adds fact to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local fact = require("fact")
func Preload(L *lua.LState) {
	L.PreloadModule("fact", Loader)
}

// Loader is the module loader function.
func Loader(L *lua.LState) int {
	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}

var api = map[string]lua.LGFunction{
	"Hostname":             Hostname,
	"OS":                   OS,
	"Platform":             Platform,
	"PlatformFamily":       PlatformFamily,
	"PlatformVersion":      PlatformVersion,
	"PlatformVersionMajor": PlatformVersionMajor,
	"PlatformVersionMinor": PlatformVersionMinor,
	"KernelVersion":        KernelVersion,
	"KernelArch":           KernelArch,
}
