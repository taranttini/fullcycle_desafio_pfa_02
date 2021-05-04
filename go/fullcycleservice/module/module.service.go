package module

import (
	"encoding/json"
	"fmt"
	"fullcycleservice/cors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const modulesBasePath = "modules"

func modulesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:

		// filtrando valores
		name := r.URL.Query().Get("name")
		active := r.URL.Query().Get("active")

		if len(name) > 0 || len(active) > 0 {

			var moduleFilter = ModuleFilter{
				NameFilter:   name,
				ActiveFilter: active,
			}
			moduleList, err := searchForModuleData(moduleFilter)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			modulesJSON, err := json.Marshal(moduleList)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(modulesJSON)
			return
		}
		// retornando todos os valores
		moduleList, err := getModules()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		modulesJSON, err := json.Marshal(moduleList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(modulesJSON)
	case http.MethodPost:
		var module Module
		err := json.NewDecoder(r.Body).Decode(&module)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = insertModule(module)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func moduleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlPathSegments := strings.Split(r.URL.Path, "modules/")
	moduleID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		module, err := getModule(moduleID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if module == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		moduleJSON, err := json.Marshal(module)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(moduleJSON)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPut:
		var module Module
		err := json.NewDecoder(r.Body).Decode(&module)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		module.ModuleID = moduleID

		err = updateModule(module)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		removeModule(moduleID)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	handleModules := http.HandlerFunc(modulesHandler)
	handleModule := http.HandlerFunc(moduleHandler)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, modulesBasePath), cors.Middleware(handleModules))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, modulesBasePath), cors.Middleware(handleModule))
}
