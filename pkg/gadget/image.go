package gadget

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"log"
	"strings"
)


func InitImageInfo(img v1.Image, imageRef string) (ImageInfo, error) {
	// Get the name
	layers, err := ListLayers(img)
	if err != nil {
		log.Fatal(err)
	}
	total_size := int64(0)
	for i := 0; i < len(layers); i++ { // && i < 5
		size := layers[i].Size
		total_size = size + total_size
	}

	name := imageRef
	
	img_info := ImageInfo{
		Name: name,
		Layers: layers,
		TotalSize: total_size,
		V1Image: img,
	}
	return img_info, nil
}

// returns a v1.Image from google container registry
func PullImage(refStr string) (v1.Image, error) {
	ref, err := name.ParseReference(refStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse reference: %w", err)
	}
	img, err := remote.Image(ref)
	if err != nil {
		return nil, fmt.Errorf("failed to pull image: %w", err)
	}
	return img, nil
}

// Returns a list of LayerInfo
func ListLayers(img v1.Image) ([]LayerInfo, error) {
	layers, err := img.Layers()
	if err != nil {
		return nil, err
	}
	result := make([]LayerInfo, len(layers))
	for i, layer := range layers {
		digest, _ := layer.Digest()
		size, _ := layer.Size()
		result[i] = LayerInfo{
			Index:  i,
			Digest: digest.String(),
			Size:   size,
		}
	}
	return result, nil
}


func GetImageFromRef(imageRef string) (v1.Image, error) {
	var img v1.Image
	var err error
	if strings.HasSuffix(strings.ToLower(imageRef), ".tar"){
		img = InspectTar(imageRef)
	} else {
		img, err = PullImage(imageRef)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	if img == nil {
		log.Fatal("Failed to read image")
		return nil, err
	}
	return img, nil
}