package dto

type UpdateNewsletterDTO struct {
	Name          string
	Text          string
	UsersLists    []int64
	MediaID       string
	ContentTypeID int64
}
