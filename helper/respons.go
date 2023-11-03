package helper

// TResponseMeta adalah struktur data yang digunakan untuk menyimpan meta informasi respons.
type TResponseMeta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// TSuccessResponse adalah struktur data yang digunakan untuk menyusun respons berhasil dengan data.
type TSuccessResponse struct {
	Meta    TResponseMeta `json:"meta"`
	Results interface{}   `json:"results"`
}

// TErrorResponse adalah struktur data yang digunakan untuk menyusun respons gagal.
type TErrorResponse struct {
	Meta TResponseMeta `json:"meta"`
}

// SuccessResponse adalah fungsi yang digunakan untuk menghasilkan respons JSON berhasil.
// Parameter message adalah pesan atau informasi tambahan tentang respons.
// Parameter data adalah data yang ingin dimasukkan dalam respons. Ini bisa berupa tipe data apa pun.
// Fungsi ini mengembalikan respons JSON yang berisi informasi meta dan data jika data tidak nil, atau respons gagal jika data nil.
func SuccessResponse(message string, data interface{}) interface{} {
	if data == nil {
		return TErrorResponse{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
		}
	} else {
		return TSuccessResponse{
			Meta: TResponseMeta{
				Success: true,
				Message: message,
			},
			Results: data,
		}
	}
}

func ErrorResponse(message string) interface{} {
	return TErrorResponse{
		Meta: TResponseMeta{
			Success: false,
			Message: message,
		},
	}
}
// ErrorResponse adalah fungsi yang digunakan untuk menghasilkan respons JSON gagal.
// Parameter message adalah pesan atau informasi tentang alasan kegagalan respons.
// Fungsi ini mengembalikan respons JSON yang berisi informasi meta dengan status kegagalan.

func FormatResponse(message string, data any) map[string]any {
	var response = map[string]any{}
	response["message"] = message
	if data != nil {
		response["data"] = data
	}
	return response
}