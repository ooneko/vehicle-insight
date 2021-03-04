package version

import (
	"github.com/emicklei/go-restful"
	"ooneko.github.com/vehicle-insight/pkg/apiserver/runtime"
	"ooneko.github.com/vehicle-insight/pkg/version"
)

func AddToContainer(container *restful.Container) {
	webservice := runtime.NewWebService("version")
	webservice.Route(webservice.GET("/version").
		To(func(request *restful.Request, response *restful.Response) {
			version := version.Get()
			response.WriteAsJson(version)
		}).Doc("API Version"))
	container.Add(webservice)
}
