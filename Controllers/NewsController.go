package Controllers

import (
	"ginPrac/Services"
	"ginPrac/Views"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NewsController struct {
	newsService *Services.NewsService
}

func NewNewsController(newsService *Services.NewsService) *NewsController {
	return &NewsController{
		newsService: newsService,
	}
}

func (dc *NewsController) GetNews(context *gin.Context) {
	request := context.Request.URL.Query()
	page, _ := strconv.Atoi(request["page"][0])
	pageSize, _ := strconv.Atoi(request["pageSize"][0])
	news, err := dc.newsService.GetNews(page, pageSize)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	v := Views.RenderNews(news)
	context.JSON(http.StatusOK, v)
}
