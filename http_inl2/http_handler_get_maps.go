package http_inl2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MapList struct {
	MapList []*MapProperty `json:"MapList"`
}

type MapProperty struct {
	MapUri  uint64 `json:"MapUri"`
	MapUrl  string `json:"MapUrl"`
	MapName string `json:"MapName"`
}

func GetMaps() string {
	maps := make([]*MapProperty, 0)
	maps = append(maps, &MapProperty{
		MapUri:  4,
		MapUrl:  "http://222.222.218.51:8888/cgi-bin/hospital/qgis_mapserv.fcgi",
		MapName: "仁爱医院4",
	})
	maps = append(maps, &MapProperty{
		MapUri:  3,
		MapUrl:  "http://222.222.218.51:8888/cgi-bin/room/qgis_mapserv.fcgi",
		MapName: "展厅3",
	})
	maps = append(maps, &MapProperty{
		MapUri:  1,
		MapUrl:  "http://222.222.218.51:8888/cgi-bin/hospital/qgis_mapserv.fcgi",
		MapName: "展厅1",
	})
	maps = append(maps, &MapProperty{
		MapUri:  2,
		MapUrl:  "http://222.222.218.51:8888/cgi-bin/prison/qgis_mapserv.fcgi",
		MapName: "鹿泉监狱2",
	})
	map_list := &MapList{
		MapList: maps,
	}

	response, _ := json.Marshal(map_list)

	return string(response)
}

func GetMapsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprint(w, GetMaps())
}
