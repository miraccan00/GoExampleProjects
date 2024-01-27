package image

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Image struct {
	ServiceName string `yaml:"service_name"`
	Server      string `yaml:"server"`
	Version     string `yaml:"stable_version"`
}

type ImageList struct {
	Images []Image `yaml:"images"`
}

func ImageGet() ImageList {
	yamlFile, err := os.ReadFile("image/version.yaml")
	if err != nil {
		fmt.Printf("Dosya okuma hatası: %v\n", err)
		return ImageList{}
	}

	var images ImageList

	err = yaml.Unmarshal(yamlFile, &images)
	if err != nil {
		fmt.Printf("YAML unmarshalling hatası: %v\n", err)
		return ImageList{}
	}

	// for _, img := range images.Images {
	// 	fmt.Printf("Image: %s, Version: %s\n", img.ServiceName, img.Version)
	// }
	return images
}

func GetLatestVersion(imageName string) string {
	latestVersion := "1.0.0" // Replace this with logic to fetch the latest version
	return latestVersion
}
