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

func ExampleCureFlora() {
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

func ExampleCureMermaid() {
	mermaid := newCureMermaid()

	// human
	fmt.Println(mermaid.name())

	mermaid.transform()

	// precure
	fmt.Println(mermaid.name())

	// Output:
	// 海藤みなみ
	// キュアマーメイド
}

func ExampleCureTwinkle() {
	twinkle := newCureTwinkle()

	// human
	fmt.Println(twinkle.name())

	twinkle.transform()

	// precure
	fmt.Println(twinkle.name())

	// Output:
	// 天ノ川きらら
	// キュアトゥインクル
}

func ExampleCureScarlett() {
	scarlett := newCureScarlett()

	// human
	fmt.Println(scarlett.name())

	scarlett.transform()

	// precure
	fmt.Println(scarlett.name())

	// Output:
	// 紅城トワ
	// キュアスカーレット
}
