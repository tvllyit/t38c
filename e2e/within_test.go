package e2e

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
	"github.com/xjem/t38c"
)

func testWithin(t *testing.T, client *t38c.Client) {
	err := client.Keys.Set("points", "point-1").Point(1, 1).Do(context.Background())
	require.NoError(t, err)

	p := geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{
		{
			{0, 0},
			{20, 0},
			{20, 20},
			{0, 20},
			{0, 0},
		},
	})
	gj, _ := geojson.Encode(p)

	err = client.Keys.Set("areas", "area-1").Geometry(gj).Do(context.Background())
	require.NoError(t, err)

	resp, err := client.Search.Within("points").
		Get("areas", "area-1").
		Format(t38c.FormatIDs).Do(context.Background())
	require.NoError(t, err)

	require.Equal(t, []string{"point-1"}, resp.IDs)
}
