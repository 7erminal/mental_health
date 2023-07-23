package controllers

import (
	"encoding/json"
	"mental_health_application/models"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"

	"time"
)

// AuthenticationController operations for Authentication
type AuthenticationController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthenticationController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("SignUp", c.SignUp)
}

// Post ...
// @Title Create
// @Description create Authentication
// @Param	body		body 	models.Authentication	true		"body for Authentication content"
// @Success 201 {object} models.Authentication
// @Failure 403 body is empty
// @router / [post]
func (c *AuthenticationController) Post() {
	var v models.Authentication
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	logs.Info("Received ", v.Password, v.Email)

	if a, err := models.GetUsersByUsername(v.Email); err == nil {
		// Compare the stored hashed password, with the hashed version of the password that was received
		if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(v.Password)); err != nil {
			// If the two passwords don't match, return a 401 status
			c.Data["json"] = err.Error()

			var resp = models.UserResponse{StatusCode: 605, User: nil, StatusDesc: "Incorrect password"}
			c.Data["json"] = resp

		} else {
			c.Ctx.Output.SetStatus(201)

			var resp = models.UserResponse{StatusCode: 200, User: a, StatusDesc: "User has been authenticated"}
			c.Data["json"] = resp
		}
	} else {
		var resp = models.UserResponse{StatusCode: 605, User: nil, StatusDesc: "Unidentified user"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// SignUp ...
// @Title SignUp
// @Description Sign up
// @Param	body		body 	models.SignUp	true		"body for SignUp content"
// @Success 201 {object} models.Users
// @Failure 403 body is empty
// @router /sign-up [post]
func (c *AuthenticationController) SignUp() {
	var v models.SignUp
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info("Received ", v.Password)

	hashedPassword, errr := bcrypt.GenerateFromPassword([]byte(v.Password), 8)

	if errr == nil {
		logs.Debug(hashedPassword)

		v.Password = string(hashedPassword)

		logs.Debug("Sending", v.Password)

		// models.Agents{AgentName: v.AgentName, BranchId: v.BranchId, IdType: v.IdType, IdNumber: v.IdNumber, IsVerified: false, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: c_by, ModifiedBy: c_by}
	}

	var addUserModel = models.Users{FullName: v.Name, Password: string(hashedPassword), Email: v.Email, DateCreated: time.Now(), Active: 1, CreatedBy: 1, ModifiedBy: 1}

	if r, err := models.AddUsers(&addUserModel); err == nil {
		c.Ctx.Output.SetStatus(201)

		// logs.Debug("Returned user is", r)

		// id, _ := strconv.ParseInt(idStr, 0, 64)
		v, err := models.GetUsersById(r)

		if err != nil {
			c.Data["json"] = err.Error()

			var resp = models.UserResponse{StatusCode: 601, User: nil, StatusDesc: "Error fetching user"}
			c.Data["json"] = resp
		} else {
			logs.Debug("Returned user is", v)

			var resp = models.UserResponse{StatusCode: 200, User: v, StatusDesc: "User created successfully"}
			c.Data["json"] = resp

			// c.Data["json"] = v
		}
	} else {
		var resp = models.UserResponse{StatusCode: 606, User: nil, StatusDesc: "Error adding user"}
		c.Data["json"] = resp

		c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}
