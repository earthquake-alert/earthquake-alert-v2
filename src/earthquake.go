package src

import (
	"encoding/xml"
	"time"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

// 震度速報、震源に関する情報、震源・震度に関する情報
type Earthquake struct {
	Row    string
	Parsed *jma.EarthquakeJmaXml
}

func ParseEarthquake(row []byte) (*Earthquake, error) {
	earthquake := new(jma.EarthquakeJmaXml)
	err := xml.Unmarshal(row, earthquake)
	if err != nil {
		return nil, err
	}

	return &Earthquake{
		Row:    string(row),
		Parsed: earthquake,
	}, nil
}

func (c *Earthquake) Assembly() error {
	// TODO: あとで
	return nil
}

func (c *Earthquake) SetImages() error {
	// TODO: あとで
	return nil
}

func (c *Earthquake) GetTitle() string {
	return ""
}

func (c *Earthquake) GetTargetDate() time.Time {
	return time.Now()
}

func (c *Earthquake) GetInfoType() jma.InfoType {
	return c.Parsed.Head.InfoType
}
