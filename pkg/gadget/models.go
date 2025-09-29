package gadget
import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/licensecheck"
)

type LayerInfo struct {
	Index  int
	Digest string
	Size   int64
}

type ImageInfo struct{
	Name string
	TotalSize int64
	Layers []LayerInfo
	V1Image v1.Image
}

type FileInfo struct {
	Path     string
	LayerIdx int
	Size     int64
	Mode     uint32
	UID      int
	GID      int
	IsLicense bool
	LicenseInfo licensecheck.Coverage
}

type ImageLicenseInfo struct {
	ID string
	Count int 
}

// Coverage struct is as follows:
// https://pkg.go.dev/github.com/google/licensecheck
// Percent, the percent of the file that matches a license
// Match[], the matches across the text in the file is an array
//type Match struct {
// 	ID    string // License identifier. (See licenses/README.md.)
// 	Type  Type   // The type of the license: BSD, MIT, etc.
// 	Start int    // Start offset of match in text; match is at text[Start:End].
// 	End   int    // End offset of match in text.
// 	IsURL bool   // Whether match is a URL.
// }


