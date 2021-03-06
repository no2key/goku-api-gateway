package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/account"
	"github.com/eolinker/goku-api-gateway/console/module/api"
)

//AddAPI 新增接口
func AddAPI(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	userID, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	apiName := httpRequest.PostFormValue("apiName")
	requestURL := httpRequest.PostFormValue("requestURL")
	requestMethod := httpRequest.PostFormValue("requestMethod")
	protocol := httpRequest.PostFormValue("protocol")
	balanceName := httpRequest.PostFormValue("balanceName")
	targetURL := httpRequest.PostFormValue("targetURL")
	targetMethod := httpRequest.PostFormValue("targetMethod")
	isFollow := httpRequest.PostFormValue("isFollow")
	timeout := httpRequest.PostFormValue("timeout")
	retryCount := httpRequest.PostFormValue("retryCount")
	groupID := httpRequest.PostFormValue("groupID")
	projectID := httpRequest.PostFormValue("projectID")
	alertValve := httpRequest.PostFormValue("alertValve")
	managerID := httpRequest.PostFormValue("managerID")
	apiType := httpRequest.PostFormValue("apiType")
	linkApis := httpRequest.PostFormValue("linkApis")
	staticResponse := httpRequest.PostFormValue("staticResponse")
	responseDataType := httpRequest.PostFormValue("responseDataType")

	if apiName == "" {
		controller.WriteError(httpResponse, "190002", "api", "[ERROR]Illegal apiName!", nil)
		return
	}
	if isFollow != "true" && isFollow != "false" && isFollow != "" {
		controller.WriteError(httpResponse, "190008", "api", "[ERROR]Illegal isFollow!", nil)
		return

	}
	if isFollow == "" {
		isFollow = "false"
	}

	aType, err := strconv.Atoi(apiType)
	if err != nil && apiType == "" {
		controller.WriteError(httpResponse, "190012", "api", "[ERROR]Illegal apiType!", err)
		return
	}
	if responseDataType != "origin" && responseDataType != "json" && responseDataType != "xml" {
		controller.WriteError(httpResponse, "190013", "api", "[ERROR]Illegal responseDataType!", err)
		return
	}
	t, err := strconv.Atoi(timeout)
	if err != nil && timeout != "" {
		controller.WriteError(httpResponse, "190010", "api", "[ERROR]Illegal timeout!", err)
		return
	}

	count, err := strconv.Atoi(retryCount)
	if err != nil && retryCount != "" {
		controller.WriteError(httpResponse, "190011", "api", "[ERROR]Illegal retryCount!", err)
		return

	}
	if t < 1 {
		controller.WriteError(httpResponse, "190010", "api", "[ERROR]Illegal timeout!", nil)
		return
	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse, "190015", "api", "[ERROR]Illegal groupID!", err)
		return

	}
	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse, "190016", "api", "[ERROR]Illegal projectID!", err)
		return

	}
	apiValve, err := strconv.Atoi(alertValve)
	if err != nil && alertValve != "" {
		controller.WriteError(httpResponse, "190017", "api", "[ERROR]Illegal alertValve!", err)
		return

	}
	mgID, err := strconv.Atoi(managerID)
	if (err != nil && managerID != "") || mgID < -1 {
		controller.WriteError(httpResponse, "190018", "api", "[ERROR]Illegal managerID!", err)
		return

	}
	if managerID == "" {
		mgID = userID
	}

	flag, id, err := api.AddAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkApis, staticResponse, responseDataType, balanceName, protocol, pjID, gID, t, count, apiValve, mgID, userID, aType)
	if !flag {

		controller.WriteError(httpResponse,
			"190000", "api", "[ERROR]URL Repeat!", err)
		return
	}

	controller.WriteResultInfo(httpResponse, "api", "apiID", id)
	return
}

