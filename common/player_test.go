package common

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"
)

func TestPlayerActiveWeapon(t *testing.T) {
	knife := NewEquipment(EqKnife)
	glock := NewEquipment(EqGlock)
	ak47 := NewEquipment(EqAK47)

	pl := NewPlayer()
	pl.RawWeapons[1] = &knife
	pl.RawWeapons[2] = &glock
	pl.RawWeapons[3] = &ak47
	pl.ActiveWeaponID = 3

	assert.Equal(t, &ak47, pl.ActiveWeapon(), "Should have AK-47 equipped")
}
func TestPlayerWeapons(t *testing.T) {
	knife := NewEquipment(EqKnife)
	glock := NewEquipment(EqGlock)
	ak47 := NewEquipment(EqAK47)

	pl := NewPlayer()
	pl.RawWeapons[1] = &knife
	pl.RawWeapons[2] = &glock
	pl.RawWeapons[3] = &ak47

	expected := []*Equipment{&knife, &glock, &ak47}
	assert.ElementsMatch(t, expected, pl.Weapons(), "Should have expected weapons")
}

func TestPlayerAlive(t *testing.T) {
	pl := NewPlayer()

	pl.Hp = 100
	assert.Equal(t, true, pl.IsAlive(), "Should be alive")

	pl.Hp = 1
	assert.Equal(t, true, pl.IsAlive(), "Should be alive")

	pl.Hp = 0
	assert.Equal(t, false, pl.IsAlive(), "Should be dead")

	pl.Hp = -10
	assert.Equal(t, false, pl.IsAlive(), "Should be dead")
}

func TestPlayerFlashed(t *testing.T) {
	pl := NewPlayer()

	assert.False(t, pl.IsBlinded(), "Should not be flashed")

	pl.FlashDuration = 2.3
	assert.True(t, pl.IsBlinded(), "Should be flashed")
}

func TestPlayer_FlashDurationTime(t *testing.T) {
	p := NewPlayer()

	assert.Equal(t, time.Duration(0), p.FlashDurationTime())

	p.FlashDuration = 2.3

	assert.Equal(t, 2300*time.Millisecond, p.FlashDurationTime())
}
