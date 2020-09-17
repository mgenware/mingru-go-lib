package mingru

import (
	"testing"

	"github.com/mgenware/mingru-go-lib/test"
)

func TestInputPlaceholdersForArray(t *testing.T) {
	test.Assert(t, InputPlaceholders(1), "?")
	test.Assert(t, InputPlaceholders(3), "?, ?, ?")
}
