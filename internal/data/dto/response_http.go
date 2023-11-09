package dto

type BaseResponse struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data"`
}

type MetaResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"OK"`
}

type LoginResponse struct {
	BonitaToken string `json:"bonita_token" example:"ed27cbeb-9953-4d77-b5a2-1f62a6c2e0bb"` // Bonita token uuidv4
	BonitaAuth  string `json:"bonita_auth" example:"C5385BFEE2969D9E46F0160C1952B0F1"`      // Bonita auth
}

type FindUserResponse struct {
	Firstname       string `json:"firstname"`
	Icon            string `json:"icon"`
	CreationDate    string `json:"creation_date"`
	UserName        string `json:"userName"`
	Title           string `json:"title"`
	CreatedByUserID string `json:"created_by_user_id"`
	Enabled         string `json:"enabled"`
	Lastname        string `json:"lastname"`
	LastConnection  string `json:"last_connection"`
	Password        string `json:"password"`
	ManagerID       string `json:"manager_id"`
	ID              string `json:"id"`
	JobTitle        string `json:"job_title"`
	LastUpdateDate  string `json:"last_update_date"`
}

type FindProcessInstanceResponse struct {
	ActivationState    string `json:"activationState" example:"ENABLED"`                                 // ENABLED, DISABLED
	Actorinitiatorid   string `json:"actorinitiatorid" example:"123456789"`                              // ID of the process instance
	ConfigurationState string `json:"configurationState" example:"RESOLVED"`                             // RESOLVED, UNRESOLVED
	DeployedBy         string `json:"deployedBy" example:"walter.bates"`                                 // User who deployed the process
	DeploymentDate     string `json:"deploymentDate" example:"2020-10-01T00:00:00.000Z"`                 // Date the process was deployed
	Description        string `json:"description" example:"This is a description of the process"`        // Description of the process
	DisplayDescription string `json:"displayDescription" example:"This is a description of the process"` // Description of the process
	DisplayName        string `json:"displayName" example:"This is the display name of the process"`     // Display name of the process
	Icon               string `json:"icon" example:"icon.png"`                                           // Icon of the process
	ID                 string `json:"id" example:"123456789"`                                            // ID of the process instance
	LastUpdateDate     string `json:"last_update_date" example:"2020-10-01T00:00:00.000Z"`               // Date the process was last updated
	Name               string `json:"name" example:"This is the name of the process"`                    // Name of the process
	Version            string `json:"version" example:"1.0"`                                             // Version of the process
}

type FindArchivedTasksResponse struct {
	DisplayDescription   string `json:"displayDescription" example:"This is a description of the task"`
	ExecutedBy           string `json:"executedBy" example:"walter.bates"` // User who executed the task
	ArchivedDate         string `json:"archivedDate" example:"2020-10-01T00:00:00.000Z"`
	RootContainerID      string `json:"rootContainerId" example:"123456789"` // ID of the process instance
	DisplayName          string `json:"displayName" example:"This is the display name of the task"`
	ExecutedBySubstitute string `json:"executedBySubstitute" example:"walter.bates"` // User who executed the task on behalf of another user
	Description          string `json:"description" example:"This is a description of the task"`
	SourceObjectID       string `json:"sourceObjectId" example:"123456789"` // ID of the process instance
	Type                 string `json:"type" example:"USER_TASK"`
	ProcessID            string `json:"processId" example:"123456789"` // ID of the process instance
	CaseID               string `json:"caseId" example:"123456789"`    // ID of the process instance
	Name                 string `json:"name" example:"This is the name of the task"`
	ReachedStateDate     string `json:"reached_state_date" example:"2020-10-01T00:00:00.000Z"`
	RootCaseID           string `json:"rootCaseId" example:"123456789"` // ID of the process instance
	ID                   string `json:"id" example:"123456789"`         // ID of the process instance
	State                string `json:"state" example:"completed"`
	ParentCaseID         string `json:"parentCaseId" example:"123456789"` // ID of the process instance
	LastUpdateDate       string `json:"last_update_date" example:"2020-10-01T00:00:00.000Z"`
}

type QueryBusinessDataResponse struct{}

type InstansiateProcessResponse struct {
	CaseID int64 `json:"caseId" example:"123456789"`
}

type FindTasksByIDResponse struct {
	DisplayDescription   string `json:"displayDescription"`
	ExecutedBySubstitute int    `json:"executedBySubstitute"`
	ProcessID            int64  `json:"processId"`
	ParentCaseID         int    `json:"parentCaseId"`
	State                string `json:"state"`
	RootContainerID      int    `json:"rootContainerId"`
	Type                 string `json:"type"`
	AssignedID           int    `json:"assigned_id"`
	AssignedDate         string `json:"assigned_date"`
	ID                   int    `json:"id"`
	ExecutedBy           int    `json:"executedBy"`
	CaseID               int    `json:"caseId"`
	Priority             string `json:"priority"`
	ActorID              int    `json:"actorId"`
	Description          string `json:"description"`
	Name                 string `json:"name"`
	ReachedStateDate     string `json:"reached_state_date"`
	RootCaseID           int    `json:"rootCaseId"`
	DisplayName          string `json:"displayName"`
	ParentTaskID         int    `json:"parentTaskId"`
	DueDate              string `json:"dueDate"`
	LastUpdateDate       string `json:"last_update_date"`
}

type FindCaseByIDResponse struct {
	ID                  string `json:"id"`
	EndDate             string `json:"end_date"`
	FailedFlowNodes     string `json:"failedFlowNodes"`
	StartedBySubstitute string `json:"startedBySubstitute"`
	Start               string `json:"start"`
	ActiveFlowNodes     string `json:"activeFlowNodes"`
	State               string `json:"state"`
	RootCaseID          string `json:"rootCaseId"`
	StartedBy           string `json:"started_by"`
	ProcessDefinitionID string `json:"processDefinitionId"`
	LastUpdateDate      string `json:"last_update_date"`
	SearchIndex1Label   string `json:"searchIndex1Label"`
	SearchIndex2Label   string `json:"searchIndex2Label"`
	SearchIndex3Label   string `json:"searchIndex3Label"`
	SearchIndex4Label   string `json:"searchIndex4Label"`
	SearchIndex5Label   string `json:"searchIndex5Label"`
	SearchIndex1Value   string `json:"searchIndex1Value"`
	SearchIndex2Value   string `json:"searchIndex2Value"`
	SearchIndex3Value   string `json:"searchIndex3Value"`
	SearchIndex4Value   string `json:"searchIndex4Value"`
	SearchIndex5Value   string `json:"searchIndex5Value"`
}

type FindTaskByNameResponse struct {
	DisplayDescription   string `json:"displayDescription"`
	ExecutedBy           string `json:"executedBy"`
	RootContainerID      string `json:"rootContainerId"`
	AssignedDate         string `json:"assigned_date"`
	DisplayName          string `json:"displayName"`
	ExecutedBySubstitute string `json:"executedBySubstitute"`
	DueDate              string `json:"dueDate"`
	Description          string `json:"description"`
	Type                 string `json:"type"`
	Priority             string `json:"priority"`
	ActorID              string `json:"actorId"`
	ProcessID            string `json:"processId"`
	CaseID               string `json:"caseId"`
	Name                 string `json:"name"`
	ReachedStateDate     string `json:"reached_state_date"`
	RootCaseID           string `json:"rootCaseId"`
	ID                   string `json:"id"`
	State                string `json:"state"`
	ParentCaseID         string `json:"parentCaseId"`
	LastUpdateDate       string `json:"last_update_date"`
	AssignedID           string `json:"assigned_id"`
}
