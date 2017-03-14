package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/EtienneR/go_sqlite_api/api"
)

var (
	server               *httptest.Server
	reader               io.Reader
	usersUrl, usersUrlId string
)

func init() {
	// Ouverture de la connexion vers la BDD SQLite
	db := api.InitDb()
	// Fermeture de la connexion vers la BDD SQLite
	defer db.Close()

	var user api.Users

	// Suppression de la table
	db.DropTable(user)
	// Création de la table
	db.CreateTable(user)

	// Création d'utilisateurs
	db.Create(&api.Users{Name: "Pierre"})
	db.Create(&api.Users{Name: "Paul"})
	db.Create(&api.Users{Name: "Jacques"})
	db.Create(&api.Users{Name: "Marie Thérèse"})

	// Démarrage du serveur HTTP
	server = httptest.NewServer(api.Handlers())

	// URL sans paramêtre et avec
	usersUrl = server.URL + "/api/v1/users"
	usersUrlId = server.URL + "/api/v1/users/5"
}

func TestCors(t *testing.T) {
	// Contenu à soumettre vide
	reader = strings.NewReader("")

	// Déclaration de la requête : type, URL, contenu
	request, err := http.NewRequest("GET", usersUrl, reader)

	// Exécution de la requête
	res, err := http.DefaultClient.Do(request)

	// Erreur si route inacessible
	if err != nil {
		t.Error(err)
	}

	// Erreur si code HTTP différent de 200
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	if res.Header.Get("Access-Control-Allow-Origin") != "*" {
		t.Error("No CORS")
	}
}
