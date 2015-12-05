package goprecure

import (
	"fmt"
)

func ExampleCureScarlett() {
	scarlett := NewCureScarlett()
	scarlett.PrintInterval = 0

	// human
	fmt.Println("Name", scarlett.Name())

	scarlett.Transform()

	// precure
	fmt.Println("Name", scarlett.Name())

	fmt.Println(Scarlett.String())
	scarlett.Exchange(Scarlett)

	fmt.Println(Hanabi.String())
	scarlett.Exchange(Hanabi)

	fmt.Println(Phoenix.String())
	scarlett.Exchange(Phoenix)

	fmt.Println(Sun.String())
	scarlett.Exchange(Sun)

	// Output:
	// Name 紅城トワ
	// プリキュア！プリンセスエンゲージ！
	// 深紅の炎のプリンセス！キュアスカーレット！
	// 強く、やさしく、美しく！
	// Go!プリンセスプリキュア！
	// 冷たい檻に閉ざされた夢、返していただきますわ！
	// お覚悟決めなさい！
	// Name キュアスカーレット
	// Scarlett
	// エクスチェンジ！モードエレガント！
	// たぎれ、炎よ！プリキュア・スカーレット・フレイム！
	// ごきげんよう
	// Hanabi
	// 花火！
	// 燃えよ、炎よ！プリキュア・スカーレット・スパーク！
	// ごきげんよう
	// Phoenix
	// スカーレットバイオリン！フェニックス！
	// 羽ばたけ炎の翼！プリキュア・フェニックス・ブレイズ！
	// ごきげんよう
	// Sun
	// モードエレガント！
	// 桜！珊瑚！銀河！サン！
	// ドレスアッププレミアム！
	// はぁッ！
	// 響け全ての力！プリキュアエクラエスポワール！
	// ごきげんよう
}
