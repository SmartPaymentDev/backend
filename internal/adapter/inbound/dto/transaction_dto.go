package dto

import "smart-payment-dev-be/internal/core/domain"

type TransactionDTO struct{}

type Transaction struct {
	Urut      int    `json:"urut"`
	CUSTID    int    `json:"custid"`
	METODE    string `json:"metode"`
	TRXDATE   string `json:"trxdate"`
	NOREFF    string `json:"noreff"`
	FIDBANK   string `json:"fidbank"`
	KDCHANNEL string `json:"kdchannel"`
	DEBET     int    `json:"debet"`
	KREDIT    int    `json:"kredit"`
	REFFBANK  string `json:"reffbank"`
	TRANSNO   string `json:"transno"`
}

type VSaldoVa struct {
	CUSTID int    `json:"custid"`
	NOCUST string `json:"nocust"`
	NMCUST string `json:"nmcust"`
	NUM2ND string `json:"num_2_nd"`
	STCUST string `json:"stcust"`
	CODE01 string `json:"code_01"`
	DESC01 string `json:"desc_01"`
	CODE02 string `json:"code_02"`
	DESC02 string `json:"desc_02"`
	CODE03 string `json:"code_03"`
	DESC03 string `json:"desc_03"`
	CODE04 string `json:"code_04"`
	DESC04 string `json:"desc_04"`
	GENUS  string `json:"genus"`
	SALDO  int    `json:"saldo"`
}

type TransactionDetail struct {
	CUSTID               int    `json:"custid"`
	NIS                  string `json:"nis"`
	Nama                 string `json:"nama"`
	NamaPanggil          string `json:"nama_panggil"`
	NoKK                 string `json:"no_kk"`
	NIK                  string `json:"nik"`
	NIKAyah              string `json:"nik_ayah"`
	NIKIbu               string `json:"nik_ibu"`
	VANO                 string `json:"vano"`
	NISN                 string `json:"nisn"`
	GENUSPlace           string `json:"genus_place"`
	Birth                string `json:"birth"`
	SekolahAsal          string `json:"sekolah_asal"`
	AnakKeDari           string `json:"anak_ke_dari"`
	Gender               string `json:"gender"`
	NamaAyah             string `json:"nama_ayah"`
	JobAyah              string `json:"job_ayah"`
	ReligiAyah           string `json:"religi_ayah"`
	EduAyah              string `json:"edu_ayah"`
	AliveAyah            string `json:"alive_ayah"`
	PenghasilanAyah      string `json:"penghasilan_ayah"`
	NamaIbu              string `json:"nama_ibu"`
	JobIbu               string `json:"job_ibu"`
	ReligiIbu            string `json:"religi_ibu"`
	EduIbu               string `json:"edu_ibu"`
	AliveIbu             string `json:"alive_ibu"`
	PenghasilanIbu       string `json:"penghasilan_ibu"`
	NamaWali             string `json:"nama_wali"`
	JobWali              string `json:"job_wali"`
	ReligiWali           string `json:"religi_wali"`
	PenghasilanWali      string `json:"penghasilan_wali"`
	AlamatWali           string `json:"alamat_wali"`
	Kemasyarakatan       string `json:"kemasyarakatan"`
	TB                   string `json:"tb"`
	BB                   string `json:"bb"`
	LK                   string `json:"lk"`
	Goldar               string `json:"goldar"`
	Penyakit             string `json:"penyakit"`
	Makanan              string `json:"makanan"`
	PenyakitKel          string `json:"penyakit_kel"`
	Asuransi             string `json:"asuransi"`
	TanggunganKeluarga   string `json:"tanggungan_keluarga"`
	PenanggungBiaya      string `json:"penanggung_biaya"`
	DaerahTinggal        string `json:"daerah_tinggal"`
	StatusTinggal        string `json:"status_tinggal"`
	KondisiLantai        string `json:"kondisi_lantai"`
	Bahasa               string `json:"bahasa"`
	Prestasi             string `json:"prestasi"`
	UnitPendidikan       string `json:"unit_pendidikan"`
	Jurusan              string `json:"jurusan"`
	Email                string `json:"email"`
	Jalan                string `json:"jalan"`
	RT                   string `json:"rt"`
	Desa                 string `json:"desa"`
	Kelurahan            string `json:"kelurahan"`
	Kecamatan            string `json:"kecamatan"`
	Kabupaten            string `json:"kabupaten"`
	Provinsi             string `json:"provinsi"`
	TanggunganKeluargaTS string `json:"tanggungan_keluarga_ts"`
	EduWali              string `json:"edu_wali"`
	ID                   string `json:"id"`
}

