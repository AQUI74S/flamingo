package security_test

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/core/security"
	"testing"
)

func TestModule_Configure(t *testing.T) {
	if err := dingo.TryModule(new(security.Module)); err != nil {
		t.Error(err)
	}
}
