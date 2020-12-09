package Kata3

func ValidarNumeroRomano(a string) bool {

	for i := 0; i < len(a); i++ {
		switch a[i] {
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
