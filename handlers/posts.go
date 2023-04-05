package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wbrijesh/console-go/models"
	"github.com/wbrijesh/console-go/database"
	"github.com/wbrijesh/console-go/utils"	
)

func ListPosts(c *fiber.Ctx) error {
	facts := []models.Post{}

	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing request body",
			"error":   err,
		})
	}

	database.DB.Db.Create(&post)

	return c.Status(200).JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	post := new(models.Post)
	id := c.Params("id")

	database.DB.Db.Where("id = ?", id).Delete(&post)

	return c.Status(200).JSON(post)
}

func EditPost(c *fiber.Ctx) error {
	post := new(models.Post)
	id := c.Params("id")

	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing request body",
			"error":   err,
		})
	}

	existingIdsInDb := []int{}

	database.DB.Db.Model(&models.Post{}).Pluck("id", &existingIdsInDb)

	if !utils.Contains(existingIdsInDb, id) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Post does not exist",
		})
	}

	database.DB.Db.Model(&post).Where("id = ?", id).Updates(post)

	return c.Status(200).JSON(post)
}
