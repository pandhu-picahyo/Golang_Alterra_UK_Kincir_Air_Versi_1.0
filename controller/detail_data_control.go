package controller

import (
	"net/http"
	"unjuk_keterampilan/config"
	"unjuk_keterampilan/model"

	"github.com/labstack/echo/v4"
)

func Detail_Data_Kincir_Controller(c_Control echo.Context) error {

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

	hitun.Energi = hitun.Efisiensi_Kincir * float32(hitun.Tinggi_Jatuh_Air) * hitun.Debit_Air * hitun.Percepatan_Gravitasi * float32(hitun.Massa_Jenis)

	return c_Control.JSON(http.StatusOK, model.Basic_Response{
		Status:  true,
		Message: "Data Detail Berhasil Ditemukan",
		Data:    hitun,
	})
}
