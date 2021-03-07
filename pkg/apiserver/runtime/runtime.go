package runtime

import (
	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	APIRootPath = "/"
)

var Container = restful.NewContainer()

func NewWebService(gv schema.GroupVersion) *restful.WebService {
	webservice := &restful.WebService{}
	webservice.Path(APIRootPath + "/" + gv.String()).
		Produces(restful.MIME_JSON)
	return webservice
}
