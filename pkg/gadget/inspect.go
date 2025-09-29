package gadget

import (
	"archive/tar"
	"io"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)
// Need to find a way to make this faster ...
func InspectFiles(layers []v1.Layer) ([]FileInfo, error) {
    const maxLicenseRead = 16 * 1024 // 16KB max amount to read
    files := make([]FileInfo, 0, 2000)

    for idx, layer := range layers {
        rc, err := layer.Uncompressed()
        if err != nil {
            return nil, err
        }
        tr := tar.NewReader(rc)

        for {
            hdr, err := tr.Next()
            if err == io.EOF {
                break
            }
            if err != nil {
                return nil, err
            }

            info := FileInfo{
                Path:  hdr.Name,
                LayerIdx: idx,
                Size:  hdr.Size,
                UID:   hdr.Uid,
                GID:   hdr.Gid,
            }

            if LooksLikeLicense(hdr.Name) && hdr.Size < 200*1024 {
                buf := make([]byte, maxLicenseRead)
                n, _ := io.ReadFull(tr, buf)
                data := buf[:n]

                if looksText(data) {
                    info.IsLicense = true
                    info.LicenseInfo = ChecksLicenseFile(data)
                }
            }

            files = append(files, info)
        }
        rc.Close()
    }
    return files, nil
}
func looksText(data []byte) bool {
    for _, b := range data {
        if b < 9 || (b > 13 && b < 32) {
            return false
        }
    }
    return true
}

