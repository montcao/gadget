package gadget

import (
	"sort"
)

func SortFiles(files []FileInfo) []FileInfo {
	sort.Slice(files, func(i, j int) bool { return files[i].Size > files[j].Size })
	return files
}

// O(n)
func GetLargestFile(files []FileInfo) FileInfo {
	if len(files) == 0 {
		return FileInfo{}
	}
	max := files[0]
	for _, f := range files[1:] {
		if f.Size > max.Size {
			max = f
		}
	}
	return max
}

// Return largest layer
func GetLargestLayer(img ImageInfo) LayerInfo {
	layers := img.Layers
	max := layers[0]
	for _, l := range layers[1:]{
		if l.Size > max.Size {
			max = l
		}
	}
	return max
}
