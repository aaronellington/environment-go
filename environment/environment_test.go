package environment_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/aaronellington/environment-go/environment"
)

func TestOrderOfOperations(t *testing.T) {
	type Config struct {
		Local   string `env:"ENV_TEST_ENV_LOCAL"`
		System  string `env:"ENV_TEST_ENV_SYSTEM"`
		Default string `env:"ENV_TEST_ENV_DEFAULT"`
	}

	t.Setenv("ENV_TEST_ENV_SYSTEM", "system_expected")

	_ = os.Chdir("test_files/test01")

	env := environment.New()

	config := Config{
		Local:   "local",
		System:  "system",
		Default: "default",
	}

	if err := env.Decode(&config); err != nil {
		t.Fatal(err)
	}

	expected := Config{
		Local:   "local_expected",
		Default: "default_expected",
		System:  "system_expected",
	}

	if !reflect.DeepEqual(config, expected) {
		t.Fatalf("Wrong Value:\n- Expected: %+v\n-   Actual: %+v", expected, config)
	}
}
