package hendler

import (
	"furniture/helper"
	modelsrequest "furniture/models/models_request"
	"furniture/service"
	res "furniture/utils/respons"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type OrderHendler interface {
	CreateOrder(h echo.Context) error
	FindAll(h echo.Context) error
	FindByID(h echo.Context) error
	Delete(h echo.Context) error
}

type OrderHendlerImpl struct {
	OrderService service.OrderService
}

func NewOrderHendler(orderService service.OrderService) OrderHendler {
    return &OrderHendlerImpl{OrderService: orderService}
}

func (h *OrderHendlerImpl) CreateOrder(ctx echo.Context) error {
	createOrder := modelsrequest.CreateOrder {}
	err := ctx.Bind(&createOrder)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.OrderService.CreateOrder(ctx, createOrder)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Order Error"))
	}

	response := res.CreateOrderToOrderResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Order Data", response))
}

func (h *OrderHendlerImpl) FindByID(ctx echo.Context) error {
	OrderId := ctx.Param("id")
	OrderIdInt, err := strconv.Atoi(OrderId)
	if err!= nil {
        return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
    }

	result, err := h.OrderService.FindByID(ctx, OrderIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
            return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Find Order Error"))
	}
	response := res.CreateOrderToOrderResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Find Order Data", response))
}

func (h *OrderHendlerImpl) FindAll(ctx echo.Context) error {
	result, err := h.OrderService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "order not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Order Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Find Order Error"))
	}

	response := res.ConvertOrderResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Find Order Data", response))
}

func (h *OrderHendlerImpl) Delete(ctx echo.Context) error {
	OrderId := ctx.Param("id")
	OrderIdInt, err := strconv.Atoi(OrderId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Client Input"))
	}

	err = h.OrderService.Delete(ctx, OrderIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Order Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Order Data Error"))
	}
	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Deleted Order Data", nil))
}