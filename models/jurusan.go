package models

type DataJurusan struct {
	NamaJ string
	Desc  string
	Foto  string
	Warna string
}


var IsiJurusan = []DataJurusan{
	{"PPLG", "Pengembangan Perangkat Lunak dan Gim", "https://web-smkn4.vercel.app/_nuxt/pplg.DhbRSnK2.png", ""},
	{"TJKT", "Teknik Jaringan Komputer dan Telekomunikasi", "https://web-smkn4.vercel.app/_nuxt/tjkt.XeO8V4_I.png", ""},
	{"TBSM", "Teknik Sepeda Motor", "https://web-smkn4.vercel.app/_nuxt/tbsm.5YZNeyvV.png", ""},
	{"DKV", "Desain Komunikasi Visual", "https://web-smkn4.vercel.app/_nuxt/dkv.Btdc-HP2.png", ""},
	{"TOI", "Teknik Otomasi Industri", "https://web-smkn4.vercel.app/_nuxt/toi.BFD6ZBmB.png", ""},
}