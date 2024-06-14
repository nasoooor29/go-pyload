package gopyload

// General types for me
type headersMap map[string]string

func (p *Pyload) Template() (any, error) {
	return GeneralPostEndpoint[any](p, "", nil)
}

