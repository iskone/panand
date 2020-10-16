package lib

import (
	"encoding/xml"
	"fmt"
)

func Xmld(s string)  {
	type m struct {
		Result
		ContentInfo ContentInfo `xml:"contentInfo"`
	}
	var sd m
fmt.Println(	xml.Unmarshal([]byte(s),&sd))
fmt.Printf("%+v",sd.ContentInfo)
}