type VSaldoVaCashless struct {
	CUSTID int    `json:"custid"`
	NOCUST string `json:"nocust"`
	NMCUST string `json:"nmcust"`
	NUM2ND string `json:"num_2_nd"`
	STCUST string `json:"stcust"`
	CODE01 string `json:"code_01"`
	CODE02 string `json:"code_02"`
	DESC02 string `json:"desc_02"`
	CODE03 string `json:"code_03"`
	DESC03 string `json:"desc_03"`
	CODE04 string `json:"code_04"`
	DESC04 string `json:"desc_04"`
	GENUS  string `json:"genus"`
	SALDO  int    `json:"saldo"`
}

type FilterTransaction struct {
	Token   string
	CustId  string
	Page    int
	PerPage int
}

type FilterTransactionDetail struct {
	Token   string
	CustId  string
	Page    int
	PerPage int
}

type TransactionRequest struct {
	NOREFF    string `json:"noreff"`
	FIDBANK   string `json:"fidbank"`
	KDCHANNEL string `json:"kdchannel"`
	KREDIT    int    `json:"kredit"`
	REFFBANK  string `json:"reffbank"`
	TRANSNO   string `json:"transno"`
}

type VSaldoVaResponse struct {
	VSaldoVa VSaldoVa `json:"v_saldo_va"`
}

type VSaldoVaCashlessResponse struct {
	VSaldoVaCashless VSaldoVaCashless `json:"v_saldo_va_cashless"`
}

type TransactionDetailResponses struct {
	TransactionDetails []TransactionDetail `json:"transaction_details"`
	Total              int                 `json:"total"`
}

type TransactionResponses struct {
	Transactions []Transaction `json:"transactions"`
	Total        int           `json:"total"`
}

func (t *TransactionDTO) ReqTransformIn(req TransactionRequest) domain.Transaction {
	return domain.Transaction{
		NOREFF:    req.NOREFF,
		FIDBANK:   req.FIDBANK,
		KDCHANNEL: req.KDCHANNEL,
		KREDIT:    req.KREDIT,
		REFFBANK:  req.REFFBANK,
		TRANSNO:   req.TRANSNO,
	}
}

func (t *TransactionDTO) GetTransactionsTransformOut(res []domain.Transaction, total int) interface{} {
	var list []Transaction

	for _, data := range res {
		list = append(list, Transaction{
			Urut:      data.Urut,
			CUSTID:    data.CUSTID,
			METODE:    data.METODE,
			TRXDATE:   data.TRXDATE,
			NOREFF:    data.NOREFF,
			FIDBANK:   data.FIDBANK,
			KDCHANNEL: data.KDCHANNEL,
			DEBET:     data.DEBET,
			KREDIT:    data.KREDIT,
			REFFBANK:  data.REFFBANK,
			TRANSNO:   data.TRANSNO,
		})
	}

	return TransactionResponses{
		Transactions: list,
		Total:        total,
	}
}

func (t *TransactionDTO) GetVSaldoVaTransformOut(res domain.VSaldoVa) interface{} {
	return VSaldoVaResponse{
		VSaldoVa: VSaldoVa{
			CUSTID: res.CUSTID,
			NOCUST: res.NOCUST,
			NMCUST: res.NMCUST,
			NUM2ND: res.NUM2ND,
			STCUST: res.STCUST,
			CODE01: res.CODE01,
			DESC01: res.DESC01,
			CODE02: res.CODE02,
			DESC02: res.DESC02,
			CODE03: res.CODE03,
			DESC03: res.DESC03,
			CODE04: res.CODE04,
			DESC04: res.DESC04,
			GENUS:  res.GENUS,
			SALDO:  res.SALDO,
		},
	}
}

