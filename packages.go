// this package is for any endpoints related to the packages in pyload
package gopyload

func (p *Pyload) DeleteFinishedPackages() ([]int64, error) {
	return GeneralPostEndpoint[[]int64](p, deleteFinishedEndpoint, nil)
}

func (p *Pyload) StatusDownloads() (any, error) {
	return GeneralPostEndpoint[any](p, statusDownloadsEndpoint, nil)
}

func (p *Pyload) StopAllDownloads() (any, error) {
	return GeneralPostEndpoint[any](p, "", nil)
}

func (p *Pyload) TogglePause() (any, error) {
	return GeneralPostEndpoint[any](p, togglePauseEndpoint, nil)
}
