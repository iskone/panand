package lib

import (
	"os"
	"testing"
)

func TestPanand_WebUploadFileRequest(t *testing.T) {
	p := testnewPan()
	r := WebUploadFileRequest{
		FileCount: 2,
		TotalSize: 0,
		UploadContentList: UploadContentList{
			Length: 1,
			UploadContentInfo: []UploadContentInfo{
				{
					ContentName:    "aaaaaaaa",
					ContentSize:    0,
					ContentDesc:    "fsfsds",
					ContentTAGList: "sdsdsdsd",
					ComlexFlag:     IBool{},
					ComlexCID:      "",
					ResCID:         IDs{},
					Digest:         "",
				},
			},
		},
		NewCatalogName:  "",
		ParentCatalogID: "",
	}
	r2, e := p.WebUploadFileRequest(r)
	if e != nil {
		t.Error(e)
	} else {
		t.Log(r2)
	}
}

func TestPanand_DownloadRequest(t *testing.T) {
	p := testnewPan()
	//https://download.mcloud.139.com:443/storageWeb/servlet/pcDownloadFile?code=0G11ZGK0d03Y026202010232129181kj27417nBbiu9ju&contentID=0G11ZGK0d03Y026202010232129181kj&dom=D960&oprChannel=22000000&filename=MTMxMiAoMSkuanBn
	//https://download.mcloud.139.com:443/storageWeb/servlet/pcDownloadFile?code=0G11ZGK0d03Y027202010150023439h007917nBbitvPH&contentID=0G11ZGK0d03Y027202010150023439h0&filename=Ym9vdA==&dom=D960<
	f, e := p.DownloadRequest(DownloadRequest{
		Operation:   IBool{true},
		ContentID:   "0G11ZGK0d03Y026202010232129181kj",
		OwnerMSISDN: "13630428126",
		FileVersion: -1,
	})
	if e != nil {
		t.Error(e)
	} else {
		t.Log(f)
	}
}
func TestPanand_PcUploadFileRequest(t *testing.T) {
	p := testnewPan()
	fi, _ := os.Stat("user.go")
	r, e := p.PcUploadFileRequest("", []UploadContentInfo{
		{
			ContentName:     "user.go",
			ContentSize:     int(fi.Size()),
			ContentDesc:     "xxsdasdasdasd",
			ComlexFlag:      IBool{},
			ComlexCID:       "",
			ResCID:          IDs{},
			Digest:          "",
			UpdateContentID: "",
			FileEtag:        0,
			FileVersion:     0,
		},
	}, int(fi.Size()))
	if e != nil {
		t.Error(e)
	} else {
		t.Log(r)
	}
}
