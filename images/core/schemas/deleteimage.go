package schemas

type DeleteImg struct {
	UrlPhoto string `json:"url_photo,omitempty" bson:"url_photo,omitempty"`
	Validate bool   `json:"validate,omitempty"`
}
