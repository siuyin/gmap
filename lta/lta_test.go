package lta

import (
	"os"
	"testing"
)

func loadTestBicycleParkingSpots(t *testing.T) []BicycleParking {
	f, err := os.Open("testdata/ltaBicycleOutput.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	return load(f)
}
func TestLoad(t *testing.T) {
	spots := loadTestBicycleParkingSpots(t)
	if n := len(spots); n != 19 {
		t.Errorf("bad len, got: %d", n)
	}
}

func TestFirstSpot(t *testing.T) {
	first := loadTestBicycleParkingSpots(t)[0]
	if d := first.Description; d != "BOON TAT STREET" {
		t.Errorf("bad description, got: %s", d)
	}
}

func TestGeoJSON(t *testing.T) {
	f, err := os.Open("testdata/ltaBicycleOutput.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	s := GeoJSON(f)
	t.Log(s)
}

func TestBicycleParkingSpots(t *testing.T) {
	t.Log(BicycleParkingSpots(1.3467093706130981, 103.77416229248047))
}