//EditAPI 编辑接口
func EditAPI(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	userID, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	apiID := httpRequest.PostFormValue("apiID")
	apiName := httpRequest.PostFormValue("apiName")
	requestURL := httpRequest.PostFormValue("requestURL")
	targetURL := httpRequest.PostFormValue("targetURL")
	requestMethod := httpRequest.PostFormValue("requestMethod")
	protocol := httpRequest.PostFormValue("protocol")
	balanceName := httpRequest.PostFormValue("balanceName")
	targetMethod := httpRequest.PostFormValue("targetMethod")
	isFollow := httpRequest.PostFormValue("isFollow")
	timeout := httpRequest.PostFormValue("timeout")
	retryCount := httpRequest.PostFormValue("retryCount")
	groupID := httpRequest.PostFormValue("groupID")
	projectID := httpRequest.PostFormValue("projectID")
	alertValve := httpRequest.PostFormValue("alertValve")
	managerID := httpRequest.PostFormValue("managerID")
	linkApis := httpRequest.PostFormValue("linkApis")
	staticResponse := httpRequest.PostFormValue("staticResponse")
	responseDataType := httpRequest.PostFormValue("responseDataType")
	if apiName == "" {
		controller.WriteError(httpResponse, "190002", "api", "[ERROR]Illegal apiName!", nil)
		return

	}
	aID, err := strconv.Atoi(apiID)
	if err != nil {
		controller.WriteError(httpResponse, "190001", "api", "[ERROR]Illegal apiID!", nil)
		return
	}
	if responseDataType != "origin" && responseDataType != "json" && responseDataType != "xml" {
		controller.WriteError(httpResponse, "190013", "api", "[ERROR]Illegal responseDataType!", err)
		return
	}

	if isFollow != "true" && isFollow != "false" && isFollow != "" {
		controller.WriteError(httpResponse, "190008", "api", "[ERROR]Illegal isFollow!", nil)
		return

	}
	if isFollow == "" {
		isFollow = "false"
	}
	t, err := strconv.Atoi(timeout)
	if err != nil && timeout != "" {
		controller.WriteError(httpResponse, "190010", "api", "[ERROR]Illegal timeout!", nil)
		return

	}
	if t < 1 {
		controller.WriteError(httpResponse, "190010", "api", "[ERROR]Illegal timeout!", nil)
		return

	}
	count, err := strconv.Atoi(retryCount)
	if err != nil && retryCount != "" {
		controller.WriteError(httpResponse, "190011", "api", "[ERROR]Illegal retryCount!", nil)
		return

	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse, "190015", "api", "[ERROR]Illegal groupID!", nil)
		return

	}
	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse, "190016", "api", "[ERROR]Illegal projectID!", nil)
		return

	}
	apiValve, err := strconv.Atoi(alertValve)
	if err != nil && alertValve != "" {
		controller.WriteError(httpResponse, "190017", "api", "[ERROR]Illegal alertValve!", nil)
		return

	}
	mgID, err := strconv.Atoi(managerID)
	if (err != nil && managerID != "") || mgID < -1 {
		controller.WriteError(httpResponse, "190018", "api", "[ERROR]Illegal managerID!", nil)
		return

	}
	if managerID == "" {
		mgID = userID
	}

	flag, err := api.EditAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkApis, staticResponse, responseDataType, balanceName, protocol, pjID, gID, t, count, apiValve, aID, mgID, userID)
	if !flag {

		controller.WriteError(httpResponse, "190000", "api", "[ERROR]apiID does not exist!", err)
		return

	}

	controller.WriteResultInfo(httpResponse, "api", "", nil)

	return
}

//GetAPIInfo 获取接口信息
func GetAPIInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationREAD)
	if e != nil {
		return
	}

	apiID := httpRequest.PostFormValue("apiID")

	aID, err := strconv.Atoi(apiID)
	if err != nil {
		controller.WriteError(httpResponse, "190001", "api", "[ERROR]Illegal apiID!", nil)
		return

	}
	flag, result, err := api.GetAPIInfo(aID)
	if !flag {
		controller.WriteError(httpResponse, "190000", "api", "[ERROR]The api does not exist!", err)
		return

	}
	controller.WriteResultInfo(httpResponse, "api", "apiInfo", result)

	return
}

