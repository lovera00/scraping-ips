package support

import (
	"fmt"

	"github.com/scrapeos/ips/bd"
	"github.com/tebeka/selenium"
)

var Driver selenium.WebDriver

//Instancia el driver
func WDInit() selenium.WebDriver {
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	Driver, err = selenium.NewRemote(caps, "")
	if err != nil {
		fmt.Println("Error al instanciar el driver, motivo: \n", err.Error())
	}
	Driver.MaximizeWindow("")
	return Driver
}
func InicioScrapeo(cedula string) {
	input_ci, err := Driver.FindElement(selenium.ByName, "nro_cic")

	if err != nil {
		fmt.Println("No encontró el elemento input_ci, motivo: ", err.Error())
		Driver.Get("http://servicios.ips.gov.py/consulta_asegurado/comprobacion_de_derecho_externo.php")
		return
	}
	input_ci.Clear()
	input_ci.SendKeys(cedula)
	btnBuscar, err := Driver.FindElement(selenium.ByName, "recuperar")
	if err != nil {
		fmt.Println("No se encontró el boton, btnBuscar")
		Driver.Get("http://servicios.ips.gov.py/consulta_asegurado/comprobacion_de_derecho_externo.php")
		return
	}
	btnBuscar.Click()
	//Datos Tabla 1
	datos_nro_documento, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[2]/tbody/tr[2]/td[2]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_nro_documento, motivo: ", err.Error())
	}
	//fmt.Println(datos_nro_documento.Text())
	datos_nombres, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[2]/tbody/tr[2]/td[3]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_nombres, motivo: ", err.Error())
	}
	//fmt.Println(datos_nombres.Text())
	datos_apellidos, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[2]/tbody/tr[2]/td[4]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_apellidos, motivo: ", err.Error())
	}
	//fmt.Println(datos_apellidos.Text())
	datos_fec_nacimiento, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[2]/tbody/tr[2]/td[5]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_fec_nacimiento, motivo: ", err.Error())
	}
	//fmt.Println(datos_fec_nacimiento.Text())
	datos_sexo, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[2]/tbody/tr[2]/td[6]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_sexo, motivo: ", err.Error())
	}
	//fmt.Println(datos_sexo.Text())
	datos_tipo_aseg, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[2]/tbody/tr[2]/td[7]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_tipo_aseg, motivo: ", err.Error())
	}
	//fmt.Println(datos_tipo_aseg.Text())
	datos_beneficiarios, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[2]/tbody/tr[2]/td[8]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_beneficiarios, motivo: ", err.Error())
	}
	//fmt.Println(datos_beneficiarios.Text())
	datos_enrolado, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[2]/tbody/tr[2]/td[9]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_enrolado, motivo: ", err.Error())
	}
	//fmt.Println(datos_enrolado.Text())
	datos_venc_fe_vida, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[2]/tbody/tr[2]/td[10]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_venc_fe_vida, motivo: ", err.Error())
	}
	//fmt.Println(datos_venc_fe_vida.Text())
	//Datos Tabla 2
	datos_nro_patronal, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[3]/tbody/tr[2]/td[1]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_nro_patronal, motivo: ", err.Error())
	}
	//fmt.Println(datos_nro_patronal.Text())
	datos_empleador, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[3]/tbody/tr[2]/td[2]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_empleador, motivo: ", err.Error())
	}
	//fmt.Println(datos_empleador.Text())
	datos_estado, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[3]/tbody/tr[2]/td[3]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_estado, motivo: ", err.Error())
	}
	//fmt.Println(datos_estado.Text())
	datos_meses_aporte, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[3]/tbody/tr[2]/td[4]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_meses_aporte, motivo: ", err.Error())
	}
	//fmt.Println(datos_meses_aporte.Text())
	datos_vencimiento, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[3]/tbody/tr[2]/td[5]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_vencimiento, motivo: ", err.Error())
	}
	//fmt.Println(datos_vencimiento.Text())
	datos_ult_per_abonado, err := Driver.FindElement(selenium.ByXPATH, "/html/body/center[2]/form/table[3]/tbody/tr[2]/td[6]")
	if err != nil {
		fmt.Println("No encontró el elemento datos_ult_per_abonado, motivo: ", err.Error())
	}
	var V_PEIPS_NRO_DOCUMENTO string
	V_PEIPS_NRO_DOCUMENTO, err = datos_nro_documento.Text()
	if err != nil {
		V_PEIPS_NRO_DOCUMENTO = ""
	}
	var V_PEIPS_NOMBRES string
	V_PEIPS_NOMBRES, err = datos_nombres.Text()
	if err != nil {
		V_PEIPS_NOMBRES = ""
	}
	var V_PEIPS_APELLIDOS string
	V_PEIPS_APELLIDOS, err = datos_apellidos.Text()
	if err != nil {
		V_PEIPS_APELLIDOS = ""
	}
	var V_PEIPS_FEC_NACIMIENTO string
	V_PEIPS_FEC_NACIMIENTO, err = datos_fec_nacimiento.Text()
	if err != nil {
		V_PEIPS_FEC_NACIMIENTO = ""
	}
	var V_PEIPS_SEXO string
	V_PEIPS_SEXO, err = datos_sexo.Text()
	if err != nil {
		V_PEIPS_SEXO = ""
	}
	var V_PEIPS_TIPO_ASEG string
	V_PEIPS_TIPO_ASEG, err = datos_tipo_aseg.Text()
	if err != nil {
		V_PEIPS_TIPO_ASEG = ""
	}
	var V_PEIPS_BENEFICIARIOS string
	V_PEIPS_BENEFICIARIOS, err = datos_beneficiarios.Text()
	if err != nil {
		V_PEIPS_BENEFICIARIOS = ""
	}
	var V_PEIPS_ENROLADO string
	V_PEIPS_ENROLADO, err = datos_enrolado.Text()
	if err != nil {
		V_PEIPS_ENROLADO = ""
	}
	var V_PEIPS_NRO_PATRONAL string
	V_PEIPS_NRO_PATRONAL, err = datos_nro_patronal.Text()
	if err != nil {
		V_PEIPS_NRO_PATRONAL = ""
	}
	var V_PEIPS_EMPLEADOR string
	V_PEIPS_EMPLEADOR, err = datos_empleador.Text()
	if err != nil {
		V_PEIPS_EMPLEADOR = ""
	}
	var V_PEIPS_MESES_APORTE string
	V_PEIPS_MESES_APORTE, err = datos_meses_aporte.Text()
	if err != nil {
		V_PEIPS_MESES_APORTE = ""
	}
	var V_PEIPS_ULT_PER_ABONADO string
	V_PEIPS_ULT_PER_ABONADO, err = datos_ult_per_abonado.Text()
	if err != nil {
		V_PEIPS_ULT_PER_ABONADO = ""
	}
	var V_PEIPS_VENC_FE_VIDA string
	V_PEIPS_VENC_FE_VIDA, err = datos_venc_fe_vida.Text()
	if err != nil {
		V_PEIPS_VENC_FE_VIDA = ""
	}
	var V_PEIPS_ESTADO_IPS string
	V_PEIPS_ESTADO_IPS, err = datos_estado.Text()
	if err != nil {
		V_PEIPS_ESTADO_IPS = ""
	}
	var V_PEIPS_VENCIMIENTO string
	V_PEIPS_VENCIMIENTO, err = datos_vencimiento.Text()
	if err != nil {
		V_PEIPS_VENCIMIENTO = ""
	}
	fmt.Println(err)
	bd.CreatePersonasIps(bd.Persona{
		PEIPS_NRO_DOCUMENTO:   V_PEIPS_NRO_DOCUMENTO,
		PEIPS_NOMBRES:         V_PEIPS_NOMBRES,
		PEIPS_APELLIDOS:       V_PEIPS_APELLIDOS,
		PEIPS_FEC_NACIMIENTO:  V_PEIPS_FEC_NACIMIENTO,
		PEIPS_SEXO:            V_PEIPS_SEXO,
		PEIPS_TIPO_ASEG:       V_PEIPS_TIPO_ASEG,
		PEIPS_BENEFICIARIOS:   V_PEIPS_BENEFICIARIOS,
		PEIPS_ENROLADO:        V_PEIPS_ENROLADO,
		PEIPS_VENC_FE_VIDA:    V_PEIPS_VENC_FE_VIDA,
		PEIPS_NRO_PATRONAL:    V_PEIPS_NRO_PATRONAL,
		PEIPS_EMPLEADOR:       V_PEIPS_EMPLEADOR,
		PEIPS_ESTADO_IPS:      V_PEIPS_ESTADO_IPS,
		PEIPS_MESES_APORTE:    V_PEIPS_MESES_APORTE,
		PEIPS_VENCIMIENTO:     V_PEIPS_VENCIMIENTO,
		PEIPS_ULT_PER_ABONADO: V_PEIPS_ULT_PER_ABONADO,
	})
	bd.UpdatePersonaIPS(bd.Persona{
		PEIPS_NRO_DOCUMENTO:   V_PEIPS_NRO_DOCUMENTO,
		PEIPS_NOMBRES:         V_PEIPS_NOMBRES,
		PEIPS_APELLIDOS:       V_PEIPS_APELLIDOS,
		PEIPS_FEC_NACIMIENTO:  V_PEIPS_FEC_NACIMIENTO,
		PEIPS_SEXO:            V_PEIPS_SEXO,
		PEIPS_TIPO_ASEG:       V_PEIPS_TIPO_ASEG,
		PEIPS_BENEFICIARIOS:   V_PEIPS_BENEFICIARIOS,
		PEIPS_ENROLADO:        V_PEIPS_ENROLADO,
		PEIPS_VENC_FE_VIDA:    V_PEIPS_VENC_FE_VIDA,
		PEIPS_NRO_PATRONAL:    V_PEIPS_NRO_PATRONAL,
		PEIPS_EMPLEADOR:       V_PEIPS_EMPLEADOR,
		PEIPS_ESTADO_IPS:      V_PEIPS_ESTADO_IPS,
		PEIPS_MESES_APORTE:    V_PEIPS_MESES_APORTE,
		PEIPS_VENCIMIENTO:     V_PEIPS_VENCIMIENTO,
		PEIPS_ULT_PER_ABONADO: V_PEIPS_ULT_PER_ABONADO,
	}, "S")

}
func ObtieneCedulas() []string {
	//cedulas := []string{"5351538", "2021268", "859788", "3865554", "925995", "1509040", "1662810", "3621170", "2122020", "1160815", "1991752"}
	cedulas := []string{"5351538"}
	return cedulas
}
