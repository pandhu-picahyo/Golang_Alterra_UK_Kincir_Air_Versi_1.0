package controller

import (
	"net/http"
	"unjuk_keterampilan/config"
	"unjuk_keterampilan/model"

	"github.com/labstack/echo/v4"
)

func Buat_Table(c_Control echo.Context) error {

	var inputData model.Data_Kincir_Air
	if err := c_Control.Bind(&inputData); err != nil {
		return c_Control.JSON(http.StatusBadRequest, model.Basic_Response{
			Status:  false,
			Message: "Format data tidak valid",
			Data:    nil,
		})
	}

	// Validasi input data
	if inputData.Efisiensi_Kincir <= 0 || inputData.Tinggi_Jatuh_Air <= 0 || inputData.Debit_Air <= 0 ||
		inputData.Percepatan_Gravitasi <= 0 || inputData.Massa_Jenis <= 0 || inputData.Modal <= 0 ||
		inputData.Umur <= 0 || inputData.Nama == "" || inputData.Alamat == "" || inputData.Pekerjaan == "" {
		return c_Control.JSON(http.StatusBadRequest, model.Basic_Response{
			Status:  false,
			Message: "Data input tidak valid",
			Data:    nil,
		})
	}

	// Hitung nilai energi
	inputData.Energi = inputData.Efisiensi_Kincir * float32(inputData.Tinggi_Jatuh_Air) * inputData.Debit_Air * inputData.Percepatan_Gravitasi * float32(inputData.Massa_Jenis)

	result := config.DB.Create(&inputData)

	if result.Error != nil {
		return c_Control.JSON(http.StatusInternalServerError, model.Basic_Response{
			Status:  false,
			Message: "Gagal membuat tabel di database",
			Data:    nil,
		})
	}
	return c_Control.JSON(http.StatusCreated, model.Basic_Response{
		Status:  true,
		Message: "Sukses menambahkan tabel di database",
		Data:    inputData,
	})
}
