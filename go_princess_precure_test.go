package go_princess_precure

import (
	"testing"
)

func AssertEquals(expect interface{}, actual interface{}, t *testing.T) {
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestTransform(t *testing.T) {
	flora := newCureFlora()

	AssertEquals("春野はるか", flora.name(), t)

	flora.transform()

	AssertEquals("キュアフローラ", flora.name(), t)

	flora.transform()

	AssertEquals("春野はるか", flora.name(), t)
}
