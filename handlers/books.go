package handlers

import (
	"readdb/models"
	"time"

	"github.com/labstack/echo/v4"
)

func handleBooks(ctx echo.Context) error {
	resMock := models.Book{
		BookIDs: []models.BookID{
			{Type: "isbn10", Value: "0141439475"},
			{Type: "isbn13", Value: "978-0141439471"},
		},
		Name:           "Frankenstein",
		Editions:       []models.Edition{{Name: "Revised Edition", Number: 3, ProfileImage: "https://m.media-amazon.com/images/I/715sAi1-d5L._SY522_.jpg"}},
		ShortBlurb:     `Mary Shelley's chilling Gothic tale was conceived when she was only eighteen, living with her lover Percy Shelley on Lake Geneva. The story of Victor Frankenstein who, obsessed with creating life itself, plunders graveyards for the material to fashion a new being, but whose botched creature sets out to destroy his maker, would become the world's most famous work of horror fiction, and remains a devastating exploration of the limits of human creativity.`,
		Rating:         4,
		TotalPageCount: 352,
		ProfileImage:   "https://m.media-amazon.com/images/I/715sAi1-d5L._SY522_.jpg",
		Authors: []models.Author{
			{Name: "Mary Shelley", Type: "Author", ProfileImage: "https://m.media-amazon.com/images/I/81FBMlQ+BOL._SY600_.jpg", ShortBio: "", AlternateNames: []models.AlternateName{}},
			{Name: "C. S. Fritz", Type: "Editor", ProfileImage: "https://m.media-amazon.com/images/S/amzn-author-media-prod/p2nn5fvd8jjbeti266ekhsmmus._SY600_.jpg", ShortBio: "", AlternateNames: []models.AlternateName{}},
			{Name: "Maurice Hindle", Type: "Editor", ProfileImage: "https://pbs.twimg.com/profile_images/520897946389209088/TL1DsUUD_400x400.jpeg", ShortBio: "", AlternateNames: []models.AlternateName{}},
		},
		ReadingLogs: []models.ReadingLog{
			{LogDate: time.Now(), Message: "Started Reading", CurrentPage: 1, Finished: false},
			{LogDate: time.Now().Add(time.Hour * 72), Message: "Finished", CurrentPage: 352, Finished: true},
		},
	}
	return ctx.JSON(200, resMock)
}
