package main

import "testing"

var (
	kubus          Kubus   = Kubus{4}
	assertVolume   float64 = 64
	assertLuas     float64 = 96
	assertkeliling float64 = 16
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Volume : %.2f", kubus.Volume())

	if kubus.Volume() != assertVolume {
		t.Errorf("Salah ..harusnya %.2f", assertVolume)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %2.f", assertLuas)

	if kubus.Luas() != assertLuas {
		t.Errorf("Salah.. harusnya %.2f", assertLuas)
	}

}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling: %.2f", kubus.Keliling())

	if kubus.Keliling() != assertkeliling {
		t.Errorf("Salah ... harusnya %.2f", assertkeliling)
	}
}
