// Package tagger is used to create individual QR tags for non-fungible items
package tagger

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/google/uuid"
	"image/png"
	"os"
	"sampo/asset/item"
	"sampo/asset/item/tag"
)

// New creates a tagger for an item, with an automatically generated UUID
func New() Tagger {
	return tagger{
		tag: tag.New(item.Item{Id: item.Id(uuid.NewString())}),
	}
}

// A Tagger creates a UUID and the associated QR code to tag a non-fungible item for storage tracking.
type Tagger interface {
	MakeQR(FilePath) error
}

type tagger struct {
	tag tag.Tag
}

func (t tagger) getTag() tag.Tag {
	return t.tag
}

// MakeQR takes the UUID created when the tagger was instantiated with New(), creates a QR code and saves it to a
// specified FilePath
func (t tagger) MakeQR(qrFileName FilePath) error {
	qrCode, err := t.createQRCode()
	if err != nil {
		return fmt.Errorf("error thrown when creating QR Code:%w", err)
	}
	err = createPNGFile(qrFileName, qrCode)
	if err != nil {
		return fmt.Errorf("error thrown when creating PNG File: %w", err)
	}
	return nil
}

func createPNGFile(qrFileName FilePath, qrCode barcode.Barcode) error {
	file, err := os.Create(string(qrFileName))
	if err != nil {
		return fmt.Errorf("could not open file to save qr: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	err = png.Encode(file, qrCode)
	if err != nil {
		return fmt.Errorf("could not save qr: %w", err)
	}
	return nil
}

func (t tagger) createQRCode() (barcode.Barcode, error) {
	qrCode, err := qr.Encode(string(t.tag.Item().Id), qr.H, qr.Auto)
	if err != nil {
		return nil, fmt.Errorf("could not encode qr: %w", err)
	}
	qrCode, err = barcode.Scale(qrCode, 600, 600)
	if err != nil {
		return nil, fmt.Errorf("could not scale qr: %w", err)
	}
	return qrCode, nil
}

type FilePath string
