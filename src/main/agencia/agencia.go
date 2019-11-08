package agencia

type Agencia struct {
	autos []Auto
}

func NewAgencia() Agencia {
	var auxAutos []Auto
	return Agencia {
		autos: auxAutos,
	}
}

func (a *Agencia) AddAuto(auto ...Auto) {

	a.autos = append(a.autos, auto...)
}

func (a *Agencia) DeleteAuto(pos int) {
	a.autos = append(a.autos[:pos], a.autos[pos+1:]...)
}

func (a *Agencia) GetAuto(pos int) Auto{
	return a.autos[pos]
}


//func (a *Agencia) GetAutoPorPatente(patente int) Auto {
//}

func (a *Agencia) GetAutos() []Auto {
	return a.autos
}