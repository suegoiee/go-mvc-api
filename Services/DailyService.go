package Services

import (
	"ginPrac/Models"
	"ginPrac/Repositories"
)

type DailyService struct {
	dailyRepository *Repositories.DailyRepository
}

func NewDailyService(dailyRepository *Repositories.DailyRepository) *DailyService {
	return &DailyService{
		dailyRepository: dailyRepository,
	}
}

func (us *DailyService) GetDailyByCode(code string) (*Models.Daily, error) {
	return us.dailyRepository.GetDaily(code)
}
