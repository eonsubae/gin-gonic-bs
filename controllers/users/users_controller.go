package users

import (
	"net/http"
	"strconv"

	"github.com/eonsubae/gin-gonic-bs/domain/users"
	"github.com/eonsubae/gin-gonic-bs/services"
	"github.com/eonsubae/gin-gonic-bs/utils/errors"
	"github.com/gin-gonic/gin"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User

	/* bind JSON1
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		// TODO: Handle json error
		fmt.Println(err)
		return
	}
	Equal logic with below code */

	/* bind JSON2 */
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		// TODO: return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: Handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewNotFoundError("user id should be number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		// TODO: Handle user reading error
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
