package server

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/mheers/k8s-secret-ui/frontend"
	"github.com/mheers/k8s-secret-ui/helpers"
	"github.com/mheers/k8s-secret-ui/k8s"
	"github.com/mheers/k8s-secret-ui/pkg/util"
	"k8s.io/klog/v2"
)

const (
	apiEndpoint       = "/api"
	configMapBaseURL  = apiEndpoint + "/configs"
	secretBaseURL     = apiEndpoint + "/secrets"
	namespacesBaseURL = apiEndpoint + "/namespaces"
	settingsBaseURL   = apiEndpoint + "/settings"
)

var (
	// EmbedFrontendFiles holds the frontend files
	EmbedFrontendFiles = frontend.Assets()
)

type Server struct {
	manager          *util.Manager
	namespaceRegexes []string
	configMapRegexes []string
	secretRegexes    []string
}

func NewServer(namespaceRegexes, configMapRegexes, secretRegexes []string) (*Server, error) {
	kubeclient, err := k8s.GetK8sClient()
	if err != nil {
		return nil, fmt.Errorf("error %s getting the k8s client", err.Error())
	}

	manager := util.NewManager(kubeclient)

	return &Server{
		namespaceRegexes: namespaceRegexes,
		configMapRegexes: configMapRegexes,
		secretRegexes:    secretRegexes,
		manager:          manager,
	}, nil
}

func (s *Server) Run() error {

	router := mux.NewRouter()

	// frontend static files
	// redirect all not-found-requests to index.html
	fileSystem404 := func(w http.ResponseWriter, r *http.Request) (doDefaultFileServe bool) {
		//if not found redirect to /
		r.URL.Path = "/"
		return true
	}

	router.HandleFunc(namespacesBaseURL, s.getNamespaces).Methods("GET")
	router.HandleFunc(configMapBaseURL+"/{cmns}", s.getConfigMapsOfNS).Methods("GET")
	router.HandleFunc(configMapBaseURL+"/{cmns}/{cmname}", s.getConfigMap).Methods("GET")
	router.HandleFunc(configMapBaseURL+"/{cmns}/{cmname}", s.updateConfigMap).Methods("PUT")
	router.HandleFunc(configMapBaseURL+"/{cmns}/{cmname}", s.deleteConfigMap).Methods("DELETE")
	router.HandleFunc(configMapBaseURL, s.createConfigMap).Methods("POST")

	router.HandleFunc(secretBaseURL+"/{secretns}", s.getSecretsOfNS).Methods("GET")
	router.HandleFunc(secretBaseURL+"/{secretns}/{secretname}", s.getSecret).Methods("GET")
	router.HandleFunc(secretBaseURL+"/{secretns}/{secretname}", s.updateSecret).Methods("PUT")
	router.HandleFunc(secretBaseURL+"/{secretns}/{secretname}", s.deleteSecret).Methods("DELETE")
	router.HandleFunc(secretBaseURL, s.createSecret).Methods("POST")

	router.PathPrefix("/").Handler(helpers.FileServerWith404(http.FS(EmbedFrontendFiles), fileSystem404))

	hostPort := ":8000"
	// allow CORS
	klog.Infof("Local endpoint is http://localhost%s", hostPort)
	err := http.ListenAndServe(hostPort, router)
	if err != nil {
		klog.Fatalf("Error %s starting the service.", err.Error())
		return err
	}

	return nil
}

func isAllowed(name string, regexes []string) bool {
	if len(regexes) == 0 {
		return false
	}

	for _, regex := range regexes {
		if matchRegex(name, regex) {
			return true
		}
	}
	return false
}

func matchRegex(name, regex string) bool {
	if regex == "*" {
		return true
	}

	if regexp.MustCompile(regex).MatchString(name) {
		return true
	}

	return false
}
