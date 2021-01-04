package govalid

import (
	"fmt"
	"strconv"
)

//IsVehiclePlate responsible for validating vehicles plates, brazil e mercosul
func IsVehiclePlate(plate string) bool {
	regex := "^([A-Z ]{3}[0-9]{4}|[A-Z ]{3}-[0-9]{4}|[A-Z ]{3}[0-9]{1}[A-Z ]{1}[0-9]{2})$"
	return Matches(plate, regex)
}

//IsVehicleRenavam responsible for validating vehicles renavam
func IsVehicleRenavam(renavam string) (isValid bool, clean string) {
	cleanNumber := ReplacePattern(renavam, "[^0-9]", "")
	clean = cleanNumber

	//Incrementing with zeros to the left to be 11 digits
	cleanNumber = fmt.Sprintf("%0*s", 11, cleanNumber)

	renavamDigit, _ := strconv.Atoi(cleanNumber[10:])
	renavamWithoutDigit := cleanNumber[0:10]
	renavamWithoutDigitReverse := Reverse(renavamWithoutDigit)

	// Multiplies renava's reverse strings by multiplier numbers
	// Example: Renava Reverse without Digit = 69488936
	// 6, 9, 4, 8, 8, 9, 3, 6
	// * * * * * * * *
	// 2, 3, 4, 5, 6, 7, 8, 9 (multiplier numbers - always the same [fixed])
	// 12 + 27 + 16 + 40 + 48 + 63 + 24 + 54
	// sum = 284

	sum := 0
	multiplier := 2
	for i := 0; i < 10; i++ {
		numeral, _ := strconv.Atoi(renavamWithoutDigitReverse[i : i+1])
		sum += numeral * multiplier

		if multiplier >= 9 {
			multiplier = 2
		} else {
			multiplier++
		}
	}

	// mod11 = 284 % 11 = 9 (rest of the division by 11)
	mod11 := sum % 11

	//Calculate 11 (Fix value) - mod11 = 11 - 9 = 2
	lastDigitCalculeted := 11 - mod11

	//lastDigit = If the previously calculated value is 10 or 11, I change it to 0
	if lastDigitCalculeted >= 10 {
		lastDigitCalculeted = 0
	}

	isValid = lastDigitCalculeted == renavamDigit
	return
}
