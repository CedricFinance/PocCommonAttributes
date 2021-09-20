package main

import (
    "flag"
    "gopkg.in/yaml.v3"
    "html/template"
    "os"
)

type CommonAttributes struct {
    Attributes []Attribute
}

type Attribute struct {
    Name string
    Description string
    Type string
}

var destination string
var config string
var templatePath string

func init() {
    flag.StringVar(&destination, "destination", "out", "The folder where the files are generated")
    flag.StringVar(&config, "config", "config.yaml", "The file containing all the common attributes")
    flag.StringVar(&templatePath, "template", "java.tmpl", "The path to the template to use for code generation")
}

func main() {
    flag.Parse()
    
    commonAttributes := CommonAttributes{}
    data, err := os.ReadFile(config)
    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal(data, &commonAttributes)
    if err != nil {
        panic(err)
    }

    templateData, err := os.ReadFile(templatePath)
    if err != nil {
        panic(err)
    }
    
    tmpl := template.Must(template.New("").Parse(string(templateData)))
    
    err = tmpl.ExecuteTemplate(os.Stdout, "", commonAttributes.Attributes)
    if err != nil {
        panic(err)
    }
}
