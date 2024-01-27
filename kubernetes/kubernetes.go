package kubernetes

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/miraccan00/blacksyriuscontroller/image" // image paketini içe aktar

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func ListDeployments() {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		fmt.Println("Kubeconfig dosyası bulunamadı.")
		return
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("Hata oluştu: %v\n", err)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Kubernetes istemci oluşturulamadı: %v\n", err)
		return
	}

	deployments, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Deployment'lar listelenemedi: %v\n", err)
		return
	}

	imageInfo := image.ImageGet()

	// fmt.Println("Deployment'lar listeleniyor...")
	for _, deployment := range deployments.Items {
		// fmt.Printf("Namespace: %s, Deployment: %s\n", deployment.Namespace, deployment.Name)
		containers := deployment.Spec.Template.Spec.Containers

		for _, container := range containers {
			imageTag, err := GetImageTag(container.Image)
			// fmt.Printf("Image Tag: %s\n", imageTag)
			if err != nil {
				fmt.Printf("Error retrieving image tag for %s: %v\n", container.Image, err)
				continue
			}
			for _, img := range imageInfo.Images {
				if img.ServiceName == container.Name {
					if img.Version != imageTag {
						fmt.Printf("Guncelleme mevcut\n")
						fmt.Printf(" Namespace: %s Image: %s - Stable Versiyon: %s, Kullanılan Versiyon: %s\n", deployment.Namespace, container.Image, img.Version, imageTag)

					} else {
						fmt.Printf("  Image: %s - Versiyon: %s\n", container.Image, imageTag)
					}
				}
			}
		}
	}
}

func GetImageTag(imageURL string) (string, error) {
	parts := strings.Split(imageURL, ":")
	if len(parts) < 2 {
		return "", fmt.Errorf("image tag not found in %v", imageURL)
	}
	return parts[len(parts)-1], nil
}
