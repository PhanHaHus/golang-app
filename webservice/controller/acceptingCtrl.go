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

func GetAcceptingHost(c echo.Context) (err error)  {
  var total int
  acceptingHosts := []model.AcceptingHosts{}
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
    // when search
    tx.Debug().Order("accepting_host_id desc").Offset(offset).Limit(Per_page).Where("name LIKE ?", "%"+query+"%").Or("description LIKE ?", "%"+query+"%").Find(&acceptingHosts).Count(&total)
  }else{
    tx.Debug().Order("accepting_host_id desc").Offset(offset).Limit(Per_page).Find(&acceptingHosts).Count(&total)
  }

  tx.Commit()
  // data response to client
  dataResp := model.ResponseObj{
    PerPage:     Per_page,
    Total:   total,
    CurrentPage:   Current_Page,
    Data: &acceptingHosts,
  }
  return c.JSON(http.StatusOK, dataResp)
}

func GetAcceptingHostById(c echo.Context) (err error){
  tx := database.MysqlConn().Begin()
	acceptingHostsId := c.Param("id")
	acceptingHost := model.AcceptingHosts{}

	if err := tx.Preload("AcceptingHost").First(&acceptingHost, acceptingHostsId).Error; err != nil {
		return c.JSON(http.StatusNotFound,map[string]string{"Message": err.Error(),"status":"false"})
	}
  tx.Commit()
	return c.JSON(http.StatusOK, &acceptingHost)
}

func PostAcceptingHost(c echo.Context) (err error) {
  	acceptingHosts := model.AcceptingHosts{}
    if err = c.Bind(&acceptingHosts); err != nil {
       return c.JSON(http.StatusInternalServerError,model.Status{StatusCode: http.StatusInternalServerError,Message: err.Error(),Status:"false"})
    }
    tx:= database.MysqlConn().Begin()
  	if err := tx.Create(&acceptingHosts).Error; err != nil {
  		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(),Status:"false"})
  	}
    tx.Commit()
    return c.JSON(http.StatusOK, &acceptingHosts)
}


func PutAcceptingHost(c echo.Context) (err error) {
  tx:= database.MysqlConn().Begin()
	acceptingHost_id := c.Param("id")
	acceptingHost := model.AcceptingHosts{}
	if err := tx.First(&acceptingHost, acceptingHost_id).Error; err != nil {
		return c.JSON(http.StatusNotFound,map[string]string{"Message": err.Error(),"status":"false"})
	}

	data_updated := model.AcceptingHosts{}
  if err = c.Bind(&data_updated); err != nil {
     return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "InternalServerError","status":"false"})
  }

	acceptingHost.Name = data_updated.Name
	acceptingHost.Password = data_updated.Password
	acceptingHost.Description = data_updated.Description
	acceptingHost.Enabled = data_updated.Enabled
	acceptingHost.CreatedById = data_updated.CreatedById

	if err := tx.Save(&acceptingHost).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(),"status":"false"})
	}
  tx.Commit()
  return c.JSON(http.StatusOK, &acceptingHost)
}

func DeleteAcceptingHost(c echo.Context) (err error){
  tx:= database.MysqlConn().Begin()
	id := c.Param("id")
	acceptingHost := model.AcceptingHosts{}
	if err := tx.First(&acceptingHost, id).Error; err != nil {
		return c.JSON(http.StatusNotFound,map[string]string{"Message": err.Error(),"status":"false"})
	}
	if err := tx.Delete(&acceptingHost).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(),"status":"false"})
	}
  tx.Commit()
	return c.JSON(http.StatusOK, map[string]string{"Message": "deleted","status":"true"})
}