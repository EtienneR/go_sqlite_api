package api

import (
	"github.com/gin-gonic/gin"
)

// La structure de données
type Users struct {
	Id   int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name string `gorm:"not null" form:"name" json:"name"`
}

// Ajouter un utilisteur
func PostUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var json Users
	c.Bind(&json)

	// Si le champ est bien saisi
	if json.Name != "" {
		// INSERT INTO "users" (name) VALUES (json.Name);
		db.Create(&json)
		// Affichage des données saisies
		c.JSON(201, gin.H{"success": json})
	} else {
		// Affichage de l'erreur
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

// Obtenir la liste de tous les utilisateurs
func GetUsers(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var users []Users
	// SELECT * FROM users
	db.Find(&users)
	// Affichage des données
	c.JSON(200, users)
}

// Obtenir un utilisateur par son id
func GetUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = id;
	db.First(&user, id)

	if user.Id != 0 {
		// Affichage des données
		c.JSON(200, user)
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

// Modifier un utilisateur
func EditUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = id;
	db.First(&user, id)

	if user.Name != "" {
		if user.Id != 0 {
			var json Users
			c.Bind(&json)

			result := Users{
				Id:   user.Id,
				Name: json.Name,
			}

			// UPDATE users SET name='json.Name' WHERE id = user.Id;
			db.Model(&user).Update("name", result.Name)
			// Affichage des données modifiées
			c.JSON(200, gin.H{"success": result})
		} else {
			// Affichage de l'erreur
			c.JSON(404, gin.H{"error": "User not found"})
		}

	} else {
		// Affichage de l'erreur
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

// Supprimer un utilisateur
func DeleteUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var user Users
	db.First(&user, id)

	if user.Id != 0 {
		// DELETE FROM users WHERE id = user.Id
		db.Delete(&user)
		// Affichage des données
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
