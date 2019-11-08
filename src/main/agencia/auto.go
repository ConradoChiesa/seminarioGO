package agencia

type Auto struct {
	patente   string
	marca     string
	modelo    int
	velocidad int
}

func NewAuto(patente, marca string, modelo int) Auto  {
	return Auto{
		patente: patente,
		marca:   marca,
		modelo:  modelo,
	}
}

func (a *Auto) Acelerar()  {
	a.velocidad += 10
}

func (a *Auto) Frenar() {
	a.velocidad -= 10
}

func (a *Auto) SetPatente(p string) {
	a.patente = p
}

func (a *Auto) SetMarca(m string) {
	a.marca = m
}

func (a *Auto) SetModelo(m int) {
	a.modelo = m
}