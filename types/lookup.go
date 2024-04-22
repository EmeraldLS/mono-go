package types

type BVNLookupRequest struct {
	BVN   string `json:"bvn"`
	Scope string `json:"scope"`
}

type BVNLookupResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    BVNLookupData `json:"data"`
}

type BVNLookupData struct {
	SessionID string            `json:"session_id"`
	Methods   []BVNLookupMethod `json:"method"`
}

type BVNLookupMethod struct {
	Method string `json:"method"`
	Hint   string `json:"hint"`
}

type BVNVerifyOTPRequest struct {
	Method      string `json:"method"`
	PhoneNumber string `json:"phone_number"`
}

type BVNVerifyOTPResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type BVNDetailsRequest struct {
	OTP string `json:"otp"`
}

type BVNDetailsResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    BVNDetailsData `json:"data"`
}

type BVNDetailsData struct {
	FirstName        string `json:"firstName"`
	MiddleName       string `json:"middleName"`
	LastName         string `json:"lastName"`
	DOB              string `json:"dob"`
	PhoneNumber1     string `json:"phoneNumber1"`
	PhoneNumber2     string `json:"phoneNumber2"`
	RegistrationDate string `json:"registrationDate"`
	Email            string `json:"email"`
	Gender           string `json:"gender"`
	LGAOfOrigin      string `json:"lgaOforigin"`
	LGAOfResidence   string `json:"lgaOfResidence"`
	MaritalStatus    string `json:"maritalStatus"`
	NIN              string `json:"nin"`
	StateOfOrigin    string `json:"stateOfOrigin"`
	StateOfResidence string `json:"stateOfResidence"`
	WatchListed      bool   `json:"watchListed"`
	PhotoID          string `json:"photoId"`
}

type BVNBankAccountsResponse struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Data    []BVNBankAccountsData `json:"data"`
}

type BVNBankAccountsData struct {
	AccountName        string                         `json:"account_name"`
	AccountNumber      string                         `json:"account_number"`
	AccountStatus      string                         `json:"account_status"`
	AccountType        string                         `json:"account_type"`
	AccountDesignation string                         `json:"account_designation"`
	Institution        BVNBankAccountsDataInstitution `json:"institution"`
}

type BVNBankAccountsDataInstitution struct {
	Name     string `json:"name"`
	Branch   string `json:"branch"`
	BankCode string `json:"bank_code"`
}

type LookupBusinessResponse struct {
	Status    string               `json:"status"`
	Message   string               `json:"message"`
	Timestamp string               `json:"timestamp"`
	Data      []LookupBusinessData `json:"data"`
}

type LookupBusinessData struct {
	State                    string `json:"state"`
	ID                       int    `json:"id"`
	ApprovedName             string `json:"approved_name"`
	HeadOfficeAddress        string `json:"head_office_address"`
	BranchAddress            string `json:"branch_address"`
	RCNumber                 string `json:"rc_number"`
	NatureOfBusinessName     string `json:"nature_of_business_name"`
	ClassificationID         string `json:"classification_id"`
	BusinessCommencementDate string `json:"business_commencement_date"`
	Classification           string `json:"classification"`
	City                     string `json:"city"`
	LGA                      string `json:"lga"`
	Active                   bool   `json:"active"`
	Email                    string `json:"email"`
	Address                  string `json:"address"`
	RegistrationDate         string `json:"registration_date"`
	RegistrationApproved     bool   `json:"registration_approved"`
	Objectives               string `json:"objectives"`
	CompanyTypeName          string `json:"company_type_name"`
	DelistingStatus          string `json:"delisting_status"`
}

type ShareholderDetailsResponse struct {
	Status    string                   `json:"status"`
	Message   string                   `json:"message"`
	Timestamp string                   `json:"timestamp"`
	Data      []ShareholderDetailsData `json:"data"`
}

