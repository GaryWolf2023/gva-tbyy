package hospitalManage

import "time"

type HrDepartment struct {
	ID                      int       `json:"id" gorm:"column:id"`
	CreateDate              time.Time `json:"createDate" gorm:"column:create_date"`
	WriteDate               time.Time `json:"writeDate" gorm:"column:write_date"`
	CreateUid               int       `json:"createUid" gorm:"column:create_uid"`
	WriteUid                int       `json:"writeUid" gorm:"column:update_uid"`
	MessageMainAttachmentId int       `json:"messageMainAttachmentId" gorm:"column:message_main_attachment_id"`
	CompanyId               int       `json:"companyId" gorm:"column:company_id"`
	ParentId                int       `json:"parentId" gorm:"column:parent_id"`
	ManagerId               int       `json:"managerId" gorm:"column:manager_id"`
	Color                   int       `json:"color" gorm:"column:color"`
	MasterDepartmentId      int       `json:"masterDepartmentId" gorm:"column:master_department_id"`
	Name                    string    `json:"name" gorm:"column:name"`
	CompleteName            string    `json:"completeName" gorm:"column:complete_name"`
	ParentPath              string    `json:"parentPath" gorm:"column:parent_path"`
	Note                    string    `json:"note" gorm:"column:note"`
	Active                  bool      `json:"active" gorm:"column:active"`
	Property                int       `json:"property" gorm:"column:property"`
	Virtual                 bool      `json:"virtual" gorm:"column:virtual"`
	Available               bool      `json:"available" gorm:"column:available"`
	DutyId                  int       `json:"dutyId" gorm:"column:duty_id"`
	Code                    string    `json:"code" gorm:"column:code"`
}
