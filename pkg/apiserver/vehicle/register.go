package vehicle

import (
	"net/http"

	"github.com/emicklei/go-restful"

	"ooneko.github.com/vehicle-insight/pkg/constants"
	"ooneko.github.com/vehicle-insight/pkg/models/vehicle"
)

var webservice = restful.WebService{}

func AddToContainer() {
	webservice.Path("/vehicles").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	webservice.Route(webservice.GET("/").
		Doc("List all vehicles").
		Param(webservice.HeaderParameter("user", "the vehicle's owner")).
		Returns(http.StatusOK, constants.OK, vehicle.Vehicle{}))

	webservice.Route(webservice.POST("/").
		Doc("Create a vehicle").
		Param(webservice.HeaderParameter("user", "the vehicle's owner")).
		Param(webservice.BodyParameter("name", "The vehicle's name")).
		Returns(http.StatusOK, constants.OK, vehicle.Vehicle{}))

	webservice.Route(webservice.GET("/{name}").
		Doc("Describe the vehicle").
		Param(webservice.HeaderParameter("user", "the vehicle's owner")).
		Param(webservice.PathParameter("name", "The vehicle's name")).
		Returns(http.StatusOK, constants.OK, vehicle.Vehicle{}))

	webservice.Route(webservice.DELETE("/{name}").
		Doc("Delete car").
		Param(webservice.HeaderParameter("user", "the vehicle's owner")).
		Param(webservice.PathParameter("name", "The vehicle's name")).
		Returns(http.StatusOK, constants.OK, vehicle.Vehicle{}))

	webservice.Route(webservice.GET("/{name}/maintenance").
		Doc("List All maintenance record").
		Param(webservice.HeaderParameter("user", "the vehicle's owner")).
		Param(webservice.PathParameter("name", "The vehicle's name")).
		Returns(http.StatusOK, constants.OK, vehicle.Vehicle{}))

	webservice.Route(webservice.POST("/{name}/maintenance").
		Doc("Create a maintenance record").
		Param(webservice.HeaderParameter("user", "the vehicle's owner")).
		Param(webservice.PathParameter("name", "The vehicle's name")).
		Returns(http.StatusOK, constants.OK, vehicle.Vehicle{}))

	webservice.Route(webservice.GET("/{name}/maintenance/{id}").
		Doc("Describe maintenance record").
		Param(webservice.HeaderParameter("user", "the vehicle's owner")).
		Param(webservice.PathParameter("name", "The vehicle's name")).
		Param(webservice.PathParameter("id", "The record's id")).
		Returns(http.StatusOK, constants.OK, vehicle.Vehicle{}))

	webservice.Route(webservice.DELETE("/{name}/maintenance/{id}").
		Doc("Delete a maintenance record").
		Param(webservice.HeaderParameter("user", "the vehicle's owner")).
		Param(webservice.PathParameter("name", "The vehicle's name")).
		Param(webservice.PathParameter("id", "The record's id")).
		Returns(http.StatusOK, constants.OK, vehicle.Vehicle{}))
}
