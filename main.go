package main

import (
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Packages struct {
	Package   string `yaml:"Package"`
	Operation string `yaml:"Operation"`
}

const (
	pkgDefinitionPath      = "packages.yaml"
	dockerfileTemplatePath = "Dockerfile.tmpl"
	dockerfileName         = "Dockerfile"
)

var selfFunc = template.FuncMap{
	"add1": add1,
}

func add1(x int) int {
	return x + 1
}

func main() {
	var pkgOps []string
	// Read data from YAML file
	f, err := os.Open(pkgDefinitionPath)
	if err != nil {
		fmt.Printf("Failed to open YAML file %s, err: %v", pkgDefinitionPath, err)
		os.Exit(1)
	}
	defer f.Close()

	var packageSet []Packages
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&packageSet)
	if err != nil {
		fmt.Printf("Failed to decode YAML, err: %v", err)
		os.Exit(1)
	}
	// convert once then giving it to the template
	for _, pkg := range packageSet {
		if pkg.Operation == "install" {
			pkgOps = append(pkgOps, pkg.Package)
		}
	}

	// Read the template file
	templateContent, err := os.ReadFile(dockerfileTemplatePath)
	if err != nil {
		fmt.Printf("Failed to read template file reading template file %s, err: %v", dockerfileTemplatePath, err)
		os.Exit(1)
	}

	// Parse the template
	dockerfileTmpl, err := template.New("DockerfileTemplate").Funcs(selfFunc).Parse(string(templateContent))
	if err != nil {
		fmt.Printf("Failed to parse the tempate, err: %v.", err)
		os.Exit(1)
	}

	dstDockerfile, err := os.Create(dockerfileName)
	if err != nil {
		fmt.Printf("Failed to create file %s, err: %v", dockerfileName, err)
		os.Exit(1)
	}
	defer dstDockerfile.Close()

	fmt.Printf("Wanted Packages: %v\n", pkgOps)
	err = dockerfileTmpl.Execute(dstDockerfile, pkgOps)
	if err != nil {
		fmt.Printf("Failed to execute the template, err: %v", err)
		os.Exit(1)
	}

	fmt.Println("Dockerfile generated successfully!")
}
