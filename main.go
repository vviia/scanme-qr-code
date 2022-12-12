package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// Create the barcode, replace https://github.com/vviia with your url
	qrCode, _ := qr.Encode("https://github.com/vviia", qr.M, qr.Auto)

	// Scale the barcode to 300x300 pixels
	qrCode, _ = barcode.Scale(qrCode, 300, 300)

	// create the output file
	file, _ := os.Create("scanme/qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}
