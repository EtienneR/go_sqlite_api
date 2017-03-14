package api

// Les imports de librairies
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Connexion à la BDD SQLite
func InitDb() *gorm.DB {
	// Ouverture de la connexion vers la BDD SQLite
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)

	if !db.HasTable(&Users{}) {
		db.CreateTable(&Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
	}

	if err != nil {
		panic(err)
	}

	return db
}

func Handlers() *gin.Engine {
	// Initialisation du serveur MUX
	r := gin.Default()

	r.Use(Cors())

	v1Users := r.Group("api/v1/users")
	{
		v1Users.POST("", PostUser)
		v1Users.GET("", GetUsers)
		v1Users.GET(":id", GetUser)
		v1Users.PUT(":id", EditUser)
		v1Users.DELETE(":id", DeleteUser)
		v1Users.OPTIONS("", OptionsUser)    // POST
		v1Users.OPTIONS(":id", OptionsUser) // PUT, DELETE
	}

	return r
}

// Activation du CORS
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
