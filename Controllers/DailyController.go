package Controllers

import (
	"fmt"
	"ginPrac/Services"

	"github.com/gin-gonic/gin"
)

type DailyController struct {
	dailyService *Services.DailyService
}

func NewDailyController(dailyService *Services.DailyService) *DailyController {
	return &DailyController{
		dailyService: dailyService,
	}
}

func (dc *DailyController) PriceToday(context *gin.Context) {
	fmt.Println("price")
	// code := context.Param("code")
	// daily, err := dc.dailyService.GetDailyByCode(code)

	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// v := Views.RenderDaily(daily)
	// context.JSON(http.StatusOK, v)
}
