package main

import "flag"

type Flags struct {
	Width     int
	Height    int
	Output    string
	Directory string
	Seed      string
	Id        string
	Grayscale bool
	Blur      int
	Webp      bool
	Count     int
}

func NewFlags() *Flags {
	width := flag.Int("w", 800, "Width of the image")
	height := flag.Int("h", 600, "Height of the image")
	output := flag.String("o", "output", "Output file name")
	directory := flag.String("d", "", "Output directory")

	seed := flag.String("s", "", "Seed for the image generation")
	id := flag.String("id", "", "ID for the image generation")
	grayscale := flag.Bool("g", false, "Generate a grayscale image")
	blur := flag.Int("b", 0, "Generate a blurred image")
	webp := flag.Bool("webp", false, "Generate a WebP image")
	count := flag.Int("c", 1, "Number of images to generate")

	flag.Parse()

	return &Flags{
		Width:     max(*width, 0),
		Height:    max(*height, 0),
		Output:    *output,
		Directory: *directory,
		Seed:      *seed,
		Id:        *id,
		Grayscale: *grayscale,
		Blur:      between(0, 10, *blur),
		Webp:      *webp,
		Count:     max(*count, 0),
	}
}

func between(lower, upper, value int) int {
	return max(lower, min(upper, value))
}
