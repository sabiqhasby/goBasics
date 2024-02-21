// kegunaan Assert adalah untuk kegunaan testing

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, 30, 20, "perhitungan volume salah")
}

// func TestHitungLuas(t *testing.T) {
// 	assert.Equal(t, kubus.Luas(), luasSeharusnya, "perhitungan luas salah")

// }

// func TestHitungKeliling(t *testing.T) {
// 	assert.Equal(t, kubus.Keliling(), kelilingSeharusnya, "perhitungan keliling salah")
// }
