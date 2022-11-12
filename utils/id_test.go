package utils

import (
	"testing"
)

func TestUuid(t *testing.T) {
	if uuid := Uuid(); uuid == "" {
		t.Errorf("uuid generate result is null")
	} else {
		t.Logf("uuid generate result: %s", uuid)
	}
}

func TestId(t *testing.T) {
	if id := Id(); id == 0 {
		t.Errorf("snowflake generate result is null")
	} else {
		t.Logf("snowflake id generate result: %d", id)
	}
}
