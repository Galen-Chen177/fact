package fact

import (
	"log"
	"testing"

	lua "github.com/yuin/gopher-lua"
)

// fact.Hostname
func TestHostname(t *testing.T) {
	state := lua.NewState()
	Preload(state)
	source := `
	local fact = require("fact")

	local hostname, err = fact.Hostname()
	if err then error(err) end
	print("Hostname:",hostname)
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
}