type ShareholderDetailsData struct {
	ID                             int             `json:"id"`
	Surname                        string          `json:"surname"`
	FirstName                      string          `json:"firstname"`
	OtherName                      string          `json:"other_name"`
	Email                          string          `json:"email"`
	Gender                         string          `json:"gneder"`
	FormerNationality              string          `json:"former_nationality"`
	Age                            int             `json:"age"`
	City                           string          `json:"city"`
	Occupation                     string          `json:"occupation"`
	RCNumber                       string          `json:"rc_number"`
	CorporationCompany             string          `json:"corporation_company"`
	State                          string          `json:"state"`
	PoBox                          string          `json:"pobox"`
	AccreditationNumber            string          `json:"accreditationnumber"`
	IsLawyer                       bool            `json:"is_lawyer"`
	IsChairman                     bool            `json:"is_chairman"`
	NumberOfSharesAlloted          int             `json:"num_shares_alloted"`
	TypeOfShares                   string          `json:"type_of_shares"`
	DOB                            string          `json:"date_of_birth"`
	Status                         string          `json:"status"`
	DateOfTermination              string          `json:"date_of_termination"`
	DateOfAppointment              string          `json:"date_of_appointment"`
	DateOfAddressChange            string          `json:"date_of_change_of_address"`
	FormerAddress                  string          `json:"former_address"`
	FormerPostal                   string          `json:"former_postal"`
	FormerFirstName                string          `json:"former_first_name"`
	FormerOtherName                string          `json:"former_other_name"`
	DateOfStatusChange             string          `json:"date_of_status_change"`
	IdentityNumber                 string          `json:"identity_number"`
	IdentityIssueState             string          `json:"identity_issue_state"`
	OtherDirectorshipDetails       string          `json:"other_directorship_details"`
	PortalUserFk                   string          `json:"portal_user_fk"`
	AffiliatesFk                   string          `json:"affiliates_fk"`
	ProcessTypeFk                  ProcessTypeFK   `json:"process_type_fk"`
	Company                        string          `json:"company"`
	SamePersonAsFk                 string          `json:"same_person_as_fk"`
	NatureOfAppOrDischarge         string          `json:"nature_of_app_or_discharge"`
	IsDesginated                   bool            `json:"is_designated"`
	EndOfAppoinment                string          `json:"end_of_appointment"`
	AppointedBy                    string          `json:"appointed_by"`
	DateOfDeedOfDischarge          string          `json:"date_of_deed_of_discharge"`
	CountryFk                      CountryFk       `json:"country_fk"`
	CountryOfResidence             string          `json:"country_of_residence"`
	IsCarriedOverFromNameAvailable string          `json:"is_carried_over_from_name_avai"`
	LGA                            string          `json:"lga"`
	CorporationRegistrationDate    string          `json:"corporation_registration_date"`
	IsCompanyDeleted               bool            `json:"is_company_deleted"`
	GovernmentOrganisationName     string          `json:"government_organisation_name"`
	ForeignOrganisationName        string          `json:"foreign_organisation_name"`
	CompanyStreetAddress           string          `json:"company_street_address"`
	CompanyState                   string          `json:"company_state"`
	CompanyCity                    string          `json:"company_city"`
	IsCorporate                    string          `json:"is_corporate"`
	CountryOfIncorporationFk       string          `json:"country_of_incorporation_fk"`
	Nationality                    string          `json:"nationality"`
	Address                        string          `json:"address"`
	Postcode                       string          `json:"postcode"`
	StreetNumber                   string          `json:"street_number"`
	AffiliatesResidentialAddress   string          `json:"affiliates_residential_address"`
	AffiliatesPSCInformation       int             `json:"affiliates_psc_information"`
	LegalOwnersOfInterests         []string        `json:"legal_owners_of_interests"`
	LegalOwnersOfVotingRights      []string        `json:"legal_owners_of_voting_rights"`
	StockExchangeSoes              []string        `json:"stock_exchange_soes"`
	ApprovedFornoticeOfPSC         string          `json:"approved_for_notice_of_psc"`
	CompanyAddress2                string          `json:"company_address2"`
	FullAddress2                   string          `json:"full_address2"`
	AffiliateTypeFk                AffiliateTypeFk `json:"affiliate_type_fk"`
}

type ProcessTypeFK struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Type        string `json:"type"`
	ProductID   string `json:"product_id"`
	BankCode    string `json:"bank_code"`
}

type CountryFk struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type AffiliateTypeFk struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PreviousBusinessAddressResponse struct {
	Status    string                      `json:"status"`
	Message   string                      `json:"message"`
	Timestamp string                      `json:"timestamp"`
	Data      PreviousBusinessAddressData `json:"data"`
}

