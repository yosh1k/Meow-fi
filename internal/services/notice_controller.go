package services

import (
	"Meow-fi/internal/database"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/controller"

	"github.com/labstack/echo"
)

type NoticeController struct {
	Interactor controller.NoticeInteractor
}

func NewNoticeController(sqlHandler interfaces.SqlHandler) *NoticeController {
	return &NoticeController{
		Interactor: controller.NoticeInteractor{
			NoticeRepository: &database.NoticeRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *NoticeController) Create(ctx echo.Context) {
	notice := models.Notice{}
	ctx.Bind(&notice)
	controller.Interactor.Add(notice)
	createdNotices := controller.Interactor.GetAllNotices()
	ctx.JSON(201, createdNotices)
	return
}
func (controller *NoticeController) UpdateNotice(t models.Notice) {

}
func (controller *NoticeController) GetNotice(id string) models.Notice {
	notice := controller.Interactor.GetNotice(id)
	return notice
}
func (controller *NoticeController) GetNoticeInfo(id string) string {
	notice := controller.Interactor.GetNoticeInfo(id)
	str := ""
	str += notice.Client.FIO + " created notice: " + notice.Containing
	return str
}
func (controller *NoticeController) Delete(id string) {
	controller.Interactor.Delete(id)
}
func (controller *NoticeController) GetAllNotices() []models.Notice {
	res := controller.Interactor.GetAllNotices()
	return res
}
