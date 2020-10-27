package main

import (
	"fmt"
	"github.com/iskone/panand/cache"
	"github.com/iskone/panand/lib"
	"github.com/kataras/iris/v12"
	"html/template"
	"strings"
	"time"
)

func newHttp(s *Server) {
	s.app = iris.Default()
	s.app.Get("/{p:path}", func(c iris.Context) {
		paths := strings.Split(c.Params().GetString("p"), "/")
		if paths[len(paths)-1] == "" {
			paths = paths[:len(paths)-1]
		}
		last := paths[len(paths)-1]
		pNode, ReqPath := cache.Get(paths[:len(paths)-1]...)
		var x = pathInfo{}
		x.C = c.FullRequestURI()
		x.Url = "//" + c.Host() + "/" + strings.Join(paths[:len(paths)-1], "/")
		x.IsEmpty = true
		find := true
		for _, v := range ReqPath {
			err := getDisk(s.PanAnd, pNode)
			if err != nil {
				c.StopWithError(503, err)
				c.EndRequest()
				return
			}
			if cn, ok := pNode.ChildNode[v]; !ok {
				find = false
				break
			} else {
				pNode = cn
			}
		}
		if find {

			if pNode.Ttl <= int(time.Now().Unix()) {
				err := getDisk(s.PanAnd, pNode)
				if err != nil {
					c.StopWithError(503, err)
					c.EndRequest()
					return
				}
			}
			fmt.Println(pNode)
			if node, ok := pNode.ChildNode[last]; ok {
				err := getDisk(s.PanAnd, node)
				if err != nil {
					c.StopWithError(503, err)
					c.EndRequest()
					return
				}
				x.IsEmpty = false
				for k, _ := range node.ChildNode {
					x.N = append(x.N, k)
				}
				for k, _ := range node.Files {
					x.F = append(x.F, k)
				}
			}
			if f, ok := pNode.Files[last]; ok {
				d, err := s.PanAnd.DownloadRequest(lib.DownloadRequest{
					ContentID: f.Id,
				})
				if err != nil {
					c.StopWithError(503, err)
					c.EndRequest()
					return
				}
				c.Redirect(d[0].Url, 302)
				c.EndRequest()
				return
			}
		}

		t, e := template.New("dir").Parse(dirTemp)
		if e != nil {
			c.StopWithError(503, e)
			c.EndRequest()
			return
		}
		e = t.Execute(c.ResponseWriter(), x)
		if e != nil {
			c.StopWithError(503, e)
			c.EndRequest()
			return
		}
	})
	s.app.Get("/favicon.ico", func(c iris.Context) {

	})
	s.app.Get("/", func(c iris.Context) {
		pNode, _ := cache.Get()
		if pNode.Ttl <= int(time.Now().Unix()) {
			err := getDisk(s.PanAnd, pNode)
			if err != nil {
				c.StopWithError(503, err)
				c.EndRequest()
				return
			}
		}
		var x pathInfo
		for k := range pNode.ChildNode {
			x.N = append(x.N, k)
		}
		for k := range pNode.Files {
			x.F = append(x.F, k)
		}
		t, e := template.New("dir").Parse(dirTemp)
		if e != nil {
			c.StopWithError(503, e)
			c.EndRequest()
			return
		}
		e = t.Execute(c.ResponseWriter(), x)
		if e != nil {
			c.StopWithError(503, e)
			c.EndRequest()
			return
		}
		return
	})
	s.app.HandleDir("/static", "./static")
	s.app.Listen(":8080")
}
func newApi(api iris.Party, s *Server) {
	api.Post("/GetDisk", func(ctx iris.Context) {
		id := ctx.URLParamDefault("id", "root")
		start := ctx.URLParamIntDefault("start", -1)
		end, err := ctx.URLParamInt("end")
		if err == nil {
			if start < 1 {
				start = 1
			}
		}
		_, _ = s.PanAnd.GetDisk(lib.GetDiskReq{
			CatalogID:       id,
			CatalogSortType: 1,
			StartNumber:     start,
			EndNumber:       end,
			CatalogType:     lib.CatalogTypeAll,
		})
	})
}
func getDisk(panand *lib.Panand, node *cache.Node) error {
	res, err := panand.GetDisk(lib.GetDiskReq{
		CatalogID:   node.Id,
		StartNumber: -1,
	})
	if err != nil {
		return err
	}
	for _, c := range res.CatalogList.CatalogInfo {
		node.Add(&cache.Node{
			Parent: node,
			Id:     c.CatalogID,
			Name:   c.CatalogName,
			File:   false,
		})
	}
	for _, f := range res.ContentList.ContentInfo {
		node.Add(&cache.Node{
			Parent: node,
			Id:     f.ContentID,
			Name:   f.ContentName,
			File:   true,
		})
	}
	node.Ttl = int(time.Now().Add(time.Minute).Unix())
	return nil
}
