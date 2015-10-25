package go_princess_precure

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	DefaultIntervalSec = 1
)

type Girl struct {
	HumanName        string
	PrecureName      string
	CastName         string
	Color            string
	CreatedDate      string
	TransformMessage string
	DressupKeys      []DressupKey
	PrintInterval    time.Duration
	PrintLine        bool
	currentState     int
}

func (g *Girl) Name() string {
	switch g.currentState {
	case 0:
		return g.HumanName
	case 1:
		return g.PrecureName
	}
	return g.HumanName
}

func (g *Girl) Transform() *Girl {
	if g.currentState == 0 {
		g.printByLine(g.TransformMessage)
	}

	g.currentState++

	if g.currentState > 1 {
		g.currentState = 0
	}

	return g
}

func (g *Girl) printByLine(message string) {
	if !g.PrintLine {
		return
	}

	for _, line := range strings.Split(message, "\n") {
		line = strings.TrimSpace(line)
		fmt.Println(line)
		time.Sleep(g.PrintInterval)
	}
}

func (g *Girl) Humanize() {
	g.currentState = 0
}

func (g *Girl) Exchange(k DressupKey) (*Girl, error) {
	if g.currentState == 0 {
		return g, errors.New("Human can not exchange!")
	}

	if !g.CanUseKey(k) {
		return g, errors.New(g.PrecureName + " can not use " + k.String())
	}

	var message = exchangeMessage(k)
	g.printByLine(message)

	return g, nil
}

func (g *Girl) CanUseKey(k DressupKey) bool {
	for _, value := range g.DressupKeys {
		if value == k {
			return true
		}
	}
	return false
}

func newGirl() *Girl {
	g := new(Girl)
	g.currentState = 0
	g.PrintInterval = DefaultIntervalSec * time.Second
	g.PrintLine = true
	return g
}

func NewCureFlora() *Girl {
	g := newGirl()
	g.HumanName = "春野はるか"
	g.PrecureName = "キュアフローラ"
	g.CastName = "嶋村侑"
	g.Color = "pink"
	g.CreatedDate = "2015-02-01"
	g.DressupKeys = []DressupKey{TransformFlora, Rose, Lily, Blossom}

	g.TransformMessage = `プリキュア！プリンセスエンゲージ！
咲き誇る花のプリンセス！キュアフローラ！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟はよろしくて？`

	return g
}

func NewCureMermaid() *Girl {
	g := newGirl()
	g.HumanName = "海藤みなみ"
	g.PrecureName = "キュアマーメイド"
	g.CastName = "浅野真澄"
	g.Color = "blue"
	g.CreatedDate = "2015-02-08"
	g.DressupKeys = []DressupKey{TransformMermaid, Ice, Bobble, Coral}

	g.TransformMessage = `プリキュア！プリンセスエンゲージ！
澄み渡る海のプリンセス！キュアマーメイド！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟はよろしくて？`

	return g
}

func NewCureTwinkle() *Girl {
	g := newGirl()
	g.HumanName = "天ノ川きらら"
	g.PrecureName = "キュアトゥインクル"
	g.CastName = "山村響"
	g.Color = "yellow"
	g.CreatedDate = "2015-02-22"
	g.DressupKeys = []DressupKey{TransformTwinkle, Luna, ShootingStar, Galaxy}

	g.TransformMessage = `プリキュア！プリンセスエンゲージ！
きらめく星のプリンセス！キュアトゥインクル！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟はよろしくて？`

	return g
}

func NewCureScarlett() *Girl {
	g := newGirl()
	g.HumanName = "紅城トワ"
	g.PrecureName = "キュアスカーレット"
	g.CastName = "沢城みゆき"
	g.Color = "red"
	g.CreatedDate = "2015-07-05"
	g.DressupKeys = []DressupKey{TransformScarlett, Fireworks, Phoenix, Sun}

	g.TransformMessage = `プリキュア！プリンセスエンゲージ！
深紅の炎のプリンセス！キュアスカーレット！
強く、やさしく、美しく！
Go!プリンセスプリキュア！
冷たい檻に閉ざされた夢、返していただきますわ！
お覚悟決めなさい！`

	return g
}
