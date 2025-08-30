package dto

import (
	"time"
)

type NewslettersStatus int8

const (
	UnKnown NewslettersStatus = iota
	// Draft черновик
	Draft
	// Pending отправляется
	Pending
	// Sent отправлено
	Sent
)

func (ns NewslettersStatus) String() string {
	switch ns {
	case Draft:
		return "Черновик"
	case Pending:
		return "Отправляется"
	case Sent:
		return "Отправлено"
	}
	return ""
}

type Newsletter struct {
	ID              int64
	RecipientsCount int64
	Text            string
	UsersLists      []int64
	UsersIds        []int64
	MediaID         string
	CreatedAt       time.Time
	SentAt          *time.Time
	Name            string
	StatusID        NewslettersStatus
	ContentType     ContentType
}
