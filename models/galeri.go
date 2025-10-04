package models

type DataGaleri struct {
	Judul string
	Foto  string
}

var IsiGaleri = []DataGaleri{
	{
		Judul: "Ekologi",
		Foto:  "https://aqinbfwyqxlynlxybhod.supabase.co/storage/v1/object/public/img_galeri/ekologi.jpg",
	},
	{
		Judul: "Hari Guru",
		Foto:  "https://aqinbfwyqxlynlxybhod.supabase.co/storage/v1/object/public/img_galeri/hari_guru.JPG",
	},
	{
		Judul: "Jalan Sehat HUT RI",
		Foto:  "https://aqinbfwyqxlynlxybhod.supabase.co/storage/v1/object/public/img_galeri/jalan_sehat.JPG",
	},
	{
		Judul: "Job Fair",
		Foto:  "https://aqinbfwyqxlynlxybhod.supabase.co/storage/v1/object/public/img_galeri/job_fair.jpg",
	},
}
