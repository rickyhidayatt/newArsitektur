package container

import (
	mediaService "bni.co.id/xpora/medias/internal/domain/service/media"
)

type DomainServiceIoC struct {
	media mediaService.MediaServiceInterface
}

func NewDomainServiceIoC(ioc RepositoryIoC) DomainServiceIoC {
	mediaSv := mediaService.NewMediaService(ioc.Media())

	return DomainServiceIoC{
		media: mediaSv,
	}
}

func (ioc DomainServiceIoC) Media() mediaService.MediaServiceInterface {
	return ioc.media
}
