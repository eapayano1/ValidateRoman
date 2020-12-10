package Kata3

func ValidarNumeroRomano(NumeroRomano string) bool {
	contadorCaracteres := 1
	for i := 0; i < len(NumeroRomano); i++ {
		if contadorCaracteres < 4 {
			if i < len(NumeroRomano)-1 {
				if NumeroRomano[i] == NumeroRomano[i+1] {
					contadorCaracteres++
				} else {
					contadorCaracteres = 1
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
