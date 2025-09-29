package gadget

import (
	"strings"
	"github.com/google/licensecheck"
)

// https://pkg.go.dev/github.com/google/licensecheck#section-documentation


// Heuristic check for license-related files
func LooksLikeLicense(path string) bool {
	l := strings.ToLower(path)

	// ignore /usr/share/common-licenses
	if strings.Contains(l, "common-licenses"){
		return false
	}

	if strings.Contains(l, "license") ||
		strings.Contains(l, "copying") ||
		strings.Contains(l, "copyright") ||
		strings.Contains(l, "notice") {
		return true
	}

	return false
}

func ChecksLicenseFile(data []byte) licensecheck.Coverage {
	// The Scan function uses a built-in license set, which is the known SPDX licenses augmented with some other commonly seen licenses. (See licenses/README.md for details about the license set.) 
	cov := licensecheck.Scan(data)
	return cov
}

func GetImageLicenses(files []FileInfo) []ImageLicenseInfo {
	counts := make(map[string]int)
	for i := 0; i < len(files); i++ { // && i < 5
		if files[i].IsLicense{
			licenseInfo := files[i].LicenseInfo
			for _, m := range licenseInfo.Match {
					counts[m.ID]++
				}
			}
	}
	licenses := make([]ImageLicenseInfo, 0, len(counts))
	for id, count := range counts {
		licenses = append(licenses,ImageLicenseInfo{ID: id, Count: count})
	}
	return licenses
}