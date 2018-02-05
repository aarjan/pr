package main

// type Client struct {
// 	ID                         int
// 	AppUUID                    string
// 	SalesPersonID              int
// 	Name                       string
// 	Type                       string
// 	LegalEntityBooking         string
// 	NumberOfEmployees          string
// 	BusinessSince              int
// 	Segment                    string
// 	LeadSource                 string
// 	ReasonForAllocationToSales string
// 	ResponsibleRM              string
// 	AcquisitionDate            string
// 	AcquisitionCohort          string
// 	// ActivityStatus             bool
// 	ChurnDate string
// 	Timestamp string // record created/updated timestamp?
// }

// // Are these needed in payment?
// // Comments                          string
// // MonthsofService                   int
// // RemainingContractEndOfMonth       string
// type Payment struct {
// 	ID                          int
// 	SubscriptionID              int
// 	Date                        string
// 	NatureOfTransaction         string
// 	ContractStatusAtPaymentDate bool // ?
// 	TotalCashReceived           float64
// 	Currency                    string
// 	RateToHKD                   float64
// 	CashReceivedHKD             float32
// 	PaymentMode                 string

// 	FirstMonthofAccrual string // accrual mean?
// 	MonthsOfAccrual     int
// 	LastMonthofAccrual  string

// 	MonthlyRevenueAllocation              int     //?
// 	CashReceivedForPaymentsToThirdParties float64 //third party?
// 	TypeofThirdParties                    string
// 	NameofThirdParty                      string //plural?
// }

// // needed?
// type ServicePayment struct {
// 	ID                                    int
// 	SubscriptionID                        int
// 	NatureofService                       string
// 	NatureofTransaction                   string
// 	CashReceivedForPaymentsToThirdParties float64
// 	TypeofThirdParties                    string
// 	NameofThirdParty                      string
// }

// type SalesPerson struct {
// 	ID          int
// 	Name        string
// 	Designation string
// 	Team        string
// }

// type Package struct {
// 	ID                int
// 	Name              string
// 	CountryofBilling  string
// 	CurrencyofBilling string
// 	PaymentTerms      string
// 	// MonthlyPrice      float64
// 	Price          float64
// 	NumberOfMonths int
// }

// type Partner struct {
// 	ID      int
// 	Name    string
// 	Type    string
// 	Country string
// }

// type Subscription struct {
// 	ID        int
// 	ClientID  int
// 	PackageID int
// 	PartnerID int

// 	StartDate       string
// 	EndDate         string
// 	MonthsofService int
// 	MonthsFree      int

// 	Status       bool
// 	PaymentTerms string

// 	ContractStatus    string
// 	ContractStartDate string
// 	ContractEndDate   string

// 	// MonthlyPayment float32 // it shows the finalized payment rate
// 	// AnnualPayment  float32

// 	ManagedAccount string // related to

// 	SpecialPromotion string

// 	ChangeOfPlanAtRenewal string // ?
// 	ChurnRationale        string //why churned

// 	NewSalesCommitment int    //?
// 	MonthsToChurn      int    // needed?
// 	Reconciliation     string //?

// 	SubscriptionAddonPrice    float64 // addon?
// 	SetupFee                  float64 // one time installation fee?
// 	ServicePaidInInstallments string  // ?

// 	ServicesNetPriceINC               float64 // all below these needed?
// 	ServicesNetPriceTR                float64
// 	ServicesNetPriceOth               float64
// 	DiscountSubsPrice                 float64
// 	DiscountSubsAddon                 float64
// 	DiscountServicePaidInInstallments float64
// 	DiscountSetupFee                  float64
// 	DiscountServiceINC                float64
// 	DiscountServiceTR                 float64
// 	DiscountServiceOth                float64
// 	OtherPromotions                   float64
// }
