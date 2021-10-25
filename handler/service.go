package handler

import "github.com/ionysoshedblom/go_scraper/service"

type MicroService struct {
	scraperService service.ScraperService
}

func NewMicroService(scraperService service.ScraperService) *MicroService {
	return &MicroService{
		scraperService: scraperService,
	}
}
