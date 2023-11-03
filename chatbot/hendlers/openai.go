package hendler

import (
	"fmt"
	usecase "furniture/chatbot/usecasee"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type FurnitureHandler struct {
	FurnitureUsecase usecase.FurnitureUsecase
}

type FurnitureResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func NewFurnitureHandler(usecase usecase.FurnitureUsecase) *FurnitureHandler {
	return &FurnitureHandler{
		FurnitureUsecase: usecase,
	}
}

func (h *FurnitureHandler) SubmitFurniture(c echo.Context) error {
	var requestData map[string]interface{}
	err := c.Bind(&requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON format")
	}

	budget, okBudget := requestData["budget"].(float64)
	kategori, okKategori := requestData["kategori"].(string)
	material, okMaterial := requestData["material"].(string)

	if !okBudget || !okKategori || !okMaterial {
		return c.JSON(http.StatusBadRequest, "Invalid request format")
	}

	// Ganti pesan yang dibuat oleh user ke sesuatu yang relevan untuk pendaftaran magang
	furnitureDetails := fmt.Sprintf("Rekomendasi Furniture dengan kategori %s dengan spesifikasi %s dan dengan budget sebesar Rp %.0f.", kategori, material, budget)


	answer, err := h.FurnitureUsecase.SubmitFurniture(furnitureDetails, kategori, material, os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error dalam pengajuan pendaftaran magang")
	}

	responseData := FurnitureResponse{
		Status: "success",
		Data:   answer,
	}

	return c.JSON(http.StatusOK, responseData)
}