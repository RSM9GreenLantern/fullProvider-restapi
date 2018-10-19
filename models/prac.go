package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Prac struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	PracID      string        `bson:"pracId" json:"pracId"`
	FirstName   string        `bson:"firstName" json:"firstName"`
	LastName    string        `bson:"lastName" json:"lastName"`
	MiddleName  string        `bson:"middleName" json:"middleName"`
	Gender      string        `bson:"gender" json:"gender"`
	Email       interface{}   `bson:"email" json:"email"`
	Salutation  interface{}   `bson:"salutation" json:"salutation"`
	Dob         time.Time     `bson:"dob" json:"dob"`
	Hatcode     string        `bson:"hatcode" json:"hatcode"`
	Ethnicity   interface{}   `bson:"ethnicity" json:"ethnicity"`
	Minority    interface{}   `bson:"minority" json:"minority"`
	Degree      string        `bson:"degree" json:"degree"`
	Taxonomies  []string      `bson:"taxonomies" json:"taxonomies"`
	Identifiers []struct {
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
	MedLicenses []struct {
		LicenseNumber  string    `bson:"licenseNumber" json:"licenseNumber"`
		State          string    `bson:"state" json:"state"`
		ExpirationDate time.Time `bson:"expirationDate" json:"expirationDate"`
	} `bson:"medLicenses" json:"medLicenses"`
	DEA         []interface{} `bson:"DEA" json:"DEA"`
	Credentials []interface{} `bson:"credentials" json:"credentials"`
	Specialty   []struct {
		SpecialtyID       string      `bson:"specialtyID" json:"specialtyID"`
		PdmName           string      `bson:"pdmName" json:"pdmName"`
		Primary           bool        `bson:"primary" json:"primary"`
		BoardStatus       interface{} `bson:"boardStatus" json:"boardStatus"`
		BoardCertIssued   interface{} `bson:"boardCertIssued" json:"boardCertIssued"`
		TakingBoards      interface{} `bson:"takingBoards" json:"takingBoards"`
		BoardCertExp      interface{} `bson:"boardCertExp" json:"boardCertExp"`
		Be5Ends           interface{} `bson:"be5ends" json:"be5ends"`
		IssuerName        interface{} `bson:"issuerName" json:"issuerName"`
		CertNumber        interface{} `bson:"certNumber" json:"certNumber"`
		UpdateDate        time.Time   `bson:"updateDate" json:"updateDate"`
		UpdateBy          interface{} `bson:"updateBy" json:"updateBy"`
		IsVoid            bool        `bson:"isVoid" json:"isVoid"`
		PracSpecialtyCert interface{} `bson:"pracSpecialtyCert" json:"pracSpecialtyCert"`
	} `bson:"specialty" json:"specialty"`
	Languages       []interface{} `bson:"languages" json:"languages"`
	StartDate       time.Time     `bson:"startDate" json:"startDate"`
	EndDate         time.Time     `bson:"endDate" json:"endDate"`
	Hours           interface{}   `bson:"hours" json:"hours"`
	Insurance       []interface{} `bson:"insurance" json:"insurance"`
	Training        []interface{} `bson:"training" json:"training"`
	PracHosp        []interface{} `bson:"pracHosp" json:"pracHosp"`
	UpdateDate      time.Time     `bson:"updateDate" json:"updateDate"`
	UpdateBy        interface{}   `bson:"updateBy" json:"updateBy"`
	IsVoid          bool          `bson:"isVoid" json:"isVoid"`
	PractitionerLoc []struct {
		Network struct {
			Carrier struct {
				ID   string `bson:"_id" json:"id"`
				Code string `bson:"" json:"code"`
				Name string `bson:"" json:"name"`
			} `bson:"" json:"carrier"`
			Product struct {
				ID   string `bson:"_id" json:"id"`
				Code string `bson:"" json:"code"`
				Name string `bson:"" json:"name"`
			} `bson:"" json:"product"`
			HealthPlan struct {
				ID   string `bson:"_id" json:"id"`
				Code string `bson:"" json:"code"`
				Name string `bson:"" json:"name"`
			} `bson:"" json:"healthPlan"`
			BusinessUnit struct {
				ID   string `bson:"_id" json:"id"`
				Code string `bson:"" json:"code"`
				Name string `bson:"" json:"name"`
			} `bson:"" json:"businessUnit"`
			Company struct {
				ID   string `bson:"_id" json:"id"`
				Code string `bson:"" json:"code"`
				Name string `bson:"" json:"name"`
			} `bson:"" json:"company"`
		} `bson:"" json:"network"`
		PracID      string `bson:"" json:"pracId"`
		PnlcID      string `bson:"" json:"pnlcId"`
		LocationID  string `bson:"" json:"locationId"`
		ProvID      string `bson:"" json:"provId"`
		Identifiers []struct {
			Type           string      `bson:"" json:"type"`
			TypeID         string      `bson:"" json:"typeId"`
			IdentifierID   string      `bson:"" json:"identifierId"`
			Identifier     string      `bson:"" json:"identifier"`
			State          interface{} `bson:"" json:"state"`
			StateName      interface{} `bson:"" json:"stateName"`
			StateCode      interface{} `bson:"" json:"stateCode"`
			StartDate      time.Time   `bson:"" json:"startDate"`
			EndDate        time.Time   `bson:"" json:"endDate"`
			ValidationDate interface{} `bson:"" json:"validationDate"`
			UpdateDate     time.Time   `bson:"updateDate" json:"updateDate"`
			UpdateBy       interface{} `bson:"updateBy" json:"updateBy"`
			IsVoid         bool        `bson:"isVoid" json:"isVoid"`
		} `bson:"identifiers" json:"identifiers"`
		Panel struct {
			ID               bson.ObjectId `bson:"_id" json:"id"`
			PracID           string        `bson:"pracId" json:"pracId"`
			ProvID           string        `bson:"provId" json:"provId"`
			LocID            string        `bson:"locId" json:"locId"`
			Location         interface{}   `bson:"location" json:"location"`
			NetID            string        `bson:"netId" json:"netId"`
			Network          interface{}   `bson:"network" json:"network"`
			StartDate        time.Time     `bson:"startDate" json:"startDate"`
			EndDate          time.Time     `bson:"endDate" json:"endDate"`
			PanelCode        string        `bson:"panelCode" json:"panelCode"`
			PanelStatus      string        `bson:"panelStatus" json:"panelStatus"`
			AgeRestriction   interface{}   `bson:"ageRestriction" json:"ageRestriction"`
			HigherUnits      interface{}   `bson:"higherUnits" json:"higherUnits"`
			LowerUnits       interface{}   `bson:"lowerUnits" json:"lowerUnits"`
			HighestValidAge  interface{}   `bson:"highestValidAge" json:"highestValidAge"`
			LowestValidAge   interface{}   `bson:"lowestValidAge" json:"lowestValidAge"`
			PanelLimit       interface{}   `bson:"panelLimit" json:"panelLimit"`
			GenderLimitation interface{}   `bson:"genderLimitation" json:"genderLimitation"`
			UpdateDate       time.Time     `bson:"updateDate" json:"updateDate"`
			UpdateBy         interface{}   `bson:"updateBy" json:"updateBy"`
		} `bson:"panel" json:"panel"`
		AcredCert         []interface{} `bson:"acredCert" json:"acredCert"`
		Insurance         interface{}   `bson:"insurance" json:"insurance"`
		PracHosp          interface{}   `bson:"pracHosp" json:"pracHosp"`
		NetworkID         string        `bson:"networkId" json:"networkId"`
		NetworkName       string        `bson:"networkName" json:"networkName"`
		BusinessUnit      string        `bson:"businessUnit" json:"businessUnit"`
		Program           string        `bson:"program" json:"program"`
		HealthPlan        string        `bson:"healthPlan" json:"healthPlan"`
		Company           string        `bson:"company" json:"company"`
		NetworkStatus     interface{}   `bson:"networkStatus" json:"networkStatus"`
		BusinessEntityID  string        `bson:"businessEntityId" json:"businessEntityId"`
		BusinessEntity    interface{}   `bson:"businessEntity" json:"businessEntity"`
		Status            string        `bson:"status" json:"status"`
		TinID             string        `bson:"tinId" json:"tinId"`
		Tin               string        `bson:"tin" json:"tin"`
		ProviderName      string        `bson:"providerName" json:"providerName"`
		TypeID            string        `bson:"typeId" json:"typeId"`
		Type              string        `bson:"type" json:"type"`
		ProvhatCode       []string      `bson:"provhatCode" json:"provhatCode"`
		Attributes        []interface{} `bson:"attributes" json:"attributes"`
		AddressID         string        `bson:"addressId" json:"addressId"`
		AddressType       string        `bson:"addressType" json:"addressType"`
		AddressLine1      string        `bson:"addressLine1" json:"addressLine1"`
		AddressLine2      interface{}   `bson:"addressLine2" json:"addressLine2"`
		AddressLine3      interface{}   `bson:"addressLine3" json:"addressLine3"`
		City              string        `bson:"city" json:"city"`
		StateCode         string        `bson:"stateCode" json:"stateCode"`
		State             string        `bson:"state" json:"state"`
		ZipCode           string        `bson:"zipCode" json:"zipCode"`
		ZipPlus4          string        `bson:"zipPlus4" json:"zipPlus4"`
		County            string        `bson:"county" json:"county"`
		PostalCode        interface{}   `bson:"postalCode" json:"postalCode"`
		CountryID         string        `bson:"countryId" json:"countryId"`
		Country           string        `bson:"country" json:"country"`
		CountryCode       interface{}   `bson:"countryCode" json:"countryCode"`
		StartDate         time.Time     `bson:"startDate" json:"startDate"`
		EndDate           time.Time     `bson:"endDate" json:"endDate"`
		IsStandardized    bool          `bson:"isStandardized" json:"isStandardized"`
		Suite             interface{}   `bson:"suite" json:"suite"`
		Source            interface{}   `bson:"source" json:"source"`
		SourceID          interface{}   `bson:"sourceID" json:"sourceID"`
		Latitude          string        `bson:"latitude" json:"latitude"`
		Longitude         string        `bson:"longitude" json:"longitude"`
		BeforeAddrStdHash interface{}   `bson:"beforeAddrStdHash" json:"beforeAddrStdHash"`
		GeoJSONPoint      struct {
			Type        string    `bson:"type" json:"type"`
			Coordinates []float64 `bson:"coordinates" json:"coordinates"`
		} `bson:"geoJSONPoint" json:"geoJSONPoint"`
		AfterAddrStdHash interface{}   `bson:"afterAddrStdHash" json:"afterAddrStdHash"`
		Primary          bool          `bson:"primary" json:"primary"`
		PrintSupress     bool          `bson:"printSupress" json:"printSupress"`
		Train            bool          `bson:"train" json:"train"`
		Bus              bool          `bson:"bus" json:"bus"`
		TransitRoute     interface{}   `bson:"transitRoute" json:"transitRoute"`
		Handicap         bool          `bson:"handicap" json:"handicap"`
		OfficeManager    interface{}   `bson:"officeManager" json:"officeManager"`
		Hours            []interface{} `bson:"hours" json:"hours"`
		Phones           []struct {
			Type         string        `bson:"type" json:"type"`
			Number       string        `bson:"number" json:"number"`
			Extension    interface{}   `bson:"extension" json:"extension"`
			Areacode     string        `bson:"areacode" json:"areacode"`
			Countrycode  string        `bson:"countrycode" json:"countrycode"`
			Foreign      string        `bson:"foreign" json:"foreign"`
			ForeignPhone interface{}   `bson:"foreignPhone" json:"foreignPhone"`
			ID           bson.ObjectId `bson:"_id" json:"id"`
			UpdateDate   time.Time     `bson:"updateDate" json:"updateDate"`
			UpdateBy     interface{}   `bson:"updateBy" json:"updateBy"`
			IsVoid       bool          `bson:"isVoid" json:"isVoid"`
		} `bson:"phones" json:"phones"`
		UpdateDate   time.Time   `bson:"updateDate" json:"updateDate"`
		UpdateBy     interface{} `bson:"updateBy" json:"updateBy"`
		IsVoid       bool        `bson:"isVoid" json:"isVoid"`
		NetworkCode  string      `bson:"networkCode" json:"networkCode"`
		ProgramCode  string      `bson:"programCode" json:"programCode"`
		Affiliations []struct {
			PnlcID   string `bson:"pnlcId" json:"pnlcId"`
			Taxonomy struct {
				Code string `bson:"code" json:"code"`
			} `bson:"taxonomy" json:"taxonomy"`
			ProviderID          string    `bson:"providerId" json:"providerId"`
			AffiliationNumber   string    `bson:"affiliationNumber" json:"affiliationNumber"`
			BusinessUnit        string    `bson:"businessUnit" json:"businessUnit"`
			Program             string    `bson:"program" json:"program"`
			Carrier             string    `bson:"carrier" json:"carrier"`
			StartDate           time.Time `bson:"startDate" json:"startDate"`
			EndDate             time.Time `bson:"endDate" json:"endDate"`
			ProviderType        string    `bson:"providerType" json:"providerType"`
			HatCode             string    `bson:"hatCode" json:"hatCode"`
			Status              string    `bson:"status" json:"status"`
			Npi                 string    `bson:"npi" json:"npi"`
			AlternateQualifier  string    `bson:"alternateQualifier" json:"alternateQualifier"`
			ClaimType           string    `bson:"claimType" json:"claimType"`
			PayClass            string    `bson:"payClass" json:"payClass"`
			PayTo               string    `bson:"payTo" json:"payTo"`
			Spec1               string    `bson:"spec1" json:"spec1"`
			IrsNumber           string    `bson:"irsNumber" json:"irsNumber"`
			GroupPracticeNumber string    `bson:"groupPracticeNumber" json:"groupPracticeNumber"`
		} `bson:"affiliations" json:"affiliations"`
	} `bson:"practitionerLoc" json:"practitionerLoc"`
}
