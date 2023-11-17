package adapter

import (
	"bpm-wrapper/internal/config"
	"bpm-wrapper/internal/data/dto"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
)

type Adapter interface {
	// Login
	Login(username string, password string) (*dto.LoginResponse, error)
	// Logout
	Logout(auth *dto.LoginResponse) error
	// StartProcess
	CreateProcessInstance(auth *dto.LoginResponse, processId string, variables map[string]interface{}) (string, error)
	// FindProcess
	FindProcess(auth *dto.LoginResponse, version string) (string, error)
	// FindCaseByID
	FindCaseByID(auth *dto.LoginResponse, caseId int64) (dto.FindCaseByIDResponse, error)
	// FindArchivedTasks
	FindArchivedTasks(auth *dto.LoginResponse, processId string) ([]dto.FindArchivedTasksResponse, error)
	// FindTasksByID
	FindTasksByID(auth *dto.LoginResponse, taskId string) (dto.FindTasksByIDResponse, error)
	// Assign Task
	AssignTask(auth *dto.LoginResponse, taskId string, userId string) error
	// Find Actor
	FindUser(auth *dto.LoginResponse, userName string) ([]dto.FindUserResponse, error)
	// Execute Task
	ExecuteTask(auth *dto.LoginResponse, taskID string, variables interface{}) error
	// QueryBusinessData
	QueryBusinessData(auth *dto.LoginResponse, query string) ([]dto.QueryBusinessDataResponse, error)
	// FindTaskByName
	FindTaskByName(auth *dto.LoginResponse, caseID int64, taskName string) (dto.FindTaskByNameResponse, error)
}

func New(client *http.Client, cfgBonita *config.BonitaConfig) Adapter {
	return &adapter{
		client:    client,
		cfgBonita: cfgBonita,
	}
}

type adapter struct {
	client    *http.Client
	cfg       *config.BonitaConfig
	cfgBonita *config.BonitaConfig
}

// FindTaskByName implements Adapter
func (a *adapter) FindTaskByName(auth *dto.LoginResponse, caseID int64, taskName string) (dto.FindTaskByNameResponse, error) {
	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s%s%s%d", host, port, "/bonita/API/bpm/humanTask?p=0&c=1&f=name=", url.QueryEscape(taskName), "&f=caseId=", caseID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dto.FindTaskByNameResponse{}, err
	}

	req.Header.Add("X-Bonita-API-Token", auth.BonitaToken)
	req.Header.Add("Cookie", "JSESSIONID="+auth.BonitaAuth)

	response, err := a.client.Do(req)

	if err != nil {
		return dto.FindTaskByNameResponse{}, err
	}

	defer response.Body.Close()

	var result []dto.FindTaskByNameResponse

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil && strings.Contains(err.Error(), "index out of range [0] with length 0") {
		return dto.FindTaskByNameResponse{}, nil
	}
	if err != nil {
		return dto.FindTaskByNameResponse{}, err
	}

	if len(result) < 1 {
		return dto.FindTaskByNameResponse{}, nil
	}

	return result[0], nil

}

// FindActor implements Adapter
func (a *adapter) FindUser(auth *dto.LoginResponse, userName string) ([]dto.FindUserResponse, error) {
	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s%s", host, port, "/bonita/API/identity/user?c=20&p=0&f=userName=", userName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []dto.FindUserResponse{}, err
	}

	req.Header.Add("X-Bonita-API-Token", auth.BonitaToken)
	req.Header.Add("Cookie", "JSESSIONID="+auth.BonitaAuth)

	response, err := a.client.Do(req)
	if err != nil {
		return []dto.FindUserResponse{}, err
	}

	defer response.Body.Close()

	var result []dto.FindUserResponse

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return []dto.FindUserResponse{}, err
	}

	return result, nil

}

// ExecuteTask implements Adapter
func (a *adapter) ExecuteTask(auth *dto.LoginResponse, taskID string, variables interface{}) error {
	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s%s%s", host, port, "/bonita/API/bpm/userTask/", taskID, "/execution?assign=true")

	body, err := json.Marshal(variables)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("X-Bonita-API-Token", auth.BonitaToken)
	req.Header.Add("Cookie", "JSESSIONID="+auth.BonitaAuth)

	response, err := a.client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return nil
}

// AssignTask implements Adapter
func (a *adapter) AssignTask(auth *dto.LoginResponse, taskId string, userId string) error {
	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s%s", host, port, "/bonita/API/bpm/userTask/", taskId)

	requestBody := dto.AssignTaskRequest{
		AssignedID: userId,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("X-Bonita-API-Token", auth.BonitaToken)
	req.Header.Add("Cookie", "JSESSIONID="+auth.BonitaAuth)

	response, err := a.client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return nil
}

// FindCaseByID implements Adapter
func (a *adapter) FindCaseByID(auth *dto.LoginResponse, caseId int64) (dto.FindCaseByIDResponse, error) {
	var caseData dto.FindCaseByIDResponse
	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s%d", host, port, "/bonita/API/bpm/case/", caseId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return caseData, err
	}

	req.Header.Add("X-Bonita-API-Token", auth.BonitaToken)
	req.Header.Add("Cookie", "JSESSIONID="+auth.BonitaAuth)

	response, err := a.client.Do(req)
	if err != nil {
		return caseData, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return dto.FindCaseByIDResponse{}, err
	}

	err = json.Unmarshal(body, &caseData)
	if err != nil {
		return dto.FindCaseByIDResponse{}, err
	}

	return caseData, nil

}

