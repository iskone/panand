package lib

import (
	"fmt"
	"github.com/iskone/itools/xml"
	"testing"
	"time"
)

func TestITime_MarshalXML(t *testing.T) {
	type x struct {
		A ITime `xml:"AA"`
	}
	x1:=&x{A: ITime{Time:time.Now()}}
	fmt.Println(xml.EncodeXmlToString(x1))
}
