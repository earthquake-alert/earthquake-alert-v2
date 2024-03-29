package src

import (
	"encoding/xml"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

type Tsunami struct {
	Row    string
	Parsed *jma.TsunamiJmaXml
}

// 気象庁の津波XML電文をパースする
func ParseTsunami(row []byte) (*Tsunami, error) {
	tsunami := new(jma.TsunamiJmaXml)
	err := xml.Unmarshal(row, tsunami)
	if err != nil {
		return nil, err
	}

	return &Tsunami{
		Row:    string(row),
		Parsed: tsunami,
	}, nil
}

// 伝聞の形式を返す
func (c *Tsunami) Status() jma.Status {
	return c.Parsed.Control.Status
}

// 伝聞が"通常の運用で発表する情報"であるかどうかを比較する
// 訓練報やテスト配信などではfalse
func (c *Tsunami) IsCommon() bool {
	s := c.Status()

	switch s {
	case jma.Common:
		return true
	default:
		return false
	}
}

// 情報携帯を返す
func (c *Tsunami) GetInfoType() jma.InfoType {
	return c.Parsed.Head.InfoType
}
