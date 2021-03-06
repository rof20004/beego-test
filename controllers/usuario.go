package controllers

import (
	"encoding/json"
	"errors"
	"mytest/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

// UsuarioController oprations for Usuario
type UsuarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsuarioController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Index url
func (c *UsuarioController) Index() {
	c.TplName = "usuario/index.tpl"
}

// Create url
func (c *UsuarioController) Create() {
	c.TplName = "usuario/create.tpl"
}

// Edit url
func (c *UsuarioController) Edit() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	u, err := models.GetUsuarioByID(id)
	if err == nil {
		c.Data["usuario"] = u
	}
	c.TplName = "usuario/edit.tpl"
}

// View url
func (c *UsuarioController) View() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	u, err := models.GetUsuarioByID(id)
	if err == nil {
		c.Data["usuario"] = u
	}
	c.TplName = "usuario/show.tpl"
}

// Post ...
// @Title Post
// @Description create Usuario
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 201 {int} models.Usuario
// @Failure 403 body is empty
// @router / [post]
func (c *UsuarioController) Post() {
	var u models.Usuario
	json.NewDecoder(c.Ctx.Request.Body).Decode(&u)

	v := validation.Validation{}
	b, err := v.Valid(&u)

	if err != nil {
		c.Data["json"] = err.Error()
	}

	if strings.TrimSpace(u.Nome) == "" {
		c.Data["json"] = err.Error()
	}

	if !b {
		c.Data["json"] = err.Error()
	} else {
		if _, err := models.AddUsuario(&u); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = u
		} else {
			c.Data["json"] = err.Error()
		}
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Usuario by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsuarioController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUsuarioByID(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Usuario
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Usuario
// @Failure 403
// @router / [get]
func (c *UsuarioController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUsuario(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Usuario
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsuarioController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	u := models.Usuario{ID: id}
	json.NewDecoder(c.Ctx.Request.Body).Decode(&u)

	v := validation.Validation{}
	b, err := v.Valid(&u)

	if err != nil {
		c.Data["json"] = err.Error()
	}

	if strings.TrimSpace(u.Nome) == "" {
		c.Data["json"] = err.Error()
	}

	if !b {
		c.Data["json"] = err.Error()
	} else {
		if err := models.UpdateUsuarioByID(&u); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Usuario
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsuarioController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteUsuario(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// UploadFile method
func (c *UsuarioController) UploadFile() {
	file, header, err := c.GetFile("arquivo") // where <<this>> is the controller and <<file>> the id of your form field

	if file != nil {
		if err != nil {
			logs.Error(err)
		}
		fileName := header.Filename
		err := c.SaveToFile("arquivo", "/home/rof20004/"+fileName)
		if err != nil {
			logs.Error(err)
		}
	}

	c.ServeJSON()
}
