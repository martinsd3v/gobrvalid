package govalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsVehiclePlate(t *testing.T) {
	t.Run("Deve ser capaz de validar uma placa", func(t *testing.T) {
		t.Helper()

		plateBrazil := "NWD4930"
		result := IsVehiclePlate(plateBrazil)
		assert.True(t, result)

		plateMercosul := "WNS8W99"
		result = IsVehiclePlate(plateMercosul)
		assert.True(t, result)

		plateInvalid := "293WSK0"
		result = IsVehiclePlate(plateInvalid)
		assert.False(t, result)
	})
}

func TestIsVehicleRenavam(t *testing.T) {
	t.Run("Deve ser capaz de validar renavam", func(t *testing.T) {
		t.Helper()

		//Renavam valid with 11 digits
		validRenavam := "70586542192"
		isValid, _ := IsVehicleRenavam(validRenavam)
		assert.True(t, isValid)

		//Renavam valid with 9 digits
		validRenavam = "163793123"
		isValid, _ = IsVehicleRenavam(validRenavam)
		assert.True(t, isValid)

		//Renavam valid with 11 digits and multiplier equal to 10
		validRenavam = "12345678900"
		isValid, _ = IsVehicleRenavam(validRenavam)
		assert.True(t, isValid)

		//Renavam valid
		validRenavam = "00123456700"
		isValid, _ = IsVehicleRenavam(validRenavam)
		assert.False(t, isValid)

	})
}
