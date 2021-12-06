package fact

import (
	"strings"

	"github.com/shirou/gopsutil/host"
	lua "github.com/yuin/gopher-lua"
)

//  fact Hostname() returns (string, err)
func Hostname(L *lua.LState) int {
	osInfo, err := host.Info()
	if err != nil {
		L.Push(lua.LString(""))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(osInfo.Hostname))
	return 1
}

//  fact OS() returns (string, err)
func OS(L *lua.LState) int {
	osInfo, err := host.Info()
	if err != nil {
		L.Push(lua.LString(""))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(osInfo.OS))
	return 1
}

//  fact Platform() returns (string, err)
func Platform(L *lua.LState) int {
	osInfo, err := host.Info()
	if err != nil {
		L.Push(lua.LString(""))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(osInfo.Platform))
	return 1
}

//  fact PlatformFamily() returns (string, err)
func PlatformFamily(L *lua.LState) int {
	osInfo, err := host.Info()
	if err != nil {
		L.Push(lua.LString(""))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(osInfo.PlatformFamily))
	return 1
}

//  fact PlatformVersion() returns (string, err)
func PlatformVersion(L *lua.LState) int {
	osInfo, err := host.Info()
	if err != nil {
		L.Push(lua.LString(""))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(osInfo.PlatformVersion))
	return 1
}

//  fact PlatformVersionMajor() returns (string, err)
func PlatformVersionMajor(L *lua.LState) int {
	osInfo, err := host.Info()
	if err != nil {
		L.Push(lua.LString(""))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	versions := strings.Split(osInfo.PlatformVersion, ".")
	if len(versions) < 2 {
		L.Push(lua.LString(""))
		L.Push(lua.LString("parse platformVersion failed"))
		return 2
	}

	L.Push(lua.LString(strings.TrimSpace(versions[0])))
	return 1
}

//  fact PlatformVersionMinor() returns (string, err)
func PlatformVersionMinor(L *lua.LState) int {
	osInfo, err := host.Info()
	if err != nil {
		L.Push(lua.LString(""))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	versions := strings.Split(osInfo.PlatformVersion, ".")
	if len(versions) < 2 {
		L.Push(lua.LString(""))
		L.Push(lua.LString("parse platformVersion failed"))
		return 2
	}

	L.Push(lua.LString(strings.TrimSpace(versions[1])))
	return 1
}

//  fact KernelVersion() returns (string, err)
func KernelVersion(L *lua.LState) int {
	osInfo, err := host.Info()
	if err != nil {
		L.Push(lua.LString(""))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(osInfo.KernelVersion))
	return 1
}

//  fact KernelArch() returns (string, err)
func KernelArch(L *lua.LState) int {
	osInfo, err := host.Info()
	if err != nil {
		L.Push(lua.LString(""))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(osInfo.KernelArch))
	return 1
}
