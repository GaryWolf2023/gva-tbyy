package request

import "time"

type StaffInfo struct {
	ID                          int       `json:"id"`
	Name                        string    `json:"name"`
	ResourceId                  int       `json:"resource_id"`
	CompanyId                   int       `json:"companyId"`
	ResourceCalendarId          int       `json:"resourceCalendarId"`
	MessageMainAttachmentId     int       `json:"messageMainAttachmentId"`
	Color                       int       `json:"color"`
	DepartmentId                int       `json:"departmentId"`
	JobId                       int       `json:"jobId"`
	AddressId                   int       `json:"addressId"`
	WorkContactId               int       `json:"workContactId"`
	WorkLocationId              int       `json:"workLocationId"`
	UserId                      int       `json:"userId"`
	ParentId                    int       `json:"parentId"`
	CoachId                     int       `json:"coachId"`
	AddressHomeId               int       `json:"addressHomeId"`
	CountryId                   int       `json:"countryId"`
	Children                    int       `json:"children"`
	CountryOfBirth              int       `json:"countryOfBirth"`
	BankAccountId               int       `json:"bankAccountId"`
	KmHomeWork                  int       `json:"kmHomeWork"`
	DepartureReasonId           int       `json:"departureReasonId"`
	CreateUid                   int       `json:"createUid"`
	WriteUid                    int       `json:"writeUid"`
	JobTitle                    string    `json:"jobTitle"`
	WorkPhone                   string    `json:"workPhone"`
	MobilePhone                 string    `json:"mobilePhone"`
	WorkEmail                   string    `json:"workEmail"`
	EmployeeType                string    `json:"employeeType"`
	Gender                      int       `json:"gender"`
	Marital                     int       `json:"marital"`
	SpouseCompleteName          string    `json:"spouseCompleteName"`
	PlaceOfBirth                string    `json:"placeOfBirth"` // 出生地
	Ssnid                       string    `json:"ssnid"`
	Sinid                       string    `json:"sinid"`
	IdentificationId            string    `json:"identificationId"` //身份证
	PassportId                  string    `json:"passportId"`
	PermitNo                    string    `json:"permitNo"`
	VisaNo                      string    `json:"visaNo"`
	Certificate                 string    `json:"certificate"`
	StudyField                  string    `json:"studyField"`
	StudySchool                 string    `json:"studySchool"` // 毕业学校
	EmergencyContact            string    `json:"emergencyContact"`
	EmergencyPhone              string    `json:"emergencyPhone"`
	Barcode                     string    `json:"barcode"`
	Pin                         string    `json:"pin"`
	SpouseBirthdate             time.Time `json:"spouseBirthdate"`
	Birthday                    time.Time `json:"birthday"` // 出生日期
	VisaExpire                  time.Time `json:"visaExpire"`
	WorkPermitExpirationDate    time.Time `json:"workPermitExpirationDate"`
	DepartureDate               time.Time `json:"departureDate"`
	AdditionalNote              string    `json:"additionalNote"`
	Notes                       string    `json:"notes"`
	DepartureDescription        string    `json:"departureDescription"`
	Active                      bool      `json:"active"`
	WorkPermitScheduledActivity bool      `json:"workPermitScheduledActivity"`
	CreateDate                  time.Time `json:"createDate"`
	WriteDate                   time.Time `json:"writeDate"`
	LastAttendanceId            int       `json:"lastAttendanceId"`
	LastCheckIn                 time.Time `json:"lastCheckIn"`
	LastCheckOut                time.Time `json:"lastCheckOut"`
	ContractId                  int       `json:"contractId"`
	Vehicle                     string    `json:"vehicle"`
	FirstContractDate           time.Time `json:"firstContractDate"`
	ContractWarning             bool      `json:"contractWarning"`
	LeaveManagerId              int       `json:"leaveManagerId"`
	HourlyCost                  int       `json:"hourlyCost"`
	PersonalMobile              string    `json:"personalMobile"`
	JoiningDate                 time.Time `json:"joiningDate"`
	IdExpiryDate                time.Time `json:"idExpiryDate"`
	PassportExpiryDate          time.Time `json:"passportExpiryDate"`
	ResignDate                  time.Time `json:"resignDate"`
	Resigned                    bool      `json:"resigned"`
	Fired                       bool      `json:"fired"`
	ExpenseManagerId            int       `json:"expenseManagerId"`
	ResourceCalendarIds         int       `json:"resourceCalendarIds"`
	Nation                      int       `json:"nation"`
	Schoolgrade                 int       `json:"schoolgrade"`
	Appellation                 int       `json:"appellation"`
	BookingNumAm                int       `json:"bookingNumAm"`
	BookingNumPm                int       `json:"bookingNumPm"`
	BookingNumNg                int       `json:"bookingNumNg"`
	PoliticsStatus              string    `json:"politicsStatus"`
	Tocompile                   string    `json:"tocompile"`
	AppellationCode             string    `json:"appellationCode"`
	Rank                        string    `json:"rank"`
	AppellationTime             time.Time `json:"appellationTime"`
	Note                        string    `json:"note"`
	Introduction                string    `json:"introduction"`
	GradeId                     int       `json:"gradeId"`
	RankId                      int       `json:"rankId"`
	NaturePosition              string    `json:"naturePosition"`
	SurgeryCert                 string    `json:"surgeryCert"`
	Inductiondate               time.Time `json:"inductiondate"`
	Performanceclass            string    `json:"performanceclass"`
	EduAppellation              string    `json:"eduAppellation"`
	EduCategory                 string    `json:"eduCategory"`
	PhysicianCert               int       `json:"physicianCert"`
	PharmacistCert              int       `json:"pharmacistCert"`
	NurseCert                   int       `json:"nurseCert"`
	OperateCert                 string    `json:"operateCert"`
	PrescriptionCert            string    `json:"prescriptionCert"`
	SpecialistId                int       `json:"specialistId"`
	NurseSpe                    string    `json:"nurseSpe"`
	SpecialistDate              time.Time `json:"specialistDate"`
	IsSpecialist                bool      `json:"isSpecialist"`
	EduSchool                   string    `json:"eduSchool"`
	SurgeryGrade                int       `json:"surgeryGrade"`
	PracticeCateg               string    `json:"practiceCateg"`
	PhysicianAidCert            int       `json:"physicianAidCert"`
	NativePlaceP                int       `json:"nativePlaceP"`
	NativePlaceC                int       `json:"nativePlaceC"`
	GraduationDate              time.Time `json:"graduationDate"`
}

