package container

import (
	"sync"
)

var ioc IoC
var iocSingleton sync.Once

type IoC struct {
	Application ApplicationServiceIoC
	Domain      DomainServiceIoC
	Repository  RepositoryIoC
}

func (ioc IoC) IsEmpty() bool {
	return (IoC{}) == ioc
}

func NewIOC() IoC {
	iocSingleton.Do(func() {
		repository := NewRepositoryIoC()
		domain := NewDomainServiceIoC(repository)
		application := NewApplicationServiceIoC(domain, repository)

		ioc = IoC{
			Application: application,
			Domain:      domain,
			Repository:  repository,
		}
	})

	return ioc
}

func Injector() IoC {
	if ioc.IsEmpty() {
		ioc = NewIOC()
	}

	return ioc
}
