package gopyload

// Errored on me
func (p *Pyload) GetServices() (any, error) {
	return GeneralPostEndpoint[any](p, getServicesEndpoint, nil)
}

func (p *Pyload) Kill() error {
	_, err := GeneralPostEndpoint[any](p, killEndpoint, nil)
	return err
}

func (p *Pyload) PauseServer() error {
	_, err := GeneralPostEndpoint[any](p, pauseServerEndpoint, nil)
	return err
}

// make sure to have session id or else you would be unauthorized
func (p *Pyload) UnpauseServer() error {
	_, err := GeneralPostEndpoint[any](p, unpauseServerEndpoint, nil)
	return err
}

func (p *Pyload) Restart() error {
	_, err := GeneralPostEndpoint[any](p, restartEndpoint, nil)
	return err
}

// i don't know what it does but it does something weird so don't use it
func (p *Pyload) RestartFailed() (any, error) {
	return GeneralPostEndpoint[any](p, restartFailedEndpoint, nil)
}
