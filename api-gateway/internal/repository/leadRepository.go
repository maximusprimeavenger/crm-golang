package repository

import (
	"github.com/gofiber/fiber"
)

func GetLead() fiber.Handler {
	return func(c *fiber.Ctx) {
		//id := c.Params("id")
		// lead, err := db.FindLeadById(id)
		// if err != nil {
		// 	c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		// 	return
		// }
		//c.JSON(lead)
	}
}
func GetLeads() fiber.Handler {
	return func(c *fiber.Ctx) {
		// 	//leads := db.FindLeads()
		// 	c.Status(200).JSON(fiber.Map{"message": "Leads found!",
		// 	//	"List": leads})
	}
}
func DeleteLead() fiber.Handler {
	return func(c *fiber.Ctx) {
		// id := c.Params("id")
		// err := db.DeleteLead(id)
		// if err != nil {
		// 	c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Error deleting lead: %v", err)})
		// 	return
		// }
		// c.Status(200).JSON(fiber.Map{"message": "Lead deleted successfully!"})
	}
}
func NewLead() fiber.Handler {
	return func(c *fiber.Ctx) {
		// leadJSON := new(models.Lead)
		// if err := c.BodyParser(leadJSON); err != nil {
		// 	c.Status(500).Send(err)
		// 	return
		// }
		// if leadJSON.Name == "" {
		// 	c.Status(500).JSON(fiber.Map{"error": "empty parameters for lead"})
		// }
		// newLead, err := lead.NewLead(leadJSON.Name, leadJSON.Email, leadJSON.Phone, leadJSON.Company)
		// if err != nil {
		// 	c.Status(500).Send(fmt.Sprintf("Error creating new lead: %v", err))
		// 	return
		// }
		// err = db.SaveLead(newLead)
		// if err != nil {
		// 	c.Status(500).Send(fmt.Sprintf("Error saving new lead: %v", err))
		// 	return
		// }
		// c.Status(200).JSON(fiber.Map{"message": "Lead created successfully!",
		// 	"Lead": []interface{}{
		// 		newLead.ID,
		// 		newLead.Name,
		// 		newLead.Phone,
		// 		newLead.Email,
		// 		newLead.Company,
		// 		newLead.CreatedAt,
		// 	},
		// })
	}
}

func UpdateLead() fiber.Handler {
	return func(c *fiber.Ctx) {
		// id := c.Params("id")
		// if id == "" {
		// 	c.Status(500).JSON(fiber.Map{"error": "empty id provided"})
		// 	return
		// }
		// lead := new(models.Lead)
		// if err := c.BodyParser(lead); err != nil {
		// 	c.Status(500).Send(fmt.Sprintf("Error parsing data: %v", err))
		// 	return
		// }
		// err := db.UpdateLead(lead.Name, lead.Email, lead.Company, lead.Phone, id)
		// if err != nil {
		// 	c.Status(500).Send(fmt.Sprintf("Error updating lead: %v", err))
		// 	return
		// }
		// updatedLead, err := db.FindLeadById(id)
		// if err != nil {
		// 	c.Status(500).Send(fmt.Sprintf("Error finding lead: %v", err))
		// 	return
		// }
		// c.Status(200).JSON(fiber.Map{"message": "Lead updated successfully!",
		// 	"Lead": []interface{}{
		// 		updatedLead.ID,
		// 		updatedLead.Name,
		// 		updatedLead.Phone,
		// 		updatedLead.Email,
		// 		updatedLead.Company,
		// 		updatedLead.CreatedAt,
		// 	},
		// })
	}
}
