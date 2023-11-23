package hospitalManage

import "time"

type HospitalEmployee struct {
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

type HrEmployee2 struct {
	ResourceId                  int       `json:"resourceId" form:"resourceId" gorm:"column:resource_id;comment:Resource;size:32;"`
	CompanyId                   int       `json:"companyId" form:"companyId" gorm:"column:company_id;comment:Company;size:32;"`
	ResourceCalendarId          int       `json:"resourceCalendarId" form:"resourceCalendarId" gorm:"column:resource_calendar_id;comment:Working Hours;size:32;"`
	MessageMainAttachmentId     int       `json:"messageMainAttachmentId" form:"messageMainAttachmentId" gorm:"column:message_main_attachment_id;comment:Main Attachment;size:32;"`
	Color                       int       `json:"color" form:"color" gorm:"column:color;comment:Color Index;size:32;"`
	DepartmentId                int       `json:"departmentId" form:"departmentId" gorm:"column:department_id;comment:Department;size:32;"`
	JobId                       int       `json:"jobId" form:"jobId" gorm:"column:job_id;comment:Job Position;size:32;"`
	AddressId                   int       `json:"addressId" form:"addressId" gorm:"column:address_id;comment:Work Address;size:32;"`
	WorkContactId               int       `json:"workContactId" form:"workContactId" gorm:"column:work_contact_id;comment:Work Contact;size:32;"`
	WorkLocationId              int       `json:"workLocationId" form:"workLocationId" gorm:"column:work_location_id;comment:Work Location;size:32;"`
	UserId                      int       `json:"userId" form:"userId" gorm:"column:user_id;comment:User;size:32;"`
	ParentId                    int       `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:Manager;size:32;"`
	CoachId                     int       `json:"coachId" form:"coachId" gorm:"column:coach_id;comment:Coach;size:32;"`
	AddressHomeId               int       `json:"addressHomeId" form:"addressHomeId" gorm:"column:address_home_id;comment:Address;size:32;"`
	CountryId                   int       `json:"countryId" form:"countryId" gorm:"column:country_id;comment:Nationality (Country);size:32;"`
	Children                    int       `json:"children" form:"children" gorm:"column:children;comment:Number of Dependent Children;size:32;"`
	CountryOfBirth              int       `json:"countryOfBirth" form:"countryOfBirth" gorm:"column:country_of_birth;comment:Country of Birth;size:32;"`
	BankAccountId               int       `json:"bankAccountId" form:"bankAccountId" gorm:"column:bank_account_id;comment:Bank Account Number;size:32;"`
	KmHomeWork                  int       `json:"kmHomeWork" form:"kmHomeWork" gorm:"column:km_home_work;comment:Home-Work Distance;size:32;"`
	DepartureReasonId           int       `json:"departureReasonId" form:"departureReasonId" gorm:"column:departure_reason_id;comment:Departure Reason;size:32;"`
	CreateUid                   int       `json:"createUid" form:"createUid" gorm:"column:create_uid;comment:Created by;size:32;"`
	WriteUid                    int       `json:"writeUid" form:"writeUid" gorm:"column:write_uid;comment:Last Updated by;size:32;"`
	Name                        string    `json:"name" form:"name" gorm:"column:name;comment:Employee Name;"`
	JobTitle                    string    `json:"jobTitle" form:"jobTitle" gorm:"column:job_title;comment:Job Title;"`
	WorkPhone                   string    `json:"workPhone" form:"workPhone" gorm:"column:work_phone;comment:Work Phone;"`
	MobilePhone                 string    `json:"mobilePhone" form:"mobilePhone" gorm:"column:mobile_phone;comment:Work Mobile;"`
	WorkEmail                   string    `json:"workEmail" form:"workEmail" gorm:"column:work_email;comment:Work Email;"`
	EmployeeType                string    `json:"employeeType" form:"employeeType" gorm:"column:employee_type;comment:Employee Type;"`
	Gender                      int       `json:"gender" form:"gender" gorm:"column:gender;comment:Gender;size:32;"`
	Marital                     int       `json:"marital" form:"marital" gorm:"column:marital;comment:Marital Status;size:32;"`
	SpouseCompleteName          string    `json:"spouseCompleteName" form:"spouseCompleteName" gorm:"column:spouse_complete_name;comment:Spouse Complete Name;"`
	PlaceOfBirth                string    `json:"placeOfBirth" form:"placeOfBirth" gorm:"column:place_of_birth;comment:Place of Birth;"`
	Ssnid                       string    `json:"ssnid" form:"ssnid" gorm:"column:ssnid;comment:SSN No;"`
	Sinid                       string    `json:"sinid" form:"sinid" gorm:"column:sinid;comment:SIN No;"`
	IdentificationId            string    `json:"identificationId" form:"identificationId" gorm:"column:identification_id;comment:Identification No;"`
	PassportId                  string    `json:"passportId" form:"passportId" gorm:"column:passport_id;comment:Passport No;"`
	PermitNo                    string    `json:"permitNo" form:"permitNo" gorm:"column:permit_no;comment:Work Permit No;"`
	VisaNo                      string    `json:"visaNo" form:"visaNo" gorm:"column:visa_no;comment:Visa No;"`
	Certificate                 string    `json:"certificate" form:"certificate" gorm:"column:certificate;comment:Certificate Level;"`
	StudyField                  string    `json:"studyField" form:"studyField" gorm:"column:study_field;comment:Field of Study;"`
	StudySchool                 string    `json:"studySchool" form:"studySchool" gorm:"column:study_school;comment:School;"`
	EmergencyContact            string    `json:"emergencyContact" form:"emergencyContact" gorm:"column:emergency_contact;comment:Contact Name;"`
	EmergencyPhone              string    `json:"emergencyPhone" form:"emergencyPhone" gorm:"column:emergency_phone;comment:Contact Phone;"`
	Barcode                     string    `json:"barcode" form:"barcode" gorm:"column:barcode;comment:Badge ID;"`
	Pin                         string    `json:"pin" form:"pin" gorm:"column:pin;comment:PIN;"`
	SpouseBirthdate             time.Time `json:"spouseBirthdate" form:"spouseBirthdate" gorm:"column:spouse_birthdate;comment:Spouse Birthdate;"`
	Birthday                    time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:Date of Birth;"`
	VisaExpire                  time.Time `json:"visaExpire" form:"visaExpire" gorm:"column:visa_expire;comment:Visa Expire Date;"`
	WorkPermitExpirationDate    time.Time `json:"workPermitExpirationDate" form:"workPermitExpirationDate" gorm:"column:work_permit_expiration_date;comment:Work Permit Expiration Date;"`
	DepartureDate               time.Time `json:"departureDate" form:"departureDate" gorm:"column:departure_date;comment:Departure Date;"`
	AdditionalNote              string    `json:"additionalNote" form:"additionalNote" gorm:"column:additional_note;comment:Additional Note;"`
	Notes                       string    `json:"notes" form:"notes" gorm:"column:notes;comment:Notes;"`
	DepartureDescription        string    `json:"departureDescription" form:"departureDescription" gorm:"column:departure_description;comment:Additional Information;"`
	Active                      bool      `json:"active" form:"active" gorm:"column:active;comment:Active;"`
	WorkPermitScheduledActivity bool      `json:"workPermitScheduledActivity" form:"workPermitScheduledActivity" gorm:"column:work_permit_scheduled_activity;comment:Work Permit Scheduled Activity;"`
	CreateDate                  time.Time `json:"createDate" form:"createDate" gorm:"column:create_date;comment:Created on;size:6;"`
	WriteDate                   time.Time `json:"writeDate" form:"writeDate" gorm:"column:write_date;comment:Last Updated on;size:6;"`
	LastAttendanceId            int       `json:"lastAttendanceId" form:"lastAttendanceId" gorm:"column:last_attendance_id;comment:Last Attendance;size:32;"`
	LastCheckIn                 time.Time `json:"lastCheckIn" form:"lastCheckIn" gorm:"column:last_check_in;comment:Check In;size:6;"`
	LastCheckOut                time.Time `json:"lastCheckOut" form:"lastCheckOut" gorm:"column:last_check_out;comment:Check Out;size:6;"`
	ContractId                  int       `json:"contractId" form:"contractId" gorm:"column:contract_id;comment:Current Contract;size:32;"`
	Vehicle                     string    `json:"vehicle" form:"vehicle" gorm:"column:vehicle;comment:Company Vehicle;"`
	FirstContractDate           time.Time `json:"firstContractDate" form:"firstContractDate" gorm:"column:first_contract_date;comment:First Contract Date;"`
	ContractWarning             bool      `json:"contractWarning" form:"contractWarning" gorm:"column:contract_warning;comment:Contract Warning;"`
	LeaveManagerId              int       `json:"leaveManagerId" form:"leaveManagerId" gorm:"column:leave_manager_id;comment:Time Off;size:32;"`
	HourlyCost                  float64   `json:"hourlyCost" form:"hourlyCost" gorm:"column:hourly_cost;comment:Hourly Cost;"`
	PersonalMobile              string    `json:"personalMobile" form:"personalMobile" gorm:"column:personal_mobile;comment:Mobile;"`
	JoiningDate                 time.Time `json:"joiningDate" form:"joiningDate" gorm:"column:joining_date;comment:Joining Date;"`
	IdExpiryDate                time.Time `json:"idExpiryDate" form:"idExpiryDate" gorm:"column:id_expiry_date;comment:Expiry Date;"`
	PassportExpiryDate          time.Time `json:"passportExpiryDate" form:"passportExpiryDate" gorm:"column:passport_expiry_date;comment:Expiry Date;"`
	ResignDate                  time.Time `json:"resignDate" form:"resignDate" gorm:"column:resign_date;comment:Resign Date;"`
	Resigned                    bool      `json:"resigned" form:"resigned" gorm:"column:resigned;comment:Resigned;"`
	Fired                       bool      `json:"fired" form:"fired" gorm:"column:fired;comment:Fired;"`
	ExpenseManagerId            int       `json:"expenseManagerId" form:"expenseManagerId" gorm:"column:expense_manager_id;comment:Expense;size:32;"`
	ResourceCalendarIds         int       `json:"resourceCalendarIds" form:"resourceCalendarIds" gorm:"column:resource_calendar_ids;comment:Working Hours;size:32;"`
	Nation                      int       `json:"nation" form:"nation" gorm:"column:nation;comment:民族;size:32;"`
	Schoolgrade                 int       `json:"schoolgrade" form:"schoolgrade" gorm:"column:schoolgrade;comment:学历;size:32;"`
	Appellation                 int       `json:"appellation" form:"appellation" gorm:"column:appellation;comment:职称;size:32;"`
	BookingNumAm                int       `json:"bookingNumAm" form:"bookingNumAm" gorm:"column:booking_num_am;comment:上午预约最大人数;size:32;"`
	BookingNumPm                int       `json:"bookingNumPm" form:"bookingNumPm" gorm:"column:booking_num_pm;comment:下午预约最大人数;size:32;"`
	BookingNumNg                int       `json:"bookingNumNg" form:"bookingNumNg" gorm:"column:booking_num_ng;comment:晚上预约最大人数;size:32;"`
	PoliticsStatus              string    `json:"politicsStatus" form:"politicsStatus" gorm:"column:politics_status;comment:政治面貌;"`
	Tocompile                   string    `json:"tocompile" form:"tocompile" gorm:"column:tocompile;comment:编制;"`
	AppellationCode             string    `json:"appellationCode" form:"appellationCode" gorm:"column:appellation_code;comment:职称证书编号;"`
	Rank                        string    `json:"rank" form:"rank" gorm:"column:rank;comment:职级;"`
	AppellationTime             time.Time `json:"appellationTime" form:"appellationTime" gorm:"column:appellation_time;comment:职称获得日期;"`
	Note                        string    `json:"note" form:"note" gorm:"column:note;comment:说明;"`
	Introduction                string    `json:"introduction" form:"introduction" gorm:"column:introduction;comment:医生简介;"`
	GradeId                     int       `json:"gradeId" form:"gradeId" gorm:"column:grade_id;comment:Grade;size:32;"`
	RankId                      int       `json:"rankId" form:"rankId" gorm:"column:rank_id;comment:Rank;size:32;"`
	NaturePosition              string    `json:"naturePosition" form:"naturePosition" gorm:"column:nature_position;comment:岗位性质;"`
	SurgeryCert                 string    `json:"surgeryCert" form:"surgeryCert" gorm:"column:surgery_cert;comment:手术资质;"`
	Inductiondate               time.Time `json:"inductiondate" form:"inductiondate" gorm:"column:inductiondate;comment:入职日期;"`
	Performanceclass            string    `json:"performanceclass" form:"performanceclass" gorm:"column:performanceclass;comment:绩效优先类别;"`
	EduAppellation              string    `json:"eduAppellation" form:"eduAppellation" gorm:"column:edu_appellation;comment:教学职称;"`
	EduCategory                 string    `json:"eduCategory" form:"eduCategory" gorm:"column:edu_category;comment:导师类型;"`
	PhysicianCert               int       `json:"physicianCert" form:"physicianCert" gorm:"column:physician_cert;comment:执业医师资质;size:32;"`
	PharmacistCert              int       `json:"pharmacistCert" form:"pharmacistCert" gorm:"column:pharmacist_cert;comment:执业药师资质;size:32;"`
	NurseCert                   int       `json:"nurseCert" form:"nurseCert" gorm:"column:nurse_cert;comment:执业护士资质;size:32;"`
	OperateCert                 string    `json:"operateCert" form:"operateCert" gorm:"column:operate_cert;comment:操作资质;"`
	PrescriptionCert            string    `json:"prescriptionCert" form:"prescriptionCert" gorm:"column:prescription_cert;comment:处方资质;"`
	SpecialistId                int       `json:"specialistId" form:"specialistId" gorm:"column:specialist_id;comment:专科护士类别;size:32;"`
	NurseSpe                    string    `json:"nurseSpe" form:"nurseSpe" gorm:"column:nurse_spe;comment:护士能级;"`
	SpecialistDate              time.Time `json:"specialistDate" form:"specialistDate" gorm:"column:specialist_date;comment:护士能级获取时间;"`
	IsSpecialist                bool      `json:"isSpecialist" form:"isSpecialist" gorm:"column:is_specialist;comment:是否为专科护士;"`
	EduSchool                   string    `json:"eduSchool" form:"eduSchool" gorm:"column:edu_school;comment:教学单位;"`
	SurgeryGrade                int       `json:"surgeryGrade" form:"surgeryGrade" gorm:"column:surgery_grade;comment:手术级别;size:32;"`
	PracticeCateg               string    `json:"practiceCateg" form:"practiceCateg" gorm:"column:practice_categ;comment:执业类别;"`
	PhysicianAidCert            int       `json:"physicianAidCert" form:"physicianAidCert" gorm:"column:physician_aid_cert;comment:执业助理医师资质;size:32;"`
	NativePlaceP                int       `json:"nativePlaceP" form:"nativePlaceP" gorm:"column:native_place_p;comment:籍贯(省);size:32;"`
	NativePlaceC                int       `json:"nativePlaceC" form:"nativePlaceC" gorm:"column:native_place_c;comment:籍贯(市);size:32;"`
	GraduationDate              time.Time `json:"graduationDate" form:"graduationDate" gorm:"column:graduation_date;comment:毕业时间;"`
}

func (h *HospitalEmployee) Table() string {
	return "hospital_employees"
}
