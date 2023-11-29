package server

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mheers/k8s-secret-ui/frontend"
	"github.com/mheers/k8s-secret-ui/helpers"
	"github.com/mheers/k8s-secret-ui/k8s"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

const (
	apiEndpoint       = "/api"
	configMapBaseURL  = apiEndpoint + "/configs"
	secretBaseURL     = apiEndpoint + "/secrets"
	namespacesBaseURL = apiEndpoint + "/namespaces"
	settingsBaseURL   = apiEndpoint + "/settings"
)

var (
	kubeclient *kubernetes.Clientset

	// EmbedFrontendFiles holds the frontend files
	EmbedFrontendFiles = frontend.Assets()
)

type Server struct {
	namespaceRegexes []string
	configMapRegexes []string
	secretRegexes    []string
}

func NewServer(namespaceRegexes, configMapRegexes, secretRegexes []string) *Server {
	return &Server{
		namespaceRegexes: namespaceRegexes,
		configMapRegexes: configMapRegexes,
		secretRegexes:    secretRegexes,
	}
}

func (s *Server) Run() error {
	var err error
	kubeclient, err = k8s.GetK8sClient()
	if err != nil {
		return fmt.Errorf("error %s getting the k8s client", err.Error())
	}

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
	klog.Infof("Endpoint is http://localhost%s", hostPort)
	err = http.ListenAndServe(hostPort,
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		klog.Fatalf("Error %s starting the service.", err.Error())
	}

	return nil
}

func isAllowed(name string, regexes []string) bool {
	if len(regexes) == 0 {
		return true
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
