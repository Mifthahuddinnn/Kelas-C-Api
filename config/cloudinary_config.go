package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

var CloudinaryURL = "cloudinary://499887531591474:HOw0nGMhlVH_bY3EeavECImE8pE@dwggdy8os"

func InitCloudinary() (*cloudinary.Cloudinary, error) {
	cloudinaryService, err := cloudinary.NewFromURL(CloudinaryURL)
	if err != nil {
		return nil, err
	}

	return cloudinaryService, nil
}
