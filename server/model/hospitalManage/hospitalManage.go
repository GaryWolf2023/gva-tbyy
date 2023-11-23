package hospitalManage

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type HrEmployee struct {
	ID                          int       `json:"id" gorm:"column:id"`
	CreateDate                  time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate                   time.Time `json:"writeDate" gorm:"column:write_date"`
	Name                        string    `json:"name" gorm:"column:name"`
	ResourceId                  int       `json:"resource_id" gorm:"column:resource_id"`
	CompanyId                   int       `json:"companyId" gorm:"column:company_id"`
	ResourceCalendarId          int       `json:"resourceCalendarId" gorm:"column:resource_calendar_id"`
	MessageMainAttachmentId     int       `json:"messageMainAttachmentId" gorm:"message_main_attachment_id"`
	Color                       int       `json:"color" gorm:"column:color"`
	DepartmentId                int       `json:"departmentId" gorm:"column:department_id"`
	JobId                       int       `json:"jobId" gorm:"column: job_id"`
	AddressId                   int       `json:"addressId" gorm:"column:address_id"`
	WorkContactId               int       `json:"workContactId" gorm:"column:work_contact_id"`
	WorkLocationId              int       `json:"workLocationId" gorm:"column:work_location_id"`
	UserId                      int       `json:"userId" gorm:"column:user_id"`
	ParentId                    int       `json:"parentId" gorm:"column:parent_id"`
	CoachId                     int       `json:"coachId" gorm:"column:coach_id"`
	AddressHomeId               int       `json:"addressHomeId" gorm:"column:address_home_id"`
	CountryId                   int       `json:"countryId" gorm:"column:country_id"`
	Children                    int       `json:"children" gorm:"column:children"`
	CountryOfBirth              int       `json:"countryOfBirth" gorm:"column:country_of_birth"`
	BankAccountId               int       `json:"bankAccountId" gorm:"column:bank_account_id"`
	KmHomeWork                  int       `json:"kmHomeWork" gorm:"column:km_home_work"`
	DepartureReasonId           int       `json:"departureReasonId" gorm:"column:departure_reason_id"`
	CreateUid                   int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid                    int       `json:"writeUid" gorm:"column:write_uid"`
	JobTitle                    string    `json:"jobTitle" gorm:"column:job_title"`
	WorkPhone                   string    `json:"workPhone" gorm:"column:work_phone"`
	MobilePhone                 string    `json:"mobilePhone" gorm:"column:mobile_phone"`
	WorkEmail                   string    `json:"workEmail" gorm:"column:work_email"`
	EmployeeType                string    `json:"employeeType" gorm:"column:employee_type"`
	Gender                      int       `json:"gender" gorm:"column:gender"`
	Marital                     int       `json:"marital" gorm:"column:marital"`
	SpouseCompleteName          string    `json:"spouseCompleteName" gorm:"column:spouse_complete_name"`
	PlaceOfBirth                string    `json:"placeOfBirth" gorm:"column:place_of_birth"` // 出生地
	Ssnid                       string    `json:"ssnid" gorm:"column:ssnid"`
	Sinid                       string    `json:"sinid" gorm:"column:sinid"`
	IdentificationId            string    `json:"identificationId" gorm:"column:identification_id"` //身份证
	PassportId                  string    `json:"passportId" gorm:"column:passport_id"`
	PermitNo                    string    `json:"permitNo" gorm:"column:permit_no"`
	VisaNo                      string    `json:"visaNo" gorm:"column:visa_no"`
	Certificate                 string    `json:"certificate" gorm:"column:certificate"`
	StudyField                  string    `json:"studyField" gorm:"column:study_field"`
	StudySchool                 string    `json:"studySchool" gorm:"column:study_school"` // 毕业学校
	EmergencyContact            string    `json:"emergencyContact" gorm:"column:emergency_contact"`
	EmergencyPhone              string    `json:"emergencyPhone" gorm:"column:emergency_phone"`
	Barcode                     string    `json:"barcode" gorm:"column:barcode"`
	Pin                         string    `json:"pin" gorm:"column:barcode"`
	SpouseBirthdate             time.Time `json:"spouseBirthdate" gorm:"column:spouse_birthdate"`
	Birthday                    time.Time `json:"birthday" gorm:"column:birthday"` // 出生日期
	VisaExpire                  time.Time `json:"visaExpire" gorm:"column:visa_expire"`
	WorkPermitExpirationDate    time.Time `json:"workPermitExpirationDate" gorm:"column:work_permit_expiration_date"`
	DepartureDate               time.Time `json:"departureDate" gorm:"column:departure_date"`
	AdditionalNote              string    `json:"additionalNote" gorm:"column:additional_note"`
	Notes                       string    `json:"notes" gorm:"column:notes"`
	DepartureDescription        string    `json:"departureDescription" gorm:"column:departure_description"`
	Active                      bool      `json:"active" gorm:"column:active"`
	WorkPermitScheduledActivity bool      `json:"workPermitScheduledActivity" gorm:"column:work_permit_scheduled_activity"`
	LastAttendanceId            int       `json:"lastAttendanceId" gorm:"column:last_attendance_id"`
	LastCheckIn                 time.Time `json:"lastCheckIn" gorm:"column:last_check_in"`
	LastCheckOut                time.Time `json:"lastCheckOut" gorm:"column:last_check_out"`
	ContractId                  int       `json:"contractId" gorm:"column:contract_id"`
	Vehicle                     string    `json:"vehicle" gorm:"column:vehicle"`
	FirstContractDate           time.Time `json:"firstContractDate" gorm:"column:first_contract_date"`
	ContractWarning             bool      `json:"contractWarning" gorm:"column:contract_warning"`
	LeaveManagerId              int       `json:"leaveManagerId" gorm:"column:leave_manager_id"`
	HourlyCost                  int       `json:"hourlyCost" gorm:"column:hourly_cost"`
	PersonalMobile              string    `json:"personalMobile" gorm:"column:personal_mobile"`
	JoiningDate                 time.Time `json:"joiningDate" gorm:"column:joining_date"`
	IdExpiryDate                time.Time `json:"idExpiryDate" gorm:"column:id_expiry_date"`
	PassportExpiryDate          time.Time `json:"passportExpiryDate" gorm:"column:passport_expiry_date"`
	ResignDate                  time.Time `json:"resignDate" gorm:"column:resign_date"`
	Resigned                    bool      `json:"resigned" gorm:"column:resigned"`
	Fired                       bool      `json:"fired" gorm:"column:fired"`
	ExpenseManagerId            int       `json:"expenseManagerId" gorm:"column:expense_manager_id"`
	ResourceCalendarIds         int       `json:"resourceCalendarIds" gorm:"column:resource_calendar_ids"`
	Nation                      int       `json:"nation" gorm:"column:nation"`
	Schoolgrade                 int       `json:"schoolgrade" gorm:"column:schoolgrade"`
	Appellation                 int       `json:"appellation" gorm:"column:appellation"`
	BookingNumAm                int       `json:"bookingNumAm" gorm:"column:booking_num_am"`
	BookingNumPm                int       `json:"bookingNumPm" gorm:"column:booking_num_pm"`
	BookingNumNg                int       `json:"bookingNumNg" gorm:"column:booking_num_ng"`
	PoliticsStatus              string    `json:"politicsStatus" gorm:"column:politics_status"`
	Tocompile                   string    `json:"tocompile" gorm:"column:tocompile"`
	AppellationCode             string    `json:"appellationCode" gorm:"column:appellation_code"`
	Rank                        string    `json:"rank" gorm:"column:rank"`
	AppellationTime             time.Time `json:"appellationTime" gorm:"column:appellation_time"`
	Note                        string    `json:"note" gorm:"column:note"`
	Introduction                string    `json:"introduction" gorm:"column:introduction"`
	GradeId                     int       `json:"gradeId" gorm:"column:grade_id"`
	RankId                      int       `json:"rankId" gorm:"column:rank_id"`
	NaturePosition              string    `json:"naturePosition" gorm:"column:nature_position"`
	SurgeryCert                 string    `json:"surgeryCert" gorm:"column:surgery_cert"`
	Inductiondate               time.Time `json:"inductiondate" gorm:"column:inductiondate"`
	Performanceclass            string    `json:"performanceclass" gorm:"column:performanceclass"`
	EduAppellation              string    `json:"eduAppellation" gorm:"column:edu_appellation"`
	EduCategory                 string    `json:"eduCategory" gorm:"column:edu_category"`
	PhysicianCert               int       `json:"physicianCert" gorm:"column:physician_cert"`
	PharmacistCert              int       `json:"pharmacistCert" gorm:"column:pharmacist_cert"`
	NurseCert                   int       `json:"nurseCert" gorm:"column:nurse_cert"`
	OperateCert                 string    `json:"operateCert" gorm:"column:operate_cert"`
	PrescriptionCert            string    `json:"prescriptionCert" gorm:"column:prescription_cert"`
	SpecialistId                int       `json:"specialistId" gorm:"column:specialist_id"`
	NurseSpe                    string    `json:"nurseSpe" gorm:"column:nurse_spe"`
	SpecialistDate              time.Time `json:"specialistDate" gorm:"column:specialist_date"`
	IsSpecialist                bool      `json:"isSpecialist" gorm:"column:is_specialist"`
	EduSchool                   string    `json:"eduSchool" gorm:"column:edu_school"`
	SurgeryGrade                int       `json:"surgeryGrade" gorm:"column:surgery_grade"`
	PracticeCateg               string    `json:"practiceCateg" gorm:"column:practice_categ"`
	PhysicianAidCert            int       `json:"physicianAidCert" gorm:"column:physician_aid_cert"`
	NativePlaceP                int       `json:"nativePlaceP" gorm:"column:native_place_p"`
	NativePlaceC                int       `json:"nativePlaceC" gorm:"column:native_place_c"`
	GraduationDate              time.Time `json:"graduationDate" gorm:"column:graduation_date"`
}
type ResPartner struct {
	ID                      int       `json:"id" gorm:"column:id"`
	CreateDate              time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate               time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid               int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid                int       `json:"writeUid" gorm:"column:update_uid"`
	CompanyId               int       `json:"companyId" gorm:"column:company_id"`
	Name                    string    `json:"name" gorm:"column:name"`
	Title                   int       `json:"title" gorm:"column:title"`
	ParentId                int       `json:"parentId" gorm:"column:parent_id"`
	UserId                  int       `json:"userId" gorm:"column:user_id"`
	StateId                 int       `json:"stateId" gorm:"column:state_id"`
	CountryId               int       `json:"countryId" gorm:"column:country_id"`
	IndustryId              int       `json:"industryId" gorm:"column:industry_id"`
	Color                   int       `json:"color" gorm:"column:color"`
	CommercialPartnerId     int       `json:"commercialPartnerId" gorm:"column:commercial_partner_id"`
	DisplayName             string    `json:"displayName" gorm:"column:display_name"`
	Ref                     string    `json:"ref" gorm:"column:ref"`
	Lang                    string    `json:"lang" gorm:"column:lang"`
	Tz                      string    `json:"tz" gorm:"column:tz"`
	Vat                     string    `json:"vat" gorm:"column:vat"`
	CompanyRegistry         string    `json:"companyRegistry" gorm:"column:company_registry"`
	Website                 string    `json:"website" gorm:"column:website"`
	Function                string    `json:"function" gorm:"column:function"`
	Type                    string    `json:"type" gorm:"column:type"`
	Street                  string    `json:"street" gorm:"column:street"`
	Street2                 string    `json:"street2" gorm:"column:street2"`
	Zip                     string    `json:"zip" gorm:"column:zip"`
	City                    string    `json:"city" gorm:"column:city"`
	Email                   string    `json:"email" gorm:"column:email"`
	Phone                   string    `json:"phone" gorm:"column:phone"`
	Mobile                  string    `json:"mobile" gorm:"column:mobile"`
	CommercialCompanyName   string    `json:"commercialCompanyName" gorm:"column:commercial_company_name"`
	CompanyName             string    `json:"companyName" gorm:"column:company_name"`
	Date                    time.Time `json:"date" gorm:"column:date"`
	Comment                 string    `json:"comment" gorm:"column:comment"`
	PartnerLatitude         float64   `json:"partnerLatitude" gorm:"column:partner_latitude"`
	PartnerLongitude        float64   `json:"partnerLongitude" gorm:"column:partner_longitude"`
	Active                  bool      `json:"active" gorm:"column:active"`
	Employee                bool      `json:"employee" gorm:"column:employee"`
	IsCompany               bool      `json:"isCompany" gorm:"column:is_company"`
	PartnerShare            bool      `json:"partnerShare" gorm:"column:partner_share"`
	MessageMainAttachmentId int       `json:"messageMainAttachmentId" gorm:"column:message_main_attachment_id"`
	MessageBounce           int       `json:"messageBounce" gorm:"column:message_bounce"`
	EmailNormalized         string    `json:"emailNormalized" gorm:"column:email_normalized"`
	SignupToken             string    `json:"signupToken" gorm:"column:signup_token"`
	SignupType              string    `json:"signupType" gorm:"column:signup_type"`
	SignupExpiration        time.Time `json:"signupExpiration" gorm:"column:signup_expiration"`
	CalendarLastNotifAck    time.Time `json:"calendarLastNotifAck" gorm:"column:calendar_last_notif_ack"`
	PartnerGid              int       `json:"partnerGid" gorm:"column:partner_gid"`
	AdditionalInfo          string    `json:"additionalInfo" gorm:"column:additional_info"`
	PhoneSanitized          string    `json:"phoneSanitized" gorm:"column:phone_sanitized"`
	SupplierRank            int       `json:"supplierRank" gorm:"column:supplier_rank"`
	CustomerRank            int       `json:"customerRank" gorm:"column:customer_rank"`
	InvoiceWarn             string    `json:"invoiceWarn" gorm:"column:invoice_warn"`
	InvoiceWarnMsg          string    `json:"invoiceWarnMsg" gorm:"column:invoice_warn_msg"`
	DebitLimit              float64   `json:"debitLimit" gorm:"column:debit_limit"`
	LastTimeEntriesChecked  time.Time `json:"lastTimeEntriesChecked" gorm:"column:last_time_entries_checked"`
	TeamId                  int       `json:"teamId" gorm:"column:team_id"`
	PickingWarn             string    `json:"pickingWarn" gorm:"column:picking_warn"`
	PickingWarnMsg          string    `json:"pickingWarnMsg" gorm:"column:picking_warn_msg"`
	SaleWarn                string    `json:"saleWarn" gorm:"column:sale_warn"`
	SaleWarnMsg             string    `json:"saleWarnMsg" gorm:"column:sale_warn_msg"`
	PurchaseWarn            string    `json:"purchaseWarn" gorm:"column:purchase_warn"`
	PurchaseWarnMsg         string    `json:"purchaseWarnMsg" gorm:"column:purchase_warn_msg"`
	CityId                  int       `json:"cityId" gorm:"column:city_id"`
	AreaId                  int       `json:"areaId" gorm:"column:area_id"`
	Town                    string    `json:"town" gorm:"column:town"`
	Village                 string    `json:"village" gorm:"column:village"`
	HouseNumber             string    `json:"houseNumber" gorm:"column:house_number"`
}
type ResPartnerBank struct {
	ID                      int       `json:"id" gorm:"column:id"`
	CreateDate              time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate               time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid               int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid                int       `json:"writeUid" gorm:"column:update_uid"`
	PartnerId               int       `json:"partnerId" gorm:"column:partner_id"`
	BankId                  int       `json:"bankId" gorm:"column:bank_id"`
	Sequence                int       `json:"sequence" gorm:"column:sequence"`
	CurrencyId              int       `json:"currencyId" gorm:"column:currency_id"`
	CompanyId               int       `json:"companyId" gorm:"column:company_id"`
	AccNumber               string    `json:"accNumber" gorm:"column:acc_number"`
	SanitizedAccNumber      string    `json:"sanitizedAccNumber" gorm:"column:sanitized_acc_number"`
	AccHolderName           string    `json:"accHolderName" gorm:"column:acc_holder_name"`
	Active                  bool      `json:"active" gorm:"column:active"`
	AllowOutPayment         bool      `json:"allowOutPayment" gorm:"column:allow_out_payment"`
	MessageMainAttachmentId int       `json:"messageMainAttachmentId" gorm:"column:message_main_attachment_id"`
}
type ResUsers struct{}
type ResCountry struct{}
type ResCountryCity struct{}
type ResCountryState struct{}
type RankRank struct{}
type ResourceCalendar struct{}
type ResourceResource struct{}
type HrContract struct {
	ID                      int       `json:"id" gorm:"column:id"`
	CreateDate              time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate               time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid               int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid                int       `json:"writeUid" gorm:"column:update_uid"`
	MessageMainAttachmentId int       `json:"messageMainAttachmentId" gorm:"column:message_main_attachment_id"`
	StructureTypeId         int       `json:"structureTypeId" gorm:"column:structure_type_id"`
	EmployeeId              int       `json:"employeeId" gorm:"column:employee_id"`
	DepartmentId            int       `json:"departmentId" gorm:"column:department_id"`
	JobId                   int       `json:"jobId" gorm:"column:job_id"`
	ResourceCalendarId      int       `json:"resourceCalendarId" gorm:"column:resource_calendar_id"`
	CompanyId               int       `json:"companyId" gorm:"column:company_id"`
	ContractTypeId          int       `json:"contractTypeId" gorm:"column:contract_type_id"`
	HrResponsibleId         int       `json:"hrResponsibleId" gorm:"column:hr_responsible_id"`
	Name                    string    `json:"name" gorm:"column:name"`
	State                   string    `json:"state" gorm:"column:state"`
	KanbanState             string    `json:"kanbanState" gorm:"column:kanban_state"`
	DateStart               time.Time `json:"dateStart" gorm:"column:date_start"`
	DateEnd                 time.Time `json:"dateEnd" gorm:"column:date_end"`
	TrialDateEnd            time.Time `json:"trialDateEnd" gorm:"column:trial_date_end"`
	Notes                   string    `json:"notes" gorm:"column:notes"`
	Wage                    float64   `json:"wage" gorm:"column:wage"`
	Active                  bool      `json:"active" gorm:"column:active"`
	TypeId                  int       `json:"typeId" gorm:"column:type_id"`
	NoticeDays              int       `json:"noticeDays" gorm:"column:notice_days"`
	StructId                int       `json:"structId" gorm:"column:struct_id"`
	SchedulePay             string    `json:"schedulePay" gorm:"column:schedule_pay"`
	Hra                     float64   `json:"hra" gorm:"column:hra"`
	TravelAllowance         float64   `json:"travelAllowance" gorm:"column:travel_allowance"`
	Da                      float64   `json:"da" gorm:"column:da"`
	MealAllowance           float64   `json:"mealAllowance" gorm:"column:meal_allowance"`
	MedicalAllowance        float64   `json:"medicalAllowance" gorm:"column:medical_allowance"`
	OtherAllowance          float64   `json:"otherAllowance" gorm:"column:other_allowance"`
	AnalyticAccountId       int       `json:"analyticAccountId" gorm:"column:analytic_account_id"`
	JournalId               int       `json:"journalId" gorm:"column:journal_id"`
	OverHour                float64   `json:"overHour" gorm:"column:over_hour"`
	OverDay                 float64   `json:"overDay" gorm:"column:over_day"`
	WorkingHours            int       `json:"workingHours" gorm:"column:working_hours"`
}
type HrDepartureReason struct {
	ID         int               `json:"id" gorm:"column:id"`
	CreateDate time.Time         `json:"createDate" gorm:"column:create_date"`
	WriteDate  time.Time         `json:"writeDate" gorm:"column:write_date"`
	Sequence   int               `json:"sequence" gorm:"column:sequence"`
	CreateUid  int               `json:"createUid" gorm:"column:create_uid"`
	WriteUid   int               `json:"writeUid" gorm:"column:update_uid"`
	Name       pgtype.JSONBCodec `json:"name" gorm:"column:name"`
}
type GradeGrade struct {
	ID          int       `json:"id" gorm:"column:id"`
	CreateDate  time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate   time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid   int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid    int       `json:"writeUid" gorm:"column:update_uid"`
	Name        string    `json:"name" gorm:"column:name"`
	Description string    `json:"description" gorm:"column:description"`
	Active      bool      `json:"active" gorm:"column:active"`
}
type HrJob struct {
	ID                      int       `json:"id" gorm:"column:id"`
	CreateDate              time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate               time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid               int       `json:"createUid" gorm:"column:create_uid;foreignKey:CreateUid"`
	WriteUid                int       `json:"writeUid" gorm:"column:write_uid;foreignKey:WriteUid"`
	MessageMainAttachmentId int       `json:"messageMainAttachmentId" gorm:"column:message_main_attachment_id"`
	Sequence                int       `json:"sequence" gorm:"column:sequence"`
	ExpectedEmployees       int       `json:"expectedEmployees" gorm:"column:expected_employees"`
	NoOfEmployee            int       `json:"noOfEmployee" gorm:"column:no_of_employee"`
	NoOfRecruitment         int       `json:"noOfRecruitment" gorm:"column:no_of_recruitment"`
	NoOfHiredEmployee       int       `json:"noOfHiredEmployee" gorm:"column:no_of_hired_employee"`
	DepartmentId            int       `json:"departmentId" gorm:"column:department_id;foreignKey:DepartmentId"`
	CompanyId               int       `json:"companyId" gorm:"column:company_id;foreignKey:CompanyId"`
	ContractTypeId          int       `json:"contractTypeId" gorm:"column:contract_type_id"`
	Name                    string    `json:"name" gorm:"column:name"`
	Description             string    `json:"description" gorm:"column:description"`
	Requirements            string    `json:"requirements" gorm:"column:requirements"`
	Active                  bool      `json:"active" gorm:"column:active"`
	AliasId                 int       `json:"aliasId" gorm:"column:alias_id"`
	AddressId               int       `json:"addressId" gorm:"column:address_id;foreignKey:AddressId"`
	ManagerId               int       `json:"managerId" gorm:"column:manager_id"`
	UserId                  int       `json:"userId" gorm:"column:user_id;foreignKey:UserId"`
	HrResponsibleId         int       `json:"hrResponsibleId" gorm:"column:hr_responsible_id"`
	Color                   int       `json:"color" gorm:"column:color"`
	NaturePost              string    `json:"naturePost" gorm:"column:nature_post"`
	CategoryPost            string    `json:"categoryPost" gorm:"column:category_post"`
	JobCode                 string    `json:"jobCode" gorm:"column:job_code"`
	Show                    string    `json:"show" gorm:"column:show"`
}
type HrAttendance struct {
	ID          int       `json:"id" gorm:"column:id"`
	CreateDate  time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate   time.Time `json:"writeDate" gorm:"column:write_date"`
	EmployeeId  int       `json:"employeeId" gorm:"column:employee_id"`
	CreateUid   int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid    int       `json:"writeUid" gorm:"column:update_uid"`
	CheckIn     time.Time `json:"checkIn" gorm:"column:check_in"`
	CheckOut    time.Time `json:"checkOut" gorm:"column:check_out"`
	WorkedHours float64   `json:"workedHours" gorm:"column:worked_hours"`
	CompanyId   int       `json:"companyId" gorm:"column:company_id"`
}
type HrEmployeePractisingcert struct {
	ID                 int       `json:"id" gorm:"column:id"`
	CreateDate         time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate          time.Time `json:"writeDate" gorm:"column:write_date"`
	Employee           int       `json:"employee" gorm:"column:employee"`
	CreateUid          int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid           int       `json:"writeUid" gorm:"column:update_uid"`
	Code               string    `json:"code" gorm:"column:code"`
	PractisingCateg    string    `json:"practisingCateg" gorm:"column:practising_categ"`
	PractisingRange    string    `json:"practisingRange" gorm:"column:practising_range"`
	PractisingLocation string    `json:"practisingLocation" gorm:"column:practising_location"`
	RegisterDate       time.Time `json:"registerDate" gorm:"column:register_date"`
	ApprovalEmployeeId int       `json:"approvalEmployeeId" gorm:"column:approval_employee_id"`
	Approval           bool      `json:"approval" gorm:"column:approval"`
	IsApproval         bool      `json:"isApproval" gorm:"column:is_approval"`
	ApprovalTime       time.Time `json:"approvalTime" gorm:"column:approval_time"`
	RevokeEmployeeId   int       `json:"revokeEmployeeId" gorm:"column:revoke_employee_id"`
	RevokeReason       string    `json:"revokeReason" gorm:"column:revoke_reason"`
	Revoke             bool      `json:"revoke" gorm:"column:revoke"`
	RevokeTime         time.Time `json:"revokeTime" gorm:"column:revoke_time"`
}
type HrEmployeeSpecialist struct {
	ID         int       `json:"id" gorm:"column:id"`
	CreateDate time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate  time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid  int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid   int       `json:"writeUid" gorm:"column:update_uid"`
	Name       string    `json:"name" gorm:"column:name"`
}
type HrWorkLocation struct {
	ID             int       `json:"id" gorm:"column:id"`
	CreateDate     time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate      time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid      int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid       int       `json:"writeUid" gorm:"column:update_uid"`
	CompanyId      int       `json:"companyId" gorm:"column:company_id"`
	AddressId      int       `json:"addressId" gorm:"column:address_id"`
	Name           string    `json:"name" gorm:"column:name"`
	LocationNumber string    `json:"locationNumber" gorm:"column:location_number"`
	Active         bool      `json:"active" gorm:"column:active"`
}
type IrAttachment struct {
	ID           int       `json:"id" gorm:"column:id"`
	CreateDate   time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate    time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid    int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid     int       `json:"writeUid" gorm:"column:update_uid"`
	ResId        int       `json:"resId" gorm:"column:res_id"`
	CompanyId    int       `json:"companyId" gorm:"column:company_id"`
	FileSize     int       `json:"fileSize" gorm:"column:file_size"`
	Name         string    `json:"name" gorm:"column:name"`
	ResModel     string    `json:"resModel" gorm:"column:res_model"`
	ResField     string    `json:"resField" gorm:"column:res_field"`
	Type         string    `json:"type" gorm:"column:type"`
	Url          string    `json:"url" gorm:"column:url"`
	AccessToken  string    `json:"accessToken" gorm:"column:access_token"`
	StoreFname   string    `json:"storeFname" gorm:"column:store_fname"`
	Checksum     string    `json:"checksum" gorm:"column:checksum"`
	Mimetype     string    `json:"mimetype" gorm:"column:mimetype"`
	Description  string    `json:"description" gorm:"column:description"`
	IndexContent string    `json:"indexContent" gorm:"column:index_content"`
	Public       bool      `json:"public" gorm:"column:public"`
	DbDatas      []byte    `json:"dbDatas" gorm:"column:db_dates"`
	OriginalId   int       `json:"originalId" gorm:"column:original_id"`
	S3flag       bool      `json:"s3Flag" gorm:"column:s3flag"`
}
type HrEmployeePrescription struct {
	ID         int       `json:"id" gorm:"column:id"`
	CreateDate time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate  time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid  int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid   int       `json:"writeUid" gorm:"column:update_uid"`
	Name       string    `json:"name" gorm:"column:name"`
	Active     bool      `json:"active" gorm:"column:active"`
}
type HrEmployeeCategory struct {
	ID         int       `json:"id" gorm:"column:id"`
	CreateDate time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate  time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid  int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid   int       `json:"writeUid" gorm:"column:update_uid"`
	Color      int       `json:"color" gorm:"column:color"`
	Name       string    `json:"name" gorm:"column:name"`
	Active     bool      `json:"active" gorm:"column:active"`
}
type HrEmployeeDocument struct {
	ID               int       `json:"id" gorm:"column:id"`
	CreateDate       time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate        time.Time `json:"writeDate" gorm:"column:write_date"`
	EmployeeRef      int       `json:"employeeRef" gorm:"column:employee_ref"`
	DocumentType     int       `json:"documentType" gorm:"column:document_type"`
	BeforeDays       int       `json:"beforeDays" gorm:"column:before_days"`
	CreateUid        int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid         int       `json:"writeUid" gorm:"column:update_uid"`
	Name             string    `json:"name" gorm:"column:name"`
	NotificationType string    `json:"notificationType" gorm:"column:notification_type"`
	ExpiryDate       time.Time `json:"expiryDate" gorm:"column:expiry_date"`
	IssueDate        time.Time `json:"issueDate" gorm:"column:issue_date"`
	Description      string    `json:"description" gorm:"column:description"`
}
type HrEmployeeFamily struct {
	ID            int       `json:"id" gorm:"column:id"`
	CreateDate    time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate     time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid     int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid      int       `json:"writeUid" gorm:"column:update_uid"`
	Job           string    `json:"job" gorm:"column:job"`
	BrithDate     time.Time `json:"brithDate" gorm:"column:brith_date"`
	RelationId    int       `json:"relationId" gorm:"column:relation_id"`
	EmployeeId    int       `json:"employeeId" gorm:"column:employee_id"`
	MemberContact string    `json:"memberContact" gorm:"column:member_contact"`
	MemberName    string    `json:"memberName" gorm:"column:member_name"`
}
type HrAnnouncement struct {
	ID                      int       `json:"id" gorm:"column:id"`
	CreateDate              time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate               time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid               int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid                int       `json:"writeUid" gorm:"column:update_uid"`
	MessageMainAttachmentId int       `json:"messageMainAttachmentId" gorm:"column:message_main_attachment_id"`
	CompanyId               int       `json:"companyId" gorm:"column:company_id"`
	Name                    string    `json:"name" gorm:"column:name"`
	State                   string    `json:"state" gorm:"column:state"`
	AnnouncementType        string    `json:"announcementType" gorm:"column:announcement_type"`
	RequestedDate           time.Time `json:"requestedDate" gorm:"column:requested_date"`
	DateStart               time.Time `json:"dateStart" gorm:"column:date_start"`
	DateEnd                 time.Time `json:"dateEnd" gorm:"column:date_end"`
	AnnouncementReason      string    `json:"announcementReason" gorm:"column:announcement_reason"`
	Announcement            string    `json:"announcement" gorm:"column:announcement"`
	IsAnnouncement          bool      `json:"isAnnouncement" gorm:"column:is_announcement"`
}

type HrApplicant struct {
	ID         int       `json:"id" gorm:"column:id"`
	CreateDate time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate  time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid  int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid   int       `json:"writeUid" gorm:"column:update_uid"`
	CampaignId int       `json:"campaignId" gorm:"column:campaign_id"`
	SourceId   int       `json:"sourceId" gorm:"column:source_id"`
	MediumId   int       `json:"mediumId" gorm:"column:medium_id"`
}

func (h *HrEmployee) Table() string {
	return "hr_employee"
}

//	func () Table() string {
//		return
//	}
func (h *HrEmployeePrescription) Table() string {
	return "hr_employee_prescription"
}
func (h *HrEmployeeCategory) Table() string {
	return "hr_employee_category"
}
func (h *HrEmployeeDocument) Table() string {
	return "hr_employee_document"
}
func (h *HrEmployeeFamily) Table() string {
	return "hr_employee_family"
}
func (h *HrApplicant) Table() string {
	return "hr_applicant"
}
