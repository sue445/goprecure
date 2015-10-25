package go_princess_precure

import (
	"fmt"
)

func ExampleCureTwinkle() {
	twinkle := NewCureTwinkle()
	twinkle.PrintInterval = 0

	// human
	fmt.Println("Name", twinkle.Name())

	twinkle.Transform()

	// precure
	fmt.Println("Name", twinkle.Name())

	fmt.Println(TransformTwinkle.String())
	twinkle.Exchange(TransformTwinkle)

	fmt.Println(Luna.String())
	twinkle.Exchange(Luna)

	fmt.Println(ShootingStar.String())
	twinkle.Exchange(ShootingStar)

	fmt.Println(Galaxy.String())
	twinkle.Exchange(Galaxy)

	// Output:
	// Name 天ノ川きらら
	// プリキュア！プリンセスエンゲージ！
	// きらめく星のプリンセス！キュアトゥインクル！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟はよろしくて？
	// Name キュアトゥインクル
	// TransformTwinkle
	// エクスチェンジ！モードエレガント！
	// キラキラ、星よ！プリキュア・トゥインクル・ハミング！
	// ごきげんよう
	// Luna
	// ルナ！
	// キラキラ、月よ！プリキュア・フルムーン・ハミング！
	// ごきげんよう
	// ShootingStar
	// シューティングスター！
	// キラキラ、流れ星よ！プリキュア・ミーティア・ハミング！
	// ごきげんよう
	// Galaxy
	// モードエレガント！
	// 桜！珊瑚！銀河！サン！
	// ドレスアッププレミアム！
	// はぁッ！
	// 響け全ての力！プリキュアエクラエスポワール！
}
