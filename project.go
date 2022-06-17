package opendsd

import (
	"encoding/json"
	"fmt"
	"io"
)

type Customer struct {
	ProjectID  int    `json:"ProjectId"`
	CustomerID int    `json:"CustomerId"`
	Role       string `json:"Role"`
	FirmName   string `json:"FirmName"`
	Name       string `json:"Name"`
}

//Review struct of marshal project
type Review struct {
	ReviewCycleID int    `json:"ReviewCycleId"`
	ReviewID      int    `json:"ReviewId"`
	Discipline    string `json:"Discipline"`
	Status        string `json:"Status"`
	DueDate       string `json:"DueDate"`
	CompletedDate string `json:"CompletedDate"`
	Performance   string `json:"Performance"`
	Name          string `json:"Name"`
	Phone         string `json:"Phone"`
	Email         string `json:"Email"`
	IsActive      bool   `json:"IsActive"`
}

//ReviewCycle struct of marshal project
type ReviewCycle struct {
	ReviewCycleID  int       `json:"ReviewCycleId"`
	CycleNum       int       `json:"CycleNum"`
	Method         string    `json:"Method"`
	Status         string    `json:"Status"`
	StatusSequence int       `json:"StatusSequence"`
	SubmitDate     string    `json:"SubmitDate"`
	DueDate        string    `json:"DueDate"`
	CloseDate      string    `json:"CloseDate"`
	Performance    string    `json:"Performance"`
	Reviews        []*Review `json:"Reviews"`
}

//SignOff struct of marshal project
type SignOff struct {
	DisciplineID          int    `json:"DisciplineId"`
	DisciplineDescription string `json:"DisciplineDescription"`
	SignedDate            string `json:"SignedDate"`
}

//ProjectApproval struct of marshal project
type ProjectApproval struct {
	JobID               int    `json:"JobId"`
	ApprovalID          int    `json:"ApprovalId"`
	Type                string `json:"Type"`
	Status              string `json:"Status"`
	Scope               string `json:"Scope"`
	Depiction           string `json:"Depiction"`
	IssuedBy            string `json:"IssuedBy"`
	IssueDate           string `json:"IssueDate"`
	FirstInspectionDate string `json:"FirstInspectionDate"`
	CompleteCancelDate  string `json:"CompleteCancelDate"`
	PermitHolder        string `json:"PermitHolder"`
	NetChangeDU         string `json:"NetChangeDU"`
	Valuation           string `json:"Valuation"`
	SquareFootage       string `json:"SquareFootage"`
}

//Job struct of marshal project
type Job struct {
	ProjectID             int               `json:"ProjectId"`
	JobID                 int               `json:"JobId"`
	Description           string            `json:"Description"`
	APN                   string            `json:"APN"`
	StreetAddress         string            `json:"StreetAddress"`
	MapReference          string            `json:"MapReference"`
	SortableStreetAddress string            `json:"SortableStreetAddress"`
	Latitude              float64           `json:"Latitude"`
	Longitude             float64           `json:"Longitude"`
	NAD83Northing         string            `json:"NAD83Northing"`
	NAD83Easting          string            `json:"NAD83Easting"`
	JobFeesSubTotal       string            `json:"JobFeesSubTotal"`
	SignOffs              []SignOff         `json:"SignOffs"`
	ApprovalInfo          []interface{}     `json:"ApprovalInfo"`
	Approvals             []ProjectApproval `json:"Approvals"`
}

//Fee struct of marshal project
type Fee struct {
	FeeID            int    `json:"FeeId"`
	Description      string `json:"Description"`
	Category         string `json:"Category"`
	Unit             string `json:"Unit"`
	QuantityRequired int    `json:"QuantityRequired"`
	QuantityPaid     int    `json:"QuantityPaid"`
	InvoiceID        int    `json:"InvoiceId"`
	InvoiceStatus    string `json:"InvoiceStatus"`
	ProjectID        int    `json:"ProjectId"`
}

// ProjectInvoice struct of marshal project
type ProjectInvoice struct {
	InvoiceID        int    `json:"InvoiceId"`
	InvoiceIssueDate string `json:"InvoiceIssueDate"`
	InvoiceStatus    string `json:"InvoiceStatus"`
	ProjectID        int    `json:"ProjectId"`
}

// Header struct of marshal project
type Header struct {
	Jurisdiction  string `json:"Jurisdiction"`
	Agency        string `json:"Agency"`
	AgencyAddress string `json:"AgencyAddress"`
	AgencyWebsite string `json:"AgencyWebsite"`
	ExtractSystem string `json:"ExtractSystem"`
	ExtractDate   string `json:"ExtractDate"`
	ExtractQuery  string `json:"ExtractQuery"`
}

// ProjectManager struct of marshal project
type ProjectManager struct {
	ProjectManagerID int    `json:"ProjectManagerId"`
	Name             string `json:"Name"`
	PhoneNum         string `json:"PhoneNum"`
	EmailAddress     string `json:"EmailAddress"`
	ActiveIndicator  bool   `json:"ActiveIndicator"`
}

// Project will be refactored
type Project struct {
	Customers             []Customer       `json:"Customers"`
	ReviewCycles          []ReviewCycle    `json:"ReviewCycles"`
	Jobs                  []Job            `json:"Jobs"`
	Fees                  []Fee            `json:"Fees"`
	Invoices              []ProjectInvoice `json:"Invoices"`
	ProjectID             int              `json:"ProjectId"`
	Title                 string           `json:"Title"`
	Scope                 string           `json:"Scope"`
	ApplicationExpiration Timestamp        `json:"ApplicationExpiration"`
	ApplicationExpired    bool             `json:"ApplicationExpired"`
	AdminHold             bool             `json:"AdminHold"`
	DevelopmentID         int              `json:"DevelopmentId"`
	DevelopmentTitle      string           `json:"DevelopmentTitle"`
	ApplicationDate       Timestamp        `json:"ApplicationDate"`
	AccountNum            string           `json:"AccountNum"`
	JobOrderNum           interface{}      `json:"JobOrderNum"`
	Header                []Header         `json:"Header"`
	ProjectManagerID      int              `json:"ProjectManagerId"`
	ProjectManager        struct {
		ProjectManagerID int    `json:"ProjectManagerId"`
		Name             string `json:"Name"`
		PhoneNum         string `json:"PhoneNum"`
		EmailAddress     string `json:"EmailAddress"`
		ActiveIndicator  bool   `json:"ActiveIndicator"`
	} `json:"ProjectManager"`
}

func DecodeProject(r io.Reader) (*Project, error) {
	var err error
	var project Project

	if err = json.NewDecoder(r).Decode(&project); err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) ProjectByID(id int) (*Project, error) {
	var err error
	var p Project

	uri := fmt.Sprintf("/project/%v", id)
	if err = c.get(uri, &p); err != nil {
		return nil, err
	}

	//	this is necessary since the API does not report a 404 on not found responses
	if p.ProjectID != id {
		return nil, APIError{
			ErrorMessage: fmt.Sprintf("Project with ID: %v could not be found.", id),
		}
	}

	return &p, nil
}
