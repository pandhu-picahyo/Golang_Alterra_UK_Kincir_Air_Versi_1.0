package controller

import (
	"net/http"
	"unjuk_keterampilan/config"
	"unjuk_keterampilan/model"

	"github.com/labstack/echo/v4"
)

func Hapus_Data_Kincir_Controller(c_Control echo.Context) error {
	id := c_Control.Param("id")

	var hitun model.Data_Kincir_Air
	result := config.DB.Where("id = ?", id).First(&hitun)

	if result.Error != nil {
		return c_Control.JSON(http.StatusNotFound, model.Basic_Response{
			Status:  false,
			Message: "Data tidak ditemukan",
			Data:    nil,
		})
	}

	result = config.DB.Delete(&hitun)

	if result.Error != nil {
		return c_Control.JSON(http.StatusInternalServerError, model.Basic_Response{
			Status:  false,
			Message: "Gagal menghapus data di database",
			Data:    nil,
		})
	}

	return c_Control.JSON(http.StatusOK, model.Basic_Response{
		Status:  true,
		Message: "Data berhasil dihapus",
		Data:    nil,
	})
}
