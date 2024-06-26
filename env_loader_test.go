package envloader_test

import (
	"os"
	"strconv"
	"testing"

	envloader "github.com/Shin-Thant/env-loader"
)

func TestEnvLoader_Result(t *testing.T) {
	PORT := 3000
	DATABASE_URL := "postgres://postgres:pwd@localhost:5432/hello"
	var unparsedField []int = nil
	STRING_VAL := "this is string"
	INT_VAL := "1000"

	os.Setenv("PORT", strconv.Itoa(PORT))
	os.Setenv("DATABASE_URL", DATABASE_URL)
	os.Setenv("StrInterfaceField", STRING_VAL)
	os.Setenv("IntInterfaceField", INT_VAL)

	type appEnv struct {
		PORT              int
		DATABASE_URL      string
		UnparsedField     []int
		StrInterfaceField interface{}
		IntInterfaceField interface{}
	}
	app := appEnv{
		UnparsedField: unparsedField,
	}
	envloader.LoadEnv(&app, nil)

	if app.PORT != PORT {
		t.Errorf("Incorrect result for PORT: got %d, want: %d\n", app.PORT, PORT)
	}
	if app.DATABASE_URL != DATABASE_URL {
		t.Errorf("Incorrect result for DATABASE_URL: got %s, want: %s\n", app.DATABASE_URL, DATABASE_URL)
	}
	if app.UnparsedField != nil {
		t.Errorf("Incorrect result for UnparsedField: got %v, want: %v\n", app.UnparsedField, unparsedField)
	}
	if app.StrInterfaceField != STRING_VAL {
		t.Errorf("Incorrect result for StrInterfaceField: got %v, want: `%s`\n", app.StrInterfaceField, STRING_VAL)
	}
	if app.IntInterfaceField != INT_VAL {
		t.Errorf("Incorrect result for IntInterfaceField: got %v, want: `%s`\n", app.IntInterfaceField, INT_VAL)
	}
}

func TestEnvLoader_LoadOptions(t *testing.T) {
	type appEnv struct {
		RandomField string
	}
	app := appEnv{}
	err := envloader.LoadEnv(&app, &envloader.LoadEnvOptions{
		EnvPath: "invalid path",
	})
	if err == nil {
		t.Error("EnvLoader should return an error.")
	}

	if app.RandomField != "" {
		t.Errorf("Incorrect result: got %s, expect empty string", app.RandomField)
	}
}
