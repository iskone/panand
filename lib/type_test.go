package lib

import (
	"bytes"
	"fmt"
	"github.com/iskone/itools/xml"
	"testing"
	"time"
)

func TestITime_MarshalXML(t *testing.T) {
	type x struct {
		A ITime `xml:"AA"`
	}
	x1 := &x{A: ITime{Time: time.Now()}}
	fmt.Println(xml.EncodeXmlToString(x1))
}
func TestITime_UnmarshalXML(t *testing.T) {
	type x struct {
		A ITime `xml:"AA"`
	}
	s := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<x>\n\t<AA>20201021154738</AA>\n</x> ")
	var x1 x

	xml.DecodeXml(bytes.NewReader(s), &x1)
	fmt.Println(x1)
}
