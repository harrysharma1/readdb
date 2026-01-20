package database

import (
	"encoding/json"
	"errors"
	"log"
	"readdb/models"
	"slices"
	"time"

	bolt "go.etcd.io/bbolt"
)

func SeedBook(db *bolt.DB) {
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
	author2 := models.Author{
		ID:           2,
		Name:         "Eiichiro oda",
		Type:         "Mangaka",
		ProfileImage: "https://static.wikia.nocookie.net/onepiece/images/3/32/Eiichiro_Oda_Infobox.png",
		ShortBio:     `Eiichiro Oda (Japanese: 尾田 栄一郎, Hepburn: Oda Eiichirō; born January 1, 1975) is a Japanese manga artist and the creator of the series One Piece, the best-selling manga in history and the best-selling comic series printed in volume. With more than 520 million tankōbon copies of One Piece in circulation worldwide, Oda is one of the best-selling fiction authors. The series' popularity resulted in Oda being named one of the artists who changed the history of manga.`,
	}
	book1 := models.Book{
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
	err := SaveBook(db, book1)
	if err != nil {
		log.Fatal(err)
	}
	book2 := models.Book{
		BookIdentifiers: []models.BookIdentifier{
			{
				Type:  "isbn10",
				Value: "0143137905",
			},
			{
				Type:  "isbn13",
				Value: "978-0143137900",
			},
		},
		Name: "The Last Man",
		ShortBlurb: `Written while Mary Shelley was in a self-imposed lockdown after the loss of her husband and children, and in the wake of intersecting crises including the climate-changing Mount Tambora eruption and a raging cholera outbreak, The Last Man (1826) is an early work of climate fiction and a prophetic depiction of environmental change.

Set in the late twenty-first century, a deadly pandemic leaves a lone survivor, and follows his journey through a post-apocalyptic world, devoid of humanity and reclaimed by nature. Rather than
give in to despair, Shelley imagines a new world where freshly-formed communities and alternative ways of being stand in for self-important politicians serving corrupt institutions, and where nature reigns mightily over humanity.

Brimming with political intrigue, The Last Man broaches partisan dysfunction, imperial warfare, refugee crises, and economic collapse―and brings the legacy of her radically progressive parents, William Godwin and Mary Wollstonecraft, to bear on present-day questions about making a better world less centred around “man.”`,
		Rating:         3,
		ProfileImage:   "https://m.media-amazon.com/images/I/81hSmVgjFPL._SY522_.jpg",
		TotalPageCount: 592,
		Authors:        []models.Author{author},
		ReadingLogs: []models.ReadingLog{
			{
				LogDate:     time.Time{},
				Message:     "Started reading",
				CurrentPage: 1,
				Finished:    false,
			},
			{
				LogDate:     time.Now().Add(72 * time.Hour),
				Message:     "Peak",
				CurrentPage: 592,
				Finished:    true,
			},
		},
	}
	err = SaveBook(db, book2)
	if err != nil {
		log.Fatal(err)
	}

	book3 := models.Book{
		BookIdentifiers: []models.BookIdentifier{
			{
				Type:  "isbn10",
				Value: "1974749924",
			},
			{
				Type:  "isbn13",
				Value: "978-1974749928",
			},
		},
		Name:           "Wanted! Eiichiro Oda Before One Piece",
		ShortBlurb:     `See another version of Luffy in the original short story that launched a legend. The creator of One Piece presents this unique collection of his earliest works. From Western to fantasy to ghost story, this special volume has it all—including “Romance Dawn,” the one-shot that would become the international megahit One Piece!`,
		Rating:         4,
		ProfileImage:   "https://m.media-amazon.com/images/I/8192fODrJIL._SY522_.jpg",
		TotalPageCount: 208,
		Authors:        []models.Author{author2},
		ReadingLogs: []models.ReadingLog{
			{
				LogDate:     time.Time{},
				Message:     "Started reading",
				CurrentPage: 1,
				Finished:    false,
			},
			{
				LogDate:     time.Now().Add(2 * time.Hour),
				Message:     "Peak",
				CurrentPage: 208,
				Finished:    true,
			},
		},
	}
	err = SaveBook(db, book3)
	if err != nil {
		log.Fatal(err)
	}
	book4 := models.Book{
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
		Authors:        []models.Author{author2},
		ReadingLogs: []models.ReadingLog{
			{
				LogDate:     time.Time{},
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
	}
	err = SaveBook(db, book4)
	if err != nil {
		log.Fatal(err)
	}
	book5 := models.Book{
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
		Authors:        []models.Author{author2},
		ReadingLogs: []models.ReadingLog{
			{
				LogDate:     time.Time{},
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
	}
	err = SaveBook(db, book5)
	if err != nil {
		log.Fatal(err)
	}
	book6 := models.Book{
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
		Authors:        []models.Author{author2},
		ReadingLogs: []models.ReadingLog{
			{
				LogDate:     time.Time{},
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
	}
	err = SaveBook(db, book6)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Seeded book")

}

func SaveBook(db *bolt.DB, book models.Book) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("books"))
		if err != nil {
			return err
		}
		id, err := nextId(b)
		if err != nil {
			return err
		}
		book.ID = id
		data, err := json.Marshal(book)
		if err != nil {
			return err
		}
		return b.Put(itob(book.ID), data)
	})
}

func GetAllBooks(db *bolt.DB) ([]models.Book, error) {
	var books []models.Book

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("books"))
		if b == nil {
			return errors.New("books does not exist")
		}

		return b.ForEach(func(k, v []byte) error {
			var book models.Book
			if err := json.Unmarshal(v, &book); err != nil {
				return err
			}
			books = append(books, book)
			return nil
		})
	})
	return books, err
}

func GetBookById(db *bolt.DB, id uint) (models.Book, error) {
	var book models.Book
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("books"))
		if b == nil {
			return errors.New("books does not exist")
		}

		json.Unmarshal(b.Get(itob(id)), &book)
		return nil
	})
	return book, err
}

func GetAllBooksByAuthorID(db *bolt.DB, authorId uint) ([]models.Book, error) {
	var books []models.Book

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("books"))
		if b == nil {
			return errors.New("books does not exist")
		}

		return b.ForEach(func(k, v []byte) error {
			var book models.Book
			if err := json.Unmarshal(v, &book); err != nil {
				return err
			}
			books = append(books, book)
			return nil
		})
	})
	books = slices.DeleteFunc(books, func(b models.Book) bool {
		for _, author := range b.Authors {
			if author.ID == authorId {
				return false
			}
		}
		return true
	})
	return books, err
}
