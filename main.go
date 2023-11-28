package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"flag"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mheers/k8s-secret-ui/frontend"
	"github.com/mheers/k8s-secret-ui/helpers"
	"github.com/mheers/k8s-secret-ui/pkg/util"
	corev1 "k8s.io/api/core/v1"
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
	masterURL  string
	kubeconfig string
	kubeclient *kubernetes.Clientset

	// EmbedFrontendFiles holds the frontend files
	EmbedFrontendFiles = frontend.Assets()
)

func main() {
	flag.Parse()
	var err error
	kubeclient, err = getK8sClient()
	if err != nil {
		klog.Fatalf("Error %s getting the k8s client.", err.Error())
	}

	// get namespaces
	namespaces := util.GetNamespaces(kubeclient)
	fmt.Println(namespaces)

	router := mux.NewRouter()

	// frontend static files
	// redirect all not-found-requests to index.html
	fileSystem404 := func(w http.ResponseWriter, r *http.Request) (doDefaultFileServe bool) {
		//if not found redirect to /
		r.URL.Path = "/"
		return true
	}

	router.HandleFunc(settingsBaseURL, getSettings).Methods("GET")

	router.HandleFunc(namespacesBaseURL, getNamespaces).Methods("GET")
	router.HandleFunc(configMapBaseURL+"/{cmns}", getConfigMapsOfNS).Methods("GET")
	router.HandleFunc(configMapBaseURL+"/{cmns}/{cmname}", getConfigMap).Methods("GET")
	router.HandleFunc(configMapBaseURL+"/{cmns}/{cmname}", updateConfigMap).Methods("PUT")
	router.HandleFunc(configMapBaseURL+"/{cmns}/{cmname}", deleteConfigMap).Methods("DELETE")
	router.HandleFunc(configMapBaseURL, createConfigMap).Methods("POST")

	router.HandleFunc(secretBaseURL+"/{secretns}", getSecretsOfNS).Methods("GET")
	router.HandleFunc(secretBaseURL+"/{secretns}/{secretname}", getSecretData).Methods("GET")
	router.HandleFunc(secretBaseURL+"/{secretns}/{secretname}", updateSecret).Methods("PUT")
	router.HandleFunc(secretBaseURL+"/{secretns}/{secretname}", deleteSecret).Methods("DELETE")
	router.HandleFunc(secretBaseURL, createSecret).Methods("POST")

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

}

func createSecret(w http.ResponseWriter, r *http.Request) {
	var secret corev1.Secret
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&secret)

	res := util.CreateSecret(kubeclient, secret)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		klog.Errorf("Error while encoding the secret: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func createConfigMap(w http.ResponseWriter, r *http.Request) {
	var configMap corev1.ConfigMap
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&configMap)

	res := util.CreateConfigMap(kubeclient, configMap)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		klog.Errorf("Error while encoding the configmap: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func deleteSecret(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	secretName := pathParams["secretname"]
	secretNS := pathParams["secretns"]

	res := util.DeleteSecret(kubeclient, secretNS, secretName)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		klog.Errorf("Error while encoding the secret: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func deleteConfigMap(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	cmName := pathParams["cmname"]
	cmNS := pathParams["cmns"]
	res := util.DeleteConfigMap(kubeclient, cmNS, cmName)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		klog.Errorf("Error while encoding the configmap: %s", err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

func updateSecret(res http.ResponseWriter, req *http.Request) {
	pathParams := mux.Vars(req)
	secretName := pathParams["secretname"]
	secretNS := pathParams["secretns"]

	var secret corev1.Secret
	err := json.NewDecoder(req.Body).Decode(&secret)
	if err != nil {
		klog.Errorf("Error while decoding the secret: %s", err.Error())
		res.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(res).Encode(util.UpdateSecret(kubeclient, secretNS, secretName, secret))
	if err != nil {
		klog.Errorf("Error while encoding the secret: %s", err.Error())
		res.Write([]byte(err.Error()))
		return
	}
}

func getSecretData(res http.ResponseWriter, req *http.Request) {
	queryParams := mux.Vars(req)
	secretName := queryParams["secretname"]
	secretNS := queryParams["secretns"]

	err := json.NewEncoder(res).Encode(util.GetSecretData(kubeclient, secretNS, secretName))
	if err != nil {
		klog.Errorf("Error while encoding the secret data: %s", err.Error())
		res.Write([]byte(err.Error()))
		return
	}
}

func getSecretsOfNS(res http.ResponseWriter, req *http.Request) {
	queryParams := mux.Vars(req)
	namespace := queryParams["secretns"]
	err := json.NewEncoder(res).Encode(util.GetSecretsOfNS(kubeclient, namespace))
	if err != nil {
		klog.Errorf("Error while encoding the secrets: %s", err.Error())
		res.Write([]byte(err.Error()))
		return
	}
}

func getConfigMapsOfNS(res http.ResponseWriter, req *http.Request) {
	queryParams := mux.Vars(req)
	namespace := queryParams["cmns"]
	err := json.NewEncoder(res).Encode(util.GetConfigMapsOfNS(kubeclient, namespace))
	if err != nil {
		klog.Errorf("Error while encoding the configmaps: %s", err.Error())
		res.Write([]byte(err.Error()))
		return
	}
}

func getNamespaces(res http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(res).Encode(util.GetNamespaces(kubeclient))
	if err != nil {
		klog.Errorf("Error while encoding the namespaces: %s", err.Error())
		res.Write([]byte(err.Error()))
		return
	}
}

func getConfigMap(res http.ResponseWriter, req *http.Request) {
	queryParams := mux.Vars(req)
	cmName := queryParams["cmname"]
	cmns := queryParams["cmns"]

	err := json.NewEncoder(res).Encode(util.GetConfigMap(kubeclient, cmns, cmName))
	if err != nil {
		klog.Errorf("Error while encoding the configmap: %s", err.Error())
		res.Write([]byte(err.Error()))
		return
	}
}

func updateConfigMap(res http.ResponseWriter, req *http.Request) {
	queryParams := mux.Vars(req)
	cmName := queryParams["cmname"]
	cmns := queryParams["cmns"]

	var configmap corev1.ConfigMap
	err := json.NewDecoder(req.Body).Decode(&configmap)
	if err != nil {
		klog.Errorf("Error while decoding the configmap: %s", err.Error())
		res.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(res).Encode(util.UpdateConfigMap(kubeclient, cmns, cmName, configmap))
	if err != nil {
		klog.Errorf("Error while encoding the configmap: %s", err.Error())
		res.Write([]byte(err.Error()))
		return
	}

}

func getSettings(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(GetSettings())
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to your kubeconfig file.")
	flag.StringVar(&masterURL, "masterurl", "", "URL of your kube-apiserver.")
}
