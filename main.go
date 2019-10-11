package main

import (
  "fmt"
  "os"
  "log"
  "image"
  "image/jpeg"
  "github.com/disintegration/gift"
  "os/exec"
)

func handleError(err error) {
  if err != nil {
    log.Fatal(err)
  }
  return
}

func main() {
  filePath := os.Args[1]
  newFilePath := os.Args[2]

  if newFilePath == "" {
    newFilePath = "generated_file"
  }
  fmt.Printf("Attempting to read image from %s\n", filePath)

  imageFile, err := os.Open(filePath)
  handleError(err)

  img, format, err := image.Decode(imageFile)
  handleError(err)

  fmt.Printf("Read a %s image of format %s \n", filePath, format)

  g := gift.New(
    gift.Contrast(10),
    gift.Sepia(100),
    gift.Colorize(80.0, 80.0, 80.0),
    gift.Pixelate(8),
  )
  dst := image.NewRGBA(g.Bounds(img.Bounds()))
  g.Draw(dst, img)

  fullNewFilePath := newFilePath + "." + format
  f, err := os.Create(fullNewFilePath)
  handleError(err)
  defer f.Close()
  jpeg.Encode(f, dst, nil)

  fmt.Printf("Wrote an image to %s \n", fullNewFilePath)

  fmt.Printf("Opening newly create image \n")

  output, err := exec.Command("open", fullNewFilePath).CombinedOutput()
  if err != nil {
    os.Stderr.WriteString(err.Error())
  }
  fmt.Println(string(output))
  return
}
