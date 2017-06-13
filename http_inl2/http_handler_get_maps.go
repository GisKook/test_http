package http_inl2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	for i := 1; i < 5; i++ {
		maps = append(maps, &MapProperty{
			MapUri:  uint64(i),
			MapUrl:  "http://222.222.218.50:8886/zhanting/map.html",
			MapName: "地图" + strconv.Itoa(i),
		})
	}
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
