package Views

import (
	"ginPrac/Models"
	"time"
)

type DailyView struct {
	Name         string  `json:"name" bson:"name"`
	Code         string  `json:"code" bson:"code"`
	OpeningPrice float32 `json:"oPrice" bson:"oPrice"`
	TradeInfo    []Daily `json:"info" bson:"info"`
}

type Daily struct {
	Price     float32   `json:"price" bson:"price"`
	Amount    int32     `json:"amount" bson:"amount"`
	TradeTime time.Time `json:"time" bson:"time"`
}

func RenderDaily(daily *Models.Daily) *DailyView {
	v := &DailyView{
		Name: daily.Name,
		Code: daily.Code,
	}

	return v
}
