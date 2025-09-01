package get_content_types

type ContentType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	ContentTypes []ContentType `json:"content_types"`
}
