package controller

import (
	"net/http"
	"unjuk_keterampilan/config"
	"unjuk_keterampilan/model"

	"github.com/labstack/echo/v4"
)

func Dapat_Data_Kincir_Control(c_Control echo.Context) error {

	var dapat_kincir []model.Data_Kincir_Air

	result := config.DB.Find(&dapat_kincir)

	if result.Error != nil {
		return c_Control.JSON(http.StatusInternalServerError, model.Basic_Response{
			Status:  false,
			Message: "Gagal mendapatkan data dari database",
			Data:    nil,
		})
	}
	return c_Control.JSON(http.StatusOK, model.Basic_Response{
		Status:  true,
		Message: "Berhasil mendapatkan data dari database",
		Data:    dapat_kincir,
	})
}
