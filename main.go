package main

import (
	"fmt"

	"github.com/scrapeos/ips/bd"
	"github.com/scrapeos/ips/support"
)

func main() {
	fmt.Println("Inició")
	bd.Setup()
	var err error
	cedulas, err := bd.GetCedulas()
	fmt.Println(err)
	Driver := support.WDInit()
	Driver.Get("http://servicios.ips.gov.py/consulta_asegurado/comprobacion_de_derecho_externo.php")
	var contador int = 0
	for _, cedula := range cedulas {
		support.InicioScrapeo(cedula.PEIPS_NRO_DOCUMENTO)
		fmt.Println("*---------------------------------------------------------------------*")
		fmt.Println(contador)
		contador++
		fmt.Println("*---------------------------------------------------------------------*")
	}
	Driver.Close()
}