type PreviousBusinessAddressData struct {
	ApprovedName    string `json:"approved_name"`
	PreviousAddress string `json:"previous_address"`
	StreetName      string `json:"street_name"`
	City            string `json:"city"`
	SubmissionDate  string `json:"submission_date"`
	ApprovalDate    string `json:"approval_date"`
	ID              int    `json:"id"`
	State           string `json:"state"`
}

type ChangeOfNameResponse struct {
	Status    string             `json:"status"`
	Message   string             `json:"message"`
	Timestamp string             `json:"timestamp"`
	Data      []ChangeOfNameData `json:"data"`
}

type ChangeOfNameData struct {
	PersistMasterID int    `json:"persist_master_id"`
	NewName         string `json:"new_name"`
	FormerName      string `json:"former_name"`
	ApprovalDate    string `json:"approval_date"`
}

type SecretResponse struct {
	Status    string          `json:"status"`
	Message   string          `json:"message"`
	Timestamp string          `json:"timestamp"`
	Data      []SecretaryData `json:"data"`
}

type SecretaryData struct {
	ID                           int64           `json:"id"`
	Surname                      string          `json:"surname"`
	Firstname                    string          `json:"firstname"`
	OtherName                    string          `json:"other_name"`
	Email                        string          `json:"email"`
	PhoneNumber                  string          `json:"phone_number"`
	Gender                       string          `json:"gender"`
	FormerNationality            string          `json:"former_nationality"`
	Age                          int64           `json:"age"`
	City                         string          `json:"city"`
	Occupation                   string          `json:"occupation"`
	FormerName                   string          `json:"former_name"`
	CorporationName              string          `json:"corporation_name"`
	RCNumber                     string          `json:"rc_number"`
	CorporationCompany           string          `json:"corporation_company"`
	State                        string          `json:"state"`
	Pobox                        string          `json:"pobox"`
	Accreditationnumber          string          `json:"accreditationnumber"`
	IsLawyer                     string          `json:"is_lawyer"`
	LastVisit                    int64           `json:"last_visit"`
	FormType                     string          `json:"form_type"`
	IsPresenter                  string          `json:"is_presenter"`
	IsChairman                   bool            `json:"is_chairman"`
	NumSharesAlloted             string          `json:"num_shares_alloted"`
	TypeOfShares                 string          `json:"type_of_shares"`
	DateOfBirth                  string          `json:"date_of_birth"`
	Status                       string          `json:"status"`
	DateOfTermination            string          `json:"date_of_termination"`
	DateOfAppointment            string          `json:"date_of_appointment"`
	DateOfChangeOfAddress        string          `json:"date_of_change_of_address"`
	FormerAddress                string          `json:"former_address"`
	FormerPostal                 string          `json:"former_postal"`
	FormerSurname                string          `json:"former_surname"`
	FormerFirstName              string          `json:"former_first_name"`
	FormerOtherName              string          `json:"former_other_name"`
	DateOfStatusChange           string          `json:"date_of_status_change"`
	IdentityNumber               string          `json:"identity_number"`
	IdentityIssueState           string          `json:"identity_issue_state"`
	OtherDirectorshipDetails     string          `json:"other_directorship_details"`
	PortalUserFk                 string          `json:"portal_user_fk"`
	AffiliatesFk                 string          `json:"affiliates_fk"`
	ProcessTypeFk                ProcessTypeFk   `json:"process_type_fk"`
	Company                      string          `json:"company"`
	SamePersonAsFk               string          `json:"same_person_as_fk"`
	NatureOfAppOrDischarge       string          `json:"nature_of_app_or_discharge"`
	IsDesignated                 string          `json:"is_designated"`
	EndOfAppointment             string          `json:"end_of_appointment"`
	AppointedBy                  string          `json:"appointed_by"`
	DateOfDeedOfDischarge        string          `json:"date_of_deed_of_discharge"`
	DateOfResolution             string          `json:"date_of_resolution"`
	CountryFk                    CountryFk       `json:"country_fk"`
	CountryOfResidence           string          `json:"country_of_residence"`
	IsCarriedOverFromNameAvai    string          `json:"is_carried_over_from_name_avai"`
	LGA                          string          `json:"lga"`
	CorporationRegistrationDate  string          `json:"corporation_registration_date"`
	IsCompanyDeleted             string          `json:"is_company_deleted"`
	GovernmentOrganisationName   string          `json:"government_organisation_name"`
	ForeignOrganisationName      string          `json:"foreign_organisation_name"`
	CompanyStreetAddress         string          `json:"company_street_address"`
	CompanyState                 string          `json:"company_state"`
	CompanyCity                  string          `json:"company_city"`
	IsCorporate                  string          `json:"is_corporate"`
	CountyOfIncorporationFk      string          `json:"county_of_incorporation_fk"`
	Nationality                  string          `json:"nationality"`
	Address                      string          `json:"address"`
	Postcode                     string          `json:"postcode"`
	StreetNumber                 string          `json:"street_number"`
	AffiliatesResidentialAddress int64           `json:"affiliates_residential_address"`
	AffiliatesPscInformation     int64           `json:"affiliates_psc_information"`
	LegalOwnersOfInterests       []string        `json:"legal_owners_of_interests"`
	LegalOwnersOfVotingRights    []string        `json:"legal_owners_of_voting_rights"`
	StockExchangeSoes            []string        `json:"stock_exchange_soes"`
	ApprovedForNoticeOfPsc       string          `json:"approved_for_notice_of_psc"`
	CompanyAddress2              string          `json:"company_address2"`
	FullAddress2                 string          `json:"full_address2"`
	AffiliateTypeFk              AffiliateTypeFk `json:"affiliate_type_fk"`
}

