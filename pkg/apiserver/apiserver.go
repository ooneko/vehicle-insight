package apiserver

import (
	"net/http"

	urlruntime "k8s.io/apimachinery/pkg/util/runtime"
	"github.com/emicklei/go-restful"
	"go.etcd.io/etcd/clientv3"
)

type APIServer struct {
	Server     *http.Server
	Container  *restful.Container
	EtcdClinet *clientv3.Client
}

func (s *APIServer) PrepareRun() error {
	s.Container = restful.NewContainer()
	//TODO Add logging filter
	s.Container.Router(restful.CurlyRouter{})
	//TODO add logstack on recover

}


func (s *APIServer)installAPIs(){
	urlruntime.Must()
}