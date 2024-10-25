package domain

type Transaction struct {
	Urut      int
	CUSTID    int
	METODE    string
	TRXDATE   string
	NOREFF    string
	FIDBANK   string
	KDCHANNEL string
	DEBET     int
	KREDIT    int
	REFFBANK  string
	TRANSNO   string
}

type VSaldoVa struct {
	CUSTID int
	NOCUST string
	NMCUST string
	NUM2ND string
	STCUST string
	CODE01 string
	DESC01 string
	CODE02 string
	DESC02 string
	CODE03 string
	DESC03 string
	CODE04 string
	DESC04 string
	GENUS  string
	SALDO  int
}

type TransactionDetail struct {
	CUSTID               int
	NIS                  string
	Nama                 string
	NamaPanggil          string
	NoKK                 string
	NIK                  string
	NIKAyah              string
	NIKIbu               string
	VANO                 string
	NISN                 string
	GENUSPlace           string
	Birth                string
	SekolahAsal          string
	AnakKeDari           string
	Gender               string
	NamaAyah             string
	JobAyah              string
	ReligiAyah           string
	EduAyah              string
	AliveAyah            string
	PenghasilanAyah      string
	NamaIbu              string
	JobIbu               string
	ReligiIbu            string
	EduIbu               string
	AliveIbu             string
	PenghasilanIbu       string
	NamaWali             string
	JobWali              string
	ReligiWali           string
	PenghasilanWali      string
	AlamatWali           string
	Kemasyarakatan       string
	TB                   string
	BB                   string
	LK                   string
	Goldar               string
	Penyakit             string
	Makanan              string
	PenyakitKel          string
	Asuransi             string
	TanggunganKeluarga   string
	PenanggungBiaya      string
	DaerahTinggal        string
	StatusTinggal        string
	KondisiLantai        string
	Bahasa               string
	Prestasi             string
	UnitPendidikan       string
	Jurusan              string
	Email                string
	Jalan                string
	RT                   string
	Desa                 string
	Kelurahan            string
	Kecamatan            string
	Kabupaten            string
	Provinsi             string
	TanggunganKeluargaTS string
	EduWali              string
	ID                   string
}

type VSaldoVaCashless struct {
	CUSTID int
	NOCUST string
	NMCUST string
	NUM2ND string
	STCUST string
	CODE01 string
	CODE02 string
	DESC02 string
	CODE03 string
	DESC03 string
	CODE04 string
	DESC04 string
	GENUS  string
	SALDO  int
}
