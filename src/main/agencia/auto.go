package agencia

type Auto struct {
	Patente   string
	Marca     string
	Modelo    int
	Velocidad int
}

func NewAuto(patente, marca string, modelo int) Auto  {
	return Auto{
		Patente: patente,
		Marca:   marca,
		Modelo:  modelo,
	}
}

func (a *Auto) Acelerar()  {
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