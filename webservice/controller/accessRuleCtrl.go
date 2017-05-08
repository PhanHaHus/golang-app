package controller
import (
  "github.com/labstack/echo"
    "log"
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

  // if exist param query and param table on url
  query := c.QueryParam("query")
  table := c.QueryParam("table")
  if query != "" && table!="" {
    switch table {
      case "applications":
          modelQuery := []model.Applications{}
          tx := database.MysqlConn().Begin()
          tx.Where("name LIKE ?", "%"+query+"%").Find(&modelQuery)
          tx.Commit()
          return c.JSON(http.StatusOK, &modelQuery)
      case "user":
          modelQuery := []model.Users{}
          tx := database.MysqlConn().Begin()
          tx.Where("name LIKE ?", "%"+query+"%").Find(&modelQuery)
          tx.Commit()
          return c.JSON(http.StatusOK, &modelQuery)
      case "group":
          modelQuery := []model.Groups{}
          tx := database.MysqlConn().Begin()
          tx.Where("name LIKE ?", "%"+query+"%").Find(&modelQuery)
          tx.Commit()
          return c.JSON(http.StatusOK, &modelQuery)
      case "device":
          modelQuery := []model.Devices{}
          tx := database.MysqlConn().Begin()
          tx.Where("name LIKE ?", "%"+query+"%").Find(&modelQuery)
          tx.Commit()
          return c.JSON(http.StatusOK, &modelQuery)
      default:
        return c.JSON(http.StatusNotFound,model.Status{StatusCode: http.StatusNotFound, Message: "Not found table!",Status:"false"})
    }
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

    log.Println("accessRules")
    log.Println(c.Bind(&accessRules))

    if err = c.Bind(&accessRules); err != nil {
       return c.JSON(http.StatusInternalServerError,model.Status{StatusCode: http.StatusInternalServerError,Message: err.Error(),Status:"false"})
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
	administrator := model.AccessRules{}
	if err := tx.First(&administrator, administrator_id).Error; err != nil {
		return c.JSON(http.StatusNotFound,model.Status{StatusCode: http.StatusNotFound, Message: err.Error(),Status:"false"})
	}

	data_updated := model.AccessRules{}
  if err = c.Bind(&data_updated); err != nil {
     return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: "InternalServerError",Status:"false"})
  }

	administrator.ApplicationId = data_updated.ApplicationId
	administrator.UserId = data_updated.UserId
	administrator.DeviceId = data_updated.DeviceId
	administrator.GroupId = data_updated.GroupId
	administrator.AccessRuleType = data_updated.AccessRuleType
	administrator.Description = data_updated.Description
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
	accessRules := model.AccessRules{}
	if err := tx.First(&accessRules, id).Error; err != nil {
		return c.JSON(http.StatusNotFound,model.Status{StatusCode: http.StatusNotFound, Message: err.Error(),Status:"false"})
	}
	if err := tx.Delete(&accessRules).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(),Status:"false"})
	}
  tx.Commit()
	return c.JSON(http.StatusOK, model.Status{StatusCode: http.StatusInternalServerError, Message:"deleted",Status:"false"})
}
