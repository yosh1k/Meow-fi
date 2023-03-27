package controller

import (
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/repo"
)

type DealInteractor struct {
	DealRepository repo.DealRepository
}

func (interactor *DealInteractor) Add(t models.Deal) {
	interactor.DealRepository.Store(t)
}
func (interactor *DealInteractor) UpdateDeal(t models.Deal) {
	interactor.DealRepository.UpdateDeal(t)
}
func (interactor *DealInteractor) GetAllDeal() []models.Deal {
	return interactor.DealRepository.Select()
}
func (interactor *DealInteractor) GetDeal(id string) models.Deal {
	return interactor.NoticeRepository.SelectById(id)
}
func (interactor *DealInteractor) GetNoticeInfo(id string) models.Deal {
	deal := interactor.DealRepository.GetDealInfo(id)
	return deal
}
func (interactor *DealInteractor) Delete(id string) {
	interactor.DealRepository.Delete(id)
}
