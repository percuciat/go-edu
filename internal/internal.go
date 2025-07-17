package internal

type internal struct {
	cfg string
}

func (i *internal) setCfg(name string) {
	i.cfg = name
}
