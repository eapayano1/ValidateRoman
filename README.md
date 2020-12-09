# "Validar Romanos"

### Problema: 
- Validar número romano

### Requerimiento: 
- El sistema debe validar números romanos como verdadero o falso.

### Criterios de aceptacion: 
-   El  sistema debe detectar un arcaísmo ortográfico.
-	El sistema debe validar como falso si un carácter se repite mas de 3 veces.
-	El sistema debe validar como falso si uno de los caracteres {‘V’,’L’,’D’}

### Escenarios de prueba: 
-	ValidarRomano(“Xi”)=>False
-	ValidarRomano(“IIII”)=>False
-	ValidarRomano(“LL”)=>False
