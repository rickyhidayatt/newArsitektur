package container

import (
	"bni.co.id/xpora/medias/internal/application"
)

type ApplicationServiceIoC struct {
	application application.Application
}

func NewApplicationServiceIoC(dsIoc DomainServiceIoC, rIoc RepositoryIoC) ApplicationServiceIoC {
	return ApplicationServiceIoC{
		application: application.New(
			dsIoc.media,
		),
	}
}

func (ioc ApplicationServiceIoC) Application() application.Application {
	return ioc.application
}
