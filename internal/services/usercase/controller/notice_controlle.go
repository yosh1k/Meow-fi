package controller

import (
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/repo"
)

type NoticeInteractor struct {
	NoticeRepository repo.NoticeRepository
}

func (interactor *NoticeInteractor) Add(t models.Notice) {
	interactor.NoticeRepository.Store(t)
}
func (interactor *NoticeInteractor) UpdateNotice(t models.Notice) {
	interactor.NoticeRepository.UpdateNotice(t)
}
func (interactor *NoticeInteractor) GetAllNotices() []models.Notice {
	return interactor.NoticeRepository.Select()
}
func (interactor *NoticeInteractor) GetNotice(id string) models.Notice {
	return interactor.NoticeRepository.SelectById(id)
}
func (interactor *NoticeInteractor) GetNoticeInfo(id string) models.Notice {
	notice := interactor.NoticeRepository.GetNoticeInfo(id)
	return notice
}
func (interactor *NoticeInteractor) Delete(id string) {
	interactor.NoticeRepository.Delete(id)
}
