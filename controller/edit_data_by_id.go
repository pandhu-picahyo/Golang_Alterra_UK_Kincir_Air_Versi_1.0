package controller

import (
	"net/http"
	"unjuk_keterampilan/config"
	"unjuk_keterampilan/model"

	"github.com/labstack/echo/v4"
)

func Edit_Data_Kincir_Controller(c_Control echo.Context) error {
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

	// Membuat instance baru untuk validasi data input
	newData := new(model.Data_Kincir_Air)
	if err := c_Control.Bind(newData); err != nil {
		return c_Control.JSON(http.StatusBadRequest, model.Basic_Response{
			Status:  false,
			Message: "Format data tidak valid",
			Data:    nil,
		})
	}

	// Validasi data input
	if newData.Efisiensi_Kincir <= 0 || newData.Tinggi_Jatuh_Air <= 0 || newData.Debit_Air <= 0 ||
		newData.Percepatan_Gravitasi <= 0 || newData.Massa_Jenis <= 0 || newData.Modal <= 0 ||
		newData.Umur <= 0 || newData.Nama == "" || newData.Alamat == "" || newData.Pekerjaan == "" {
		return c_Control.JSON(http.StatusBadRequest, model.Basic_Response{
			Status:  false,
			Message: "Data input tidak valid",
			Data:    nil,
		})
	}

	// Update data dengan nilai baru
	hitun.Nama = newData.Nama
	hitun.Umur = newData.Umur
	hitun.Alamat = newData.Alamat
	hitun.Pekerjaan = newData.Pekerjaan
	hitun.Modal = newData.Modal
	hitun.Efisiensi_Kincir = newData.Efisiensi_Kincir
	hitun.Tinggi_Jatuh_Air = newData.Tinggi_Jatuh_Air
	hitun.Debit_Air = newData.Debit_Air
	hitun.Percepatan_Gravitasi = newData.Percepatan_Gravitasi
	hitun.Massa_Jenis = newData.Massa_Jenis

	// Menghitung ulang energi berdasarkan nilai yang diperbarui
	hitun.Energi = hitun.Efisiensi_Kincir * float32(hitun.Tinggi_Jatuh_Air) * hitun.Debit_Air * hitun.Percepatan_Gravitasi * float32(hitun.Massa_Jenis)

	result = config.DB.Save(&hitun)

	if result.Error != nil {
		return c_Control.JSON(http.StatusInternalServerError, model.Basic_Response{
			Status:  false,
			Message: "Gagal memperbarui data di database",
			Data:    nil,
		})
	}

	return c_Control.JSON(http.StatusOK, model.Basic_Response{
		Status:  true,
		Message: "Data berhasil diperbarui",
		Data:    hitun,
	})
}