type ProcessTypeFk struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
	Type        string `json:"type"`
	ProductID   string `json:"product_id"`
	BankCode    string `json:"bank_code"`
}

type BanksListingResponse struct {
	Status    string           `json:"status"`
	Message   string           `json:"message"`
	Timestamp string           `json:"timestamp"`
	Data      BanksListingData `json:"data"`
}

type BanksListingData struct {
	Banks []BanksListingBanks `json:"bank"`
}

type BanksListingBanks struct {
	Name     string `json:"name"`
	BankCode string `json:"bank_code"`
	NipCode  string `json:"nip_code"`
}

type HouseAddressVerificationRequest struct {
	MeterNumber string `json:"meter_number"`
	Address     string `json:"address"`
}

type HouseAddressVerificationResponse struct {
	Status    string                       `json:"status"`
	Message   string                       `json:"message"`
	Timestamp string                       `json:"timestamp"`
	Data      HouseAddressVerificationData `json:"data"`
}

type HouseAddressVerificationData struct {
	Verified        bool   `json:"verified"`
	HouseAddress    string `json:"house_address"`
	HouseOwner      string `json:"house_owner"`
	ConfidenceLevel string `json:"confidence_level"`
	DiscoCode       string `json:"disco_code"`
}

type InternationalPassportRequest struct {
	PassportNumber string `json:"passport_number"`
	LastName       string `json:"last_name"`
}

type InternationalPassportResponse struct {
	Status    string                    `json:"status"`
	Message   string                    `json:"message"`
	Timestamp string                    `json:"timestamp"`
	Data      InternationalPassportData `json:"data"`
	Number    string                    `json:"number"`
}

