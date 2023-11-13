package response

import (
	"github.com/shopspring/decimal"
	"time"
)

// 查看普通信息
type BasicStaffInfo struct {
	ID                          int             `json:"id" gorm:"column:id"`
	Name                        string          `json:"name" gorm:"column:name"`
	ResourceId                  int             `json:"resource_id" gorm:"column:resource_id"`
	CompanyId                   int             `json:"companyId" gorm:"column:company_id"`
	ResourceCalendarId          int             `json:"resourceCalendarId" gorm:"column:resource_calendar_id"`
	MessageMainAttachmentId     int             `json:"messageMainAttachmentId" gorm:"message_main_attachment_id"`
	Color                       int             `json:"color" gorm:"column:color"`
	DepartmentId                int             `json:"departmentId" gorm:"column:department_id"`
	JobId                       int             `json:"jobId" gorm:"column: job_id"`
	AddressId                   int             `json:"addressId" gorm:"column:address_id"`
	WorkContactId               int             `json:"workContactId" gorm:"column:work_contact_id"`
	WorkLocationId              int             `json:"workLocationId" gorm:"column:work_location_id"`
	UserId                      int             `json:"userId" gorm:"column:user_id"`
	ParentId                    int             `json:"parentId" gorm:"column:parent_id"`
	CoachId                     int             `json:"coachId" gorm:"column:coach_id"`
	AddressHomeId               int             `json:"addressHomeId" gorm:"column:address_home_id"`
	CountryId                   int             `json:"countryId" gorm:"column:country_id"`
	Children                    int             `json:"children" gorm:"column:children"`
	CountryOfBirth              int             `json:"countryOfBirth" gorm:"column:country_of_birth"`
	BankAccountId               int             `json:"bankAccountId" gorm:"column:bank_account_id"`
	KmHomeWork                  int             `json:"kmHomeWork" gorm:"column:km_home_work"`
	DepartureReasonId           int             `json:"departureReasonId" gorm:"column:departure_reason_id"`
	CreateUid                   int             `json:"createUid" gorm:"column:create_uid"`
	WriteUid                    int             `json:"writeUid" gorm:"column:write_uid"`
	JobTitle                    string          `json:"jobTitle" gorm:"column:job_title"`
	WorkPhone                   string          `json:"workPhone" gorm:"column:work_phone"`
	MobilePhone                 string          `json:"mobilePhone" gorm:"column:mobile_phone"`
	WorkEmail                   string          `json:"workEmail" gorm:"column:work_email"`
	EmployeeType                string          `json:"employeeType" gorm:"column:employee_type"`
	Gender                      int             `json:"gender" gorm:"column:gender"`
	Marital                     int             `json:"marital" gorm:"column:marital"`
	SpouseCompleteName          string          `json:"spouseCompleteName" gorm:"column:spouse_complete_name"`
	PlaceOfBirth                string          `json:"placeOfBirth" gorm:"column:place_of_birth"` // 出生地
	Ssnid                       string          `json:"ssnid" gorm:"column:ssnid"`
	Sinid                       string          `json:"sinid" gorm:"column:sinid"`
	IdentificationId            string          `json:"identificationId" gorm:"column:identification_id"` //身份证
	PassportId                  string          `json:"passportId" gorm:"column:passport_id"`
	PermitNo                    string          `json:"permitNo" gorm:"column:permit_no"`
	VisaNo                      string          `json:"visaNo" gorm:"column:visa_no"`
	Certificate                 string          `json:"certificate" gorm:"column:certificate"`
	StudyField                  string          `json:"studyField" gorm:"column:study_field"`
	StudySchool                 string          `json:"studySchool" gorm:"column:study_school"` // 毕业学校
	EmergencyContact            string          `json:"emergencyContact" gorm:"column:emergency_contact"`
	EmergencyPhone              string          `json:"emergencyPhone" gorm:"column:emergency_phone"`
	Barcode                     string          `json:"barcode" gorm:"column:barcode"`
	Pin                         string          `json:"pin" gorm:"column:barcode"`
	SpouseBirthdate             time.Time       `json:"spouseBirthdate" gorm:"column:spouse_birthdate"`
	Birthday                    time.Time       `json:"birthday" gorm:"column:birthday"` // 出生日期
	VisaExpire                  time.Time       `json:"visaExpire" gorm:"column:visa_expire"`
	WorkPermitExpirationDate    time.Time       `json:"workPermitExpirationDate" gorm:"column:work_permit_expiration_date"`
	DepartureDate               time.Time       `json:"departureDate" gorm:"column:departure_date"`
	AdditionalNote              string          `json:"additionalNote" gorm:"column:additional_note"`
	Notes                       string          `json:"notes" gorm:"column:notes"`
	DepartureDescription        string          `json:"departureDescription" gorm:"column:departure_description"`
	Active                      bool            `json:"active" gorm:"column:active"`
	WorkPermitScheduledActivity bool            `json:"workPermitScheduledActivity" gorm:"column:work_permit_scheduled_activity"`
	CreateDate                  time.Time       `json:"createDate" gorm:"column:create_date"`
	WriteDate                   time.Time       `json:"writeDate" gorm:"column:write_date"`
	LastAttendanceId            int             `json:"lastAttendanceId" gorm:"column:last_attendance_id"`
	LastCheckIn                 time.Time       `json:"lastCheckIn" gorm:"column:last_check_in"`
	LastCheckOut                time.Time       `json:"lastCheckOut" gorm:"column:last_check_out"`
	ContractId                  int             `json:"contractId" gorm:"column:contract_id"`
	Vehicle                     string          `json:"vehicle" gorm:"column:vehicle"`
	FirstContractDate           time.Time       `json:"firstContractDate" gorm:"column:first_contract_date"`
	ContractWarning             bool            `json:"contractWarning" gorm:"column:contract_warning"`
	LeaveManagerId              int             `json:"leaveManagerId" gorm:"column:leave_manager_id"`
	HourlyCost                  decimal.Decimal `json:"hourlyCost" gorm:"column:hourly_cost"`
	PersonalMobile              string          `json:"personalMobile" gorm:"column:personal_mobile"`
	JoiningDate                 time.Time       `json:"joiningDate" gorm:"column:joining_date"`
	IdExpiryDate                time.Time       `json:"idExpiryDate" gorm:"column:id_expiry_date"`
	PassportExpiryDate          time.Time       `json:"passportExpiryDate" gorm:"column:passport_expiry_date"`
	ResignDate                  time.Time       `json:"resignDate" gorm:"column:resign_date"`
	Resigned                    bool            `json:"resigned" gorm:"column:resigned"`
	Fired                       bool            `json:"fired" gorm:"column:fired"`
	ExpenseManagerId            int             `json:"expenseManagerId" gorm:"column:expense_manager_id"`
	ResourceCalendarIds         int             `json:"resourceCalendarIds" gorm:"column:resource_calendar_ids"`
	Nation                      int             `json:"nation" gorm:"column:nation"`
	Schoolgrade                 int             `json:"schoolgrade" gorm:"column:schoolgrade"`
	Appellation                 int             `json:"appellation" gorm:"column:appellation"`
	BookingNumAm                int             `json:"bookingNumAm" gorm:"column:booking_num_am"`
	BookingNumPm                int             `json:"bookingNumPm" gorm:"column:booking_num_pm"`
	BookingNumNg                int             `json:"bookingNumNg" gorm:"column:booking_num_ng"`
	PoliticsStatus              string          `json:"politicsStatus" gorm:"column:politics_status"`
	Tocompile                   string          `json:"tocompile" gorm:"column:tocompile"`
	AppellationCode             string          `json:"appellationCode" gorm:"column:appellation_code"`
	Rank                        string          `json:"rank" gorm:"column:rank"`
	AppellationTime             time.Time       `json:"appellationTime" gorm:"column:appellation_time"`
	Note                        string          `json:"note" gorm:"column:note"`
	Introduction                string          `json:"introduction" gorm:"column:introduction"`
	GradeId                     int             `json:"gradeId" gorm:"column:grade_id"`
	RankId                      int             `json:"rankId" gorm:"column:rank_id"`
	NaturePosition              string          `json:"naturePosition" gorm:"column:nature_position"`
	SurgeryCert                 string          `json:"surgeryCert" gorm:"column:surgery_cert"`
	Inductiondate               time.Time       `json:"inductiondate" gorm:"column:inductiondate"`
	Performanceclass            string          `json:"performanceclass" gorm:"column:performanceclass"`
	EduAppellation              string          `json:"eduAppellation" gorm:"column:edu_appellation"`
	EduCategory                 string          `json:"eduCategory" gorm:"column:edu_category"`
	PhysicianCert               int             `json:"physicianCert" gorm:"column:physician_cert"`
	PharmacistCert              int             `json:"pharmacistCert" gorm:"column:pharmacist_cert"`
	NurseCert                   int             `json:"nurseCert" gorm:"column:nurse_cert"`
	OperateCert                 string          `json:"operateCert" gorm:"column:operate_cert"`
	PrescriptionCert            string          `json:"prescriptionCert" gorm:"column:prescription_cert"`
	SpecialistId                int             `json:"specialistId" gorm:"column:specialist_id"`
	NurseSpe                    string          `json:"nurseSpe" gorm:"column:nurse_spe"`
	SpecialistDate              time.Time       `json:"specialistDate" gorm:"column:specialist_date"`
	IsSpecialist                bool            `json:"isSpecialist" gorm:"column:is_specialist"`
	EduSchool                   string          `json:"eduSchool" gorm:"column:edu_school"`
	SurgeryGrade                int             `json:"surgeryGrade" gorm:"column:surgery_grade"`
	PracticeCateg               string          `json:"practiceCateg" gorm:"column:practice_categ"`
	PhysicianAidCert            int             `json:"physicianAidCert" gorm:"column:physician_aid_cert"`
	NativePlaceP                int             `json:"nativePlaceP" gorm:"column:native_place_p"`
	NativePlaceC                int             `json:"nativePlaceC" gorm:"column:native_place_c"`
	GraduationDate              time.Time       `json:"graduationDate" gorm:"column:graduation_date"`
}
