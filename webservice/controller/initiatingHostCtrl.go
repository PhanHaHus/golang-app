package controller

import (
	_ "encoding/json"
	_ "log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	database "../db"
	model "../model"
	"github.com/labstack/echo"
)

func GetInitiatingHost(c echo.Context) (err error) {
	var total int
	initiatingHosts := []model.InitiatingHosts{}
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

	stringQr := toStringQueryInitialHost(queryParams) // get query string filter

	//calculate offset
	var offset = (Current_Page - 1) * Per_page
	tx := database.MysqlConn().Begin()
	if query != "" {
		// when search
		tx.Debug().Order("initiating_host_id desc").Offset(offset).Limit(Per_page).Where("name LIKE ?", "%"+query+"%").Or("description LIKE ?", "%"+query+"%").Find(&initiatingHosts).Count(&total)
	} else {
		tx.Debug().Order("initiating_host_id desc").Offset(offset).Limit(Per_page).Find(&initiatingHosts).Count(&total)
	}
	if stringQr != "" {
		//when type both
		tx.Debug().Order("initiating_host_id desc").Offset(offset).Limit(Per_page).Where(stringQr).Find(&initiatingHosts).Count(&total)
	}

	tx.Commit()
	// data response to client
	dataResp := model.ResponseObj{
		PerPage:     Per_page,
		Total:       total,
		CurrentPage: Current_Page,
		Data:        &initiatingHosts,
	}
	return c.JSON(http.StatusOK, dataResp)
}

// get param from url and map to query string
func toStringQueryInitialHost(queryParams url.Values) string {
	stringQr := ""
	id_search := ""
	name_search := ""
	description_search := ""
	s := []string{}

	if queryParams["id_search"] != nil {
		id_search = queryParams["id_search"][0]
	}
	if queryParams["name_search"] != nil {
		name_search = queryParams["name_search"][0]
	}
	if queryParams["description_search"] != nil {
		description_search = queryParams["description_search"][0]
	}
	//push to array s
	if id_search != "" {
		s = append(s, " initiating_host_id='"+id_search+"'")
	}
	if name_search != "" {
		s = append(s, "initiating_hosts.name LIKE '%"+name_search+"%'")
	}
	if description_search != "" {
		s = append(s, "initiating_hosts.description LIKE '%"+description_search+"%'")
	}

	//array join with 'and' ->to string for query
	if len(s) > 0 {
		stringQr = strings.Join(s, " and ") //result query string
	}
	return stringQr

}

func GetInitiatingHostById(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	initiatingHostsId := c.Param("id")
	initiatingHost := model.InitiatingHosts{}

	if err := tx.First(&initiatingHost, initiatingHostsId).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"Message": err.Error(), "status": "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &initiatingHost)
}

func PostInitiatingHost(c echo.Context) (err error) {
	initiatingHosts := model.InitiatingHosts{}
	if err = c.Bind(&initiatingHosts); err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
	}
	tx := database.MysqlConn().Begin()
	if err := tx.Create(&initiatingHosts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &initiatingHosts)
}

func PutInitiatingHost(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	initiatingHost_id := c.Param("id")
	initiatingHost := model.InitiatingHosts{}
	if err := tx.First(&initiatingHost, initiatingHost_id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"Message": err.Error(), "status": "false"})
	}

	data_updated := model.InitiatingHosts{}
	if err = c.Bind(&data_updated); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "InternalServerError", "status": "false"})
	}

	initiatingHost.Name = data_updated.Name
	initiatingHost.Description = data_updated.Description
	initiatingHost.Enabled = data_updated.Enabled
	initiatingHost.CreatedById = data_updated.CreatedById

	if err := tx.Save(&initiatingHost).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(), "status": "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &initiatingHost)
}

func DeleteInitiatingHost(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	id := c.Param("id")
	initiatingHost := model.InitiatingHosts{}
	if err := tx.First(&initiatingHost, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"Message": err.Error(), "status": "false"})
	}
	if err := tx.Delete(&initiatingHost).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(), "status": "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, map[string]string{"Message": "deleted", "status": "true"})
}