// GetAPIIDList 获取接口ID列表
func GetAPIIDList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationREAD)
	if e != nil {
		return
	}

	httpRequest.ParseForm()
	projectID := httpRequest.Form.Get("projectID")
	groupID := httpRequest.Form.Get("groupID")
	keyword := httpRequest.Form.Get("keyword")
	condition := httpRequest.Form.Get("condition")
	idsStr := httpRequest.Form.Get("ids")

	pjID, e := strconv.Atoi(projectID)
	if e != nil {
		controller.WriteError(httpResponse, "190016", "api", "[ERROR]Illegal projectID!", e)
		return
	}
	gID, e := strconv.Atoi(groupID)
	if e != nil {
		if groupID != "" {
			controller.WriteError(httpResponse, "190015", "api", "[ERROR]Illegal groupID!", e)
			return
		}
		gID = -1
	}
	op, e := strconv.Atoi(condition)
	if e != nil {
		if condition != "" {
			controller.WriteError(httpResponse, "190019", "api", "[ERROR]Illegal condition!", e)
			return
		}
	}
	ids := make([]int, 0)
	json.Unmarshal([]byte(idsStr), &ids)

	_, result, _ := api.GetAPIIDList(pjID, gID, keyword, op, ids)

	// controller.WriteResultInfo(httpResponse, "api", "apiList", result)
	controller.WriteResultInfoWithPage(httpResponse, "api", "apiIDList", result, &controller.PageInfo{
		ItemNum:  len(result),
		TotalNum: len(result),
	})
	return
}

//GetAPIList 获取接口列表
func GetAPIList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationREAD)
	if e != nil {
		return
	}

	httpRequest.ParseForm()
	projectID := httpRequest.Form.Get("projectID")
	groupID := httpRequest.Form.Get("groupID")
	keyword := httpRequest.Form.Get("keyword")
	condition := httpRequest.Form.Get("condition")
	idsStr := httpRequest.Form.Get("ids")
	page := httpRequest.Form.Get("page")
	pageSize := httpRequest.Form.Get("pageSize")

	p, e := strconv.Atoi(page)
	if e != nil {
		p = 1
	}
	pSize, e := strconv.Atoi(pageSize)
	if e != nil {
		pSize = 15
	}

	pjID, e := strconv.Atoi(projectID)
	if e != nil {
		controller.WriteError(httpResponse, "190016", "api", "[ERROR]Illegal projectID!", e)
		return
	}
	gID, e := strconv.Atoi(groupID)
	if e != nil {
		if groupID != "" {
			controller.WriteError(httpResponse, "190015", "api", "[ERROR]Illegal groupID!", e)
			return
		}
		gID = -1
	}
	op, e := strconv.Atoi(condition)
	if e != nil {
		if condition != "" {
			controller.WriteError(httpResponse, "190019", "api", "[ERROR]Illegal condition!", e)
			return
		}
	}
	result := make([]map[string]interface{}, 0)
	ids := make([]int, 0)
	json.Unmarshal([]byte(idsStr), &ids)

	_, result, count, _ := api.GetAPIList(pjID, gID, keyword, op, p, pSize, ids)

	// controller.WriteResultInfo(httpResponse, "api", "apiList", result)
	controller.WriteResultInfoWithPage(httpResponse, "api", "apiList", result, &controller.PageInfo{
		ItemNum:  len(result),
		TotalNum: count,
		Page:     p,
		PageSize: pSize,
	})
	return
}

