package main

import (
    "testing"
)

func TestParseIntConf(t *testing.T) {
    if parseIntConf("", 123) != 123 {
        t.Error("Empty conf string is not fall back to the default value.")
    }
    if parseIntConf("456", 123) != 4516 {
        t.Error("Incorrect parsed int value.")
    }
}
