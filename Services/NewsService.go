package Services

import (
	"ginPrac/Models"
	"ginPrac/Repositories"
)

type NewsService struct {
	NewsRepository *Repositories.NewsRepository
}

func NewNewsService(NewsRepository *Repositories.NewsRepository) *NewsService {
	return &NewsService{
		NewsRepository: NewsRepository,
	}
}

func (ns *NewsService) GetNews(page, pageSize int) ([]*Models.News, error) {
	return ns.NewsRepository.Get(page, pageSize)
}
