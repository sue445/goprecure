package goprecure

import (
	"fmt"
	"testing"
)

func ExampleGirl_Name() {
	flora := NewCureFlora()
	flora.PrintLine = false

	// human
	fmt.Println("Name", flora.Name())

	flora.Transform()

	// precure
	fmt.Println("Name", flora.Name())

	// Output:
	// Name 春野はるか
	// Name キュアフローラ
}

func TestTransform(t *testing.T) {
	flora := NewCureFlora()
	flora.PrintInterval = 0
	flora.PrintLine = false

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

func ExampleGirl_Transform() {
	flora := NewCureFlora()
	flora.PrintInterval = 0

	// human
	fmt.Println("Name", flora.Name())

	flora.Transform()

	// precure
	fmt.Println("Name", flora.Name())

	// Output:
	// Name 春野はるか
	// プリキュア！プリンセスエンゲージ！
	// 咲き誇る花のプリンセス！キュアフローラ！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟はよろしくて？
	// Name キュアフローラ
}

func TestHumanize(t *testing.T) {
	flora := NewCureFlora()
	flora.PrintInterval = 0
	flora.PrintLine = false

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

func ExampleGirl_Humanize() {
	flora := NewCureFlora()
	flora.PrintLine = false

	// human
	fmt.Println("Name", flora.Name())

	flora.Transform()

	// precure
	fmt.Println("Name", flora.Name())

	flora.Humanize()

	// human
	fmt.Println("Name", flora.Name())

	// Output:
	// Name 春野はるか
	// Name キュアフローラ
	// Name 春野はるか
}

func TestExchange(t *testing.T) {
	flora := NewCureFlora()
	flora.PrintLine = false

	flora.Transform()

	flora.Exchange(Flora)
}

func ExampleGirl_Exchange_human() {
	flora := NewCureFlora()

	_, err := flora.Exchange(Flora)
	fmt.Println(err)

	// Output:
	// Human can not exchange!
}

func ExampleGirl_Exchange_precure() {
	flora := NewCureFlora()
	flora.PrintInterval = 0
	flora.PrintLine = false

	flora.Transform()

	flora.PrintLine = true
	flora.Exchange(Flora)

	// Output:
	// エクスチェンジ！モードエレガント！
	// 舞え、花よ！プリキュア・フローラル・トルビヨン！
	// ごきげんよう
}

func ExampleGirl_Exchange_invalidDressupKey() {
	flora := NewCureFlora()
	flora.PrintInterval = 0
	flora.PrintLine = false

	flora.Transform()

	_, err := flora.Exchange(Mermaid)
	fmt.Println(err)

	// Output:
	// キュアフローラ can not use Mermaid
}

func TestCanUseKey(t *testing.T) {
	flora := NewCureFlora()

	actual := flora.CanUseKey(Flora)

	if !actual {
		t.Errorf("expect flora.canUseKey(Flora) == true")
	}

	actual = flora.CanUseKey(Mermaid)

	if actual {
		t.Errorf("expect flora.canUseKey(Mermaid) == false")
	}
}

func ExampleGirl_CanUseKey() {
	flora := NewCureFlora()

	ret1 := flora.CanUseKey(Flora)
	fmt.Println("flora can use Flora ?:", ret1)

	ret2 := flora.CanUseKey(Mermaid)
	fmt.Println("flora can use Mermaid ?:", ret2)

	// Output:
	// flora can use Flora ?: true
	// flora can use Mermaid ?: false
}

func ExampleNewCureFlora() {
	flora := NewCureFlora()

	fmt.Println("HumanName", flora.HumanName)
	fmt.Println("PrecureName", flora.PrecureName)
	fmt.Println("CastName", flora.CastName)
	fmt.Println("Color", flora.Color)
	fmt.Println("CreatedDate", flora.CreatedDate)
	fmt.Println("DressupKeys", flora.DressupKeys)

	// Output:
	// HumanName 春野はるか
	// PrecureName キュアフローラ
	// CastName 嶋村侑
	// Color pink
	// CreatedDate 2015-02-01
	// DressupKeys [Flora Rose Lily Sakura Royal]
}

func ExampleNewCureMermaid() {
	mermaid := NewCureMermaid()

	fmt.Println("HumanName", mermaid.HumanName)
	fmt.Println("PrecureName", mermaid.PrecureName)
	fmt.Println("CastName", mermaid.CastName)
	fmt.Println("Color", mermaid.Color)
	fmt.Println("CreatedDate", mermaid.CreatedDate)
	fmt.Println("DressupKeys", mermaid.DressupKeys)

	// Output:
	// HumanName 海藤みなみ
	// PrecureName キュアマーメイド
	// CastName 浅野真澄
	// Color blue
	// CreatedDate 2015-02-08
	// DressupKeys [Mermaid Ice Bobble Sango Royal]
}

func ExampleNewCureTwinkle() {
	twinkle := NewCureTwinkle()

	fmt.Println("HumanName", twinkle.HumanName)
	fmt.Println("PrecureName", twinkle.PrecureName)
	fmt.Println("CastName", twinkle.CastName)
	fmt.Println("Color", twinkle.Color)
	fmt.Println("CreatedDate", twinkle.CreatedDate)
	fmt.Println("DressupKeys", twinkle.DressupKeys)

	// Output:
	// HumanName 天ノ川きらら
	// PrecureName キュアトゥインクル
	// CastName 山村響
	// Color yellow
	// CreatedDate 2015-02-22
	// DressupKeys [Twinkle Luna ShootingStar Ginga Royal]
}

func ExampleNewCureScarlett() {
	scarlett := NewCureScarlett()

	fmt.Println("HumanName", scarlett.HumanName)
	fmt.Println("PrecureName", scarlett.PrecureName)
	fmt.Println("CastName", scarlett.CastName)
	fmt.Println("Color", scarlett.Color)
	fmt.Println("CreatedDate", scarlett.CreatedDate)
	fmt.Println("DressupKeys", scarlett.DressupKeys)

	// Output:
	// HumanName 紅城トワ
	// PrecureName キュアスカーレット
	// CastName 沢城みゆき
	// Color red
	// CreatedDate 2015-07-05
	// DressupKeys [Scarlett Hanabi Phoenix Sun Royal]
}
