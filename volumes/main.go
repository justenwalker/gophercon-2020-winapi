package main

import (
	"fmt"
	"log"
)

func main() {
	vis, err := EnumerateVolumeInfo()
	if err != nil {
		log.Fatalln("EnumerateVolumeInfo error", err)
	}
	for _, vi := range vis {
		fmt.Println("Volume:", vi.Volume)
		fmt.Println("DOS Device:", vi.DOSDevice)
		if len(vi.Paths) > 0 {
			fmt.Println("Paths:")
			for _, path := range vi.Paths {
				fmt.Println(" -", path)
			}
		}
		fmt.Println("")
	}
}
