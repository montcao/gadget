package cmd

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/montcao/gadget/pkg/gadget"
)

// add the clear command to the root command
func addInspectCommand(root *cobra.Command) {
	inspectCmd := NewInspect()
	root.AddCommand(inspectCmd)
}

func NewInspect() *cobra.Command {
	inspectCmd := &cobra.Command{
		Use:   "inspect",
		Short: "Programatically inspect an image",
		RunE: func(cmd *cobra.Command, args []string) error {

				banner := gadget.CreateBanner()
				fmt.Println(banner)

				loading := gadget.CreateLoadingBar()
				fmt.Println(loading)
				// Testing
				//imageRef := "gcr.io/distroless/static:latest"
				//imageRef = "cerog/hashtopolis-nvidia-agent-lite12.0:v1.0.1"
				imageRef := args[0]

				// if tar, run a local check. Add more features and configurations for this later
				img, err := gadget.GetImageFromRef(imageRef)
				if err != nil {
					log.Fatal(err)
				}
				img_info, err := gadget.InitImageInfo(img, imageRef)
				v1_layers, err := img_info.V1Image.Layers()
				files, err := gadget.InspectFiles(v1_layers)
				if err != nil {
					log.Fatal(err)
				}
				
				// Name
				name := img_info.Name
				image_size := float64(img_info.TotalSize)/1024.0/1024.0 // MB
				num_layers := len(img_info.Layers)

				// Sort files by size
				files = gadget.SortFiles((files))


				
				// Print output

				// Table head
				head_t := gadget.CreateTable()
				head_t.Headers("Name", "Size", "# Layers")
				head_t.Row(name,  fmt.Sprintf("%.2f MB", image_size), fmt.Sprintf("%d", num_layers))
				fmt.Println(head_t)
			
				
				// Biggest layer
				fmt.Println("Largest layer: ")
				l_t := gadget.CreateTable()
				layer_info := gadget.GetLargestLayer(img_info)
				l_t.Headers("Digest",  "Size", "Layer Idx")
				layer_size := fmt.Sprintf("%.2f MB", float64(layer_info.Size)/1024.0/1024.0) // MB
				l_t.Row(layer_info.Digest, layer_size, fmt.Sprintf("%d",layer_info.Index))
				fmt.Println(l_t)
				// Top 5 files 
				fmt.Println("Top 20 largest files:")
				file_t := gadget.CreateTable()
				file_t.Headers("Path", "Size", "Layer Index")
				for i := 0; i < len(files) && i < 20; i++ { // && i < 5
					mb := fmt.Sprintf("%.2f MB" , (float64(files[i].Size) / 1024.0 / 1024.0))
					idx := fmt.Sprintf("%d", files[i].LayerIdx)
					file_t.Row(files[i].Path, mb, idx)
				}
				
				fmt.Println(file_t)

				
				// Print License Candidates
				fmt.Println("Image License Info (beta):")
				licenses_info := gadget.GetImageLicenses(files)
				license_t := gadget.CreateTable()
				license_t.Headers("License", "Count")
				for _, license := range licenses_info {
					license_t.Row(license.ID, fmt.Sprintf("%d", license.Count))
				}
				fmt.Println(license_t)
				
			return nil
		},
	}
	return inspectCmd
}

