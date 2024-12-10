package functions

import (
	"testing"

	"github.com/kolukattai/kurl/boot"
)

func TestCall(t *testing.T) {
	boot.UpdateConfig("config.json", "example")
	Call("example-get-api.md")
}
