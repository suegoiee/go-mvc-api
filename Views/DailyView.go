package Views

import (
	"ginPrac/Models"
)

type DailyView struct {
	Name         string         `json:"name" bson:"name"`
	Code         string         `json:"code" bson:"code"`
	OpeningPrice float32        `json:"oPrice" bson:"oPrice"`
	TradeInfo    []Models.Daily `json:"info" bson:"info"`
}

// func RenderDaily(daily *Models.Daily) *DailyView {
// 	v := &DailyView{
// 		Name: daily.Name,
// 		Code: daily.Code,
// 	}

// 	return v
// }
