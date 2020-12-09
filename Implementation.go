package Kata3

func ValidarNumeroRomano(NumeroRomano string) bool {

	ContadorRepeticiones := 1
	for i := 0; i < len(NumeroRomano); i++ {
		if ContadorRepeticiones < 4 {
			if i < len(NumeroRomano)-1 {
				if NumeroRomano[i] == NumeroRomano[i+1] {
					ContadorRepeticiones++
				} else {
					ContadorRepeticiones = 1
				}
			}
		} else {
			return false
		}
		switch NumeroRomano[i] {
		case 'I':
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
