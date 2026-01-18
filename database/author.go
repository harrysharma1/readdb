package database

import (
	"encoding/json"
	"errors"
	"log"
	"readdb/models"

	bolt "go.etcd.io/bbolt"
)

func SeedAuthor(db *bolt.DB) {
	author := models.Author{
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
	err := SaveAuthor(db, author)
	if err != nil {
		log.Fatal(err)

	}
	log.Print("Seeded author")
}

func SaveAuthor(db *bolt.DB, author models.Author) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("authors"))
		if err != nil {
			return err
		}
		id, err := nextId(b)
		if err != nil {
			return err
		}
		author.ID = id
		data, err := json.Marshal(author)
		if err != nil {
			return err
		}
		return b.Put(itob(author.ID), data)
	})
}

func GetAllAuthors(db *bolt.DB) ([]models.Author, error) {
	var authors []models.Author

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("authors"))
		if b == nil {
			return errors.New("authors does not exist")
		}

		return b.ForEach(func(k, v []byte) error {
			var author models.Author
			if err := json.Unmarshal(v, &author); err != nil {
				return err
			}
			authors = append(authors, author)
			return nil
		})
	})
	return authors, err
}