type InternationalPassportData struct {
	Number       string `json:"number"`
	ReferenceID  string `json:"reference_id"`
	IssuedDate   string `json:"issued_date"`
	ExpiryDate   string `json:"expiry_date"`
	DocumentType string `json:"document_type"`
	IssuedAt     string `json:"issued_at"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MiddleName   string `json:"middle_name"`
	Dob          string `json:"dob"`
	Gender       string `json:"gender"`
	Photo        string `json:"photo"`
	Signature    string `json:"signature"`
	Mobile       string `json:"mobile"`
}

type TINRequest struct {
	Number  string `json:"number"`
	Channel string `json:"channel"`
}

type TINResponse struct {
	Status    string  `json:"status"`
	Message   string  `json:"message"`
	Timestamp string  `json:"timestamp"`
	Data      TINData `json:"data"`
}

type TINData struct {
	Email              string `json:"email"`
	ID                 string `json:"id"`
	Name               string `json:"name"`
	OrganizationNumber string `json:"organization_number"`
	PhoneNumber        string `json:"phone_number"`
	RegistrationNumber string `json:"registration_number"`
	StreetName         string `json:"street_name"`
	TaxOffice          string `json:"tax_office"`
	TinSource          string `json:"tin_source"`
	TinType            string `json:"tin_type"`
	Search             string `json:"search"`
	CacRegNumber       string `json:"cac_reg_number"`
	TaxpayerName       string `json:"taxpayer_name"`
	Number             string `json:"number"`
}

type NINRequest struct {
	NIN string `json:"nin"`
}

type NINResponse struct {
	Status    string  `json:"status"`
	Message   string  `json:"message"`
	Timestamp string  `json:"timestamp"`
	Data      NINData `json:"data"`
}

type NINData struct {
	Birthcountry     string `json:"birthcountry"`
	Birthdate        string `json:"birthdate"`
	Birthlga         string `json:"birthlga"`
	Birthstate       string `json:"birthstate"`
	Educationallevel string `json:"educationallevel"`
	Email            string `json:"email"`
	Employmentstatus string `json:"employmentstatus"`
	Firstname        string `json:"firstname"`
	Gender           string `json:"gender"`
	Heigth           string `json:"heigth"`
	Maritalstatus    string `json:"maritalstatus"`
	Middlename       string `json:"middlename"`
	Nin              string `json:"nin"`
	NokAddress1      string `json:"nok_address1"`
	NokAddress2      string `json:"nok_address2"`
	NokFirstname     string `json:"nok_firstname"`
	NokLGA           string `json:"nok_lga"`
	NokMiddlename    string `json:"nok_middlename"`
	NokPostalcode    string `json:"nok_postalcode"`
	NokState         string `json:"nok_state"`
	NokSurname       string `json:"nok_surname"`
	NokTown          string `json:"nok_town"`
	Ospokenlang      string `json:"ospokenlang"`
	Pfirstname       string `json:"pfirstname"`
	Photo            string `json:"photo"`
	Pmiddlename      string `json:"pmiddlename"`
	Profession       string `json:"profession"`
	Psurname         string `json:"psurname"`
	Religion         string `json:"religion"`
	ResidenceAddress string `json:"residence_address"`
	ResidenceLGA     string `json:"residence_lga"`
	ResidenceState   string `json:"residence_state"`
	ResidenceTown    string `json:"residence_town"`
	Residencestatus  string `json:"residencestatus"`
	SelfOriginLGA    string `json:"self_origin_lga"`
	SelfOriginPlace  string `json:"self_origin_place"`
	SelfOriginState  string `json:"self_origin_state"`
	Signature        string `json:"signature"`
	SpokenLanguage   string `json:"spoken_language"`
	Surname          string `json:"surname"`
	Telephoneno      string `json:"telephoneno"`
	Title            string `json:"title"`
	Userid           string `json:"userid"`
	Vnin             string `json:"vnin"`
	CentralID        string `json:"central_iD"`
	TrackingID       string `json:"tracking_id"`
}

type DriverLicenseRequest struct {
	LicenseNumber string `json:"license_number"`
	DOB           string `json:"date_of_birth"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
}

type DriverLicenseResponse struct {
	Status    string            `json:"status"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	Data      DriverLicenseData `json:"data"`
}

type DriverLicenseData struct {
	Gender       string `json:"gender"`
	Photo        string `json:"photo"`
	LicenseNo    string `json:"license_no"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MiddleName   string `json:"middle_name"`
	IssuedDate   string `json:"issued_date"`
	ExpiryDate   string `json:"expiry_date"`
	StateOfIssue string `json:"state_ofIssue"`
	BirthDate    string `json:"birth_date"`
}

type AccountNumberRequest struct {
	NipCode       string `json:"nip_code" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
}

type AccountNumberResponse struct {
	Status    string            `json:"status"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	Data      AccountNumberData `json:"data"`
}

type AccountNumberData struct {
	Name          string            `json:"name"`
	AccountNumber string            `json:"account_number"`
	Bvn           string            `json:"bvn"`
	Bank          AccountNumberBank `json:"bank"`
}

type AccountNumberBank struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreditHistoryRequest struct {
	BVN string `json:"bvn"`
}

