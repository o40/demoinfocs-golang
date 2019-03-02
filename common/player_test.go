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

	pl := newPlayer()
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

	pl := newPlayer()
	pl.RawWeapons[1] = &knife
	pl.RawWeapons[2] = &glock
	pl.RawWeapons[3] = &ak47

	expected := []*Equipment{&knife, &glock, &ak47}
	assert.ElementsMatch(t, expected, pl.Weapons(), "Should have expected weapons")
}

func TestPlayerAlive(t *testing.T) {
	pl := newPlayer()

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
	pl := newPlayer()

	assert.False(t, pl.IsBlinded(), "Should not be flashed")

	pl.FlashDuration = 2.3
	pl.FlashTick = 50
	*pl.ingameTick = 128
	assert.True(t, pl.IsBlinded(), "Should be flashed")
}

func TestPlayerFlashed_FlashDuration_Over(t *testing.T) {
	pl := newPlayer()

	pl.FlashDuration = 1.9
	pl.FlashTick = 128
	*pl.ingameTick = 128 * 3
	assert.False(t, pl.IsBlinded(), "Should not be flashed")
}

func TestPlayer_FlashDurationTime_Default(t *testing.T) {
	pl := newPlayer()

	assert.Equal(t, time.Duration(0), pl.FlashDurationTime())
}

func TestPlayer_FlashDurationTime(t *testing.T) {
	pl := newPlayer()

	pl.FlashDuration = 2.3

	assert.Equal(t, 2300*time.Millisecond, pl.FlashDurationTime())
}

func TestPlayer_FlashDurationTimeRemaining_Default(t *testing.T) {
	pl := newPlayer()

	assert.Equal(t, time.Duration(0), pl.FlashDurationTimeRemaining())
}

func TestPlayer_FlashDurationTimeRemaining(t *testing.T) {
	pl := newPlayer()

	pl.FlashDuration = 3
	pl.FlashTick = 128
	*pl.ingameTick = 128 * 2
	assert.Equal(t, 2*time.Second, pl.FlashDurationTimeRemaining())
}

func TestPlayer_FlashDurationTimeRemaining_Zero(t *testing.T) {
	pl := newPlayer()

	pl.FlashDuration = 3
	pl.FlashTick = 128
	*pl.ingameTick = 128 * 4
	assert.Equal(t, time.Duration(0), pl.FlashDurationTimeRemaining())
}

func TestPlayer_FlashDurationTimeRemaining_FlashDuration_Over(t *testing.T) {
	pl := newPlayer()

	pl.FlashDuration = 1
	pl.FlashTick = 128
	*pl.ingameTick = 128 * 4
	assert.Equal(t, time.Duration(0), pl.FlashDurationTimeRemaining())
}

func newPlayer() *Player {
	tick := 0
	return NewPlayer(128, &tick)
}
