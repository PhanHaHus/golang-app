package controller

import (
	_ "encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	database "../db"
	model "../model"
	"github.com/labstack/echo"
)

func GetAllAccessRules(c echo.Context) (err error) {
	var total int
	accessrules := []model.AccessRules{}
	paginateParams := model.NewPaginateParams()

	//set value default
	Current_Page := paginateParams.CurrentPage
	Per_page := paginateParams.PerPage //set value default
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
	stringQr := toStringQueryAccessRule(queryParams) // get query string filter

	//calculate offset
	var offset = (Current_Page - 1) * Per_page

	tx := database.MysqlConn().Begin()
	if query != "" {
		// when search
		tx.Debug().Order("access_rule_id desc").Offset(offset).Limit(Per_page).Joins(" join applications ON applications.application_id = access_rules.application_id").Joins(" join groups ON groups.group_id = access_rules.group_id").Joins(" join devices ON devices.device_id = access_rules.device_id").Joins(" join users ON users.user_id = access_rules.user_id").Where("users.name LIKE ?", "%"+query+"%").Or("applications.name LIKE ?", "%"+query+"%").Or("devices.name LIKE ?", "%"+query+"%").Or("access_rules.description LIKE ?", "%"+query+"%").Or("groups.name LIKE ?", "%"+query+"%").Preload("Application").Preload("Device").Preload("User").Preload("Group").Preload("CreatedByUser").Find(&accessrules).Count(&total)
	} else {
		// eager load relationship
		tx.Debug().Order("access_rule_id desc").Offset(offset).Limit(Per_page).Preload("Application").Preload("User").Preload("Device").Preload("Group").Preload("CreatedByUser").Find(&accessrules).Count(&total)
	}

	if stringQr != "" {
		//when type both
		tx.Debug().Joins("join applications ON applications.application_id = access_rules.application_id").Joins(" join groups ON groups.group_id = access_rules.group_id").Joins(" join devices ON devices.device_id = access_rules.device_id").Joins(" join users ON users.user_id = access_rules.user_id").Joins(" join administrators ON administrators.administrator_id = access_rules.created_by_id").Order("access_rule_id desc").Offset(offset).Limit(Per_page).Where(stringQr).Preload("Application").Preload("User").Preload("Device").Preload("Group").Preload("CreatedByUser").Find(&accessrules).Count(&total)
	}

	tx.Commit()
	// data response to client
	dataResp := model.ResponseObj{
		PerPage:     Per_page,
		Total:       total,
		CurrentPage: Current_Page,
		Data:        &accessrules,
	}
	return c.JSON(http.StatusOK, dataResp)
}

// get param from url and map to query string
func toStringQueryAccessRule(queryParams url.Values) string {
	stringQr := ""
	id_search := ""
	application_search := ""
	description_search := ""
	group_search := ""
	user_search := ""
	device_search := ""
	byuser_search := ""
	access_rule_type := ""
	s := []string{}
	if queryParams["application_search"] != nil {
		application_search = queryParams["application_search"][0]
	}
	if queryParams["id_search"] != nil {
		id_search = queryParams["id_search"][0]
	}
	if queryParams["description_search"] != nil {
		description_search = queryParams["description_search"][0]
	}
	if queryParams["group_search"] != nil {
		group_search = queryParams["group_search"][0]
	}
	if queryParams["user_search"] != nil {
		user_search = queryParams["user_search"][0]
	}
	if queryParams["device_search"] != nil {
		device_search = queryParams["device_search"][0]
	}
	if queryParams["byuser_search"] != nil {
		byuser_search = queryParams["byuser_search"][0]
	}
	if queryParams["access_rule_type"] != nil {
		access_rule_type = queryParams["access_rule_type"][0]
	}
	//push to array s
	if application_search != "" {
		s = append(s, "applications.name LIKE '%"+application_search+"%'")
	}
	if id_search != "" {
		s = append(s, " access_rules.access_rule_id='"+id_search+"'")
	}
	if description_search != "" {
		s = append(s, "access_rules.description LIKE '%"+description_search+"%'")
	}
	if user_search != "" {
		s = append(s, "users.name LIKE '%"+user_search+"%'")
	}
	if group_search != "" {
		s = append(s, "groups.name LIKE '%"+group_search+"%'")
	}
	if device_search != "" {
		s = append(s, "devices.name LIKE '%"+device_search+"%'")
	}
	if byuser_search != "" {
		s = append(s, "administrators.name LIKE '%"+byuser_search+"%'")
	}
	if access_rule_type != "" {
		s = append(s, " access_rules.access_rule_type='"+access_rule_type+"'")
	}
	//array join with 'and' ->to string for query
	if len(s) > 0 {
		stringQr = strings.Join(s, " and ") //result query string
	}
	return stringQr
}

func SearchACLCtrl(c echo.Context) (err error) {
	// if exist param query and param table on url
	query := c.QueryParam("query")
	table := c.QueryParam("table")
	if query != "" && table != "" {
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
			return c.JSON(http.StatusNotFound, model.Status{StatusCode: http.StatusNotFound, Message: "Not found table!", Status: "false"})
		}
	}

	return c.JSON(http.StatusNotFound, model.Status{StatusCode: http.StatusNotFound, Message: err.Error(), Status: "false"})
}

func GetAccessRuleById(c echo.Context) (err error) {
	accessrulesId := c.Param("id")
	accessrules := model.AccessRules{}
	tx := database.MysqlConn().Begin()
	if err := tx.Preload("Application").Preload("User").Preload("Device").Preload("Group").Preload("CreatedByUser").First(&accessrules, accessrulesId).Error; err != nil {
		return c.JSON(http.StatusNotFound, model.Status{StatusCode: http.StatusNotFound, Message: err.Error(), Status: "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &accessrules)
}

func PostAccessRule(c echo.Context) (err error) {

	accessrules := model.AccessRules{}
	if err = c.Bind(&accessrules); err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
	}
	log.Println("accessrules")
	log.Println(&accessrules)
	tx := database.MysqlConn().Begin()
	if err := tx.Create(&accessrules).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &accessrules)
}

func PutAccessRule(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	administrator_id := c.Param("id")
	administrator := model.AccessRules{}
	if err := tx.First(&administrator, administrator_id).Error; err != nil {
		return c.JSON(http.StatusNotFound, model.Status{StatusCode: http.StatusNotFound, Message: err.Error(), Status: "false"})
	}

	data_updated := model.AccessRules{}
	if err = c.Bind(&data_updated); err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: "InternalServerError", Status: "false"})
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
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, &administrator)
}

func DeleteAccessRule(c echo.Context) (err error) {
	tx := database.MysqlConn().Begin()
	id := c.Param("id")
	accessRules := model.AccessRules{}
	if err := tx.First(&accessRules, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, model.Status{StatusCode: http.StatusNotFound, Message: err.Error(), Status: "false"})
	}
	if err := tx.Delete(&accessRules).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Status{StatusCode: http.StatusInternalServerError, Message: err.Error(), Status: "false"})
	}
	tx.Commit()
	return c.JSON(http.StatusOK, model.Status{StatusCode: http.StatusInternalServerError, Message: "deleted", Status: "true"})
}
