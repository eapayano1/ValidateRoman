## "Validar Romano"

### Requerimiento: 
- El sistema debe validar si un numero es romano o no.
### Criterios de aceptación: 
-	El sistema debe detectar un arcaísmo ortográfico.
-	El sistema debe validar si un carácter de un numero romano se ha repetido mas de 3 veces.
-	El sistema debe validar si uno de los caracteres {‘V’,’L’,’D’} se ha repetido mas de una vez.
### Escenarios de prueba:
- ValidarRomano(“iXX”)=>false
- ValidarRomano(“IIII”)=>false
- ValidarRomano(“VV”)=>false
