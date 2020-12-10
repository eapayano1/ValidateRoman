package Kata3

func ValidarNumeroRomano(NumeroRomano string) bool {
	contadorCaracteres := 1
	condition := true
	CaracteresSinRestriccion := []byte{'I', 'X', 'C', 'M'}
	CaracteresConRestriccion := []byte{'V', 'L', 'D'}
	for i := 0; i < len(NumeroRomano); i++ {
		for x := 0; x < len(CaracteresSinRestriccion); x++ {
			if NumeroRomano[i] == CaracteresSinRestriccion[x] {
				if contadorCaracteres < 4 {
					if i < len(NumeroRomano)-1 {
						condition = false
						if NumeroRomano[i] == NumeroRomano[i+1] {
							contadorCaracteres++
						} else {
							contadorCaracteres = 1
						}
					}
				} else {
					return false
				}
			}
		}

		if condition {
			for k := 0; k < len(CaracteresConRestriccion); k++ {
				if NumeroRomano[i] == CaracteresConRestriccion[k] {
					if i < len(NumeroRomano)-1 {
						if NumeroRomano[i] == NumeroRomano[i+1] {
							return false
						}
					}
				}
			}
		}
		switch NumeroRomano[i] {
		case 'I':
			break
		case 'V':
			break
		case 'X':
			break
		case 'L':
			break
		case 'C':
			break
		case 'D':
			break
		case 'M':
			break
		default:
			return false
		}
	}
	condition = true
	return true
}
