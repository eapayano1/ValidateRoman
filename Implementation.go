package KataValidarRomano

func ValidarNumeroRomano(NumeroRomano string) bool {

	ContRepeticiones := 1
	condition := true
	RomanosSinRestriccion := []byte{'I', 'X', 'C', 'M'}
	RomanosConRestriccion := []byte{'V', 'L', 'D'}
	for i := 0; i < len(NumeroRomano); i++ { //Validar arcaismo

		for x := 0; x < len(RomanosSinRestriccion); x++ {
			if NumeroRomano[i] == RomanosSinRestriccion[x] {
				if ContRepeticiones < 4 {
					if i < len(NumeroRomano)-1 {
						condition = false
						if NumeroRomano[i] == NumeroRomano[i+1] {
							ContRepeticiones++
						} else {
							ContRepeticiones = 1
						}
					}
				} else {
					return false
				}
			}

		}
		if condition {

			for k := 0; k < len(RomanosConRestriccion); k++ {
				if NumeroRomano[i] == RomanosConRestriccion[k] {
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
		condition = true
	}

	return true
}
