package gobrvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCPF(t *testing.T) {
	t.Run("Should be able to validate CPF", func(t *testing.T) {
		var (
			validCPF   string
			invalidCPF string
			isValid    bool
		)

		//CPF validate algorithm
		validCPF = "701.220.940-26"
		isValid, _ = IsCPF(validCPF)
		assert.True(t, isValid)

		//CPF missing characters
		invalidCPF = "701.220.940-2"
		isValid, _ = IsCPF(invalidCPF)
		assert.False(t, isValid)

		//CPF invalidate algorithm
		invalidCPF = "04278262088"
		isValid, _ = IsCPF(invalidCPF)
		assert.False(t, isValid)

		//CPF sequence of numbers
		invalidCPF = "11111111111"
		isValid, _ = IsCPF(invalidCPF)
		assert.False(t, isValid)

	})
}

func TestIsCNPJ(t *testing.T) {
	t.Run("Should be able to validate CNPJ", func(t *testing.T) {
		var (
			validCNPJ   string
			invalidCNPJ string
			isValid     bool
		)

		//CNPJ validate algorithm
		validCNPJ = "33.756.021/0001-90"
		isValid, _ = IsCNPJ(validCNPJ)
		assert.True(t, isValid)

		//CNPJ missing characters
		invalidCNPJ = "33.756.021/0001-9"
		isValid, _ = IsCNPJ(invalidCNPJ)
		assert.False(t, isValid)

		//CNPJ invalidate algorithm
		invalidCNPJ = "33.756.021/0001-74"
		isValid, _ = IsCNPJ(invalidCNPJ)
		assert.False(t, isValid)

		//CNPJ sequence of numbers
		invalidCNPJ = "11111111111111"
		isValid, _ = IsCNPJ(invalidCNPJ)
		assert.False(t, isValid)
	})
}

func TestIsCPForCNPJ(t *testing.T) {
	t.Run("Should be able to validate CPF and CNPJ", func(t *testing.T) {
		var (
			validCPF  string
			validCNPJ string
			isValid   bool
		)

		//CPF validate algorithm
		validCPF = "701.220.940-26"
		isValid, _ = IsCPForCNPJ(validCPF)
		assert.True(t, isValid)

		//CNPJ invalidate algorithm
		validCNPJ = "33.756.021/0001-90"
		isValid, _ = IsCPForCNPJ(validCNPJ)
		assert.True(t, isValid)

	})
}

func TestIsCNH(t *testing.T) {
	t.Run("Should be able to validate CNH", func(t *testing.T) {
		var (
			validCNH   string
			invalidCNH string
			isValid    bool
		)

		//CNH validate algorithm
		validCNH = "04616181089"
		isValid, _ = IsCNH(validCNH)
		assert.True(t, isValid)

		//CNH validate algorithm
		validCNH = "87222700600"
		isValid, _ = IsCNH(validCNH)
		assert.True(t, isValid)

		//CNH validate algorithm
		validCNH = "45991167705"
		isValid, _ = IsCNH(validCNH)
		assert.True(t, isValid)

		//CNH validate algorithm
		validCNH = "00067600300"
		isValid, _ = IsCNH(validCNH)
		assert.True(t, isValid)

		//CNH missing characters
		invalidCNH = "0324189377"
		isValid, _ = IsCNH(invalidCNH)
		assert.False(t, isValid)

		//CNH invalidate algorithm
		invalidCNH = "03241893755"
		isValid, _ = IsCNH(invalidCNH)
		assert.False(t, isValid)

		//CNH sequence of numbers
		invalidCNH = "11111111111"
		isValid, _ = IsCNH(invalidCNH)
		assert.False(t, isValid)

	})
}
