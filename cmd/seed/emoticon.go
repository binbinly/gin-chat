package seed

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"gin-chat/internal/model"
	"gin-chat/pkg/client/http"
	"gin-chat/pkg/dbs"
)

var url = "https://raw.githubusercontent.com/zhaoolee/ChineseBQB/master/chinesebqb_github.json"

type result struct {
	Status int        `json:"status"`
	Info   string     `json:"info"`
	Data   []dataList `json:"data"`
}

type dataList struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Url      string `json:"url"`
}

// SyncBQB 同步 https://github.com/zhaoolee/ChineseBQB 表情包入库
func SyncBQB() error {
	client := http.NewRestyClient()
	rsp, err := client.Get(context.Background(), url)
	if err != nil {
		return err
	}
	var rs result
	if err = json.Unmarshal(rsp, &rs); err != nil {
		return err
	}
	if rs.Status != 1000 {
		return errors.New(rs.Info)
	}

	var emo []model.EmoticonModel
	for i, datum := range rs.Data {
		nameStart := strings.LastIndex(datum.Name, "-")
		nameEnd := strings.LastIndex(datum.Name, ".")
		if nameEnd == -1 {
			nameEnd = len(datum.Name)
		}
		if nameStart > nameEnd {
			nameStart = 0
		}
		cat := datum.Category[strings.LastIndex(datum.Category, "_")+1:]
		if cat == "BQB" {
			cat = datum.Category[strings.Index(datum.Category, "_")+1:]
		}
		cat = strings.Replace(cat, "BQB", "", 1)

		emo = append(emo, model.EmoticonModel{
			PriID:    model.PriID{ID: i + 1},
			Category: cat,
			Name:     datum.Name[nameStart+1 : nameEnd],
			Url:      datum.Url,
		})
	}
	//先清空表
	dbs.DB.Exec("truncate emoticon")
	return dbs.DB.CreateInBatches(emo, 200).Error
}