// BatchEditAPIGroup 批量修改接口分组
func BatchEditAPIGroup(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	apiIDList := httpRequest.PostFormValue("apiIDList")
	groupID := httpRequest.PostFormValue("groupID")
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse, "190015", "api", "[ERROR]Illegal groupID!", nil)
		return
	}
	flag, result, err := api.BatchEditAPIGroup(strings.Split(apiIDList, ","), gID)
	if !flag {
		controller.WriteError(httpResponse, "190015", "api", result, err)
		return
	}
	controller.WriteResultInfo(httpResponse, "api", "", nil)

	return
}

//BatchDeleteAPI 批量删除接口
func BatchDeleteAPI(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	apiIDList := httpRequest.PostFormValue("apiIDList")

	flag, result, err := api.BatchDeleteAPI(apiIDList)
	if !flag {

		controller.WriteError(httpResponse, "190000", "api", result, err)
		return

	}

	controller.WriteResultInfo(httpResponse, "api", "", nil)
	return
}

//GetAPIManagerList 获取接口负责人列表
func GetAPIManagerList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationREAD)
	if e != nil {
		return
	}

	flag, result, err := account.GetUserListWithPermission("apiManagement", "edit")
	if !flag {
		controller.WriteError(httpResponse, "190000", "api", err.Error(), err)
		return
	}
	controller.WriteResultInfo(httpResponse, "api", "userList", result)
	return
}

//CopyAPI 复制接口
func CopyAPI(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	userID, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationAPI, controller.OperationEDIT)
	if e != nil {
		return
	}

	apiID := httpRequest.PostFormValue("apiID")
	apiName := httpRequest.PostFormValue("apiName")
	requestURL := httpRequest.PostFormValue("requestURL")
	targetURL := httpRequest.PostFormValue("targetURL")
	requestMethod := httpRequest.PostFormValue("requestMethod")
	protocol := httpRequest.PostFormValue("protocol")
	balanceName := httpRequest.PostFormValue("balanceName")
	targetMethod := httpRequest.PostFormValue("targetMethod")
	isFollow := httpRequest.PostFormValue("isFollow")
	groupID := httpRequest.PostFormValue("groupID")
	projectID := httpRequest.PostFormValue("projectID")
	if apiName == "" {
		controller.WriteError(httpResponse, "190002", "api", "[ERROR]Illegal apiName!", nil)
		return

	}
	aID, err := strconv.Atoi(apiID)
	if err != nil {
		controller.WriteError(httpResponse, "190001", "api", "[ERROR]Illegal apiID!", nil)
		return
	}

	if isFollow != "true" && isFollow != "false" && isFollow != "" {
		controller.WriteError(httpResponse, "190008", "api", "[ERROR]Illegal isFollow!", nil)
		return

	}
	if isFollow == "" {
		isFollow = "false"
	}
	gID, err := strconv.Atoi(groupID)
	if err != nil {
		controller.WriteError(httpResponse, "190015", "api", "[ERROR]Illegal groupID!", nil)
		return

	}
	pjID, err := strconv.Atoi(projectID)
	if err != nil {
		controller.WriteError(httpResponse, "190016", "api", "[ERROR]Illegal projectID!", nil)
		return

	}
	flag, apiInfo, err := api.GetAPIInfo(aID)
	if !flag {
		controller.WriteError(httpResponse, "190000", "api", "[ERROR]apiID does not exist!", nil)
		return
	}
	linkApis, _ := json.Marshal(apiInfo.LinkAPIs)
	flag, id, err := api.AddAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, string(linkApis), apiInfo.StaticResponse, apiInfo.ResponseDataType, balanceName, protocol, pjID, gID, apiInfo.Timeout, apiInfo.RetryConut, apiInfo.Valve, apiInfo.ManagerID, userID, apiInfo.APIType)
	if !flag {
		controller.WriteError(httpResponse, "190000", "api", "[ERROR]Fail to add api!", err)
		return

	}

	controller.WriteResultInfo(httpResponse, "api", "apiID", id)
	return
}
