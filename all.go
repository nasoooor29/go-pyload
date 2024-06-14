package gopyload


func (p *Pyload) StatusServer() (any, error) {
	return GeneralPostEndpoint[any](p, "", nil)
}

func (p *Pyload) ToggleReconnect() (any, error) {
	return GeneralPostEndpoint[any](p, "", nil)
}

