package dotnev

import (
	"os"
	"testing"
)

func TestLoadDotenv(T *testing.T) {
	test := Default{"PASSWORD", "secret"}

	LoadDotenv(test)
	if os.Getenv("PASSWORD") != "secret" {
		T.Fatalf("Got %q", os.Getenv("PASSWORD"))
	}

	os.Setenv("PASSWORD", "")

	test = Default{"password", "secret"}

	LoadDotenv(test)
	if os.Getenv("PASSWORD") != "secret" {
		T.Fatalf("Got %q", os.Getenv("PASSWORD"))
	}

	os.Setenv("PASSWORD", "")

	test = Default{"password", "secret"}
	test2 := Default{"login", "admin"}

	LoadDotenv(test, test2)
	if os.Getenv("LOGIN") != "admin" {
		T.Fatalf("Got %q", os.Getenv("LOGIN"))
	}

}
