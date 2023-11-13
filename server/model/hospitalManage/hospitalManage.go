package hospitalManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type HrEmployee struct {
	ID                          int       `json:"id" gorm:"column:id"`
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
	CreateDate                  time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate                   time.Time `json:"writeDate" gorm:"column:write_date"`
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
	global.GVA_MODEL
	CreateUid               int       `json:"createUid" gorm:"create_uid"`
	WriteUid                int       `json:"writeUid" gorm:"update_uid"`
	CompanyId               int       `json:"companyId" gorm:"company_id"`
	Name                    string    `json:"name" gorm:"name"`
	Title                   int       `json:"title" gorm:"title"`
	ParentId                int       `json:"parentId" gorm:"parent_id"`
	UserId                  int       `json:"userId" gorm:"user_id"`
	StateId                 int       `json:"stateId" gorm:"state_id"`
	CountryId               int       `json:"countryId" gorm:"country_id"`
	IndustryId              int       `json:"industryId" gorm:"industry_id"`
	Color                   int       `json:"color" gorm:"color"`
	CommercialPartnerId     int       `json:"commercialPartnerId" gorm:"commercial_partner_id"`
	DisplayName             string    `json:"displayName" gorm:"display_name"`
	Ref                     string    `json:"ref" gorm:"ref"`
	Lang                    string    `json:"lang" gorm:"lang"`
	Tz                      string    `json:"tz" gorm:"tz"`
	Vat                     string    `json:"vat" gorm:"vat"`
	CompanyRegistry         string    `json:"companyRegistry" gorm:"company_registry"`
	Website                 string    `json:"website" gorm:"website"`
	Function                string    `json:"function" gorm:"function"`
	Type                    string    `json:"type" gorm:"type"`
	Street                  string    `json:"street" gorm:"street"`
	Street2                 string    `json:"street2" gorm:"street2"`
	Zip                     string    `json:"zip" gorm:"zip"`
	City                    string    `json:"city" gorm:"city"`
	Email                   string    `json:"email" gorm:"email"`
	Phone                   string    `json:"phone" gorm:"phone"`
	Mobile                  string    `json:"mobile" gorm:"mobile"`
	CommercialCompanyName   string    `json:"commercialCompanyName" gorm:"commercial_company_name"`
	CompanyName             string    `json:"companyName" gorm:"company_name"`
	Date                    time.Time `json:"date" gorm:"date"`
	Comment                 string    `json:"comment" gorm:"comment"`
	PartnerLatitude         float64   `json:"partnerLatitude" gorm:"partner_latitude"`
	PartnerLongitude        float64   `json:"partnerLongitude" gorm:"partner_longitude"`
	Active                  bool      `json:"active" gorm:"active"`
	Employee                bool      `json:"employee" gorm:"employee"`
	IsCompany               bool      `json:"isCompany" gorm:"is_company"`
	PartnerShare            bool      `json:"partnerShare" gorm:"partner_share"`
	MessageMainAttachmentId int       `json:"messageMainAttachmentId" gorm:"message_main_attachment_id"`
	MessageBounce           int       `json:"messageBounce" gorm:"message_bounce"`
	EmailNormalized         string    `json:"emailNormalized" gorm:"email_normalized"`
	SignupToken             string    `json:"signupToken" gorm:"signup_token"`
	SignupType              string    `json:"signupType" gorm:"signup_type"`
	SignupExpiration        time.Time `json:"signupExpiration" gorm:"signup_expiration"`
	CalendarLastNotifAck    time.Time `json:"calendarLastNotifAck" gorm:"calendar_last_notif_ack"`
	PartnerGid              int       `json:"partnerGid" gorm:"partner_gid"`
	AdditionalInfo          string    `json:"additionalInfo" gorm:"additional_info"`
	PhoneSanitized          string    `json:"phoneSanitized" gorm:"phone_sanitized"`
	SupplierRank            int       `json:"supplierRank" gorm:"supplier_rank"`
	CustomerRank            int       `json:"customerRank" gorm:"customer_rank"`
	InvoiceWarn             string    `json:"invoiceWarn" gorm:"invoice_warn"`
	InvoiceWarnMsg          string    `json:"invoiceWarnMsg" gorm:"invoice_warn_msg"`
	DebitLimit              float64   `json:"debitLimit" gorm:"debit_limit"`
	LastTimeEntriesChecked  time.Time `json:"lastTimeEntriesChecked" gorm:"last_time_entries_checked"`
	TeamId                  int       `json:"teamId" gorm:"team_id"`
	PickingWarn             string    `json:"pickingWarn" gorm:"picking_warn"`
	PickingWarnMsg          string    `json:"pickingWarnMsg" gorm:"picking_warn_msg"`
	SaleWarn                string    `json:"saleWarn" gorm:"sale_warn"`
	SaleWarnMsg             string    `json:"saleWarnMsg" gorm:"sale_warn_msg"`
	PurchaseWarn            string    `json:"purchaseWarn" gorm:"purchase_warn"`
	PurchaseWarnMsg         string    `json:"purchaseWarnMsg" gorm:"purchase_warn_msg"`
	CityId                  int       `json:"cityId" gorm:"city_id"`
	AreaId                  int       `json:"areaId" gorm:"area_id"`
	Town                    string    `json:"town" gorm:"town"`
	Village                 string    `json:"village" gorm:"village"`
	HouseNumber             string    `json:"houseNumber" gorm:"house_number"`
}
type ResPartnerBank struct {
	global.GVA_MODEL
	CreateUid               int    `json:"createUid" gorm:"create_uid"`
	WriteUid                int    `json:"writeUid" gorm:"update_uid"`
	PartnerId               int    `json:"partnerId" gorm:"partner_id"`
	BankId                  int    `json:"bankId" gorm:"bank_id"`
	Sequence                int    `json:"sequence" gorm:"sequence"`
	CurrencyId              int    `json:"currencyId" gorm:"currency_id"`
	CompanyId               int    `json:"companyId" gorm:"company_id"`
	AccNumber               string `json:"accNumber" gorm:"acc_number"`
	SanitizedAccNumber      string `json:"sanitizedAccNumber" gorm:"sanitized_acc_number"`
	AccHolderName           string `json:"accHolderName" gorm:"acc_holder_name"`
	Active                  bool   `json:"active" gorm:"active"`
	AllowOutPayment         bool   `json:"allowOutPayment" gorm:"allow_out_payment"`
	MessageMainAttachmentId int    `json:"messageMainAttachmentId" gorm:"message_main_attachment_id"`
}
type ResCompany struct{}
type ResUsers struct{}
type ResCountry struct{}
type ResCountryCity struct{}
type ResCountryState struct{}
type RankRank struct{}
type ResourceCalendar struct{}
type ResourceResource struct{}
type HrContract struct {
	global.GVA_MODEL
	CreateUid               int       `json:"createUid" gorm:"create_uid"`
	WriteUid                int       `json:"writeUid" gorm:"update_uid"`
	MessageMainAttachmentId int       `json:"messageMainAttachmentId" gorm:"message_main_attachment_id"`
	StructureTypeId         int       `json:"structureTypeId" gorm:"structure_type_id"`
	EmployeeId              int       `json:"employeeId" gorm:"employee_id"`
	DepartmentId            int       `json:"departmentId" gorm:"department_id"`
	JobId                   int       `json:"jobId" gorm:"job_id"`
	ResourceCalendarId      int       `json:"resourceCalendarId" gorm:"resource_calendar_id"`
	CompanyId               int       `json:"companyId" gorm:"company_id"`
	ContractTypeId          int       `json:"contractTypeId" gorm:"contract_type_id"`
	HrResponsibleId         int       `json:"hrResponsibleId" gorm:"hr_responsible_id"`
	Name                    string    `json:"name" gorm:"name"`
	State                   string    `json:"state" gorm:"state"`
	KanbanState             string    `json:"kanbanState" gorm:"kanban_state"`
	DateStart               time.Time `json:"dateStart" gorm:"date_start"`
	DateEnd                 time.Time `json:"dateEnd" gorm:"date_end"`
	TrialDateEnd            time.Time `json:"trialDateEnd" gorm:"trial_date_end"`
	Notes                   string    `json:"notes" gorm:"notes"`
	Wage                    float64   `json:"wage" gorm:"wage"`
	Active                  bool      `json:"active" gorm:"active"`
	TypeId                  int       `json:"typeId" gorm:"type_id"`
	NoticeDays              int       `json:"noticeDays" gorm:"notice_days"`
	StructId                int       `json:"structId" gorm:"struct_id"`
	SchedulePay             string    `json:"schedulePay" gorm:"schedule_pay"`
	Hra                     float64   `json:"hra" gorm:"hra"`
	TravelAllowance         float64   `json:"travelAllowance" gorm:"travel_allowance"`
	Da                      float64   `json:"da" gorm:"da"`
	MealAllowance           float64   `json:"mealAllowance" gorm:"meal_allowance"`
	MedicalAllowance        float64   `json:"medicalAllowance" gorm:"medical_allowance"`
	OtherAllowance          float64   `json:"otherAllowance" gorm:"other_allowance"`
	AnalyticAccountId       int       `json:"analyticAccountId" gorm:"analytic_account_id"`
	JournalId               int       `json:"journalId" gorm:"journal_id"`
	OverHour                float64   `json:"overHour" gorm:"over_hour"`
	OverDay                 float64   `json:"overDay" gorm:"over_day"`
	WorkingHours            int       `json:"workingHours" gorm:"working_hours"`
}
type HrDepartment struct {
	global.GVA_MODEL
	CreateUid               int    `json:"createUid" gorm:"create_uid"`
	WriteUid                int    `json:"writeUid" gorm:"update_uid"`
	MessageMainAttachmentId int    `json:"messageMainAttachmentId" gorm:"message_main_attachment_id"`
	CompanyId               int    `json:"companyId" gorm:"company_id"`
	ParentId                int    `json:"parentId" gorm:"parent_id"`
	ManagerId               int    `json:"managerId" gorm:"manager_id"`
	Color                   int    `json:"color" gorm:"color"`
	MasterDepartmentId      int    `json:"masterDepartmentId" gorm:"master_department_id"`
	Name                    string `json:"name" gorm:"name"`
	CompleteName            string `json:"completeName" gorm:"complete_name"`
	ParentPath              string `json:"parentPath" gorm:"parent_path"`
	Note                    string `json:"note" gorm:"note"`
	Active                  bool   `json:"active" gorm:"active"`
	Property                int    `json:"property" gorm:"property"`
	Virtual                 bool   `json:"virtual" gorm:"virtual"`
	Available               bool   `json:"available" gorm:"available"`
	DutyId                  int    `json:"dutyId" gorm:"duty_id"`
	Code                    string `json:"code" gorm:"code"`
}
type HrDepartureReason struct {
	global.GVA_MODEL
	Sequence  int               `json:"sequence" gorm:"sequence"`
	CreateUid int               `json:"createUid" gorm:"create_uid"`
	WriteUid  int               `json:"writeUid" gorm:"update_uid"`
	Name      pgtype.JSONBCodec `json:"name" gorm:"name"`
}
type GradeGrade struct {
	global.GVA_MODEL
	CreateUid   int    `json:"createUid" gorm:"create_uid"`
	WriteUid    int    `json:"writeUid" gorm:"update_uid"`
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
	Active      bool   `json:"active" gorm:"active"`
}
type HrJob struct {
	global.GVA_MODEL
	CreateUid               int    `json:"createUid" gorm:"create_uid"`
	WriteUid                int    `json:"writeUid" gorm:"update_uid"`
	MessageMainAttachmentId int    `json:"messageMainAttachmentId" gorm:"message_main_attachment_id"`
	Sequence                int    `json:"sequence" gorm:"sequence"`
	ExpectedEmployees       int    `json:"expectedEmployees" gorm:"expected_employees"`
	NoOfEmployee            int    `json:"noOfEmployee" gorm:"no_of_employee"`
	NoOfRecruitment         int    `json:"noOfRecruitment" gorm:"no_of_recruitment"`
	NoOfHiredEmployee       int    `json:"noOfHiredEmployee" gorm:"no_of_hired_employee"`
	DepartmentId            int    `json:"departmentId" gorm:"department_id"`
	CompanyId               int    `json:"companyId" gorm:"company_id"`
	ContractTypeId          int    `json:"contractTypeId" gorm:"contract_type_id"`
	Name                    string `json:"name" gorm:"name"`
	Description             string `json:"description" gorm:"description"`
	Requirements            string `json:"requirements" gorm:"requirements"`
	Active                  bool   `json:"active" gorm:"active"`
	AliasId                 int    `json:"aliasId" gorm:"alias_id"`
	AddressId               int    `json:"addressId" gorm:"address_id"`
	ManagerId               int    `json:"managerId" gorm:"manager_id"`
	UserId                  int    `json:"userId" gorm:"user_id"`
	HrResponsibleId         int    `json:"hrResponsibleId" gorm:"hr_responsible_id"`
	Color                   int    `json:"color" gorm:"color"`
	NaturePost              string `json:"naturePost" gorm:"nature_post"`
	CategoryPost            string `json:"categoryPost" gorm:"category_post"`
	JobCode                 string `json:"jobCode" gorm:"job_code"`
	Show                    string `json:"show" gorm:"show"`
}
type HrAttendance struct {
	global.GVA_MODEL
	EmployeeId  int       `json:"employeeId" gorm:"employee_id"`
	CreateUid   int       `json:"createUid" gorm:"create_uid"`
	WriteUid    int       `json:"writeUid" gorm:"update_uid"`
	CheckIn     time.Time `json:"checkIn" gorm:"check_in"`
	CheckOut    time.Time `json:"checkOut" gorm:"check_out"`
	WorkedHours float64   `json:"workedHours" gorm:"worked_hours"`
	CompanyId   int       `json:"companyId" gorm:"company_id"`
}
type HrEmployeePractisingcert struct {
	global.GVA_MODEL
	Employee           int       `json:"employee" gorm:"employee"`
	CreateUid          int       `json:"createUid" gorm:"create_uid"`
	WriteUid           int       `json:"writeUid" gorm:"update_uid"`
	Code               string    `json:"code" gorm:"code"`
	PractisingCateg    string    `json:"practisingCateg" gorm:"practising_categ"`
	PractisingRange    string    `json:"practisingRange" gorm:"practising_range"`
	PractisingLocation string    `json:"practisingLocation" gorm:"practising_location"`
	RegisterDate       time.Time `json:"registerDate" gorm:"register_date"`
	ApprovalEmployeeId int       `json:"approvalEmployeeId" gorm:"approval_employee_id"`
	Approval           bool      `json:"approval" gorm:"approval"`
	IsApproval         bool      `json:"isApproval" gorm:"is_approval"`
	ApprovalTime       time.Time `json:"approvalTime" gorm:"approval_time"`
	RevokeEmployeeId   int       `json:"revokeEmployeeId" gorm:"revoke_employee_id"`
	RevokeReason       string    `json:"revokeReason" gorm:"revoke_reason"`
	Revoke             bool      `json:"revoke" gorm:"revoke"`
	RevokeTime         time.Time `json:"revokeTime" gorm:"revoke_time"`
}
type HrEmployeeSpecialist struct {
	global.GVA_MODEL
	CreateUid int    `json:"createUid" gorm:"create_uid"`
	WriteUid  int    `json:"writeUid" gorm:"update_uid"`
	Name      string `json:"name" gorm:"name"`
}
type HrWorkLocation struct {
	global.GVA_MODEL
	CreateUid      int    `json:"createUid" gorm:"create_uid"`
	WriteUid       int    `json:"writeUid" gorm:"update_uid"`
	CompanyId      int    `json:"companyId" gorm:"company_id"`
	AddressId      int    `json:"addressId" gorm:"address_id"`
	Name           string `json:"name" gorm:"name"`
	LocationNumber string `json:"locationNumber" gorm:"location_number"`
	Active         bool   `json:"active" gorm:"active"`
}
type IrAttachment struct {
	global.GVA_MODEL
	CreateUid    int    `json:"createUid" gorm:"create_uid"`
	WriteUid     int    `json:"writeUid" gorm:"update_uid"`
	ResId        int    `json:"resId" gorm:"res_id"`
	CompanyId    int    `json:"companyId" gorm:"company_id"`
	FileSize     int    `json:"fileSize" gorm:"file_size"`
	Name         string `json:"name" gorm:"name"`
	ResModel     string `json:"resModel" gorm:"res_model"`
	ResField     string `json:"resField" gorm:"res_field"`
	Type         string `json:"type" gorm:"type"`
	Url          string `json:"url" gorm:"url"`
	AccessToken  string `json:"accessToken" gorm:"access_token"`
	StoreFname   string `json:"storeFname" gorm:"store_fname"`
	Checksum     string `json:"checksum" gorm:"checksum"`
	Mimetype     string `json:"mimetype" gorm:"mimetype"`
	Description  string `json:"description" gorm:"description"`
	IndexContent string `json:"indexContent" gorm:"index_content"`
	Public       bool   `json:"public" gorm:"public"`
	DbDatas      []byte `json:"dbDatas" gorm:"db_dates"`
	OriginalId   int    `json:"originalId" gorm:"original_id"`
	S3flag       bool   `json:"s3Flag" gorm:"s3flag"`
}
type HrEmployeePrescription struct {
	global.GVA_MODEL
	CreateUid int    `json:"createUid" gorm:"create_uid"`
	WriteUid  int    `json:"writeUid" gorm:"update_uid"`
	Name      string `json:"name" gorm:"name"`
	Active    bool   `json:"active" gorm:"active"`
}
type HrEmployeeCategory struct {
	global.GVA_MODEL
	CreateUid int    `json:"createUid" gorm:"create_uid"`
	WriteUid  int    `json:"writeUid" gorm:"update_uid"`
	Color     int    `json:"color" gorm:"color"`
	Name      string `json:"name" gorm:"name"`
	Active    bool   `json:"active" gorm:"active"`
}
type HrEmployeeDocument struct {
	global.GVA_MODEL
	EmployeeRef      int       `json:"employeeRef" gorm:"employee_ref"`
	DocumentType     int       `json:"documentType" gorm:"document_type"`
	BeforeDays       int       `json:"beforeDays" gorm:"before_days"`
	CreateUid        int       `json:"createUid" gorm:"create_uid"`
	WriteUid         int       `json:"writeUid" gorm:"update_uid"`
	Name             string    `json:"name" gorm:"name"`
	NotificationType string    `json:"notificationType" gorm:"notification_type"`
	ExpiryDate       time.Time `json:"expiryDate" gorm:"expiry_date"`
	IssueDate        time.Time `json:"issueDate" gorm:"issue_date"`
	Description      string    `json:"description" gorm:"description"`
}
type HrEmployeeFamily struct {
	global.GVA_MODEL
	CreateUid     int       `json:"createUid" gorm:"create_uid"`
	WriteUid      int       `json:"writeUid" gorm:"update_uid"`
	Job           string    `json:"job" gorm:"job"`
	BrithDate     time.Time `json:"brithDate" gorm:"brith_date"`
	RelationId    int       `json:"relationId" gorm:"relation_id"`
	EmployeeId    int       `json:"employeeId" gorm:"employee_id"`
	MemberContact string    `json:"memberContact" gorm:"member_contact"`
	MemberName    string    `json:"memberName" gorm:"member_name"`
}
type HrAnnouncement struct {
	global.GVA_MODEL
	CreateUid               int       `json:"createUid" gorm:"create_uid"`
	WriteUid                int       `json:"writeUid" gorm:"update_uid"`
	MessageMainAttachmentId int       `json:"messageMainAttachmentId" gorm:"message_main_attachment_id"`
	CompanyId               int       `json:"companyId" gorm:"company_id"`
	Name                    string    `json:"name" gorm:"name"`
	State                   string    `json:"state" gorm:"state"`
	AnnouncementType        string    `json:"announcementType" gorm:"announcement_type"`
	RequestedDate           time.Time `json:"requestedDate" gorm:"requested_date"`
	DateStart               time.Time `json:"dateStart" gorm:"date_start"`
	DateEnd                 time.Time `json:"dateEnd" gorm:"date_end"`
	AnnouncementReason      string    `json:"announcementReason" gorm:"announcement_reason"`
	Announcement            string    `json:"announcement" gorm:"announcement"`
	IsAnnouncement          bool      `json:"isAnnouncement" gorm:"is_announcement"`
}

type HrApplicant struct {
	global.GVA_MODEL
	CreateUid  int `json:"createUid" gorm:"create_uid"`
	WriteUid   int `json:"writeUid" gorm:"update_uid"`
	CampaignId int `json:"campaignId" gorm:"campaign_id"`
	SourceId   int `json:"sourceId" gorm:"source_id"`
	MediumId   int `json:"mediumId" gorm:"medium_id"`
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
