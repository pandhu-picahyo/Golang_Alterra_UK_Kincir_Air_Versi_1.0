package route

import (
	"unjuk_keterampilan/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Inisiasi_Rute(e *echo.Echo) *echo.Echo {

	// untuk memunculkan data logger
	e.Use(middleware.Logger())

	// Endpoint untuk mengambil data semua kincir air
	e.GET("/kincirs", controller.Dapat_Data_Kincir_Control)

	// Endpoint untuk menambahkan data kincir air baru
	e.POST("/kincirs", controller.Buat_Table)

	// Endpoint untuk mendapatkan detail data kincir air berdasarkan ID
	e.GET("/kincirs/:id", controller.Detail_Data_Kincir_Controller)

	// Endpoint untuk menghitung energi berdasarkan ID
	e.GET("/kincirs/:id/energi", controller.Hitung_Energi_Controller)

	// Endpoint untuk mengupdate data kincir air berdasarkan ID
	e.PUT("/kincirs/:id", controller.Update_Data_kincir_Controller)

	// Endpoint untuk menghapus data kincir air berdasarkan ID
	e.DELETE("/kincirs/:id", controller.Hapus_Data_Kincir_Controller)

	// Endpoint untuk mengambil data kincir air dengan modal di atas suatu nilai tertentu
	e.GET("/kincirs/modal/:value", controller.Dapat_Data_Kincir_Dengan_Modal_Atas)

	// Endpoint untuk mengambil data kincir dengan umur di bawah nilai tertentu
	e.GET("/kincirs/umur/:value", controller.Dapat_Data_Kincir_Dengan_Umur_Bawah)

	// Endpoint untuk mengambil data kincir dengan energi di atas nilai tertentu
	e.GET("/kincirs/energi/:value", controller.Dapat_Data_Kincir_Dengan_Energi_Atas)

	// Endpoint mengambil data kincir dengan umur dan modal di kisaran angka tertentu
	e.GET("/kincirs/umur-modal/:umur/:modal", controller.Dapat_Data_Kincir_Dengan_Umur_Diatas_Modal_Dibawah)

	// Endpoint untuk mengambil data kincir dengan umur dan energi di rentang nilai tertentu
	e.GET("/kincirs/umur-energi/:umur/:energi", controller.Dapat_Data_Kincir_Dengan_Umur_Diatas_Energi_Diatas)

	// Endpoint untuk mengedit data kincir berdasarkan ID
	e.PUT("/kincirs/:id", controller.Edit_Data_Kincir_Controller)

	return e
}
