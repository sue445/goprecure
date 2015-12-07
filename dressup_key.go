package goprecure

import (
	"fmt"
)

type DressupKey int

const (
	Flora DressupKey = iota
	Mermaid
	Twinkle
	Scarlet
	Rose
	Ice
	Luna
	Hanabi // Fireworks
	Lily
	Bubble
	ShootingStar
	Phoenix
	Sakura // Cherry Blossom
	Sango  // Coral
	Ginga  // Galaxy
	Sun
	Royal
)

func exchangeMessage(k DressupKey) string {
	switch k {
	case Flora:
		return `エクスチェンジ！モードエレガント！
舞え、花よ！プリキュア・フローラル・トルビヨン！
ごきげんよう`

	case Mermaid:
		return `エクスチェンジ！モードエレガント！
高鳴れ、海よ！プリキュア・マーメイド・リップル！
ごきげんよう`

	case Twinkle:
		return `エクスチェンジ！モードエレガント！
キラキラ、星よ！プリキュア・トゥインクル・ハミング！
ごきげんよう`

	case Scarlet:
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

	case Hanabi:
		return `花火！
燃えよ、炎よ！プリキュア・スカーレット・スパーク！
ごきげんよう`

	case Lily:
		return `リリィ！
舞え、百合よ！プリキュア・リース・トルビヨン！
ごきげんよう`

	case Bubble:
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

	case Sakura:
		fallthrough
	case Sango:
		fallthrough
	case Ginga:
		fallthrough
	case Sun:
		return `モードエレガント！
桜！珊瑚！銀河！サン！
ドレスアッププレミアム！
はぁッ！
響け全ての力！プリキュアエクラエスポワール！
ごきげんよう`

	case Royal:
		return `モードエレガント！ロイヤル！
ドレスアップロイヤル！
はぁっ！
響け遥か彼方へ！プリキュアグランプランタン！
ブルーミング
ごきげんよう`
	}

	panic(fmt.Sprintf("Unknown DressUpKey: %d", k))
}
