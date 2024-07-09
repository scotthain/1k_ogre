/*
Copyright © 2024 Scott Hain <elejia@gmail.com>
*/
package main

import (
	"fmt"
	"os"
)

// / This represents our lovely artifact
type artifact struct {
	Name         string     `json:"name"`
	Version      string     `json:"versions"`
	Source       string     `json:"source"`
	Dependencies []artifact `json:"dependencies"`
}

func main() {
	// incoming artifact
	var inc_artifact artifact
	inc_artifact.Name = "loggerbot"
	inc_artifact.Version = "1.5.4"
	inc_artifact.Source = "http://github.com/somethwwer/loggerbot.zip"
	inc_artifact.Dependencies = nil

	// dependencies
	var deps = flense(inc_artifact)
	inc_artifact.Dependencies = append(inc_artifact.Dependencies, deps...)

	print("WHAT IS ", inc_artifact.Dependencies[0].Name)

	dat, err := os.ReadFile("./fixtures/tarball.")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat))
}

// first thing to be called after receiving an object
func parse_artifact(inc_artifact artifact) artifact {

	return artifact{}
}

func flense(inc_artifact artifact) []artifact {
	// magic happens, slice them open
	//incoming_artifact.

	var dependencies []artifact

	var dep1 = artifact{Name: "colorfultext", Version: "2.3.0", Source: "http://github.com/somethwwer/colorfultext.zip", Dependencies: nil}
	dependencies = append(dependencies, dep1)
	var dep2 = artifact{Name: "loggerbot", Version: "1.5.4", Source: "http://github.com/somethwwer/loggerbot.zip", Dependencies: nil}
	dependencies = append(dependencies, dep2)

	return dependencies
}
