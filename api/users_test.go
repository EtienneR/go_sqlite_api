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

func TestPostUser(t *testing.T) {
	// Contenu à soumettre
	userJson := `{"name": "dennis"}`

	// Contenu à soumettre au bon format
	reader = strings.NewReader(userJson)

	// Déclaration de la requête : type, URL, contenu
	req, err := http.NewRequest("POST", usersUrl, reader)
	// Requête de type JSON
	req.Header.Set("Content-Type", "application/json")

	// Exécution de la requête
	res, err := http.DefaultClient.Do(req)

	// Erreur si route inacessible
	if err != nil {
		t.Error(err)
	}

	// Erreur si code HTTP différent de 201
	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	// Test sur le format des données
	if res.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("No JSON Content-Type")
	}
}

// FetchAll
func TestGetUsers(t *testing.T) {
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

	// Test sur le format des données
	if res.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("No JSON Content-Type")
	}
}

// FetchOne
func TestGetUser(t *testing.T) {
	// Contenu à soumettre vide
	reader = strings.NewReader("")

	// Déclaration de la requête : type, URL, contenu
	request, err := http.NewRequest("GET", usersUrlId, reader)

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

	// Test sur le format des données
	if res.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("No JSON Content-Type")
	}
}

func TestEditUser(t *testing.T) {
	// Contenu à soumettre
	userJson := `{"name": "mark"}`

	// Contenu à soumettre au bon format
	reader = strings.NewReader(userJson)

	// Déclaration de la requête : type, URL, contenu
	request, err := http.NewRequest("PUT", usersUrlId, reader)
	// Requête de type JSON
	request.Header.Set("Content-Type", "application/json")

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

	// Test sur le format des données
	if res.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("No JSON Content-Type")
	}
}

func TestDeleteUser(t *testing.T) {
	// Contenu à soumettre vide
	reader = strings.NewReader("")

	// Déclaration de la requête : type, URL, contenu
	request, err := http.NewRequest("DELETE", usersUrlId, reader)

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
}

func TestOptionsUser(t *testing.T) {
	// Contenu à soumettre vide
	reader = strings.NewReader("")

	// Déclaration de la requête : type, URL, contenu
	request, err := http.NewRequest("OPTIONS", usersUrl, reader)

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

	// Test sur le Access-Control-Allow-Methods
	if res.Header.Get("Access-Control-Allow-Methods") != "DELETE, POST, PUT" {
		t.Error("Access-Control-Allow-Methods false")
	}
}

func TestOptionsUserId(t *testing.T) {
	// Contenu à soumettre vide
	reader = strings.NewReader("")

	// Déclaration de la requête : type, URL, contenu
	request, err := http.NewRequest("OPTIONS", usersUrlId, reader)

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

	// Test sur le Access-Control-Allow-Methods
	if res.Header.Get("Access-Control-Allow-Methods") != "DELETE, POST, PUT" {
		t.Error("Access-Control-Allow-Methods false")
	}
}
