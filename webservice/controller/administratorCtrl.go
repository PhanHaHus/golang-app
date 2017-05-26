package controller

import (
	_ "encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	database "../db"
	model "../model"
	"github.com/labstack/echo"
)

func GetAllAdmin(c echo.Context) (err error) {
	var total int
	administrators := []model.Administrators{}
	// acceptingHost := model.AcceptingHost{}
	paginateParams := model.NewPaginateParams()
	//set value default
	Current_Page := paginateParams.CurrentPage
	Per_page := paginateParams.PerPage
	// if exist param current page from url
	curr := c.QueryParam("current_page")
	if curr != "" {
		curr, _ := strconv.Atoi(curr) //string to int
		Current_Page = curr
	}
	// if exist param per_page from url
	per_page_params := c.QueryParam("per_page")
	if per_page_params != "" {
		per_page_params, _ := strconv.Atoi(per_page_params) //string to int
		Per_page = per_page_params
	}

	var queryParams = c.QueryParams()
	query := ""
	if queryParams["query"] != nil {
		query = queryParams["query"][0]
	}

	stringQr := toStringQuery(queryParams) // get query string filter

	// calculate offset
	var offset = (Current_Page - 1) * Per_page
	tx := database.MysqlConn().Begin()
	if query != "" {
		// when search
		tx.Debug().Order("administrator_id desc").Offset(offset).Limit(Per_page).Where("name LIKE ?", "%"+query+"%").Or("email LIKE ?", "%"+query+"%").Or("permission = ?", ""+query+"").Preload("AcceptingHost").Find(&administrators).Count(&total)
	} else if stringQr != "" {
		//when filter
		tx.Debug().Joins("JOIN accepting_hosts ON accepting_hosts.accepting_host_id = administrators.accepting_host_id").Order("administrator_id desc").Offset(offset).Limit(Per_page).Where(stringQr).Preload("AcceptingHost").Find(&administrators).Count(&total)
	} else {
		//when load page
		tx.Debug().Order("administrator_id desc").Offset(offset).Limit(Per_page).Preload("AcceptingHost").Find(&administrators).Count(&total)
	}

	if stringQr != "" && query != "" {
		//when type both
		tx.Debug().Joins("JOIN accepting_hosts ON accepting_hosts.accepting_host_id = administrators.accepting_host_id").Order("administrator_id desc").Offset(offset).Limit(Per_page).Where("name LIKE ?", "%"+query+"%").Where(stringQr).Find(&administrators).Count(&total)
	}

	tx.Commit()
	// data response to client
	dataResp := model.ResponseObj{
		PerPage:     Per_page,
		Total:       total,
		CurrentPage: Current_Page,
		Data:        &administrators,
	}
	return c.JSON(http.StatusOK, dataResp)
}

// get param from url and map to query string
func toStringQuery(queryParams url.Values) string {
	stringQr := ""
	email_search := ""
	id_search := ""
	name_search := ""
	permission := ""
	accepting_host_name := ""
	s := []string{}
	if queryParams["email_search"] != nil {
		email_search = queryParams["email_search"][0]
	}
	if queryParams["id_search"] != nil {
		id_search = queryParams["id_search"][0]
	}
	if queryParams["name_search"] != nil {
		name_search = queryParams["name_search"][0]
	}
	if queryParams["permission"] != nil {
		permission = queryParams["permission"][0]
	}
	if queryParams["accepting_host_name"] != nil {
		accepting_host_name = queryParams["accepting_host_name"][0]
	}
	//push to array s
	if email_search != "" {
		s = append(s, "email LIKE '%"+email_search+"%'")
	}
	if id_search != "" {
		s = append(s, " administrator_id='"+id_search+"'")
	}
	if name_search != "" {
		s = append(s, "administrators.name LIKE '%"+name_search+"%'")
	}
	if permission != "" {
		s = append(s, "permission ='"+permission+"'")
	}
	if accepting_host_name != "" {
		s = append(s, "accepting_hosts.name LIKE '%"+accepting_host_name+"%'")
	}
	//array join with 'and' ->to string for query
	if len(s) > 0 {
		stringQr = strings.Join(s, " and ") //result query string
	}

	return stringQr

}

func GetAdminById(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	administratorsId := c.Param("id")
	administrator := model.Administrators{}

	if err := tx.Preload("AcceptingHost").First(&administrator, administratorsId).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"Message": err.Error(), "status": "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &administrator)
}

func PostAdmin(c echo.Context) (err error) {
	administrators := model.Administrators{}
	if err = c.Bind(&administrators); err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
	}
	tx := database.MysqlConn().Begin()
	if err := tx.Create(&administrators).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &administrators)
}

func PutAdmin(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	administrator_id := c.Param("id")
	administrator := model.Administrators{}
	if err := tx.First(&administrator, administrator_id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"Message": err.Error(), "status": "false"})
	}

	data_updated := model.Administrators{}
	if err = c.Bind(&data_updated); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "InternalServerError", "status": "false"})
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
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(), "status": "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &administrator)
}

func DeleteAdmin(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	id := c.Param("id")
	reminder := model.Administrators{}
	if err := tx.First(&reminder, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"Message": err.Error(), "status": "false"})
	}
	if err := tx.Delete(&reminder).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(), "status": "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, map[string]string{"Message": "deleted", "status": "true"})
}
