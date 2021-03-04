package runtime

import "github.com/emicklei/go-restful"

const (
	APIRootPath = "/"
)

var Container = restful.NewContainer()

func NewWebService(path string) *restful.WebService {
	webservice := &restful.WebService{}
	webservice.Path(APIRootPath + "/" + path).
		Produces(restful.MIME_JSON)
	return webservice
}
