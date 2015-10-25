package go_princess_precure

import (
	"fmt"
)

type DressupKey int

const (
	TransformFlora    DressupKey = iota // 変身ドレスアップキー（キュアフローラ）
	TransformMermaid                    // 変身ドレスアップキー（キュアマーメイド）
	TransformTwinkle                    // 変身ドレスアップキー（キュアトゥインクル）
	TransformScarlett                   // 変身ドレスアップキー（キュアスカーレット）
	Rose                                // エレガントローズドレスアップキー
	Ice                                 // エレガントアイスドレスアップキー
	Luna                                // エレガントルナドレスアップキー
	Fireworks                           // エレガントハナビドレスアップキー
	Lily                                // ミラクルリリィドレスアップキー
	Bobble                              // ミラクルバブルドレスアップキー
	ShootingStar                        // ミラクルシューティングスタードレスアップキー
	Phoenix                             // ミラクルフェニックスドレスアップキー
	Blossom                             // プレミアムサクラドレスアップキー
	Coral                               // プレミアムサンゴドレスアップキー
	Galaxy                              // プレミアムギンガドレスアップキー
	Sun                                 // プレミアムサンドレスアップキー
)

func exchangeMessage(k DressupKey) string {
	switch k {
	case TransformFlora:
		return `エクスチェンジ！モードエレガント！
舞え、花よ！プリキュア・フローラル・トルビヨン！
ごきげんよう`

	case TransformMermaid:
		return `エクスチェンジ！モードエレガント！
高鳴れ、海よ！プリキュア・マーメイド・リップル！
ごきげんよう`

	case TransformTwinkle:
		return `エクスチェンジ！モードエレガント！
キラキラ、星よ！プリキュア・トゥインクル・ハミング！
ごきげんよう`

	case TransformScarlett:
		return `エクスチェンジ！モードエレガント！
たぎれ、炎よ！プリキュア・スカーレット・フレイム！
ごきげんよう`

	case Rose:
		return `ローズ！
舞え、薔薇よ！プリキュア・ローズ・トルビヨン！
ごきげんよう`

	case Ice:
		return `アイス！
高鳴れ、氷よ！プリキュア・フローズン・リップル！
ごきげんよう`

	case Luna:
		return `ルナ！
キラキラ、月よ！プリキュア・フルムーン・ハミング！
ごきげんよう`

	case Fireworks:
		return `花火！
燃えよ、炎よ！プリキュア・スカーレット・スパーク！
ごきげんよう`

	case Lily:
		return `リリィ！
舞え、百合よ！プリキュア・リース・トルビヨン！
ごきげんよう`

	case Bobble:
		return `バブル！
高まれ、泡よ！プリキュア・バブルリップル！
ごきげんよう`

	case ShootingStar:
		return `シューティングスター！
キラキラ、流れ星よ！プリキュア・ミーティア・ハミング！
ごきげんよう`

	case Phoenix:
		return `スカーレットバイオリン！フェニックス！
羽ばたけ炎の翼！プリキュア・フェニックス・ブレイズ！
ごきげんよう`

	case Blossom:
		fallthrough
	case Coral:
		fallthrough
	case Galaxy:
		fallthrough
	case Sun:
		return `モードエレガント！
桜！珊瑚！銀河！サン！
ドレスアッププレミアム！
はぁッ！
響け全ての力！プリキュアエクラエスポワール！`
	}

	panic(fmt.Sprintf("Unknown DressUpKey: %d", k))
}
