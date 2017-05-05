package controller
import (
  "github.com/labstack/echo"
  _  "log"
  "net/http"
  _ "encoding/json"
  "strconv"
	database "../db"
  model "../model"
)

func GetAllAccessRules(c echo.Context) (err error)  {
  var total int
  tx := database.MysqlConn().Begin()
  accessrules := []model.AccessRules{}
  paginateParams := model.NewPaginateParams()

  //set value default
  Current_Page := paginateParams.CurrentPage
  Per_page := paginateParams.PerPage //set value default
  // if exist param current page from url
  curr:=c.QueryParam("current_page")
  if curr != "" {
      curr, _ := strconv.Atoi(curr) //string to int
      Current_Page = curr
  }

  // if exist param per_page from url
  per_page_params:=c.QueryParam("per_page")
  if per_page_params != "" {
      per_page_params, _ := strconv.Atoi(per_page_params) //string to int
      Per_page = per_page_params
  }

  //calculate offset
  var offset = (Current_Page - 1) * Per_page
  // eager load relationship
	tx.Order("access_rule_id desc").Offset(offset).Limit(Per_page).Preload("Application").Preload("User").Preload("Device").Preload("Group").Preload("CreatedByUser").Find(&accessrules).Count(&total)
  tx.Commit()
  // data response to client
  dataResp := model.ResponseObj{
    PerPage:     Per_page,
    Total:   total,
    CurrentPage:   Current_Page,
    Data: &accessrules,
  }
  return c.JSON(http.StatusOK, dataResp)
}

func SearchACLCtrl(c echo.Context) (err error)  {

  applications := model.Applications{}
  //user := model.Users{}

  // if exist param current page from url
  query := c.QueryParam("query")
  table := c.QueryParam("type")
  if query != "" {
    tx := database.MysqlConn().Begin()
    tx.Where("name LIKE ?", "%"+query+"%").Find(&applications)
    tx.Commit()
    return c.JSON(http.StatusOK, &applications)
  }

  return c.JSON(http.StatusNotFound,model.Status{StatusCode: http.StatusNotFound, Message: err.Error(),Status:"false"})
}

func GetAccessRuleById(c echo.Context) (err error){
	accessrulesId := c.Param("id")
	accessrules := model.AccessRules{}
  tx := database.MysqlConn().Begin()
	if err := tx.Preload("Application").Preload("User").Preload("Device").Preload("Group").Preload("CreatedByUser").First(&accessrules, accessrulesId).Error; err != nil {
		return c.JSON(http.StatusNotFound,model.Status{StatusCode: http.StatusNotFound, Message: err.Error(),Status:"false"})
	}
  tx.Commit()
	return c.JSON(http.StatusOK, &accessrules)
}

func PostAccessRule(c echo.Context) (err error) {
    tx:= database.MysqlConn().Begin()
  	accessRules := model.AccessRules{}
    if err = c.Bind(&accessRules); err != nil {
       return c.JSON(http.StatusInternalServerError,model.Status{StatusCode: http.StatusInternalServerError,Message: "InternalServerError",Status:"false"})
    }

  	if err := tx.Create(&accessRules).Error; err != nil {
  		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(),Status:"false"})
  	}
    tx.Commit()
    return c.JSON(http.StatusOK, &accessRules)
}


func PutAccessRule(c echo.Context) (err error) {
  tx:= database.MysqlConn().Begin()
	administrator_id := c.Param("id")
	administrator := model.Administrators{}
	if err := tx.First(&administrator, administrator_id).Error; err != nil {
		return c.JSON(http.StatusNotFound,model.Status{StatusCode: http.StatusNotFound, Message: err.Error(),Status:"false"})
	}

	data_updated := model.Administrators{}
  if err = c.Bind(&data_updated); err != nil {
     return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: "InternalServerError",Status:"false"})
  }

	administrator.Name = data_updated.Name
	administrator.Email = data_updated.Email
	administrator.Password = data_updated.Password
	administrator.Description = data_updated.Description
	administrator.AcceptingHostId = data_updated.AcceptingHostId
	administrator.Enabled = data_updated.Enabled
	administrator.CreatedById = data_updated.CreatedById

	if err := tx.Save(&administrator).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(),Status:"false"})
	}
  tx.Commit()
  return c.JSON(http.StatusOK, &administrator)
}

func DeleteAccessRule(c echo.Context) (err error){
  tx:= database.MysqlConn().Begin()
	id := c.Param("id")
	reminder := model.Administrators{}
	if err := tx.First(&reminder, id).Error; err != nil {
		return c.JSON(http.StatusNotFound,model.Status{StatusCode: http.StatusNotFound, Message: err.Error(),Status:"false"})
	}
	if err := tx.Delete(&reminder).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(),Status:"false"})
	}
  tx.Commit()
	return c.JSON(http.StatusOK, model.Status{StatusCode: http.StatusInternalServerError, Message:"deleted",Status:"false"})
}
