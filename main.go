package main

import (
	"bufio"
	"fmt"
	cursor "github.com/ProductionPanic/go-cursor"
	lg "github.com/charmbracelet/lipgloss"
	term "golang.org/x/term"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

// getArgs function to get command line arguments
func getArgs() []string {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		os.Exit(0)
	}
	return args
}

// getImage function to read an image file
func getImage(filePath string) image.Image {
	allowed_extensions := []string{".png", ".jpg", ".jpeg"}
	allowed := false
	for _, ext := range allowed_extensions {
		if ext == filePath[len(filePath)-len(ext):] {
			allowed = true
		}
	}
	if !allowed {
		fmt.Println("Error: File extension not allowed")
		os.Exit(1)
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}
	return img
}

// displayImage function to display the image
func displayImage(img image.Image) {
	max_w, max_h, err := term.GetSize(0)
	max_w /= 2
	if err != nil {
		fmt.Println("Error: Could not get terminal size")
		os.Exit(1)
	}

	imageWidth := img.Bounds().Dx()
	imageHeight := img.Bounds().Dy()

	pixels_per_w := math.Ceil(float64(imageWidth) / float64(max_w))
	pixels_per_h := math.Ceil(float64(imageHeight) / float64(max_h))
	biggest := math.Max(pixels_per_w, pixels_per_h)

	image_width_in_terminal := int(float64(imageWidth) / biggest)
	image_height_in_terminal := int(float64(imageHeight) / biggest)

	pixel := lg.NewStyle().Width(2).Height(1)
	var output []string
	for y := 0; y < image_height_in_terminal; y++ {
		var row []string
		y_ := int(float64(y) * biggest)
		for x := 0; x < image_width_in_terminal; x++ {
			x_ := int(float64(x) * biggest)
			r, g, b, _ := img.At(x_, y_).RGBA()
			hex := fmt.Sprintf("#%02x%02x%02x", r/256, g/256, b/256)
			pixel.Background(lg.Color(hex))
			row = append(row, pixel.Render(" "))
		}
		output = append(output, lg.JoinHorizontal(lg.Left, row...))
	}
	cursor.ClearScreen()
	cursor.Top()
	cursor.Hide()
	tw, th, _ := term.GetSize(0)
	fmt.Println(lg.Place(tw, th, lg.Center, lg.Center, lg.JoinVertical(lg.Top, output...)))
}

// waitForInput function to wait for user input
func waitForInput() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// resetTerminal function to reset the terminal
func resetTerminal() {
	cursor.Show()
	cursor.ClearScreen()
	cursor.Top()
}

/**
 * Imshow is a cli tool to display images in the terminal.
 */
func main() {
	// get args
	args := getArgs()
	if len(args) == 0 {
		fmt.Println("Error: No file path provided")
		os.Exit(1)
	}
	// get image
	img := getImage(args[0])

	// display image
	displayImage(img)

	// wait for user input
	waitForInput()

	// reset terminal
	resetTerminal()
}

func printHelp() {
	fmt.Println("Usage: imshow <file> (file must be a png or jpg)")
}
