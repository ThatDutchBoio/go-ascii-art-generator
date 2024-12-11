package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"strconv"
)

func AppendCharacterToFile(filename string, char rune) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(string(char))
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
func getChar(brightness int) string {
	brightness = brightness / 10
	fmt.Println("brightness", brightness)
	chars := []string{" ", ".", ":", "-", "=", "+", "*", "#", "%", "@"}
	if brightness >= len(chars) {
		return "@"
	}

	switch brightness {
	case 0:
		return " "
	case 1:
		return "."
	}

	return chars[brightness]
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please two arguments, format:")
		fmt.Println("go run main.go <image_path> <resolution (parse every n pixel)>")
		return
	}
	resolution, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Please provide a valid resolution")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	bounds := img.Bounds()
	// var maxBrightness float64 = 42774765
	var art = make([][]string, bounds.Max.Y)
	var highest float64 = 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y += resolution {
		for x := bounds.Min.X; x < bounds.Max.X; x += resolution {
			r, g, b, _ := img.At(x, y).RGBA()

			var brightness float64 = (float64(100) / 65535) * ((float64(r) + float64(g) + float64(b)) / 3)
			if brightness > highest {
				highest = brightness
			}
			fmt.Println(brightness)

			art[y] = append(art[y], getChar(int(brightness)))
			AppendCharacterToFile("output.txt", rune(getChar(int(brightness))[0]))
			// fmt.Printf("x=%d, y=%d, r=%d, g=%d, b=%d, a=%d\n", x, y, r, g, b, a)

		}
		AppendCharacterToFile("output.txt", rune('\n'))
	}
	fmt.Println(highest)
	fmt.Println(art)
}
