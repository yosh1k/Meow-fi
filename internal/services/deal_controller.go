package services
	

	import (
		"Meow-fi/internal/database"
		"Meow-fi/internal/database/interfaces"
		"Meow-fi/internal/models"
		"Meow-fi/internal/services/usercase/controller"
	

		"github.com/labstack/echo"
	)
	

	type DealController struct {
		Interactor controller.DealInteractor
	}
	

	func NewDealController(sqlHandler interfaces.SqlHandler) *DealController {
		return &DealController{
			Interactor: controller.DealInteractor{
				DealRepository: &database.DealRepository{
					SqlHandler: sqlHandler,
				},
			},
		}
	}
	

	func (controller *DealController) Create(ctx echo.Context) {
		Deal := models.Deal{}
		ctx.Bind(&Deal)
		controller.Interactor.Add(Deal)
		createdDeals := controller.Interactor.GetAllDeals()
		ctx.JSON(201, createdDeals)
		return
	}
	func (controller *DealController) UpdateDeal(t models.Deal) {
	

	}
	func (controller *DealController) GetDeal(id string) models.Deal {
		Deal := controller.Interactor.GetDeal(id)
		return Deal
	}
	func (controller *DealController) GetDealInfo(id string) string {
		Deal := controller.Interactor.GetDealInfo(id)
		str := ""
		str += Deal.Client.FIO + " created Deal: " + Deal.Containing
		return str
	}
	func (controller *DealController) Delete(id string) {
		controller.Interactor.Delete(id)
	}
	func (controller *DealController) GetAllDeals() []models.Deal {
		res := controller.Interactor.GetAllDeals()
		return res
	}

