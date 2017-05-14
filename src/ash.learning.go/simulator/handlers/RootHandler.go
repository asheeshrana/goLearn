package handlers

import (
	"net/http"
	"os"
	"strings"
	"text/template"

	"ash.learning.go/configuration"
	"ash.learning.go/log"
)

//Test is just a test structure
type Test struct {
}

//RootHandler is provides default implementation for any incoming request
var RootHandler = func(writer http.ResponseWriter, request *http.Request) {
	var templateName string
	urlPath := strings.Trim(request.URL.Path, "/")
	log.Info.Println("urlpath = " + urlPath)

	urlParts := strings.Split(urlPath, "/")
	if urlParts == nil || len(urlParts) < 1 {
		templateName = "noTemplateFound"
	} else {
		lastPart := strings.Trim(urlParts[len(urlParts)-1], " ")
		if len(lastPart) > 0 {
			templateName = lastPart
		} else {
			templateName = "noTemplateFound"
		}
	}
	log.Info.Println("templateName = " + templateName)
	templatesDir := configuration.GetConfiguration().TemplatesDir
	templateFile := templatesDir + "/" + templateName + ".tmpl"

	if _, err := os.Stat(templateFile); os.IsNotExist(err) {
		log.Error.Println("Template file " + templateFile + " does not exist")
		templateFile = templatesDir + "/" + "noTemplateFound.tmpl"
		templateName = "noTemplateFound"
		log.Info.Println("Updated Template file " + templateFile)
	}

	templateToReturn := template.Must(template.New(templateName + ".tmpl").ParseFiles(templateFile))
	err := templateToReturn.Execute(writer, Test{})
	if err != nil {
		log.Info.Printf("Error = %v", err)
	}

	//writer.Write([]byte("URL Path = " + templateName))
}
