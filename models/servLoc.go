package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type ServLoc struct {
	ID               bson.ObjectId      `bson:"_id" json:"id"`
	BusinessEntityID string      `bson:"businessEntityId" json:"businessEntityId"`
	LocationID       string      `bson:"locationId" json:"locationId"`
	ProvID           string      `bson:"provId" json:"provId"`
	BusinessEntity   string      `bson:"businessEntity" json:"businessEntity"`
	Status           interface{} `bson:"status" json:"status"`
	TinID            string      `bson:"tinId" json:"tinId"`
	Tin              string      `bson:"tin" json:"tin"`
	ProviderName     string      `bson:"providerName" json:"providerName"`
	Identifiers      []struct {
		Type           string      `bson:"type" json:"type"`
		TypeID         string      `bson:"typeId" json:"typeId"`
		IdentifierID   string      `bson:"identifierId" json:"identifierId"`
		Identifier     string      `bson:"identifier" json:"identifier"`
		State          interface{} `bson:"state" json:"state"`
		StateName      interface{} `bson:"stateName" json:"stateName"`
		StateCode      interface{} `bson:"stateCode" json:"stateCode"`
		StartDate      interface{} `bson:"startDate" json:"startDate"`
		EndDate        interface{} `bson:"endDate" json:"endDate"`
		ValidationDate interface{} `bson:"validationDate" json:"validationDate"`
		UpdateDate     time.Time   `bson:"updateDate" json:"updateDate"`
		UpdateBy       interface{} `bson:"updateBy" json:"updateBy"`
		IsVoid         bool        `bson:"isVoid" json:"isVoid"`
	} `bson:"identifiers" json:"identifiers"`
	ProvhatCode    []string    `bson:"provhatCode" json:"provhatCode"`
	MedicareNumber interface{} `bson:"medicareNumber" json:"medicareNumber"`
	TypeID         string      `bson:"typeId" json:"typeId"`
	Type           string      `bson:"type" json:"type"`
	AcredCert      []struct {
		AccredCertID string      `bson:"accredCertId" json:"accredCertId"`
		AttributeID  string      `bson:"attributeId" json:"attributeId"`
		Type         string      `bson:"type" json:"type"`
		TypeID       string      `bson:"typeId" json:"typeId"`
		Number       interface{} `bson:"number" json:"number"`
		Level        interface{} `bson:"level" json:"level"`
		Identifier   interface{} `bson:"identifier" json:"identifier"`
		Status       string      `bson:"status" json:"status"`
		State        interface{} `bson:"state" json:"state"`
		StateName    interface{} `bson:"stateName" json:"stateName"`
		StateCode    interface{} `bson:"stateCode" json:"stateCode"`
		UpdateDate   time.Time   `bson:"updateDate" json:"updateDate"`
		UpdateBy     interface{} `bson:"updateBy" json:"updateBy"`
		IsVoid       bool        `bson:"isVoid" json:"isVoid"`
	} `bson:"acredCert" json:"acredCert"`
	Phones []struct {
		Type         string      `bson:"type" json:"type"`
		Number       string      `bson:"number" json:"number"`
		Extension    interface{} `bson:"extension" json:"extension"`
		Areacode     string      `bson:"areacode" json:"areacode"`
		Countrycode  string      `bson:"countrycode" json:"countrycode"`
		Foreign      string      `bson:"foreign" json:"foreign"`
		ForeignPhone interface{} `bson:"foreignPhone" json:"foreignPhone"`
		ID           bson.ObjectId      `bson:"_id" json:"id"`
		UpdateDate   time.Time   `bson:"updateDate" json:"updateDate"`
		UpdateBy     interface{} `bson:"updateBy" json:"updateBy"`
		IsVoid       bool        `bson:"isVoid" json:"isVoid"`
	} `bson:"phones" json:"phones"`
	AddressID         string      `bson:"addressId" json:"addressId"`
	AddressType       string      `bson:"addressType" json:"addressType"`
	AddressLine1      string      `bson:"addressLine1" json:"addressLine1"`
	AddressLine2      interface{} `bson:"addressLine2" json:"addressLine2"`
	AddressLine3      interface{} `bson:"addressLine3" json:"addressLine3"`
	City              string      `bson:"city" json:"city"`
	StateCode         string      `bson:"stateCode" json:"stateCode"`
	State             string      `bson:"state" json:"state"`
	ZipCode           string      `bson:"zipCode" json:"zipCode"`
	ZipPlus4          interface{} `bson:"zipPlus4" json:"zipPlus4"`
	County            string      `bson:"county" json:"county"`
	PostalCode        interface{} `bson:"postalCode" json:"postalCode"`
	CountryID         string      `bson:"countryId" json:"countryId"`
	Country           string      `bson:"country" json:"country"`
	CountryCode       interface{} `bson:"countryCode" json:"countryCode"`
	StartDate         time.Time   `bson:"startDate" json:"startDate"`
	EndDate           time.Time   `bson:"endDate" json:"endDate"`
	IsStandardized    bool        `bson:"isStandardized" json:"isStandardized"`
	Suite             interface{} `bson:"suite" json:"suite"`
	Source            interface{} `bson:"source" json:"source"`
	SourceID          interface{} `bson:"sourceID" json:"sourceID"`
	Latitude          string      `bson:"latitude" json:"latitude"`
	Longitude         string      `bson:"longitude" json:"longitude"`
	BeforeAddrStdHash interface{} `bson:"beforeAddrStdHash" json:"beforeAddrStdHash"`
	GeoJSONPoint      struct {
		Type        string    `bson:"type" json:"type"`
		Coordinates []float64 `bson:"coordinates" json:"coordinates"`
	} `bson:"geoJSONPoint" json:"geoJSONPoint"`
	AfterAddrStdHash interface{} `bson:"afterAddrStdHash" json:"afterAddrStdHash"`
	Primary          bool        `bson:"primary" json:"primary"`
	PrintSupress     bool        `bson:"printSupress" json:"printSupress"`
	Train            bool        `bson:"train" json:"train"`
	Bus              bool        `bson:"bus" json:"bus"`
	TransitRoute     interface{} `bson:"transitRoute" json:"transitRoute"`
	Handicap         bool        `bson:"handicap" json:"handicap"`
	OfficeManager    interface{} `bson:"officeManager" json:"officeManager"`
	Hours            []struct {
		Day               string      `bson:"day" json:"day"`
		Is24Hours         bool        `bson:"is24Hours" json:"is24Hours"`
		Open              string      `bson:"open" json:"open"`
		Close             string      `bson:"close" json:"close"`
		ByAppointmentOnly bool        `bson:byAppointmentOnly"" json:"byAppointmentOnly"`
		UpdateDate        time.Time   `bson:"updateDate" json:"updateDate"`
		UpdateBy          interface{} `bson:"updateBy" json:"updateBy"`
		IsVoid            bool        `bson:"isVoid" json:"isVoid"`
	} `bson:"hours" json:"hours"`
	Addresses []struct {
		AddressID         string      `bson:"addressId" json:"addressId"`
		AddressType       string      `bson:"addressType" json:"addressType"`
		AddressLine1      string      `bson:"addressLine1" json:"addressLine1"`
		AddressLine2      interface{} `bson:"addressLine2" json:"addressLine2"`
		AddressLine3      interface{} `bson:"addressLine3" json:"addressLine3"`
		City              string      `bson:"city" json:"city"`
		StateCode         string      `bson:"stateCode" json:"stateCode"`
		State             string      `bson:"state" json:"state"`
		ZipCode           string      `bson:"zipCode" json:"zipCode"`
		ZipPlus4          interface{} `bson:"zipPlus4" json:"zipPlus4"`
		County            string      `bson:"county" json:"county"`
		PostalCode        interface{} `bson:"postalCode" json:"postalCode"`
		CountryID         string      `bson:"countryId" json:"countryId"`
		Country           string      `bson:"country" json:"country"`
		CountryCode       interface{} `bson:"countryCode" json:"countryCode"`
		StartDate         time.Time   `bson:"startDate" json:"startDate"`
		EndDate           time.Time   `bson:"endDate" json:"endDate"`
		IsStandardized    bool        `bson:"isStandardized" json:"isStandardized"`
		Suite             interface{} `bson:"suite" json:"suite"`
		Source            interface{} `bson:"source" json:"source"`
		SourceID          interface{} `bson:"sourceID" json:"sourceID"`
		Latitude          string      `bson:"latitude" json:"latitude"`
		Longitude         string      `bson:"longitude" json:"longitude"`
		BeforeAddrStdHash interface{} `bson:"beforeAddrStdHash" json:"beforeAddrStdHash"`
		GeoJSONPoint      struct {
			Type        string    `bson:"type" json:"type"`
			Coordinates []float64 `bson:"coordinates" json:"coordinates"`
		} `bson:"geoJSONPoint" json:"geoJSONPoint"`
		AfterAddrStdHash interface{} `bson:"afterAddrStdHash" json:"afterAddrStdHash"`
		UpdateDate       time.Time   `bson:"updateDate" json:"updateDate"`
		UpdateBy         interface{} `bson:"updateBy" json:"updateBy"`
		IsVoid           bool        `bson:"isVoid" json:"isVoid"`
	} `bson:"addresses" json:"addresses"`
	Attributes    []interface{} `bson:"attributes" json:"attributes"`
	NetworkID     string        `bson:"networkId" json:"networkId"`
	NetworkName   string        `bson:"networkName" json:"networkName"`
	BusinessUnit  string        `bson:"businessUnit" json:"businessUnit"`
	Program       string        `bson:"program" json:"program"`
	HealthPlan    string        `bson:"healthPlan" json:"healthPlan"`
	Company       string        `bson:"company" json:"company"`
	NetworkStatus interface{}   `bson:"networkStatus" json:"networkStatus"`
	PracLoc       []interface{} `bson:"pracLoc" json:"pracLoc"`
	InsertDate    time.Time     `bson:"insertDate" json:"insertDate"`
	UpdateDate    time.Time     `bson:"updateDate" json:"updateDate"`
	UpdateBy      time.Time     `bson:"updateBy" json:"updateBy"`
}