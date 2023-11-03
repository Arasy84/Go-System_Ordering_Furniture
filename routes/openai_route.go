package routes

import (

	hendler "furniture/chatbot/hendlers"
	usecase "furniture/chatbot/usecasee"

	"github.com/labstack/echo/v4"
)
func OpenAI(e *echo.Echo) {
	furnitureUsecase := usecase.NewFurnitureUsecase()
	furnitureHandler := hendler.NewFurnitureHandler(furnitureUsecase)
	e.POST("/recommendation", furnitureHandler.SubmitFurniture)
}