// Find Task by ID implements Adapter
func (a *adapter) FindTasksByID(auth *dto.LoginResponse, taskId string) (dto.FindTasksByIDResponse, error) {
	var task dto.FindTasksByIDResponse
	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s%s", host, port, "/bonita/API/bpm/task/", taskId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return task, err
	}

	req.Header.Add("X-Bonita-API-Token", auth.BonitaToken)
	req.Header.Add("Cookie", "JSESSIONID="+auth.BonitaAuth)

	response, err := a.client.Do(req)
	if err != nil {
		return task, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return task, err
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		return task, err
	}

	return task, nil
}

// Login implements Adapter
func (a *adapter) Login(username string, password string) (*dto.LoginResponse, error) {

	value := url.Values{
		"username": {a.cfgBonita.Username},
		"password": {a.cfgBonita.Password},
	}

	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s", host, port, "/bonita/loginservice")

	response, err := a.client.PostForm(url, value)
	if err != nil {
		return nil, err
	}

	var bonita_token string
	var bonita_auth string

	retrieveCookie := response.Cookies()

	for _, cookie := range retrieveCookie {
		if cookie.Name == "X-Bonita-API-Token" {
			bonita_token = cookie.Value
		}
		if cookie.Name == "JSESSIONID" {
			bonita_auth = cookie.Value
		}
	}

	bonitaResponse := dto.LoginResponse{
		BonitaToken: bonita_token,
		BonitaAuth:  bonita_auth,
	}

	return &bonitaResponse, nil
}

// Logout implements Adapter
func (a *adapter) Logout(auth *dto.LoginResponse) error {

	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s", host, port, "/bonita/logoutservice?redirect=false")

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("X-Bonita-API-Token", auth.BonitaToken)
	request.Header.Set("Cookie", "JSESSIONID="+auth.BonitaAuth)
	if err != nil {
		return err
	}

	_, err = a.client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

// CreateProcessInstance implements Adapter
func (a *adapter) CreateProcessInstance(auth *dto.LoginResponse, processId string, variables map[string]interface{}) (string, error) {
	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s", host, port, "/bonita/API/bpm/process/"+processId+"/instantiation")

	request, err := http.NewRequest("POST", url, nil)
	request.Header.Set("X-Bonita-API-Token", auth.BonitaToken)
	request.Header.Set("Cookie", "JSESSIONID="+auth.BonitaAuth)
	if err != nil {
		return "", err
	}

	response, err := a.client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseBody dto.InstansiateProcessResponse

	err = json.Unmarshal(responseData, &responseBody)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(responseBody.CaseID)), nil

}

// FindProcess implements Adapter
func (a *adapter) FindProcess(auth *dto.LoginResponse, version string) (string, error) {
	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s%s", host, port, "/bonita/API/bpm/process?p=0&c=20&f=name=TMS&f=version=", version)
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("X-Bonita-API-Token", auth.BonitaToken)
	request.Header.Set("Cookie", "JSESSIONID="+auth.BonitaAuth)
	if err != nil {
		return "", err
	}

	response, err := a.client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseBody []dto.FindProcessInstanceResponse

	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return "", err
	}

	return responseBody[0].ID, nil

}

func (a *adapter) FindArchivedTasks(auth *dto.LoginResponse, processId string) ([]dto.FindArchivedTasksResponse, error) {
	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s", host, port, "/bonita/API/bpm/archivedTask?p=0&c=20")
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("X-Bonita-API-Token", auth.BonitaToken)
	request.Header.Set("Cookie", "JSESSIONID="+auth.BonitaAuth)
	if err != nil {
		return nil, err
	}

	response, err := a.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		panic(err.Error())
	}

	var responseBody []dto.FindArchivedTasksResponse

	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil

}

// QueryBusinessData implements Adapter
func (a *adapter) QueryBusinessData(auth *dto.LoginResponse, query string) ([]dto.QueryBusinessDataResponse, error) {
	host := a.cfgBonita.Host
	port := a.cfgBonita.Port
	url := fmt.Sprintf("%s:%s%s", host, port, "/bonita/API/bdm/businessData/com.company.model.Ticket?q="+query)
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("X-Bonita-API-Token", auth.BonitaToken)
	request.Header.Set("Cookie", "JSESSIONID="+auth.BonitaAuth)
	if err != nil {
		return nil, err
	}

	response, err := a.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseBody []dto.QueryBusinessDataResponse

	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		return nil, err
	}

	return responseBody, nil

}