type BasicInfo struct {
	ID                 int       `json:"id" bind:"required" msg:"用户名不能为空"`
	Name               string    `json:"name" bind:"required" msg:"用户名不能为空"`
	JobTitle           string    `json:"jobTitle"`
	MobilePhone        string    `json:"mobilePhone"`
	WorkEmail          string    `json:"workEmail"`
	WorkPhone          string    `json:"workPhone"`
	Nation             int       `json:"nation"`
	CountryId          int       `json:"countryId"`
	Gender             int       `json:"gender"`
	Birthday           time.Time `json:"birthday"`
	PlaceOfBirth       string    `json:"placeOfBirth"`
	CountryOfBirth     int       `json:"countryOfBirth"`
	NativePlaceP       int       `json:"nativePlaceP"`
	NativePlaceC       int       `json:"nativePlaceC"`
	PoliticsStatus     string    `json:"politicsStatus"`
	AddressHomeId      int       `json:"addressHomeId"`
	BankAccountID      int       `json:"bankAccountId"`
	KmHomeWork         int       `json:"kmHomeWork"`
	IdentificationId   string    `json:"identificationId"`
	IdExpiryDate       time.Time `json:"idExpiryDate"`
	PassportId         string    `json:"passportId"`
	Marital            int       `json:"marital"`
	SpouseCompleteName string    `json:"spouseCompleteName"`
	SpouseBirthdate    time.Time `json:"spouseBirthdate"`
	Children           int       `json:"children"`
	EmergencyContact   string    `json:"emergencyContact"`
	EmergencyPhone     string    `json:"emergencyPhone"`
	Certificate        string    `json:"certificate"`
	SchoolGrade        int       `json:"schoolGrade"`
	GraduationDate     time.Time `json:"graduationDate"`
	StudyField         string    `json:"studyField"`
	StudySchool        string    `json:"studySchool"`
	VisaNo             string    `json:"visaNo"`
	PermitNo           string    `json:"permitNo"`
	VisaExpire         time.Time `json:"visaExpire"`
}

type SystemConfig struct {
	ID           int    `json:"id"`
	BookingNumAm int    `json:"bookingNumAm"`
	BookingNumPm int    `json:"bookingNumPm"`
	BookingNumNg int    `json:"bookingNumNg"`
	Introduction string `json:"introduction"`
}

type Qualification struct {
	Appellation      int       `json:"appellation"`
	AppellationCode  int       `json:"appellationCode"`
	AppellationTime  time.Time `json:"appellationTime"`
	EduAppellation   string    `json:"eduAppellation"`
	EduSchool        string    `json:"eduSchool"`
	EduCategory      string    `json:"eduCategory"`
	PracticeCateg    string    `json:"practiceCateg"`
	PhysicianCert    int       `json:"physicianCert"`
	PhysicianAidCert int       `json:"physicianAidCert"`
	PharmacistCert   int       `json:"pharmacistCert"`
	NurseCert        int       `json:"nurseCert"`
	OperateCert      string    `json:"operateCert"`
	SurgeryCert      string    `json:"surgeryCert"`
	SurgeryGrade     int       `json:"surgeryGrade"`
}
