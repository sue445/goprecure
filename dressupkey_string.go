// generated by stringer -type DressupKey dressup_key.go; DO NOT EDIT

package goprecure

import "fmt"

const _DressupKey_name = "FloraMermaidTwinkleScarlettRoseIceLunaHanabiLilyBobbleShootingStarPhoenixSakuraSangoGingaSun"

var _DressupKey_index = [...]uint8{0, 5, 12, 19, 27, 31, 34, 38, 44, 48, 54, 66, 73, 79, 84, 89, 92}

func (i DressupKey) String() string {
	if i < 0 || i >= DressupKey(len(_DressupKey_index)-1) {
		return fmt.Sprintf("DressupKey(%d)", i)
	}
	return _DressupKey_name[_DressupKey_index[i]:_DressupKey_index[i+1]]
}
