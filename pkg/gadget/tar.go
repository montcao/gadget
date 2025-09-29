// functionality to locally analyze a tar
package gadget 
import (
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)
func InspectTar(imageRef string) v1.Image {
	img, err := tarball.ImageFromPath(imageRef, nil)
	if err != nil {
		panic(err)
	}
	return img
}