package hendler

import (
	"fmt"
	"furniture/helper"
	modelsrequest "furniture/models/models_request"
	"furniture/service"
	res "furniture/utils/respons"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	UserCreate(ctx echo.Context) error
	UserLogin(ctx echo.Context) error
	UserUpdate(ctx echo.Context) error
	UserDelete(ctx echo.Context) error
	UserGetById(ctx echo.Context) error
	UserGetAll(ctx echo.Context) error
}

type HandlerUser struct {
	Service service.UserService
}

func NewHandlerUser(userService service.UserService) HandlerUser {
	return HandlerUser{Service: userService}
}

func (h *HandlerUser) UserCreate(ctx echo.Context) error {
	userCreateRequest := modelsrequest.UserCreateRequest{}
	err := ctx.Bind(&userCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.Service.CreateUser(ctx, userCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		if strings.Contains(err.Error(), "email already exists") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exists"))
		}
		fmt.Println(result)
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}
	response := res.UserDomainToUserResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign Up", response))
}

func (h *HandlerUser) UserLogin(ctx echo.Context) error {
	UserLoginRequest := modelsrequest.UserLogin{}
	err := ctx.Bind(&UserLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid User Login"))
	}
	response, err := h.Service.Login(ctx, UserLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		if strings.Contains(err.Error(), "invalid email or password") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Email or Password"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign In Error"))
	}
	UserLogin := res.UserDomainToUserLoginResponse(response)

	token, err := helper.GenerateUserToken(&UserLogin, uint(response.ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	UserLogin.Token = token

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign In", UserLogin))
}

func (h *HandlerUser) UserUpdate(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	userUpdateRequest := modelsrequest.UserUpdate{}
	err = ctx.Bind(&userUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.Service.Update(ctx, userUpdateRequest, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update User Error"))
	}

	response := res.UserDomainToUserResponse(result)
	fmt.Print(result)
	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated User Data", response))
}

func (h *HandlerUser) UserDelete(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = h.Service.Delete(ctx, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete User Data Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Deleted User Data", nil))
}

func (h *HandlerUser) UserGetById(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := h.Service.GetId(ctx, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data Error"))
	}

	response := res.UserDomainToUserResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get User Data", response))
}

func (h *HandlerUser) UserGetAll(ctx echo.Context) error {
	result, err := h.Service.GetAll(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All User Data Error"))
	}

	response := res.ConvertUserResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get All User Data", response))
}
