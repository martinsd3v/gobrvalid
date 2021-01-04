package govalid

import (
	"fmt"
	"strconv"
)

//IsCPF responsible for validating CPF
func IsCPF(cpf string) (isValid bool, clean string) {
	cleanCPF := ReplacePattern(cpf, "[^0-9]", "")
	clean = cleanCPF

	//Checking number of digits
	if len(cleanCPF) != 11 {
		isValid = false
		return
	}

	//Checking if they are not the same numbers in sequence
	if Matches(cleanCPF, "([0-9])[1]{10}") {
		isValid = false
		return
	}

	//Algorithm for validation
	for t := 9; t < 11; t++ {
		c, d := 0, 0

		for d, c = 0, 0; c < t; c++ {
			a, _ := strconv.Atoi(cleanCPF[c : c+1])
			d += a * ((t + 1) - c)
		}

		d = ((10 * d) % 11) % 10

		b, _ := strconv.Atoi(cleanCPF[c : c+1])

		if b != d {
			isValid = false
			return
		}
	}

	isValid = true
	return
}

//IsCNPJ responsible for validating CNPJ
func IsCNPJ(cnpj string) (isValid bool, clean string) {
	cleanCNPJ := ReplacePattern(cnpj, "[^0-9]", "")
	clean = cleanCNPJ

	//Checking number of digits
	if len(cleanCNPJ) != 14 {
		isValid = false
		return
	}

	//Checking if they are not the same numbers in sequence
	if Matches(cleanCNPJ, "([0-9])[1]{13}") {
		isValid = false
		return
	}

	//Algorithm for validation
	for t := 12; t < 14; t++ {
		c, d, p := 0, 0, 0

		for d, p, c = 0, t-7, 0; c < t; c++ {
			a, _ := strconv.Atoi(cleanCNPJ[c : c+1])
			d += a * p

			if p < 3 {
				p = 9
			} else {
				p--
			}
		}

		d = ((10 * d) % 11) % 10

		b, _ := strconv.Atoi(cleanCNPJ[c : c+1])

		if b != d {
			isValid = false
			return
		}
	}

	isValid = true
	return
}

//IsCPForCNPJ responsible for validating CNPJ and CNPJ
func IsCPForCNPJ(doc string) (isValid bool, clean string) {
	iPF, cPF := IsCNPJ(doc)
	iPJ, cPJ := IsCPF(doc)

	if iPF {
		clean = cPF
		isValid = true
	} else if iPJ {
		clean = cPJ
		isValid = true
	}

	return
}

//IsCNH responsible for validating CNH
func IsCNH(cnh string) (isValid bool, clean string) {
	cleanCNH := ReplacePattern(cnh, "[^0-9]", "")

	//Checking number of digits
	if len(cleanCNH) != 11 {
		isValid = false
		return
	}

	//Checking if they are not the same numbers in sequence
	if Matches(cleanCNH, "([0-9])[1]{10}") {
		isValid = false
		return
	}

	//Algorithm for validation
	p := cleanCNH[0:9]

	sum := 0
	acc := 9
	for i := range p {
		n, _ := strconv.Atoi(p[i : i+1])
		sum += n * acc
		acc--
	}

	base := 0
	digit1 := sum % 11
	if digit1 == 10 {
		base = -2
	}
	if digit1 > 9 {
		digit1 = 0
	}

	sum = 0
	acc = 1
	for i := range p {
		n, _ := strconv.Atoi(p[i : i+1])
		sum += n * acc
		acc++
	}

	var digit2 int
	if (sum%11)+base < 0 {
		digit2 = 11 + (sum % 11) + base
	}
	if (sum%11)+base >= 0 {
		digit2 = (sum % 11) + base
	}
	if digit2 > 9 {
		digit2 = 0
	}

	digitsAlg := fmt.Sprintf("%d%d", digit1, digit2)
	digitsCnh := cleanCNH[9:11]

	isValid = digitsAlg == digitsCnh
	return
}
