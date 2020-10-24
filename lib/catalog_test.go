package lib

import (
	"fmt"
	"testing"
)

func TestPanand_GetDisk(t *testing.T) {
	p:=testnewPan()
	r,e:=p.GetDisk(GetDiskReq{
		CatalogID:           "root",
		EntryShareCatalogID: "",
		FilterType:          0,
		CatalogSortType:     0,
		ContentType:         0,
		ContentSortType:     0,
		SortDirection:       IBool{false},
		StartNumber:         -1,
		EndNumber:           0,
		ChannelList:         "",
		CatalogType:         0,
		Path:                "",
	})
	if e!=nil {
		t.Error(e)
	} else {
		t.Log(r)
	}
}
func TestPanand_CreatCatalogExt(t *testing.T) {
	p:=testnewPan()
	r,e:=p.CreatCatalogExt(CreateCatalogExtReq{
		ParentCatalogID: "0G11ZGK0d03Y00019700101000000001",
		NewCatalogName:  "test123",
		CatalogType:     CatalogTypeCommon,
		Path:            "",
		ManualRename:    0,
	})
	fmt.Println(r,e)
}
func TestPanand_CopyContentCatalog(t *testing.T) {
	p:=testnewPan()
	r,e:=p.CopyContentCatalog("0G11ZGK0d03Y027202010150005116fa",[]string{"0G11ZGK0d03Y026202010232129181kj"},nil)
	if e!=nil {
		t.Error(e)
	} else {
		t.Log(r)
	}
}
func TestPanand_UpdateCatalogInfo(t *testing.T) {
	p:=testnewPan()
	r,e:=p.UpdateCatalogInfo("0G11ZGK0d03Y027202010150005116fa","bbbbbTest",CatalogTypeCommon,"")
	if e!=nil {
		t.Error(e)
	} else {
		t.Log(r)
	}
}
func TestPanand_MoveContentCatalog(t *testing.T) {
	p:=testnewPan()
	r,e:=p.MoveContentCatalog("0G11ZGK0d03Y06820201023205044m5r",nil,[]string{"0G11ZGK0d03Y027202010150008036ul"})
	if e!=nil {
		t.Error(e)
	} else {
		t.Log(r)
	}
}
