package main 

import "testing"

var(
	kubus				Kubus	= Kubus{4}
	volumeSeharusnya	float64 = 64
	luasSeharusnya		float64 = 96
	kelilingSeharusnya  float64 = 48
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Volume : %.2f", kubus.Volume()) // digunakan untuk memunculkan log, sama seperti *printf*

	if kubus.Volume() != volumeSeharusnya {
		t.Errorf("Salah! harusnya %.2f", volumeSeharusnya) // memunculkan log dengan keterangan fail pada testing
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", kubus.Luas())

	if kubus.Luas() != luasSeharusnya {
		t.Errorf("Salah! harusnya %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", kubus.Keliling())

	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("Salah! harusnya %.2f", kelilingSeharusnya)
	}
}

/*		
	Untuk menjalankan program :
		go test kubus.go kubus_test.go -v // -v(verbose) menampilkan semua output log pada saat pengujian

*/