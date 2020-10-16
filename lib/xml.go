package lib

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)
func encodeXml(i interface{}) (*bytes.Buffer,error) {
	buf:=&bytes.Buffer{}
	buf.WriteString(xml.Header)
	enc:=xml.NewEncoder(buf)
	enc.Indent("","\t")
	if e:=enc.Encode(i);e!=nil{
		return nil, e
	}
	return buf,nil
}
func encodeXmlToString(i interface{}) (string,error) {
	buf,e:=encodeXml(i)
	if e!=nil{
		return "", e
	}
	return buf.String(),nil
}
func decodeXml(r io.Reader,v interface{}) error {
	dec:=xml.NewDecoder(r)
	return dec.Decode(v)
}

func Xmld(s string)  {
	type m struct {
		Result
		ContentInfo ContentInfo `xml:"contentInfo"`
	}
	var sd Result
fmt.Println(	xml.Unmarshal([]byte(s),&sd))
fmt.Printf("%#v",sd)
}
