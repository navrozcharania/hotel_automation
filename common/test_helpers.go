package common

import (
	"log"
	"testing"
)

func AssertNil(t *testing.T, val interface{}) {
	if val != nil {
		t.Fail()
	}
}

func AssertTrue(t *testing.T, val bool) {
	if !val {
		log.Fatal("Expected true got false")
		t.Fail()
	}
}

func AssertFalse(t *testing.T, val bool) {
	if val {
		log.Fatal("Expected false got true")
		t.Fail()
	}
}

func AssertEqualInt64(t *testing.T, val1, val2 int64) {
	if val1 != val2 {
		t.Fail()
	}
}