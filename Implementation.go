package KataValidarRomano

func ValidarNumeroRomano(NumeroRomano string) bool {

	ContRepeticiones := 1
	for i := 0; i < len(NumeroRomano); i++ { //Validar arcaismo
		if ContRepeticiones < 4 {
			if i < len(NumeroRomano)-1 {
				if NumeroRomano[i] == NumeroRomano[i+1] {
					ContRepeticiones++
				} else {
					ContRepeticiones = 1
				}
			}
		} else {
			return false
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

	return true
}
