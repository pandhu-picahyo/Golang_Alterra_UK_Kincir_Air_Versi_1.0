package model

type Data_Kincir_Air struct {
	Id                   int     `json:"id"`
	Nama                 string  `json:"nama"`
	Umur                 int     `json:"umur"`
	Alamat               string  `json:"alamat"`
	Pekerjaan            string  `json:"pekerjaan"`
	Modal                int     `json:"modal"`
	Efisiensi_Kincir     float32 `json:"efisiensi_kincir"`
	Tinggi_Jatuh_Air     int     `json:"tinggi_jatuh_air"`
	Debit_Air            float32 `json:"debit_air"`
	Percepatan_Gravitasi float32 `json:"percepatan_gravitasi"`
	Massa_Jenis          int     `json:"massa_jenis"`
	Energi               float32 `json:"energi"`
}
