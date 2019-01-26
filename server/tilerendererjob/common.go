package tilerendererjob

import (
	"mapserver/app"
	"mapserver/coords"
	"mapserver/mapblockparser"
	"strconv"

	"github.com/sirupsen/logrus"
)

func getTileKey(tc *coords.TileCoords) string {
	return strconv.Itoa(tc.X) + "/" + strconv.Itoa(tc.Y) + "/" + strconv.Itoa(tc.Zoom)
}

func renderMapblocks(ctx *app.App, jobs chan *coords.TileCoords, mblist []*mapblockparser.MapBlock) int {
	tileRenderedMap := make(map[string]bool)
	tilecount := 0

	for i := 12; i >= 1; i-- {
		for _, mb := range mblist {
			//13
			tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)

			//12-1
			tc = tc.ZoomOut(13 - i)

			key := getTileKey(tc)

			if tileRenderedMap[key] {
				continue
			}

			tileRenderedMap[key] = true

			fields := logrus.Fields{
				"X":       tc.X,
				"Y":       tc.Y,
				"Zoom":    tc.Zoom,
				"LayerId": tc.LayerId,
			}
			logrus.WithFields(fields).Debug("Tile removing (z12-1)")

			//remove tile
			ctx.Objectdb.RemoveTile(tc)
		}
	}

	tileRenderedMap = make(map[string]bool)

	for i := 12; i >= 1; i-- {
		for _, mb := range mblist {
			//13
			tc := coords.GetTileCoordsFromMapBlock(mb.Pos, ctx.Config.Layers)

			//12-1
			tc = tc.ZoomOut(13 - i)

			key := getTileKey(tc)

			if tileRenderedMap[key] {
				continue
			}

			tileRenderedMap[key] = true

			fields := logrus.Fields{
				"X":       tc.X,
				"Y":       tc.Y,
				"Zoom":    tc.Zoom,
				"LayerId": tc.LayerId,
			}
			logrus.WithFields(fields).Debug("Dispatching tile rendering (z12-1)")

			tilecount++

			//dispatch re-render
			jobs <- tc
		}
	}

	return tilecount
}
