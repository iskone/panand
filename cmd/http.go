package cmd

import (
	"github.com/iskone/panand/lib"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func newHttp(s *Server) {
	s.app = iris.Default()
}
func newApi(api iris.Party, s *Server) {
	api.Post("/GetDisk", func(ctx context.Context) {
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