func (t *TransactionDTO) GetTransactionDetailsTransformOut(res []domain.TransactionDetail, total int) interface{} {
	var list []TransactionDetail

	for _, data := range res {
		list = append(list, TransactionDetail{
			CUSTID:               data.CUSTID,
			NIS:                  data.NIS,
			Nama:                 data.Nama,
			NamaPanggil:          data.NamaPanggil,
			NoKK:                 data.Email,
			NIK:                  data.NIK,
			NIKAyah:              data.NIKAyah,
			NIKIbu:               data.NIKIbu,
			VANO:                 data.VANO,
			NISN:                 data.NISN,
			GENUSPlace:           data.GENUSPlace,
			Birth:                data.Birth,
			SekolahAsal:          data.SekolahAsal,
			AnakKeDari:           data.AnakKeDari,
			Gender:               data.Gender,
			NamaAyah:             data.NamaAyah,
			JobAyah:              data.JobAyah,
			ReligiAyah:           data.ReligiAyah,
			EduAyah:              data.EduAyah,
			AliveAyah:            data.AliveAyah,
			PenghasilanAyah:      data.PenghasilanAyah,
			NamaIbu:              data.NamaIbu,
			JobIbu:               data.JobIbu,
			ReligiIbu:            data.ReligiIbu,
			EduIbu:               data.EduIbu,
			AliveIbu:             data.AliveIbu,
			PenghasilanIbu:       data.PenghasilanIbu,
			NamaWali:             data.NamaWali,
			JobWali:              data.JobWali,
			ReligiWali:           data.ReligiWali,
			PenghasilanWali:      data.PenghasilanWali,
			AlamatWali:           data.AlamatWali,
			Kemasyarakatan:       data.Kemasyarakatan,
			TB:                   data.TB,
			BB:                   data.BB,
			LK:                   data.LK,
			Goldar:               data.Goldar,
			Penyakit:             data.Penyakit,
			Makanan:              data.Makanan,
			PenyakitKel:          data.PenyakitKel,
			Asuransi:             data.Asuransi,
			TanggunganKeluarga:   data.TanggunganKeluarga,
			PenanggungBiaya:      data.PenanggungBiaya,
			DaerahTinggal:        data.DaerahTinggal,
			StatusTinggal:        data.StatusTinggal,
			KondisiLantai:        data.KondisiLantai,
			Bahasa:               data.Bahasa,
			Prestasi:             data.Prestasi,
			UnitPendidikan:       data.UnitPendidikan,
			Jurusan:              data.Jurusan,
			Email:                data.Email,
			Jalan:                data.Jalan,
			RT:                   data.Email,
			Desa:                 data.Desa,
			Kelurahan:            data.Kelurahan,
			Kecamatan:            data.Kecamatan,
			Kabupaten:            data.Kabupaten,
			Provinsi:             data.Provinsi,
			TanggunganKeluargaTS: data.TanggunganKeluargaTS,
			EduWali:              data.EduWali,
			ID:                   data.ID,
		})
	}

	return TransactionDetailResponses{
		TransactionDetails: list,
		Total:              total,
	}
}

func (t *TransactionDTO) GetVSaldoVaCashlessTransformOut(res domain.VSaldoVaCashless) interface{} {
	return VSaldoVaCashlessResponse{
		VSaldoVaCashless: VSaldoVaCashless{
			CUSTID: res.CUSTID,
			NOCUST: res.NOCUST,
			NMCUST: res.NMCUST,
			NUM2ND: res.NUM2ND,
			STCUST: res.STCUST,
			CODE01: res.CODE01,
			CODE02: res.CODE02,
			DESC02: res.DESC02,
			CODE03: res.CODE03,
			DESC03: res.DESC03,
			CODE04: res.CODE04,
			DESC04: res.DESC04,
			GENUS:  res.GENUS,
			SALDO:  res.SALDO,
		},
	}
}
