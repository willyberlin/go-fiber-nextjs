package rooms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/rooms", getRooms)
	app.Get("/rooms/:id", getSingleRoom)
	app.Post("/rooms", createRoom)
	app.Put("/rooms/:id", updateRoom)
	app.Delete("/rooms/:id", deleteRoom)
	app.Post("/rooms/:id/togglebooking", toggleRoomBooking)
}

func getRooms(c *fiber.Ctx) error {
	return c.JSON(GetAllRooms())
}

func getSingleRoom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid room ID"})
	}
	room := GetRoom(id)
	if room == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Room not found"})
	}
	return c.JSON(room)
}

func createRoom(c *fiber.Ctx) error {
	var input Room
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	room := CreateRoom(input.Name, input.Description, input.HasDesks, input.DesksCount)
	return c.Status(fiber.StatusCreated).JSON(room)
}

func updateRoom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid room ID"})
	}
	var input Room
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	room := UpdateRoom(id, input.Name, input.Description, input.HasDesks, input.DesksCount)
	if room == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Room not found"})
	}
	return c.JSON(room)
}

func deleteRoom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid room ID"})
	}
	if !DeleteRoom(id) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Room not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func toggleRoomBooking(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid room ID"})
	}
	room, err := ToggleBooking(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(room)
}
