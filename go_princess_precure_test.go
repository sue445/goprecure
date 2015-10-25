package go_princess_precure

import (
	"fmt"
	"testing"
)

func TestTransform(t *testing.T) {
	flora := NewCureFlora()
	flora.transform_interval = 0

	if flora.Name() != "春野はるか" {
		t.Errorf("actual=%s, expect=%s", flora.Name(), "春野はるか")
	}

	flora.Transform()

	if flora.Name() != "キュアフローラ" {
		t.Errorf("actual=%s, expect=%s", flora.Name(), "キュアフローラ")
	}

	flora.Transform()

	if flora.Name() != "春野はるか" {
		t.Errorf("actual=%s, expect=%s", flora.Name(), "春野はるか")
	}
}

func TestHumanize(t *testing.T) {
	flora := NewCureFlora()
	flora.transform_interval = 0

	flora.Humanize()

	if flora.Name() != "春野はるか" {
		t.Errorf("actual=%s, expect=%s", flora.Name(), "春野はるか")
	}

	flora.Transform()

	flora.Humanize()

	if flora.Name() != "春野はるか" {
		t.Errorf("actual=%s, expect=%s", flora.Name(), "春野はるか")
	}
}

func ExampleCureFlora() {
	flora := NewCureFlora()
	flora.transform_interval = 0

	// human
	fmt.Println(flora.Name())

	flora.Transform()

	// precure
	fmt.Println(flora.Name())

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
	mermaid := NewCureMermaid()
	mermaid.transform_interval = 0

	// human
	fmt.Println(mermaid.Name())

	mermaid.Transform()

	// precure
	fmt.Println(mermaid.Name())

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
	twinkle := NewCureTwinkle()
	twinkle.transform_interval = 0

	// human
	fmt.Println(twinkle.Name())

	twinkle.Transform()

	// precure
	fmt.Println(twinkle.Name())

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
	scarlett := NewCureScarlett()
	scarlett.transform_interval = 0

	// human
	fmt.Println(scarlett.Name())

	scarlett.Transform()

	// precure
	fmt.Println(scarlett.Name())

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
	flora := NewCureFlora()
	flora.transform_interval = 0

	flora.Transform()

	flora.Exchange(TransformFlora)
}

func TestCanUseKey(t *testing.T) {
	flora := NewCureFlora()

	actual := flora.CanUseKey(TransformFlora)

	if !actual {
		t.Errorf("expect flora.canUseKey(TransformFlora) == true")
	}

	actual = flora.CanUseKey(TransformMermaid)

	if actual {
		t.Errorf("expect flora.canUseKey(TransformMermaid) == false")
	}
}
