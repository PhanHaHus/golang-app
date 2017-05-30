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

func GetAllApplications(c echo.Context) (err error) {
	var total int
	applications := []model.Applications{}
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

	stringQr := toStringQueryApplications(queryParams) // get query string filter

	// calculate offset
	var offset = (Current_Page - 1) * Per_page
	tx := database.MysqlConn().Begin()
	if query != "" {
		// when search
		tx.Debug().Order("application_id desc").Offset(offset).Limit(Per_page).Where("applications.name LIKE ?", "%"+query+"%").Preload("AcceptingHost").Find(&applications).Count(&total)
	} else {
		//when load page
		tx.Debug().Order("application_id desc").Offset(offset).Limit(Per_page).Preload("AcceptingHost").Find(&applications).Count(&total)
	}

	if stringQr != "" {
		//when type both
		tx.Debug().Joins("JOIN accepting_hosts ON accepting_hosts.accepting_host_id = applications.accepting_host_id").Order("application_id desc").Offset(offset).Limit(Per_page).Preload("AcceptingHost").Where("applications.name LIKE ?", "%"+query+"%").Where(stringQr).Find(&applications).Count(&total)
	}

	tx.Commit()
	// data response to client
	dataResp := model.ResponseObj{
		PerPage:     Per_page,
		Total:       total,
		CurrentPage: Current_Page,
		Data:        &applications,
	}
	return c.JSON(http.StatusOK, dataResp)
}

// get param from url and map to query string
func toStringQueryApplications(queryParams url.Values) string {
	stringQr := ""
	id_search := ""
	name_search := ""
	ip_search := ""
	port_search := ""
	hostname_search := ""
	type_search := ""
	accepting_host_name := ""
	s := []string{}
	if queryParams["type_search"] != nil {
		type_search = queryParams["type_search"][0]
	}
	if queryParams["id_search"] != nil {
		id_search = queryParams["id_search"][0]
	}
	if queryParams["name_search"] != nil {
		name_search = queryParams["name_search"][0]
	}
	if queryParams["ip_search"] != nil {
		ip_search = queryParams["ip_search"][0]
	}
	if queryParams["port_search"] != nil {
		port_search = queryParams["port_search"][0]
	}
	if queryParams["hostname_search"] != nil {
		hostname_search = queryParams["hostname_search"][0]
	}
	if queryParams["accepting_host_name"] != nil {
		accepting_host_name = queryParams["accepting_host_name"][0]
	}
	//push to array s
	if type_search != "" {
		s = append(s, " applications.application_type='"+type_search+"'")
	}
	if id_search != "" {
		s = append(s, " application_id='"+id_search+"'")
	}
	if name_search != "" {
		s = append(s, "applications.name LIKE '%"+name_search+"%'")
	}
	if ip_search != "" {
		s = append(s, "applications.ip LIKE '%"+ip_search+"%'")
	}
	if port_search != "" {
		s = append(s, "applications.port LIKE '%"+port_search+"%'")
	}
	if hostname_search != "" {
		s = append(s, "applications.host_name LIKE '%"+hostname_search+"%'")
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

//search for select option in html
func SearchAppCtrl(c echo.Context) (err error) {
	// if exist param query and param table on url
	query := c.QueryParam("query")
	table := c.QueryParam("table")
	if query != "" && table != "" {
		switch table {
		case "accepting_hosts":
			modelQuery := []model.AcceptingHosts{}
			tx := database.MysqlConn().Begin()
			tx.Where("name LIKE ?", "%"+query+"%").Find(&modelQuery)
			tx.Commit()
			return c.JSON(http.StatusOK, &modelQuery)

		default:
			return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: "Not found table!", Status: "false"})
		}
	}

	return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
}

func GetApplicationsById(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	applicationsId := c.Param("id")
	applications := model.Applications{}

	if err := tx.Preload("AcceptingHost").First(&applications, applicationsId).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"Message": err.Error(), "status": "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &applications)
}

func PostApplications(c echo.Context) (err error) {
	applications := model.Applications{}
	if err = c.Bind(&applications); err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
	}
	tx := database.MysqlConn().Begin()
	if err := tx.Create(&applications).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &applications)
}

func PutApplications(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	application_id := c.Param("id")
	applications := model.Applications{}
	if err := tx.First(&applications, application_id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"Message": err.Error(), "status": "false"})
	}

	data_updated := model.Applications{}
	if err = c.Bind(&data_updated); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "InternalServerError", "status": "false"})
	}

	applications.Name = data_updated.Name
	applications.Description = data_updated.Description
	applications.Ip = data_updated.Ip
	applications.Port = data_updated.Port
	applications.HostName = data_updated.HostName
	applications.IsValidUserRequired = data_updated.IsValidUserRequired
	applications.IsValidDeviceRequired = data_updated.IsValidDeviceRequired
	applications.ApplicationType = data_updated.ApplicationType
	applications.AcceptingHostId = data_updated.AcceptingHostId
	applications.Enabled = data_updated.Enabled
	applications.CreatedById = data_updated.CreatedById

	if err := tx.Save(&applications).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(), "status": "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &applications)
}

func DeleteApplications(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	id := c.Param("id")
	reminder := model.Applications{}
	if err := tx.First(&reminder, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"Message": err.Error(), "status": "false"})
	}
	if err := tx.Delete(&reminder).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error(), "status": "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, map[string]string{"Message": "deleted", "status": "true"})
}
