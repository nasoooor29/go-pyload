// any endpoints related to information
package gopyload

import "gopyload/models"

func (p *Pyload) GetFreeSpace() (int64, error) {
	return GeneralPostEndpoint[int64](p, freeSpaceEndpoint, nil)
}

func (p *Pyload) GetAvailablePremiumAccounts() ([]string, error) {
	return GeneralPostEndpoint[[]string](p, getAccountTypesEndpoint, nil)
}

func (p *Pyload) GetAllHooksInfo() (map[string]any, error) {
	return GeneralPostEndpoint[map[string]any](p, getAllInfoEndpoint, nil)
}

func (p *Pyload) GetConfig() (models.ConfigModel, error) {
	return GeneralPostEndpoint[models.ConfigModel](p, getConfigEndpoint, nil)
}

func (p *Pyload) GetConfigDict() (models.ConfigDictModel, error) {
	return GeneralPostEndpoint[models.ConfigDictModel](p, getConfigDictEndpoint, nil)
}

func (p *Pyload) GetPluginConfig() (models.PluginConfigModel, error) {
	return GeneralPostEndpoint[models.PluginConfigModel](p, getPluginConfigEndpoint, nil)
}

func (p *Pyload) GetPluginConfigDictEndpoint() (models.PluginConfigDictModel, error) {
	return GeneralPostEndpoint[models.PluginConfigDictModel](p, getConfigDictEndpoint, nil)
}

func (p *Pyload) GetServerVersionEndpoint() (string, error) {
	return GeneralPostEndpoint[string](p, getServerVersionEndpoint, nil)
}


func (p *Pyload) GetServerVersion() (string, error) {
	return GeneralPostEndpoint[string](p, getServerVersionEndpoint, nil)
}

func (p *Pyload) IsCaptchaWaiting() (bool, error) {
	return GeneralPostEndpoint[bool](p, isCaptchaWaitingEndpoint, nil)
}

func (p *Pyload) IsTimeDownload() (bool, error) {
	return GeneralPostEndpoint[bool](p, isTimeDownloadEndpoint, nil)
}

func (p *Pyload) IsTimeReconnect() (bool, error) {
	return GeneralPostEndpoint[bool](p, isTimeReconnectEndpoint, nil)
}
