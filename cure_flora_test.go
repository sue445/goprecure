package goprecure

import (
	"fmt"
)

func ExampleCureFlora() {
	flora := NewCureFlora()
	flora.PrintInterval = 0

	// human
	fmt.Println("Name", flora.Name())

	flora.Transform()

	// precure
	fmt.Println("Name", flora.Name())

	fmt.Println(Flora.String())
	flora.Exchange(Flora)

	fmt.Println(Rose.String())
	flora.Exchange(Rose)

	fmt.Println(Lily.String())
	flora.Exchange(Lily)

	fmt.Println(Sakura.String())
	flora.Exchange(Sakura)

	// Output:
	// Name 春野はるか
	// プリキュア！プリンセスエンゲージ！
	// 咲き誇る花のプリンセス！キュアフローラ！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟はよろしくて？
	// Name キュアフローラ
	// Flora
	// エクスチェンジ！モードエレガント！
	// 舞え、花よ！プリキュア・フローラル・トルビヨン！
	// ごきげんよう
	// Rose
	// ローズ！
	// 舞え、薔薇よ！プリキュア・ローズ・トルビヨン！
	// ごきげんよう
	// Lily
	// リリィ！
	// 舞え、百合よ！プリキュア・リース・トルビヨン！
	// ごきげんよう
	// Sakura
	// モードエレガント！
	// 桜！珊瑚！銀河！サン！
	// ドレスアッププレミアム！
	// はぁッ！
	// 響け全ての力！プリキュアエクラエスポワール！
	// ごきげんよう
}
