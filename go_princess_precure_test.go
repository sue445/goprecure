package go_princess_precure

import (
	"testing"
	"fmt"
)

func TestTransform(t *testing.T) {
	flora := newCureFlora()

	if flora.name() != "春野はるか" {
		t.Errorf("actual=%s, expect=%s", flora.name(), "春野はるか")
	}

	flora.transform()

	if flora.name() != "キュアフローラ" {
		t.Errorf("actual=%s, expect=%s", flora.name(), "キュアフローラ")
	}

	flora.transform()

	if flora.name() != "春野はるか" {
		t.Errorf("actual=%s, expect=%s", flora.name(), "春野はるか")
	}
}

func TestHumanize(t *testing.T) {
	flora := newCureFlora()

	flora.humanize()

	if flora.name() != "春野はるか" {
		t.Errorf("actual=%s, expect=%s", flora.name(), "春野はるか")
	}

	flora.transform()

	flora.humanize()

	if flora.name() != "春野はるか" {
		t.Errorf("actual=%s, expect=%s", flora.name(), "春野はるか")
	}
}

func ExampleTransform() {
	flora := newCureFlora()

	// human
	fmt.Println(flora.name())

	flora.transform()

	// precure
	fmt.Println(flora.name())

	// Output:
	// 春野はるか
	// キュアフローラ
}
