package apiserver

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"go.uber.org/zap"
	"ooneko.github.com/vehicle-insight/cmd/app/options"
	"ooneko.github.com/vehicle-insight/pkg/apiserver/config"
	"ooneko.github.com/vehicle-insight/pkg/apiserver/version"
)

type APIServer struct {
	Server    *http.Server
	Container *restful.Container

	config config.Config
	Logger *zap.Logger
}

func NewMainAPIServer(opt *options.APIServer) *APIServer {
	return &APIServer{
		config: opt.Config,
		Logger: opt.Logger,
	}
}

func (s *APIServer) PrepareRun() error {
	s.Container = restful.NewContainer()
	//TODO Add logging filter
	s.Container.Router(restful.CurlyRouter{})
	//TODO add logstack on recover
	s.installAPIs()
	return nil
}

func (s *APIServer) installAPIs() {
	version.AddToContainer(s.Container)

	// vehicle api vehicle CRUD for users

	// rule api, add rule for vehicle

	// record api, Add/Remove a record for vehicle

}

func (s *APIServer) Run() error {
	s.PrepareRun()
	return nil
}
