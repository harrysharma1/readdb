package database

import (
	"encoding/json"
	"errors"
	"log"
	"readdb/models"
	"time"

	bolt "go.etcd.io/bbolt"
)

func SeedSeries(db *bolt.DB) {
	author := models.Author{
		ID:           2,
		Name:         "Eiichiro oda",
		Type:         "Mangaka",
		ProfileImage: "https://static.wikia.nocookie.net/onepiece/images/3/32/Eiichiro_Oda_Infobox.png",
		ShortBio:     `Eiichiro Oda (Japanese: 尾田 栄一郎, Hepburn: Oda Eiichirō; born January 1, 1975) is a Japanese manga artist and the creator of the series One Piece, the best-selling manga in history and the best-selling comic series printed in volume. With more than 520 million tankōbon copies of One Piece in circulation worldwide, Oda is one of the best-selling fiction authors. The series' popularity resulted in Oda being named one of the artists who changed the history of manga.`,
	}

	book1 := models.Book{
		ID: 4,
		BookIdentifiers: []models.BookIdentifier{
			{
				Type:  "isbn10",
				Value: "1569319014",
			},
			{
				Type:  "isbn13",
				Value: "978-1569319017",
			},
		},
		Name: "One Piece, Vol. 1: Romance Dawn",
		ShortBlurb: `As a child, Monkey D. Luffy dreamed of becoming King of the Pirates. But his life changed when he accidentally gained the power to stretch like rubber...at the cost of never being able to swim again! Years later, Luffy sets off in search of the One Piece, said to be the greatest treasure in the world...

As a child, Monkey D. Luffy was inspired to become a pirate by listening to the tales of the buccaneer "Red-Haired" Shanks. But his life changed when Luffy accidentally ate the Gum-Gum Devil Fruit and gained the power to stretch like rubber...at the cost of never being able to swim again! Years later, still vowing to become the king of the pirates, Luffy sets out on his adventure...one guy alone in a rowboat, in search of the legendary "One Piece," said to be the greatest treasure in the world...`,
		Rating:         4,
		ProfileImage:   "https://m.media-amazon.com/images/I/91NxYvUNf6L._SY522_.jpg",
		TotalPageCount: 208,
		Authors:        []models.Author{author},
		ReadingLogs: []models.ReadingLog{
			{
				LogDate:     time.Now(),
				Message:     "Started reading",
				CurrentPage: 1,
				Finished:    false,
			},
			{
				LogDate:     time.Now().Add(2 * time.Hour),
				Message:     "Peak",
				CurrentPage: 216,
				Finished:    true,
			},
		},
		BookBuyLinks: []models.BookBuyLink{
			{
				Retailer: "Waterstones",
				URL:      "https://www.waterstones.com/book/one-piece-vol-1/eiichiro-oda/9781569319017",
			},
			{
				Retailer: "Amazon",
				URL:      "https://www.amazon.com/One-Piece-Vol-Romance-Dawn/dp/1569319014",
			},
		},
	}

	book2 := models.Book{
		ID: 5,
		BookIdentifiers: []models.BookIdentifier{
			{
				Type:  "isbn10",
				Value: "159116057X",
			},
			{
				Type:  "isbn13",
				Value: "978-1591160571",
			},
		},
		Name: "One Piece, Vol. 2: Buggy the Clown",
		ShortBlurb: `As a child, Monkey D. Luffy dreamed of becoming King of the Pirates. But his life changed when he accidentally gained the power to stretch like rubber…at the cost of never being able to swim again! Years, later, Luffy sets off in search of the “One Piece,” said to be the greatest treasure in the world...

As a kid, Monkey D. Luffy vowed to become King of the Pirates and find the legendary treasure called the "One Piece." The enchanted Gum-Gum Fruit has given Luffy the power to stretch like rubber--and his new crewmate, the infamous pirate hunter Roronoa Zolo, strikes fear into the hearts of other buccaneers! But what chance does one rubber guy stand against Nami, a thief so tough she specializes in robbing pirates...or Captain Buggy, a fiendish pirate lord whose weird, clownish appearance conceals even weirder powers? It's pirate vs. pirate in the second swashbuckling volume of One Piece!`,
		Rating:         4,
		ProfileImage:   "https://m.media-amazon.com/images/I/61tLcPnnddL._SY445_SX342_ML2_.jpg",
		TotalPageCount: 200,
		Authors:        []models.Author{author},
		ReadingLogs: []models.ReadingLog{
			{
				LogDate:     time.Now(),
				Message:     "Started reading",
				CurrentPage: 1,
				Finished:    false,
			},
			{
				LogDate:     time.Now().Add(2 * time.Hour),
				Message:     "Peak",
				CurrentPage: 200,
				Finished:    true,
			},
		},
		BookBuyLinks: []models.BookBuyLink{
			{
				Retailer: "Waterstones",
				URL:      "https://www.waterstones.com/book/one-piece-vol-2/eiichiro-oda/9781591160571",
			},
			{
				Retailer: "Amazon",
				URL:      "https://www.amazon.com/One-Piece-Vol-Buggy-Clown/dp/159116057X",
			},
		},
	}

	book3 := models.Book{
		ID: 6,
		BookIdentifiers: []models.BookIdentifier{
			{
				Type:  "isbn10",
				Value: "1591161843",
			},
			{
				Type:  "isbn13",
				Value: "978-1591161844",
			},
		},
		Name: "One Piece, Vol. 3: Don't Get Fooled Again",
		ShortBlurb: `As a child, Monkey D. Luffy dreamed of becoming King of the Pirates. But his life changed when he accidentally gained the power to stretch like rubber…at the cost of never being able to swim again! Years, later, Luffy sets off in search of the “One Piece,” said to be the greatest treasure in the world...

Sure, lots of people say they want to be the King of the Pirates, but how many have the guts to do what it takes? When Monkey D. Luffy first set out to sea in a leaky rowboat, he had no idea what might lie over the horizon. Now he's got a crew--sort of--in the form of swordsman Roronoa Zolo and treasure-hunting thief Nami. If he wants to prove himself on the high seas, Luffy will have to defeat the weird pirate lord Buggy the Clown. He'll have to find a map to the Grand Line, the sea route where the toughest pirates sail. And he'll have to face the Dread Captain Usopp, who claims to be a notorious pirate captain...but frankly, Usopp says a lot of things...`,
		Rating:         4,
		ProfileImage:   "https://m.media-amazon.com/images/I/81sVMmbkafL._SY522_.jpg",
		TotalPageCount: 200,
		Authors:        []models.Author{author},
		ReadingLogs: []models.ReadingLog{
			{
				LogDate:     time.Now(),
				Message:     "Started reading",
				CurrentPage: 1,
				Finished:    false,
			},
			{
				LogDate:     time.Now().Add(2 * time.Hour),
				Message:     "Peak",
				CurrentPage: 200,
				Finished:    true,
			},
		},
		BookBuyLinks: []models.BookBuyLink{
			{
				Retailer: "Waterstones",
				URL:      "https://www.waterstones.com/book/one-piece-vol-3/eiichiro-oda/9781591161844",
			},
			{
				Retailer: "Amazon",
				URL:      "https://www.amazon.com/One-Piece-Vol-Fooled-Again/dp/1591161843",
			},
		},
	}
	series := models.Series{
		PartOfASeries: true,
		Name:          "One Piece",
		OtherBooksInSeries: []models.Book{
			book1,
			book2,
			book3,
		},
	}
	err := SaveSeries(db, series)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Seeded series")
}

func SaveSeries(db *bolt.DB, series models.Series) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("series"))
		if err != nil {
			return err
		}
		id, err := nextId(b)
		if err != nil {
			return err
		}
		series.ID = id
		data, err := json.Marshal(series)
		if err != nil {
			return err
		}
		return b.Put(itob(series.ID), data)
	})
}

func GetAllSeries(db *bolt.DB) ([]models.Series, error) {
	var series []models.Series

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("series"))
		if b == nil {
			return errors.New("series does not exist")
		}

		return b.ForEach(func(k, v []byte) error {
			var seriesSingle models.Series
			if err := json.Unmarshal(v, &seriesSingle); err != nil {
				return err
			}
			series = append(series, seriesSingle)
			return nil
		})
	})
	return series, err
}

func GetSeriesById(db *bolt.DB, id uint) (models.Series, error) {
	var series models.Series
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("series"))
		if b == nil {
			return errors.New("series does not exist")
		}

		json.Unmarshal(b.Get(itob(id)), &series)
		return nil
	})
	return series, err
}
