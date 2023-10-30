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

type AdminHandler interface {
	AdminRegister(ctx echo.Context) error
	AdminLogin(ctx echo.Context) error
	AdminUpdate(ctx echo.Context) error
	AdminDelete(ctx echo.Context) error
	AdminGetById(ctx echo.Context) error
	AdminGetAll(ctx echo.Context) error
}

type HandlerAdmin struct {
	Service service.AdminService
}

func NewHandlerAdmin(adminService service.AdminService) HandlerAdmin {
	return HandlerAdmin{Service: adminService}
}

func (h *HandlerAdmin) AdminRegister(ctx echo.Context) error {
	AdminCreateRequest := modelsrequest.AdminCreateRequest{}
	err := ctx.Bind(&AdminCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.Service.Create(ctx, AdminCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}
		if strings.Contains(err.Error(), "email already exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign Up", response))
}

func (h *HandlerAdmin) AdminLogin(ctx echo.Context) error {
	AdminLoginRequest := modelsrequest.AdminLoginRequest{}
	err := ctx.Bind(&AdminLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Admin Login"))
	}
	response, err := h.Service.Login(ctx, AdminLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		if strings.Contains(err.Error(), "invalid email or password") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Email or Password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign In Error"))
	}

	adminLoginResponse := res.AdminDomainToAdminLoginResponse(response)

	token, err := helper.GenerateAdminToken(&adminLoginResponse, uint(response.ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	adminLoginResponse.Token = token

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign In", adminLoginResponse))
}

func (h *HandlerAdmin) AdminUpdate(ctx echo.Context) error {
	AdminId := ctx.Param("id")
	AdminIdInt, err := strconv.Atoi(AdminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}
	AdminUpdateRequest := modelsrequest.AdminUpdateRequest{}
	err = ctx.Bind(&AdminUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.Service.Update(ctx, AdminUpdateRequest, AdminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Admin Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated Admin Data", response))
}

func (h *HandlerAdmin) AdminDelete(ctx echo.Context) error {
	AdminId := ctx.Param("id")
	AdminIdInt, err := strconv.Atoi(AdminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = h.Service.Delete(ctx, AdminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Admin Data Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Deleted Admin Data", nil))
}

func (h *HandlerAdmin) AdminGetById(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := h.Service.GetId(ctx, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Admin Data Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get Admin Data", response))
}

func (h *HandlerAdmin) AdminGetAll(ctx echo.Context) error {
	result, err := h.Service.GetAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "admins not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admins Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Admins Data Error"))
	}

	response := res.ConvertAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get All Admin Data", response))
}
