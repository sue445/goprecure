package go_princess_precure

import (
	"fmt"
	"testing"
)

func TestTransform(t *testing.T) {
	flora := newCureFlora()
	flora.transform_interval = 0

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
	flora.transform_interval = 0

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
	flora.transform_interval = 0

	// human
	fmt.Println(flora.name())

	flora.transform()

	// precure
	fmt.Println(flora.name())

	// Output:
	// 春野はるか
	// プリキュア！プリンセスエンゲージ！
	// 咲き誇る花のプリンセス！キュアフローラ！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟はよろしくて？
	// キュアフローラ
}

func ExampleCureMermaid() {
	mermaid := newCureMermaid()
	mermaid.transform_interval = 0

	// human
	fmt.Println(mermaid.name())

	mermaid.transform()

	// precure
	fmt.Println(mermaid.name())

	// Output:
	// 海藤みなみ
	// プリキュア！プリンセスエンゲージ！
	// 澄み渡る海のプリンセス！キュアマーメイド！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟はよろしくて？
	// キュアマーメイド
}

func ExampleCureTwinkle() {
	twinkle := newCureTwinkle()
	twinkle.transform_interval = 0

	// human
	fmt.Println(twinkle.name())

	twinkle.transform()

	// precure
	fmt.Println(twinkle.name())

	// Output:
	// 天ノ川きらら
	// プリキュア！プリンセスエンゲージ！
	// きらめく星のプリンセス！キュアトゥインクル！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟はよろしくて？
	// キュアトゥインクル
}

func ExampleCureScarlett() {
	scarlett := newCureScarlett()
	scarlett.transform_interval = 0

	// human
	fmt.Println(scarlett.name())

	scarlett.transform()

	// precure
	fmt.Println(scarlett.name())

	// Output:
	// 紅城トワ
	// プリキュア！プリンセスエンゲージ！
	// 深紅の炎のプリンセス！キュアスカーレット！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟決めなさい！
	// キュアスカーレット
}

func TestExchange(t *testing.T) {
	flora := newCureFlora()
	flora.transform_interval = 0

	flora.transform()

	flora.exchange(TransformFlora)
}

func TestCanUseKey(t *testing.T) {
	flora := newCureFlora()

	actual := flora.canUseKey(TransformFlora)

	if !actual {
		t.Errorf("expect flora.canUseKey(TransformFlora) == true")
	}

	actual = flora.canUseKey(TransformMermaid)

	if actual {
		t.Errorf("expect flora.canUseKey(TransformMermaid) == false")
	}
}