type CreditHistoryResponse struct {
	Status    string            `json:"status"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	Data      CreditHistoryData `json:"data"`
}

type CreditHistoryData struct {
	Providers     []string             `json:"providers"`
	Profile       CreditHistoryProfile `json:"profile"`
	CreditHistory []CreditHistory      `json:"credit_history"`
}

type CreditHistory struct {
	Institution string                 `json:"institution"`
	History     []CreditHistoryHistory `json:"history"`
}

type CreditHistoryHistory struct {
	DateOpened         string                           `json:"date_opened"`
	OpeningBalance     int64                            `json:"opening_balance"`
	Currency           string                           `json:"currency"`
	PerformanceStatus  string                           `json:"performance_status"`
	Tenor              int64                            `json:"tenor"`
	ClosedDate         string                           `json:"closed_date"`
	LoanStatus         string                           `json:"loan_status"`
	RepaymentFrequency string                           `json:"repayment_frequency"`
	RepaymentAmount    int64                            `json:"repayment_amount"`
	RepaymentSchedule  []CreditHistoryRepaymentSchedule `json:"repayment_schedule"`
}

type CreditHistoryRepaymentSchedule struct {
	Date   string `json:"date"`
	Status string `json:"status"`
}

type CreditHistoryProfile struct {
	FullName        string                        `json:"full_name"`
	Dob             string                        `json:"dob"`
	AddressHistory  []CreditHistoryAddressHistory `json:"address_history"`
	EmailAddresses  []string                      `json:"email_addresses"`
	PhoneNumbers    []string                      `json:"phone_numbers"`
	Gender          string                        `json:"gender"`
	Identifications []Identification              `json:"identifications"`
}

type CreditHistoryAddressHistory struct {
	Address      string `json:"address"`
	Type         string `json:"type"`
	DateReported string `json:"date_reported"`
}

type CreditHistoryIdentification struct {
	Type string `json:"type"`
	No   string `json:"no"`
}

type MarshupRequest struct {
	NIN string `json:"nin"`
	BVN string `json:"bvn"`
	DOB string `json:"date_of_birth"`
}

type MarshupPersonalInformation struct {
	Title         string `json:"title"`
	FirstName     string `json:"first_name"`
	MiddleName    string `json:"middle_name"`
	Surname       string `json:"surname"`
	Gender        string `json:"gender"`
	DOB           string `json:"dob"`
	BirthDate     string `json:"birth_date"`
	BirthCountry  string `json:"birth_country"`
	BirthState    string `json:"birth_state"`
	BirthLGA      string `json:"birth_lga"`
	MaritalStatus string `json:"marital_status"`
	Email         string `json:"email"`
	TelephoneNo   string `json:"telephone_no"`
	Occupation    string `json:"occupation"`
	LGAOfOrigin   string `json:"lga_of_origin"`
	StateOfOrigin string `json:"state_of_origin"`
	WatchListed   bool   `json:"watch_listed"`
}

type MarshupIdentificationNumbers struct {
	NIN string `json:"nin"`
	BVN string `json:"bvn"`
}

type MarshupResidenceInformation struct {
	Address         string `json:"address"`
	Town            string `json:"town"`
	LGA             string `json:"lga"`
	State           string `json:"state"`
	ResidenceStatus string `json:"residence_status"`
}

type MarshupBiometrics struct {
	Photo string `json:"photo"`
}

type MarshupData struct {
	PersonalInformation   MarshupPersonalInformation   `json:"personal_information"`
	IdentificationNumbers MarshupIdentificationNumbers `json:"identification_numbers"`
	ResidenceInformation  MarshupResidenceInformation  `json:"residence_information"`
	Biometrics            MarshupBiometrics            `json:"biometrics"`
}

type MarshupResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Timestamp string      `json:"timestamp"`
	Data      MarshupData `json:"data"`
}

type DirectorsResponse struct {
	Status    string          `json:"status"`
	Message   string          `json:"message"`
	Timestamp string          `json:"timestamp"`
	Data      []DirectorsData `json:"data"`
}

type DirectorsData struct {
	ID                           int64           `json:"id"`
	Surname                      string          `json:"surname"`
	Firstname                    string          `json:"firstname"`
	OtherName                    string          `json:"other_name"`
	Email                        string          `json:"email"`
	PhoneNumber                  string          `json:"phone_number"`
	Gender                       string          `json:"gender"`
	FormerNationality            string          `json:"former_nationality"`
	Age                          int64           `json:"age"`
	City                         string          `json:"city"`
	Occupation                   string          `json:"occupation"`
	FormerName                   string          `json:"former_name"`
	CorporationName              string          `json:"corporation_name"`
	RCNumber                     string          `json:"rc_number"`
	CorporationCompany           string          `json:"corporation_company"`
	State                        string          `json:"state"`
	Pobox                        string          `json:"pobox"`
	Accreditationnumber          string          `json:"accreditationnumber"`
	IsLawyer                     string          `json:"is_lawyer"`
	LastVisit                    int64           `json:"last_visit"`
	FormType                     string          `json:"form_type"`
	IsPresenter                  string          `json:"is_presenter"`
	IsChairman                   string          `json:"is_chairman"`
	NumSharesAlloted             int64           `json:"num_shares_alloted"`
	TypeOfShares                 string          `json:"type_of_shares"`
	DateOfBirth                  string          `json:"date_of_birth"`
	Status                       string          `json:"status"`
	DateOfTermination            string          `json:"date_of_termination"`
	DateOfAppointment            string          `json:"date_of_appointment"`
	DateOfChangeOfAddress        string          `json:"date_of_change_of_address"`
	FormerAddress                string          `json:"former_address"`
	FormerPostal                 string          `json:"former_postal"`
	FormerSurname                string          `json:"former_surname"`
	FormerFirstName              string          `json:"former_first_name"`
	FormerOtherName              string          `json:"former_other_name"`
	DateOfStatusChange           string          `json:"date_of_status_change"`
	IdentityNumber               string          `json:"identity_number"`
	IdentityIssueState           string          `json:"identity_issue_state"`
	OtherDirectorshipDetails     string          `json:"other_directorship_details"`
	PortalUserFk                 string          `json:"portal_user_fk"`
	AffiliatesFk                 string          `json:"affiliates_fk"`
	ProcessTypeFk                ProcessTypeFk   `json:"process_type_fk"`
	Company                      string          `json:"company"`
	SamePersonAsFk               string          `json:"same_person_as_fk"`
	NatureOfAppOrDischarge       string          `json:"nature_of_app_or_discharge"`
	IsDesignated                 string          `json:"is_designated"`
	EndOfAppointment             string          `json:"end_of_appointment"`
	AppointedBy                  string          `json:"appointed_by"`
	DateOfDeedOfDischarge        string          `json:"date_of_deed_of_discharge"`
	DateOfResolution             string          `json:"date_of_resolution"`
	CountryFk                    CountryFk       `json:"country_fk"`
	CountryOfResidence           string          `json:"country_of_residence"`
	IsCarriedOverFromNameAvai    string          `json:"is_carried_over_from_name_avai"`
	LGA                          string          `json:"lga"`
	CorporationRegistrationDate  string          `json:"corporation_registration_date"`
	IsCompanyDeleted             string          `json:"is_company_deleted"`
	GovernmentOrganisationName   string          `json:"government_organisation_name"`
	ForeignOrganisationName      string          `json:"foreign_organisation_name"`
	CompanyStreetAddress         string          `json:"company_street_address"`
	CompanyState                 string          `json:"company_state"`
	CompanyCity                  string          `json:"company_city"`
	IsCorporate                  string          `json:"is_corporate"`
	CountyOfIncorporationFk      string          `json:"county_of_incorporation_fk"`
	Nationality                  string          `json:"nationality"`
	Address                      string          `json:"address"`
	Postcode                     string          `json:"postcode"`
	StreetNumber                 string          `json:"street_number"`
	AffiliatesResidentialAddress string          `json:"affiliates_residential_address"`
	AffiliatesPscInformation     int64           `json:"affiliates_psc_information"`
	LegalOwnersOfInterests       []string        `json:"legal_owners_of_interests"`
	LegalOwnersOfVotingRights    []string        `json:"legal_owners_of_voting_rights"`
	StockExchangeSoes            []string        `json:"stock_exchange_soes"`
	ApprovedForNoticeOfPsc       string          `json:"approved_for_notice_of_psc"`
	CompanyAddress2              string          `json:"company_address2"`
	FullAddress2                 string          `json:"full_address2"`
	AffiliateTypeFk              AffiliateTypeFk `json:"affiliate_type_fk"`
}
