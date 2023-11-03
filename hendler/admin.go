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

// AdminRegister digunakan untuk mendaftarkan admin baru.
func (h *HandlerAdmin) AdminRegister(ctx echo.Context) error {
	// Bind request ke objek AdminCreateRequest
	AdminCreateRequest := modelsrequest.AdminCreateRequest{}
	err := ctx.Bind(&AdminCreateRequest)
	if err != nil {
		// Jika binding gagal, kirim respons dengan status 400 (Bad Request)
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	// Coba mendaftarkan admin baru
	result, err := h.Service.Create(ctx, AdminCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			// Jika validasi gagal, kirim respons dengan status 400 (Bad Request)
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}
		if strings.Contains(err.Error(), "email already exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}

	// Konversi hasil operasi ke respons admin dan kirim respons dengan status 201 (Created)
	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign Up", response))
}

// AdminLogin digunakan untuk proses login admin.
func (h *HandlerAdmin) AdminLogin(ctx echo.Context) error {
	// AdminLogin digunakan untuk proses login admin.
	AdminLoginRequest := modelsrequest.AdminLoginRequest{}
	err := ctx.Bind(&AdminLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Admin Login"))
	}
	
    // Coba proses login admin
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

	// Konversi hasil operasi ke respons login admin dan kirim respons dengan status 201 (Created)
	adminLoginResponse := res.AdminDomainToAdminLoginResponse(response)

	token, err := helper.GenerateAdminToken(&adminLoginResponse, uint(response.ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}
	// Generate token untuk admin dan sertakan dalam respons
	adminLoginResponse.Token = token

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign In", adminLoginResponse))
}

// AdminUpdate digunakan untuk memperbarui data admin.
func (h *HandlerAdmin) AdminUpdate(ctx echo.Context) error {
	// Ambil ID admin dari parameter URL
	AdminId := ctx.Param("id")
	AdminIdInt, err := strconv.Atoi(AdminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}
	// Bind request ke objek AdminUpdateRequest
	AdminUpdateRequest := modelsrequest.AdminUpdateRequest{}
	err = ctx.Bind(&AdminUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	// Coba memperbarui data admin berdasarkan ID
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

	// Konversi hasil operasi ke respons admin yang diperbarui dan kirim respons dengan status 201 (Created)
	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated Admin Data", response))
}

// AdminDelete digunakan untuk menghapus data admin.
func (h *HandlerAdmin) AdminDelete(ctx echo.Context) error {
	// Ambil ID admin dari parameter URL
	AdminId := ctx.Param("id")
	AdminIdInt, err := strconv.Atoi(AdminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	// Coba menghapus data admin berdasarkan ID
	err = h.Service.Delete(ctx, AdminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Admin Data Error"))
	}

	 // Kirim respons dengan status 201 (Created) setelah berhasil menghapus admin
	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Deleted Admin Data", nil))
}

// AdminGetById digunakan untuk mendapatkan data admin berdasarkan ID.
func (h *HandlerAdmin) AdminGetById(ctx echo.Context) error {
	// Ambil ID admin dari parameter URL
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	// Coba mendapatkan data admin berdasarkan ID
	result, err := h.Service.GetId(ctx, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}
		// Jika terjadi kesalahan lain, kirim respons dengan status 500 (Internal Server Error)
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Admin Data Error"))
	}

	 // Konversi hasil operasi ke respons admin dan kirim respons dengan status 201 (Created)
	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get Admin Data", response))
}

// AdminGetAll digunakan untuk mendapatkan semua data admin.
func (h *HandlerAdmin) AdminGetAll(ctx echo.Context) error {
	// Coba mendapatkan semua data admin
	result, err := h.Service.GetAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "admins not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admins Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Admins Data Error"))
	}

	// Konversi hasil operasi ke respons admin dan kirim respons dengan status 201 (Created)
	response := res.ConvertAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get All Admin Data", response))
}
