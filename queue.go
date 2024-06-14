// this file for the queues and collectors endpoints
package gopyload

import "gopyload/models"


func (p *Pyload) GetQueues() ([]models.QueueModel, error) {
	return GeneralGetEndpoint[[]models.QueueModel](p, getQueueEndpoint)
}

func (p *Pyload) GetQueueData() ([]models.QueueDataModel, error) {
	return GeneralGetEndpoint[[]models.QueueDataModel](p, getQueueEndpoint)
}

func (p *Pyload) GetCollectors() ([]any, error) {
	return GeneralGetEndpoint[[]any](p, getCollectorEndpoint)
}

func (p *Pyload) GetCollectorsData() ([]any, error) {
	return GeneralGetEndpoint[[]any](p, getCollectorDataEndpoint)
}
