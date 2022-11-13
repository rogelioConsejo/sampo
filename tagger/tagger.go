package tagger

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/google/uuid"
	"image/png"
	"os"
	"sampo/inventory"
	"sampo/inventory/tag"
)

func New() Tagger {
	return tagger{
		tag: tag.New(inventory.Item{Id: inventory.Id(uuid.NewString())}),
	}
}

type Tagger interface {
	Tag() tag.Tag
	MakeQR() (FilePath, error)
}

type tagger struct {
	tag tag.Tag
}

func (t tagger) MakeQR() (FilePath, error) {
	qrFileName := FilePath(t.tag.Item().Id + ".png")
	qrCode, err := qr.Encode(string(t.tag.Item().Id), qr.H, qr.Auto)
	if err != nil {
		return "", fmt.Errorf("could not encode qr: %w", err)
	}
	qrCode, err = barcode.Scale(qrCode, 600, 600)
	if err != nil {
		return "", fmt.Errorf("could not scale qr: %w", err)
	}
	file, err := os.Create(string(qrFileName))
	if err != nil {
		return "", fmt.Errorf("could not open file to save qr: %w", err)
	}
	defer file.Close()
	err = png.Encode(file, qrCode)
	if err != nil {
		return "", fmt.Errorf("could not save qr: %w", err)
	}
	return qrFileName, nil
}

func (t tagger) Tag() tag.Tag {
	return t.tag
}

type FilePath string
