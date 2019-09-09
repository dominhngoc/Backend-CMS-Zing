package response

import "backend-cms-zing/CmsZing/models"

type Response struct {
	Code            string
	Message         interface{}
	Data            interface{}
	DetailErrorCode interface{}
	Success         string
}
type ArtistSwaggerResponse struct {
	Code            string
	Message         interface{}
	Data           []models.ArtistsSwagger
	DetailErrorCode interface{}
	Success         string
}