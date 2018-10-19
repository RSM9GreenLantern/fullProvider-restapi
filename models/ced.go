package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type CED struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Network struct {
		Carrier struct {
			Code string `bson:"code" json:"code"`
			Name string `bson:"name" json:"name"`
		} `bson:"carrier" json:"carrier"`
		Product struct {
			Code string `bson:"code" json:"code"`
			Name string `bson:"name" json:"name"`
		} `bson:"product" json:"product"`
		HealthPlan struct {
			Code string `bson:"code" json:"code"`
			Name string `bson:"name" json:"name"`
		} `bson:"healthPlan" json:"healthPlan"`
		BusinessUnit struct {
			Code string `bson:"code" json:"code"`
			Name string `bson:"name" json:"name"`
		} `bson:"businessUnit" json:"businessUnit"`
		Company struct {
			Code string `bson:"code" json:"code"`
			Name string `bson:"name" json:"name"`
		} `bson:"company" json:"company"`
		NetworkCycles []struct {
			Status    string    `bson:"" json:"status"`
			StartDate time.Time `bson:"" json:"startDate"`
		} `bson:"networkCycles" json:"networkCycles"`
	} `bson:"network" json:"network"`
	EntityID                    string `bson:"entityId" json:"entityId"`
	EntityType                  string `bson:"entityType" json:"entityType"`
	NationalProviderIdentifiers []struct {
		Number        string    `bson:"number" json:"number"`
		Type          string    `bson:"type" json:"type"`
		EffectiveDate time.Time `bson:"effectiveDate" json:"effectiveDate"`
	} `bson:"nationalProviderIdentifiers" json:"nationalProviderIdentifiers"`
	UpdateTime time.Time `bson:"updateTime" json:"updateTime"`
	HatCode    string    `bson:"hatCode" json:"hatCode"`
	Locations  []struct {
		LocationID string `bson:"locationId" json:"locationId"`
		Address    struct {
			AddressLine1 string `bson:"addressLine1" json:"addressLine1"`
			AddressLine2 string `bson:"addressLine2" json:"addressLine2"`
			City         string `bson:"city" json:"city"`
			StateCode    string `bson:"stateCode" json:"stateCode"`
			State        string `bson:"state" json:"state"`
			ZipCode      string `bson:"zipCode" json:"zipCode"`
			ZipPlus4     string `bson:"zipPlus4" json:"zipPlus4"`
			County       string `bson:"county" json:"county"`
			Country      string `bson:"country" json:"country"`
			Latitude     string `bson:"latitude" json:"latitude"`
			Longitude    string `bson:"longitude" json:"longitude"`
			GeoCoords    struct {
				Type        string    `bson:"type" json:"type"`
				Coordinates []float64 `bson:"coordinates" json:"coordinates"`
			} `bson:"geoCoords" json:"geoCoords"`
		} `bson:"address" json:"address"`
		Name     string `bson:"name" json:"name"`
		Handicap string `bson:"handicap" json:"handicap"`
		Phone    struct {
			Phone string `bson:"phone" json:"phone"`
			Fax   string `bson:"fax" json:"fax"`
		} `bson:"phone" json:"phone"`
		OfficeHours struct {
			Monday struct {
				Open  string `bson:"open" json:"open"`
				Close string `bson:"close" json:"close"`
			} `bson:"monday" json:"monday"`
			Tuesday struct {
				Open  string `bson:"open" json:"open"`
				Close string `bson:"close" json:"close"`
			} `bson:"tuesday" json:"tuesday"`
			Wednesday struct {
				Open  string `bson:"open" json:"open"`
				Close string `bson:"close" json:"close"`
			} `bson:"wednesday" json:"wednesday"`
			Thursday struct {
				Open  string `bson:"open" json:"open"`
				Close string `bson:"close" json:"close"`
			} `bson:"thursday" json:"thursday"`
			Friday struct {
				Open  string `bson:"open" json:"open"`
				Close string `bson:"close" json:"close"`
			} `bson:"friday" json:"friday"`
		} `bson:"officeHours" json:"officeHours"`
		Primary       string `bson:"primary" json:"primary"`
		NetworkCycles []struct {
			NetworkCycleIDRef string    `bson:"networkCycleIdRef" json:"networkCycleIdRef"`
			StartDate         time.Time `bson:"startDate" json:"startDate"`
		} `bson:"networkCycles" json:"networkCycles"`
		PublicTransportation struct {
			Train        string `bson:"train" json:"train"`
			Bus          string `bson:"bus" json:"bus"`
			TransitRoute string `bson:"transitRoute" json:"transitRoute"`
		} `bson:"publicTransportation,omitempty" json:"publicTransportation,omitempty"`
	} `bson:"" json:"locations"`
	CarrierID    string `bson:"carrierId" json:"carrierId"`
	ProviderName string `bson:"providerName" json:"providerName"`
	ProviderType string `bson:"providerType" json:"providerType"`
	TaxIDNumber  string `bson:"taxIdNumber" json:"taxIdNumber"`
}