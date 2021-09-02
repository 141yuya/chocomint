package create_user_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/141yuya/go-clean-architecture/di"
	"github.com/141yuya/go-clean-architecture/domain/entities"
	"github.com/141yuya/go-clean-architecture/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
)

func setup() {
	sqlHandler := infrastructure.NewSqlHandler()
	sqlHandler.DB.AutoMigrate(entities.User{})
	sqlHandler.DB.Exec("TRUNCATE TABLE users")
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}

func TestCreateUser(t *testing.T) {
	sqlHandler := infrastructure.NewSqlHandler()
	router := infrastructure.NewRouter(di.InitUserController(sqlHandler))
	engine := router.SetupRouter()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := bytes.NewBufferString("{\"FirstName\":\"foo\",\"LastName\":\"bar\"}")
	c.Request, _ = http.NewRequest("POST", "/users", body)
	c.Request.Header.Add("Content-Type", binding.MIMEJSON)
	engine.ServeHTTP(w, c.Request)

	assert.Equal(t, 201, w.Code)
}

func teardown() {

}
