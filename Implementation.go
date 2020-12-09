package KataValidarRomano

func ValidarNumeroRomano(NumeroRomano string) bool {

	for i := 0; i < len(NumeroRomano); i++ { //Validar arcaismo
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
