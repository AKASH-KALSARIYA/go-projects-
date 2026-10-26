package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Punlisher string `json:"publisher"`
}

func (r *Repository) CreateBook(c *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"Message": "Internal sever error!! request is Failed!!"})
		return err
	}

	err := r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"Message": "request is Failed!! Could not create book!!"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"Message": "Book created successfully!!"})
	return nil
}

func (r *Repository) GetBooks(c *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
	context.Status(http.StatusBadRequest).JSON(
		&fiber.Map{"Message": "request is Failed!! Could not get books!!"
	})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"Message": "Books retrieved successfully!!", "data": bookModels})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_books/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBooksById)
	api.Get("/books", r.GetBooks)
	api.Put("/update_books/:id", r.UpdateBook)
	api.Patch("/update_book/:id", r.UpdateBookById)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8000")

}
