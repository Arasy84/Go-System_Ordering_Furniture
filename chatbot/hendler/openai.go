package hendler

// import (
// 	"fmt"
// 	"furniture/chatbot/usecase"
// 	"net/http"
// 	"os"

// 	"github.com/labstack/echo/v4"
// )

// type FurnitureHandler struct {
// 	FurnitureUsecase usecase.FurnitureUsecase
// }

// type FurnitureResponse struct {
// 	Status string `json:"status"`
// 	Data   string `json:"data"`
// }

// func NewFurnitureHandler(usecase usecase.FurnitureUsecase) *FurnitureHandler {
// 	return &FurnitureHandler{
// 		FurnitureUsecase: usecase,
// 	}
// }

// func (h *FurnitureHandler) SubmitFurniture(c echo.Context) error {
// 	var requestData map[string]interface{}
// 	err := c.Bind(&requestData)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, "Invalid JSON format")
// 	}

// 	name, okName := requestData["name"].(string)
// 	email, okEmail := requestData["email"].(string)
// 	userInput, okUserInput := requestData["userInput"].(string)

// 	if !okName || !okEmail || !okUserInput {
// 		return c.JSON(http.StatusBadRequest, "Invalid request format")
// 	}

// 	// Ganti pesan yang dibuat oleh user ke sesuatu yang relevan untuk pendaftaran magang
// 	furnuitureDetails := fmt.Sprintf("Pendaftaran furniture oleh %s (email: %s) - Furniture di%s Hanya memerlukan identitas diri kamu dan Kuota, Karena nantinya kamu akan mengisi formulir yang telah disediakan oleh pt kami.", name, email, userInput)

// 	answer, err := h.FurnitureUsecase.SubmitFurniture(furnuitureDetails, name, email, os.Getenv("OPENAI_API_KEY"))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, "Error dalam pengajuan pendaftaran magang")
// 	}

// 	responseData := FurnitureResponse{
// 		Status: "success",
// 		Data:   answer,
// 	}

// 	return c.JSON(http.StatusOK, responseData)
// }
