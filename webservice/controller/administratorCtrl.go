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

func GetAllAdmin(c echo.Context) (err error)  {
  var total int
  administrators := []model.Administrators{}
  // acceptingHost := model.AcceptingHost{}
  paginateParams := model.NewPaginateParams()
  //set value default
  Current_Page := paginateParams.CurrentPage
  Per_page := paginateParams.PerPage
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
  query := c.QueryParam("query")

  //calculate offset
  var offset = (Current_Page - 1) * Per_page
  tx := database.MysqlConn().Begin()
  if query != "" {
    tx.Debug().Order("administrator_id desc").Offset(offset).Limit(Per_page).Where("name LIKE ?", "%"+query+"%").Preload("AcceptingHost").Find(&administrators).Count(&total)
  }else{
    tx.Debug().Order("administrator_id desc").Offset(offset).Limit(Per_page).Preload("AcceptingHost").Find(&administrators).Count(&total)
  }

  tx.Commit()
  // data response to client
  dataResp := model.ResponseObj{
    PerPage:     Per_page,
    Total:   total,
    CurrentPage:   Current_Page,
    Data: &administrators,
  }
  return c.JSON(http.StatusOK, dataResp)
}

func GetAdminById(c echo.Context) (err error){
  tx := database.MysqlConn().Begin()
	administratorsId := c.Param("id")
	administrator := model.Administrators{}

	if err := tx.Preload("AcceptingHost").First(&administrator, administratorsId).Error; err != nil {
		return c.JSON(http.StatusNotFound,map[string]string{"Message": err.Error(),"status":"false"})
	}
  tx.Commit()
	return c.JSON(http.StatusOK, &administrator)
}

func PostAdmin(c echo.Context) (err error) {
    tx:= database.MysqlConn().Begin()
  	administrators := model.Administrators{}
    if err = c.Bind(&administrators); err != nil {
       return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "InternalServerError","status":"false"})
    }

  	if err := tx.Create(&administrators).Error; err != nil {
  		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(),"status":"false"})
  	}
    tx.Commit()
    return c.JSON(http.StatusOK, &administrators)
}


func PutAdmin(c echo.Context) (err error) {
  tx:= database.MysqlConn().Begin()
	administrator_id := c.Param("id")
	administrator := model.Administrators{}
	if err := tx.First(&administrator, administrator_id).Error; err != nil {
		return c.JSON(http.StatusNotFound,map[string]string{"Message": err.Error(),"status":"false"})
	}

	data_updated := model.Administrators{}
  if err = c.Bind(&data_updated); err != nil {
     return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "InternalServerError","status":"false"})
  }

	administrator.Name = data_updated.Name
	administrator.Email = data_updated.Email
	administrator.Password = data_updated.Password
	administrator.Description = data_updated.Description
	administrator.AcceptingHostId = data_updated.AcceptingHostId
	administrator.Enabled = data_updated.Enabled
	administrator.Permission = data_updated.Permission
	administrator.CreatedById = data_updated.CreatedById

	if err := tx.Save(&administrator).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(),"status":"false"})
	}
  tx.Commit()
  return c.JSON(http.StatusOK, &administrator)
}

func DeleteAdmin(c echo.Context) (err error){
  tx:= database.MysqlConn().Begin()
	id := c.Param("id")
	reminder := model.Administrators{}
	if err := tx.First(&reminder, id).Error; err != nil {
		return c.JSON(http.StatusNotFound,map[string]string{"Message": err.Error(),"status":"false"})
	}
	if err := tx.Delete(&reminder).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(),"status":"false"})
	}
  tx.Commit()
	return c.JSON(http.StatusOK, map[string]string{"Message": "deleted","status":"true"})
}

func SearchAdminCtrl(c echo.Context) (err error) {
  query := c.QueryParam("query")
  if query != ""  {
      modelQuery := []model.Administrators{}
      tx := database.MysqlConn().Begin()
      tx.Where("name LIKE ?", "%"+query+"%").Find(&modelQuery)
      tx.Commit()
      return c.JSON(http.StatusOK, &modelQuery)
  }
	return c.JSON(http.StatusNotFound,model.Status{StatusCode: http.StatusNotFound, Message: err.Error(),Status:"false"})
}
