# Saludos en Go

Este paquete proporciona una forma simple deobtener saludos personalizados en go

## Instalación
Ejecutar el siguiente comando para instalar el paquete:
```bash
go get -u github.com/Haden0101/gretting-modularizacion
```

## Uso
Aqui tienes un ejemplo de cómo utilizar el paquete en tu código:

```go
package main

import (
	"fmt"
	"log"

	modularizacion "github.com/Haden0101/gretting-modularizacion"
)

func main() {

	messages, err := modularizacion.Hello("Isaias")

	if err != nil {
		log.Fatal("Ocurrio un error:", err)
	}

	fmt.Println(messages)
}
```
Este ejemplo importa el paquete de github.com/Haden0101/gretting-modularizacion y llama a la función Hello(),
saudo personzalizado. Si ocurre un error, se imprimeun mensaje de error.