package webservicehandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webservice/k8sapicall"

	"k8s.io/apimachinery/pkg/api/errors"
)

type Jsonbody struct {
	Action string `json:"action"`
	Name   string `json:"name"`
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this path is invalid: %s!\n", r.URL.Path[1:])
	fmt.Fprintln(w, "valid paths are:")
	fmt.Fprintln(w, "/getpod")
	fmt.Fprintln(w, "/getnamespace")
	fmt.Fprintln(w, "/createnamespace")
	fmt.Fprintln(w, "/deletenamespace")
}

func Getpod(w http.ResponseWriter, r *http.Request) {

	podList, _ := k8sapicall.GetPods()
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(podList)

}

func Namespace(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "only POST request supported")
		return
	}

	body := Jsonbody{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		panic(err)
	}
	if body.Action != "get" && body.Name == "" {
		fmt.Fprintf(w, "name can not be empty for action: %s", body.Action)
		return
	}
	switch body.Action {
	case "get":
		Getnamespace(w, r)
	case "create":
		CreateNamespace(w, r, body.Name)
	case "delete":
		DeleteNamespace(w, r, body.Name)
	default:
		fmt.Fprintf(w, "This action is not supported: %s ", body.Action)
	}
}

func Getnamespace(w http.ResponseWriter, r *http.Request) {

	namespaceList, _ := k8sapicall.GetNamespace()
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(namespaceList)
}

func CreateNamespace(w http.ResponseWriter, r *http.Request, name string) {
	namespace, err := k8sapicall.CreateNamespace(name)
	if errors.IsAlreadyExists(err) {
		fmt.Fprintf(w, "namespace already exist: %s", "test")
	} else {
		fmt.Fprintf(w, "namespace created: %s", namespace.GetName())
	}
}

func DeleteNamespace(w http.ResponseWriter, r *http.Request, name string) {
	err := k8sapicall.DeleteNamespace(name)
	if err != nil {
		fmt.Fprintf(w, "error occured while deleting namespace: %v", err)
	} else {
		fmt.Fprintf(w, "namespace deleted: %s", "test")
	}
}
