package agencia

type Auto struct {
	Patente   string `json: patente`
	Marca     string `json: marca`
	Modelo    int    `json: modelo`
	Velocidad int    `json: velocidad`
}

func NewAuto(patente, marca string, modelo int) Auto {
	return Auto{
		Patente: patente,
		Marca:   marca,
		Modelo:  modelo,
	}
}

func (a *Auto) Acelerar() {
	a.Velocidad += 10
}

func (a *Auto) Frenar() {
	a.Velocidad -= 10
}

func (a *Auto) SetPatente(p string) {
	a.Patente = p
}

func (a *Auto) SetMarca(m string) {
	a.Marca = m
}

func (a *Auto) SetModelo(m int) {
	a.Modelo = m
}

/*func (a *Auto) AddCar(ctx *gin.Context) {
body, err := ioutil.ReadAll(ctx.Request.Body)
if err != nil {
	(...)
}

var car Auto
if err = json.Unmarshal(body, &car); err != nil {
	(...)
}

a.carDatabase.AddCar(car)

ctx.JSON(http.StatusCreated, gin.H{
	"message": "car added",
})
*/

//}
