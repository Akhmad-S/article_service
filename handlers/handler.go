package handlers

import "github.com/uacademy/blogpost/article_service/storage"

type Handler struct {
	Stg storage.StorageI
}
