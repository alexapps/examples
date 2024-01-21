package main

import (
	"os"
	"testing"
)

func TestGetVarDefault(t *testing.T) {
	defaultValue := "--Test--"
	variableName := "TEST_ENV_VALUE"

	if v := getVar(variableName, defaultValue); v != defaultValue {
		t.Fail()
	}
}

func TestGetVarEnv(t *testing.T) {
	defaultValue := "--Test--"
	variableName := "TEST_ENV_VALUE"
	originValue := "SUPPER+TEST"

	if err := os.Setenv(variableName, originValue); err != nil {
		t.Fail()
	}

	if v := getVar(variableName, defaultValue); v != originValue {
		t.Fail()
	}
}

func TestAvg(t *testing.T) {
	xs := []float64{98, 93, 77, 82, 83}

	if avg(xs) != 86.6 {
		t.Fail()
	}
}

func TestAvgEmpty(t *testing.T) {
	xs := []float64{}

	if avg(xs) != 0 {
		t.Fail()
	}
}
