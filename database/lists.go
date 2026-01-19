package database

import (
	"encoding/json"
	"errors"
	"log"
	"readdb/models"
	"time"

	bolt "go.etcd.io/bbolt"
)

func SeedList(db *bolt.DB) {
	author := models.Author{
		ID:           1,
		Name:         "Mary Shelley",
		Type:         "Author",
		ProfileImage: "https://cdn-test.poetryfoundation.org/cdn-cgi/image/w=2292,h=1500,q=80/content/images/f32d3878f890870597a16f89778191bfdc1ca678.jpeg",
		ShortBio:     `Mary Wollstonecraft Shelley (born August 30, 1797, London, England—died February 1, 1851, London) was an English Romantic novelist best known as the author of Frankenstein (1818), a seminal work of Romanticism and a Gothic horror classic that is also considered to be one of the first science-fiction novels.`,
		AlternateNames: []models.AlternateName{
			{
				Name: "Mary Wollstonecraft Godwin",
			},
		},
	}
	book := models.Book{
		BookIdentifiers: []models.BookIdentifier{
			{
				Type:  "isbn10",
				Value: "1548256528",
			},
			{
				Type:  "isbn13",
				Value: "978-1548256524",
			},
		},
		Name: "Frankenstein",
		ShortBlurb: `Mary Shelley's classic novel, Frankenstein; or The Modern Prometheus, is a story about a young scientist Victor Frankenstein and his grotesque sapient creation through unorthodox science.
Odin's Library Classics is dedicated to bringing the world the best of humankind's literature from throughout the ages. Carefully selected, each work is unabridged from classic works of fiction, nonfiction, poetry, or drama.`,
		Rating:         4,
		ProfileImage:   "https://m.media-amazon.com/images/I/81JUcKDq9aL._SX522_.jpg",
		TotalPageCount: 352,
		Authors:        []models.Author{author},
		ReadingLogs: []models.ReadingLog{
			{
				LogDate:     time.Time{},
				Message:     "Started reading",
				CurrentPage: 1,
				Finished:    false,
			},
			{
				LogDate:     time.Time{}.Add(28 * time.Hour),
				Message:     "Oh poor wretch",
				CurrentPage: 128,
				Finished:    false,
			},
			{
				LogDate:     time.Now().Add(72 * time.Hour),
				Message:     "Peak",
				CurrentPage: 352,
				Finished:    true,
			},
		},
	}
	list := models.List{
		Name: "classics",
		Books: []models.Book{
			book,
		},
	}
	err := SaveList(db, list)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Seeded List")
}

func SaveList(db *bolt.DB, list models.List) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("lists"))
		if err != nil {
			return err
		}
		id, err := nextId(b)
		if err != nil {
			return err
		}
		list.ID = id
		data, err := json.Marshal(list)
		if err != nil {
			return err
		}
		return b.Put(itob(list.ID), data)
	})
}

func GetAllLists(db *bolt.DB) ([]models.List, error) {
	var lists []models.List
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("lists"))
		if b == nil {
			return errors.New("lists does not exist")
		}

		return b.ForEach(func(k, v []byte) error {
			var list models.List
			if err := json.Unmarshal(v, &list); err != nil {
				return err
			}
			lists = append(lists, list)
			return nil
		})
	})
	return lists, err
}
