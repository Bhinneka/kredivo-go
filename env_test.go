package kredivo

import (
	"testing"
)

func TestEnv(t *testing.T) {

	t.Run("Test Env SandBox", func(t *testing.T) {

		sandboxEnv := SandBox

		if sandboxEnv.String() != "https://sandbox.kredivo.com/kredivo/v2" {
			t.Errorf("Sandbox Env Invalid %s", sandboxEnv.String())
		}
	})

	t.Run("Test Env Production", func(t *testing.T) {

		productionEnv := Production

		if productionEnv.String() != "https://api.kredivo.com/kredivo/v2" {
			t.Errorf("Production Env Invalid %s", productionEnv.String())
		}
	})

}
