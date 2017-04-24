package controller
import (
  "github.com/labstack/echo"
    _ "log"
  "net/http"
  _ "encoding/json"
	database "../db"
  model "../model"
)

func GetAllAdmin(c echo.Context) error  {
  MiddlewareJWT(c)
  tx := database.MysqlConn().Begin()
  administrators := []model.Administrators{}
	tx.Order("administrators.administrator_id desc").Limit(10).Find(&administrators)
  tx.Commit()
  return c.JSON(http.StatusOK, &administrators)
}



func GetAdminById(c echo.Context) (err error){
  tx := database.MysqlConn().Begin()
	administratorsId := c.Param("id")
	administrator := model.Administrators{}

	if err := tx.First(&administrator, administratorsId).Error; err != nil {
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
