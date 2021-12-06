# fact 

## Usage

```lua
local fact = require("fact")

-- fact.Hostname
local hostname, err = fact.Hostname()
if err then error(err) end
print("Hostname:",hostname)

-- fact.OS
local os, err = fact.OS()
if err then error(err) end
print("OS:",os)

-- fact.Platform
local Platform, err = fact.Platform()
if err then error(err) end
print("Platform:",Platform)

-- fact.PlatformFamily
local PlatformFamily, err = fact.PlatformFamily()
if err then error(err) end
print("PlatformFamily:",PlatformFamily)

-- fact.PlatformVersion
local PlatformVersion, err = fact.PlatformVersion()
if err then error(err) end
print("PlatformVersion:",PlatformVersion)

-- fact.PlatformVersionMajor
local PlatformVersionMajor, err = fact.PlatformVersionMajor()
if err then error(err) end
print("PlatformVersionMajor:",PlatformVersionMajor)

-- fact.PlatformVersionMinor
local PlatformVersionMinor, err = fact.PlatformVersionMinor()
if err then error(err) end
print("PlatformVersionMinor:",PlatformVersionMinor)

-- fact.KernelVersion
local KernelVersion, err = fact.KernelVersion()
if err then error(err) end
print("KernelVersion:",KernelVersion)

-- fact.KernelArch
local KernelArch, err = fact.KernelArch()
if err then error(err) end
print("KernelArch:",KernelArch)
```