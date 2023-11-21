package hospitalManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

// ResCompany 结构体
type ResCompany struct {
	global.GVA_MODEL
	Name                                        string         `json:"name" form:"name" gorm:"column:name;comment:;"`
	PartnerId                                   int            `json:"partnerId" form:"partnerId" gorm:"column:partner_id;comment:;size:32;"`
	CurrencyId                                  int            `json:"currencyId" form:"currencyId" gorm:"column:currency_id;comment:;size:32;"`
	Sequence                                    int            `json:"sequence" form:"sequence" gorm:"column:sequence;comment:;size:32;"`
	CreateDate                                  *time.Time     `json:"createDate" form:"createDate" gorm:"column:create_date;comment:;size:6;"`
	ParentId                                    int            `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:Parent Company;size:32;"`
	PaperformatId                               int            `json:"paperformatId" form:"paperformatId" gorm:"column:paperformat_id;comment:Paper format;size:32;"`
	ExternalReportLayoutId                      int            `json:"externalReportLayoutId" form:"externalReportLayoutId" gorm:"column:external_report_layout_id;comment:Document Template;size:32;"`
	CreateUid                                   int            `json:"createUid" form:"createUid" gorm:"column:create_uid;comment:Created by;size:32;"`
	WriteUid                                    int            `json:"writeUid" form:"writeUid" gorm:"column:write_uid;comment:Last Updated by;size:32;"`
	Email                                       string         `json:"email" form:"email" gorm:"column:email;comment:Email;"`
	Phone                                       string         `json:"phone" form:"phone" gorm:"column:phone;comment:Phone;"`
	Mobile                                      string         `json:"mobile" form:"mobile" gorm:"column:mobile;comment:Mobile;"`
	BaseOnboardingCompanyState                  string         `json:"baseOnboardingCompanyState" form:"baseOnboardingCompanyState" gorm:"column:base_onboarding_company_state;comment:State of the onboarding company step;"`
	Font                                        string         `json:"font" form:"font" gorm:"column:font;comment:Font;"`
	PrimaryColor                                string         `json:"primaryColor" form:"primaryColor" gorm:"column:primary_color;comment:Primary Color;"`
	SecondaryColor                              string         `json:"secondaryColor" form:"secondaryColor" gorm:"column:secondary_color;comment:Secondary Color;"`
	LayoutBackground                            string         `json:"layoutBackground" form:"layoutBackground" gorm:"column:layout_background;comment:Layout Background;"`
	ReportFooter                                string         `json:"reportFooter" form:"reportFooter" gorm:"column:report_footer;comment:Report Footer;"`
	ReportHeader                                string         `json:"reportHeader" form:"reportHeader" gorm:"column:report_header;comment:Company Tagline;"`
	CompanyDetails                              string         `json:"companyDetails" form:"companyDetails" gorm:"column:company_details;comment:Company Details;"`
	Active                                      bool           `json:"active" form:"active" gorm:"column:active;comment:Active;"`
	WriteDate                                   *time.Time     `json:"writeDate" form:"writeDate" gorm:"column:write_date;comment:Last Updated on;size:6;"`
	LogoWeb                                     []byte         `json:"logoWeb" form:"logoWeb" gorm:"column:logo_web;comment:Logo Web;"`
	NomenclatureId                              int            `json:"nomenclatureId" form:"nomenclatureId" gorm:"column:nomenclature_id;comment:Nomenclature;size:32;"`
	ResourceCalendarId                          int            `json:"resourceCalendarId" form:"resourceCalendarId" gorm:"column:resource_calendar_id;comment:Default Working Hours;size:32;"`
	HrPresenceControlEmailAmount                int            `json:"hrPresenceControlEmailAmount" form:"hrPresenceControlEmailAmount" gorm:"column:hr_presence_control_email_amount;comment:# emails to send;size:32;"`
	HrPresenceControlIpList                     string         `json:"hrPresenceControlIpList" form:"hrPresenceControlIpList" gorm:"column:hr_presence_control_ip_list;comment:Valid IP addresses;"`
	OvertimeCompanyThreshold                    int            `json:"overtimeCompanyThreshold" form:"overtimeCompanyThreshold" gorm:"column:overtime_company_threshold;comment:Tolerance Time In Favor Of Company;size:32;"`
	OvertimeEmployeeThreshold                   int            `json:"overtimeEmployeeThreshold" form:"overtimeEmployeeThreshold" gorm:"column:overtime_employee_threshold;comment:Tolerance Time In Favor Of Employee;size:32;"`
	AttendanceKioskDelay                        int            `json:"attendanceKioskDelay" form:"attendanceKioskDelay" gorm:"column:attendance_kiosk_delay;comment:Attendance Kiosk Delay;size:32;"`
	AttendanceKioskMode                         string         `json:"attendanceKioskMode" form:"attendanceKioskMode" gorm:"column:attendance_kiosk_mode;comment:Attendance Mode;"`
	AttendanceBarcodeSource                     string         `json:"attendanceBarcodeSource" form:"attendanceBarcodeSource" gorm:"column:attendance_barcode_source;comment:Barcode Source;"`
	OvertimeStartDate                           *time.Time     `json:"overtimeStartDate" form:"overtimeStartDate" gorm:"column:overtime_start_date;comment:Extra Hours Starting Date;"`
	HrAttendanceOvertime                        bool           `json:"hrAttendanceOvertime" form:"hrAttendanceOvertime" gorm:"column:hr_attendance_overtime;comment:Count Extra Hours;"`
	PartnerGid                                  int            `json:"partnerGid" form:"partnerGid" gorm:"column:partner_gid;comment:Company database ID;size:32;"`
	IapEnrichAutoDone                           bool           `json:"iapEnrichAutoDone" form:"iapEnrichAutoDone" gorm:"column:iap_enrich_auto_done;comment:Enrich Done;"`
	SnailmailColor                              bool           `json:"snailmailColor" form:"snailmailColor" gorm:"column:snailmail_color;comment:Color;"`
	SnailmailCover                              bool           `json:"snailmailCover" form:"snailmailCover" gorm:"column:snailmail_cover;comment:Add a Cover Page;"`
	SnailmailDuplex                             bool           `json:"snailmailDuplex" form:"snailmailDuplex" gorm:"column:snailmail_duplex;comment:Both sides;"`
	PaymentProviderOnboardingState              string         `json:"paymentProviderOnboardingState" form:"paymentProviderOnboardingState" gorm:"column:payment_provider_onboarding_state;comment:State of the onboarding payment provider step;"`
	PaymentOnboardingPaymentMethod              string         `json:"paymentOnboardingPaymentMethod" form:"paymentOnboardingPaymentMethod" gorm:"column:payment_onboarding_payment_method;comment:Selected onboarding payment method;"`
	MessageMainAttachmentId                     int            `json:"messageMainAttachmentId" form:"messageMainAttachmentId" gorm:"column:message_main_attachment_id;comment:Main Attachment;size:32;"`
	FiscalyearLastDay                           int            `json:"fiscalyearLastDay" form:"fiscalyearLastDay" gorm:"column:fiscalyear_last_day;comment:Fiscalyear Last Day;size:32;"`
	TransferAccountId                           int            `json:"transferAccountId" form:"transferAccountId" gorm:"column:transfer_account_id;comment:Inter-Banks Transfer Account;size:32;"`
	ChartTemplateId                             int            `json:"chartTemplateId" form:"chartTemplateId" gorm:"column:chart_template_id;comment:Chart Template;size:32;"`
	DefaultCashDifferenceIncomeAccountId        int            `json:"defaultCashDifferenceIncomeAccountId" form:"defaultCashDifferenceIncomeAccountId" gorm:"column:default_cash_difference_income_account_id;comment:Cash Difference Income Account;size:32;"`
	DefaultCashDifferenceExpenseAccountId       int            `json:"defaultCashDifferenceExpenseAccountId" form:"defaultCashDifferenceExpenseAccountId" gorm:"column:default_cash_difference_expense_account_id;comment:Cash Difference Expense Account;size:32;"`
	AccountJournalSuspenseAccountId             int            `json:"accountJournalSuspenseAccountId" form:"accountJournalSuspenseAccountId" gorm:"column:account_journal_suspense_account_id;comment:Journal Suspense Account;size:32;"`
	AccountJournalPaymentDebitAccountId         int            `json:"accountJournalPaymentDebitAccountId" form:"accountJournalPaymentDebitAccountId" gorm:"column:account_journal_payment_debit_account_id;comment:Journal Outstanding Receipts Account;size:32;"`
	AccountJournalPaymentCreditAccountId        int            `json:"accountJournalPaymentCreditAccountId" form:"accountJournalPaymentCreditAccountId" gorm:"column:account_journal_payment_credit_account_id;comment:Journal Outstanding Payments Account;size:32;"`
	AccountJournalEarlyPayDiscountGainAccountId int            `json:"accountJournalEarlyPayDiscountGainAccountId" form:"accountJournalEarlyPayDiscountGainAccountId" gorm:"column:account_journal_early_pay_discount_gain_account_id;comment:Cash Discount Write-Off Gain Account;size:32;"`
	AccountJournalEarlyPayDiscountLossAccountId int            `json:"accountJournalEarlyPayDiscountLossAccountId" form:"accountJournalEarlyPayDiscountLossAccountId" gorm:"column:account_journal_early_pay_discount_loss_account_id;comment:Cash Discount Write-Off Loss Account;size:32;"`
	AccountSaleTaxId                            int            `json:"accountSaleTaxId" form:"accountSaleTaxId" gorm:"column:account_sale_tax_id;comment:Default Sale Tax;size:32;"`
	AccountPurchaseTaxId                        int            `json:"accountPurchaseTaxId" form:"accountPurchaseTaxId" gorm:"column:account_purchase_tax_id;comment:Default Purchase Tax;size:32;"`
	CurrencyExchangeJournalId                   int            `json:"currencyExchangeJournalId" form:"currencyExchangeJournalId" gorm:"column:currency_exchange_journal_id;comment:Exchange Gain or Loss Journal;size:32;"`
	IncomeCurrencyExchangeAccountId             int            `json:"incomeCurrencyExchangeAccountId" form:"incomeCurrencyExchangeAccountId" gorm:"column:income_currency_exchange_account_id;comment:Gain Exchange Rate Account;size:32;"`
	ExpenseCurrencyExchangeAccountId            int            `json:"expenseCurrencyExchangeAccountId" form:"expenseCurrencyExchangeAccountId" gorm:"column:expense_currency_exchange_account_id;comment:Loss Exchange Rate Account;size:32;"`
	PropertyStockAccountInputCategId            int            `json:"propertyStockAccountInputCategId" form:"propertyStockAccountInputCategId" gorm:"column:property_stock_account_input_categ_id;comment:Input Account for Stock Valuation;size:32;"`
	PropertyStockAccountOutputCategId           int            `json:"propertyStockAccountOutputCategId" form:"propertyStockAccountOutputCategId" gorm:"column:property_stock_account_output_categ_id;comment:Output Account for Stock Valuation;size:32;"`
	PropertyStockValuationAccountId             int            `json:"propertyStockValuationAccountId" form:"propertyStockValuationAccountId" gorm:"column:property_stock_valuation_account_id;comment:Account Template for Stock Valuation;size:32;"`
	IncotermId                                  int            `json:"incotermId" form:"incotermId" gorm:"column:incoterm_id;comment:Default incoterm;size:32;"`
	AccountOpeningMoveId                        int            `json:"accountOpeningMoveId" form:"accountOpeningMoveId" gorm:"column:account_opening_move_id;comment:Opening Journal Entry;size:32;"`
	AccountDefaultPosReceivableAccountId        int            `json:"accountDefaultPosReceivableAccountId" form:"accountDefaultPosReceivableAccountId" gorm:"column:account_default_pos_receivable_account_id;comment:Default PoS Receivable Account;size:32;"`
	ExpenseAccrualAccountId                     int            `json:"expenseAccrualAccountId" form:"expenseAccrualAccountId" gorm:"column:expense_accrual_account_id;comment:Expense Accrual Account;size:32;"`
	RevenueAccrualAccountId                     int            `json:"revenueAccrualAccountId" form:"revenueAccrualAccountId" gorm:"column:revenue_accrual_account_id;comment:Revenue Accrual Account;size:32;"`
	AutomaticEntryDefaultJournalId              int            `json:"automaticEntryDefaultJournalId" form:"automaticEntryDefaultJournalId" gorm:"column:automatic_entry_default_journal_id;comment:Automatic Entry Default Journal;size:32;"`
	AccountFiscalCountryId                      int            `json:"accountFiscalCountryId" form:"accountFiscalCountryId" gorm:"column:account_fiscal_country_id;comment:Fiscal Country;size:32;"`
	TaxCashBasisJournalId                       int            `json:"taxCashBasisJournalId" form:"taxCashBasisJournalId" gorm:"column:tax_cash_basis_journal_id;comment:Cash Basis Journal;size:32;"`
	AccountCashBasisBaseAccountId               int            `json:"accountCashBasisBaseAccountId" form:"accountCashBasisBaseAccountId" gorm:"column:account_cash_basis_base_account_id;comment:Base Tax Received Account;size:32;"`
	FiscalyearLastMonth                         string         `json:"fiscalyearLastMonth" form:"fiscalyearLastMonth" gorm:"column:fiscalyear_last_month;comment:Fiscalyear Last Month;"`
	BankAccountCodePrefix                       string         `json:"bankAccountCodePrefix" form:"bankAccountCodePrefix" gorm:"column:bank_account_code_prefix;comment:Prefix of the bank accounts;"`
	CashAccountCodePrefix                       string         `json:"cashAccountCodePrefix" form:"cashAccountCodePrefix" gorm:"column:cash_account_code_prefix;comment:Prefix of the cash accounts;"`
	EarlyPayDiscountComputation                 string         `json:"earlyPayDiscountComputation" form:"earlyPayDiscountComputation" gorm:"column:early_pay_discount_computation;comment:Cash Discount Tax Reduction;"`
	TransferAccountCodePrefix                   string         `json:"transferAccountCodePrefix" form:"transferAccountCodePrefix" gorm:"column:transfer_account_code_prefix;comment:Prefix of the transfer accounts;"`
	TaxCalculationRoundingMethod                string         `json:"taxCalculationRoundingMethod" form:"taxCalculationRoundingMethod" gorm:"column:tax_calculation_rounding_method;comment:Tax Calculation Rounding Method;"`
	AccountSetupBankDataState                   string         `json:"accountSetupBankDataState" form:"accountSetupBankDataState" gorm:"column:account_setup_bank_data_state;comment:State of the onboarding bank data step;"`
	AccountSetupFyDataState                     string         `json:"accountSetupFyDataState" form:"accountSetupFyDataState" gorm:"column:account_setup_fy_data_state;comment:State of the onboarding fiscal year step;"`
	AccountSetupCoaState                        string         `json:"accountSetupCoaState" form:"accountSetupCoaState" gorm:"column:account_setup_coa_state;comment:State of the onboarding charts of account step;"`
	AccountSetupTaxesState                      string         `json:"accountSetupTaxesState" form:"accountSetupTaxesState" gorm:"column:account_setup_taxes_state;comment:State of the onboarding Taxes step;"`
	AccountOnboardingInvoiceLayoutState         string         `json:"accountOnboardingInvoiceLayoutState" form:"accountOnboardingInvoiceLayoutState" gorm:"column:account_onboarding_invoice_layout_state;comment:State of the onboarding invoice layout step;"`
	AccountOnboardingSaleTaxState               string         `json:"accountOnboardingSaleTaxState" form:"accountOnboardingSaleTaxState" gorm:"column:account_onboarding_sale_tax_state;comment:State of the onboarding sale tax step;"`
	AccountInvoiceOnboardingState               string         `json:"accountInvoiceOnboardingState" form:"accountInvoiceOnboardingState" gorm:"column:account_invoice_onboarding_state;comment:State of the account invoice onboarding panel;"`
	AccountDashboardOnboardingState             string         `json:"accountDashboardOnboardingState" form:"accountDashboardOnboardingState" gorm:"column:account_dashboard_onboarding_state;comment:State of the account dashboard onboarding panel;"`
	TermsType                                   string         `json:"termsType" form:"termsType" gorm:"column:terms_type;comment:Terms & Conditions format;"`
	AccountSetupBillState                       string         `json:"accountSetupBillState" form:"accountSetupBillState" gorm:"column:account_setup_bill_state;comment:State of the onboarding bill step;"`
	QuickEditMode                               string         `json:"quickEditMode" form:"quickEditMode" gorm:"column:quick_edit_mode;comment:Quick encoding;"`
	PeriodLockDate                              *time.Time     `json:"periodLockDate" form:"periodLockDate" gorm:"column:period_lock_date;comment:Journals Entries Lock Date;"`
	FiscalyearLockDate                          *time.Time     `json:"fiscalyearLockDate" form:"fiscalyearLockDate" gorm:"column:fiscalyear_lock_date;comment:All Users Lock Date;"`
	TaxLockDate                                 *time.Time     `json:"taxLockDate" form:"taxLockDate" gorm:"column:tax_lock_date;comment:Tax Return Lock Date;"`
	AccountOpeningDate                          *time.Time     `json:"accountOpeningDate" form:"accountOpeningDate" gorm:"column:account_opening_date;comment:Opening Entry;"`
	InvoiceTerms                                string         `json:"invoiceTerms" form:"invoiceTerms" gorm:"column:invoice_terms;comment:Default Terms and Conditions;"`
	InvoiceTermsHtml                            string         `json:"invoiceTermsHtml" form:"invoiceTermsHtml" gorm:"column:invoice_terms_html;comment:Default Terms and Conditions as a Web page;"`
	ExpectsChartOfAccounts                      bool           `json:"expectsChartOfAccounts" form:"expectsChartOfAccounts" gorm:"column:expects_chart_of_accounts;comment:Expects a Chart of Accounts;"`
	AngloSaxonAccounting                        bool           `json:"angloSaxonAccounting" form:"angloSaxonAccounting" gorm:"column:anglo_saxon_accounting;comment:Use anglo-saxon accounting;"`
	QrCode                                      bool           `json:"qrCode" form:"qrCode" gorm:"column:qr_code;comment:Display QR-code on invoices;"`
	InvoiceIsEmail                              bool           `json:"invoiceIsEmail" form:"invoiceIsEmail" gorm:"column:invoice_is_email;comment:Email by default;"`
	InvoiceIsPrint                              bool           `json:"invoiceIsPrint" form:"invoiceIsPrint" gorm:"column:invoice_is_print;comment:Print by default;"`
	AccountUseCreditLimit                       bool           `json:"accountUseCreditLimit" form:"accountUseCreditLimit" gorm:"column:account_use_credit_limit;comment:Sales Credit Limit;"`
	AccountOnboardingCreateInvoiceStateFlag     bool           `json:"accountOnboardingCreateInvoiceStateFlag" form:"accountOnboardingCreateInvoiceStateFlag" gorm:"column:account_onboarding_create_invoice_state_flag;comment:Account Onboarding Create Invoice State Flag;"`
	TaxExigibility                              bool           `json:"taxExigibility" form:"taxExigibility" gorm:"column:tax_exigibility;comment:Use Cash Basis;"`
	AccountStorno                               bool           `json:"accountStorno" form:"accountStorno" gorm:"column:account_storno;comment:Storno accounting;"`
	ExpenseJournalId                            int            `json:"expenseJournalId" form:"expenseJournalId" gorm:"column:expense_journal_id;comment:Default Expense Journal;size:32;"`
	CompanyExpenseJournalId                     int            `json:"companyExpenseJournalId" form:"companyExpenseJournalId" gorm:"column:company_expense_journal_id;comment:Default Company Expense Journal;size:32;"`
	ProjectTimeModeId                           int            `json:"projectTimeModeId" form:"projectTimeModeId" gorm:"column:project_time_mode_id;comment:Project Time Unit;size:32;"`
	TimesheetEncodeUomId                        int            `json:"timesheetEncodeUomId" form:"timesheetEncodeUomId" gorm:"column:timesheet_encode_uom_id;comment:Timesheet Encoding Unit;size:32;"`
	InternalProjectId                           int            `json:"internalProjectId" form:"internalProjectId" gorm:"column:internal_project_id;comment:Internal Project;size:32;"`
	InvoiceIsSnailmail                          bool           `json:"invoiceIsSnailmail" form:"invoiceIsSnailmail" gorm:"column:invoice_is_snailmail;comment:Send by Post;"`
	LeaveTimesheetTaskId                        int            `json:"leaveTimesheetTaskId" form:"leaveTimesheetTaskId" gorm:"column:leave_timesheet_task_id;comment:Time Off Task;size:32;"`
	InternalTransitLocationId                   int            `json:"internalTransitLocationId" form:"internalTransitLocationId" gorm:"column:internal_transit_location_id;comment:Internal Transit Location;size:32;"`
	StockMailConfirmationTemplateId             int            `json:"stockMailConfirmationTemplateId" form:"stockMailConfirmationTemplateId" gorm:"column:stock_mail_confirmation_template_id;comment:Email Template confirmation picking;size:32;"`
	AnnualInventoryDay                          int            `json:"annualInventoryDay" form:"annualInventoryDay" gorm:"column:annual_inventory_day;comment:Day of the month;size:32;"`
	AnnualInventoryMonth                        string         `json:"annualInventoryMonth" form:"annualInventoryMonth" gorm:"column:annual_inventory_month;comment:Annual Inventory Month;"`
	StockMoveEmailValidation                    bool           `json:"stockMoveEmailValidation" form:"stockMoveEmailValidation" gorm:"column:stock_move_email_validation;comment:Email Confirmation picking;"`
	StockSmsConfirmationTemplateId              int            `json:"stockSmsConfirmationTemplateId" form:"stockSmsConfirmationTemplateId" gorm:"column:stock_sms_confirmation_template_id;comment:SMS Template;size:32;"`
	StockMoveSmsValidation                      bool           `json:"stockMoveSmsValidation" form:"stockMoveSmsValidation" gorm:"column:stock_move_sms_validation;comment:SMS Confirmation;"`
	HasReceivedWarningStockSms                  bool           `json:"hasReceivedWarningStockSms" form:"hasReceivedWarningStockSms" gorm:"column:has_received_warning_stock_sms;comment:Has Received Warning Stock Sms;"`
	QuotationValidityDays                       int            `json:"quotationValidityDays" form:"quotationValidityDays" gorm:"column:quotation_validity_days;comment:Default Quotation Validity (Days);size:32;"`
	SaleQuotationOnboardingState                string         `json:"saleQuotationOnboardingState" form:"saleQuotationOnboardingState" gorm:"column:sale_quotation_onboarding_state;comment:State of the sale onboarding panel;"`
	SaleOnboardingOrderConfirmationState        string         `json:"saleOnboardingOrderConfirmationState" form:"saleOnboardingOrderConfirmationState" gorm:"column:sale_onboarding_order_confirmation_state;comment:State of the onboarding confirmation order step;"`
	SaleOnboardingSampleQuotationState          string         `json:"saleOnboardingSampleQuotationState" form:"saleOnboardingSampleQuotationState" gorm:"column:sale_onboarding_sample_quotation_state;comment:State of the onboarding sample quotation step;"`
	SaleOnboardingPaymentMethod                 string         `json:"saleOnboardingPaymentMethod" form:"saleOnboardingPaymentMethod" gorm:"column:sale_onboarding_payment_method;comment:Sale onboarding selected payment method;"`
	PortalConfirmationSign                      bool           `json:"portalConfirmationSign" form:"portalConfirmationSign" gorm:"column:portal_confirmation_sign;comment:Online Signature;"`
	PortalConfirmationPay                       bool           `json:"portalConfirmationPay" form:"portalConfirmationPay" gorm:"column:portal_confirmation_pay;comment:Online Payment;"`
	SecurityLead                                float64        `json:"securityLead" form:"securityLead" gorm:"column:security_lead;comment:Sales Safety Days;"`
	SaleOrderTemplateId                         int            `json:"saleOrderTemplateId" form:"saleOrderTemplateId" gorm:"column:sale_order_template_id;comment:Default Sale Template;size:32;"`
	PoLock                                      string         `json:"poLock" form:"poLock" gorm:"column:po_lock;comment:Purchase Order Modification;"`
	PoDoubleValidation                          string         `json:"poDoubleValidation" form:"poDoubleValidation" gorm:"column:po_double_validation;comment:Levels of Approvals;"`
	PoDoubleValidationAmount                    pgtype.Numeric `json:"poDoubleValidationAmount" form:"poDoubleValidationAmount" gorm:"column:po_double_validation_amount;comment:Double validation amount;"`
	PoLead                                      float64        `json:"poLead" form:"poLead" gorm:"column:po_lead;comment:Purchase Lead Time;"`
	DaysToPurchase                              float64        `json:"daysToPurchase" form:"daysToPurchase" gorm:"column:days_to_purchase;comment:Days to Purchase;"`
	HealthCare                                  string         `json:"healthCare" form:"healthCare" gorm:"column:health_care;comment:医保机构编码;"`
	CityId                                      int            `json:"cityId" form:"cityId" gorm:"column:city_id;comment:市(地区、州);size:32;"`
	AreaId                                      int            `json:"areaId" form:"areaId" gorm:"column:area_id;comment:县(区);size:32;"`
	Town                                        string         `json:"town" form:"town" gorm:"column:town;comment:乡(镇、街道办事处);"`
	Village                                     string         `json:"village" form:"village" gorm:"column:village;comment:村(街、路弄等);"`
	HouseNumber                                 string         `json:"houseNumber" form:"houseNumber" gorm:"column:house_number;comment:门牌号码;"`
	HospLv                                      string         `json:"hospLv" form:"hospLv" gorm:"column:hosp_lv;comment:医院等级;"`
	DclaSouc                                    string         `json:"dclaSouc" form:"dclaSouc" gorm:"column:dcla_souc;comment:机构类别;"`
	FixmedinsType                               string         `json:"fixmedinsType" form:"fixmedinsType" gorm:"column:fixmedins_type;comment:定点医疗机构类型;"`
	InsuOpen                                    string         `json:"insuOpen" form:"insuOpen" gorm:"column:insu_open;comment:开通情况;"`
	PointOfSaleUpdateStockQuantities            string         `json:"pointOfSaleUpdateStockQuantities" form:"pointOfSaleUpdateStockQuantities" gorm:"column:point_of_sale_update_stock_quantities;comment:Update quantities in stock;"`
	PointOfSaleUseTicketQrCode                  bool           `json:"pointOfSaleUseTicketQrCode" form:"pointOfSaleUseTicketQrCode" gorm:"column:point_of_sale_use_ticket_qr_code;comment:Use QR code on ticket;"`
	Version                                     string         `json:"version" form:"version" gorm:"column:version;comment:版本;"`
	RegistrationAuth                            string         `json:"registrationAuth" form:"registrationAuth" gorm:"column:registration_auth;comment:注册机构;"`
	EnvRelevant                                 string         `json:"envRelevant" form:"envRelevant" gorm:"column:env_relevant;comment:相关环境;"`
	CompetentAuth                               string         `json:"competentAuth" form:"competentAuth" gorm:"column:competent_auth;comment:主管机构;"`
	StatusAuth                                  string         `json:"statusAuth" form:"statusAuth" gorm:"column:status_auth;comment:注册状态;"`
}

// TableName ResCompany 表名
func (ResCompany) TableName() string {
	return "res_company"
}
