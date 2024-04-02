package services

import (
	"errors"
	"io/ioutil"
)

type IconService struct {
}

func (s *IconService) GetIconURL(os, icon string) ([]byte, error) {
	iconPath := ""
	switch os {
	case "windows":
		iconPath = "./images/lowres/" + icon + ".png"
	case "linux":
		iconPath = "./images/lowres/" + icon + ".png"
	case "macos":
		iconPath = "./images/highres/" + icon + ".png"
	case "android":
		iconPath = "./images/highres/" + icon + ".png"
	case "ios":
		iconPath = "./images/highres/" + icon + ".png"
	case "web":
		iconPath = "./images/highres/" + icon + ".png"
	}

	if iconPath == "" {
		return nil, errors.New("invalid os")
	}

	img, err := ioutil.ReadFile(iconPath)
	if err != nil {
		return nil, err
	}

	return img, nil
}
