package dto

import (
	"encoding/json"

	"github.com/adibhauzan/crons/internal/domain/entity"
)

type ClaimGetResponse struct {
	ID                          int             `json:"id"`
	Number                      string          `json:"number"`
	TransactionDate             string          `json:"transaction_date"`
	ReferenceNumber             string          `json:"reference_number"`
	XFrom                       string          `json:"x_from"`
	InsurancePartnerCode        string          `json:"insurance_partner_code"`
	InsuranceTypeCode           string          `json:"insurance_type_code"`
	InsuranceProductCode        string          `json:"insurance_product_code"`
	ClaimLetterNumber           string          `json:"claim_letter_number"`
	ClaimLetterDate             string          `json:"claim_letter_date"`
	DateOfIncident              string          `json:"date_of_incident"`
	Remarks                     json.RawMessage `json:"remarks"`
	Description                 string          `json:"description"`
	Amount                      float64         `json:"amount"`
	StatusInsurancePartner      int             `json:"status_insurance_partner"`
	UPBID                       int             `json:"upb_id"`
	UPB                         string          `json:"upb"`
	UPB_City                    string          `json:"upb_city"`
	StatusSendEMail             int             `json:"status_send_email"`
	StatusBlast                 int             `json:"status_blast"`
	Approve                     int             `json:"approve"`
	Approve2                    int             `json:"approve2"`
	CreatedBy                   string          `json:"created_by"`
	CreatedAt                   string          `json:"created_at"`
	UpdatedBy                   string          `json:"updated_by"`
	UpdatedAt                   string          `json:"updated_at"`
	ApprovedBy                  string          `json:"approved_by"`
	ApprovedAt                  string          `json:"approved_at"`
	Approved2By                 string          `json:"approved2_by"`
	Approved2At                 string          `json:"approved2_at"`
	Premium                     float64         `json:"premium"`
	InsurancePartner            string          `json:"insurance_partner"`
	InsurancePartnerEmail       string          `json:"insurance_partner_email"`
	InsuranceType               string          `json:"insurance_type"`
	Customer                    string          `json:"customer"`
	Branch                      string          `json:"branch"`
	Insured                     string          `json:"insured"`
	InsuredAddress              string          `json:"insured_address"`
	InsuredDateOfBirth          string          `json:"insured_date_of_birth"`
	InsuredPhoneNumber          string          `json:"insured_phone_number"`
	NoDebitur                   string          `json:"no_debitur"`
	StartDate                   string          `json:"start_date"`
	EndDate                     string          `json:"end_date"`
	JUP                         float64         `json:"jup"`
	NoPolicy                    string          `json:"no_policy"`
	LeaderConsortium            string          `json:"leader_consortium"`
	PaidAmount                  float64         `json:"paid_amount"`
	PaidDate                    string          `json:"paid_date"`
	Submitter                   string          `json:"submitter"`
	StatusLack                  int             `json:"status_lack"`
	StatusInsuranceReceipt      int             `json:"status_insurance_receipt"`
	StatusArrears               int             `json:"status_arrears"`
	Arrears                     float64         `json:"arrears"`
	StatusPenalty               int             `json:"status_penalty"`
	Penalty                     float64         `json:"penalty"`
	StatusEscalation            int             `json:"status_escalation"`
	Escalation                  float64         `json:"escalation"`
	ReferenceDate               string          `json:"reference_date"`
	AmountRevision              float64         `json:"amount_revision"`
	DateInsurancePartner        string          `json:"date_insurance_partner"`
	FileInsurancePartner        string          `json:"file_insurance_partner"`
	StatusClaimSubmission       int             `json:"status_claim_submission"`
	StatusPaid                  int             `json:"status_paid"`
	ClaimApprove                int             `json:"claim_approve"`
	ClaimApprove2               int             `json:"claim_approve2"`
	ClaimApprovedBy             string          `json:"claim_approved_by"`
	ClaimApprovedAt             string          `json:"claim_approved_at"`
	ClaimApproved2By            string          `json:"claim_approved2_by"`
	ClaimApproved2At            string          `json:"claim_approved2_at"`
	DeletedBy                   string          `json:"deleted_by"`
	DeletedAt                   string          `json:"deleted_at"`
	DeleteStatus                int             `json:"delete_status"`
	InsuranceProduct            string          `json:"insurance_product"`
	LeaderConsorsium            string          `json:"leader_consorsium"`
	TotalRecord                 int             `json:"total_record"`
	NumberLackOfDocumentLetter  string          `json:"number_lack_of_document_letter"`
	XTimeStamp                  string          `json:"x_timestamp"`
	InsurancePartnerAddress     string          `json:"insurance_partner_address"`
	UPBCode                     string          `json:"upb_code"`
	IsLeader                    int             `json:"is_leader"`
	SharePercentage             float64         `json:"share_percentage"`
	ShareAmount                 float64         `json:"share_amount"`
	BankID                      int             `json:"bank_id"`
	BankName                    string          `json:"bank_name"`
	AccountNumber               string          `json:"account_number"`
	AccountName                 string          `json:"account_name"`
	InsuredEmail                string          `json:"insured_email"`
	SubmitterPhone              string          `json:"submitter_phone"`
	SubmitterEmail              string          `json:"submitter_email"`
	StatusDocument              int             `json:"status_document"`
	InsuredIDCard               string          `json:"insured_id_card"`
	DocumentPolicy              string          `json:"document_policy"`
	LetterOfDischargeFile       string          `json:"letter_of_discharge_file"`
	LetterOfDischargeDate       string          `json:"letter_of_discharge_date"`
	LetterOfDischargeStatus     int             `json:"letter_of_discharge_status"`
	LetterOfDischargeSigned     string          `json:"letter_of_discharge_signed"`
	LetterOfDischargeSignedDate string          `json:"letter_of_discharge_signed_date"`
	Note                        string          `json:"note"`
	Documents                   []struct {
		ID                          int    `json:"id"`
		ClaimNumber                 string `json:"claim_number"`
		DocumentID                  int    `json:"document_id"`
		Description                 string `json:"description"`
		Files                       string `json:"file"`
		ReceiptNumber               string `json:"receipt_number"`
		ReceiptNumberFiles          string `json:"receipt_number_file"`
		RegisterNumber              string `json:"register_number"`
		StatusFiles                 int    `json:"status_file"`
		StatusReceive               int    `json:"status_receive"`
		ReceiptNumberDate           string `json:"receipt_number_date"`
		StatusPhysicalDocument      int    `json:"status_physical_document"`
		StatusApprove               int    `json:"status_approve"`
		XTimeStamp                  string `json:"x_time_stamp"`
		ReasonFile                  string `json:"reason_file"`
		ReasonPhysicalDocument      string `json:"reason_physical_document"`
		ReasonApprove               string `json:"reason_approve"`
		Document                    string `json:"document"`
		ClaimReceiptNumber          string `json:"claim_receipt_number"`
		ClaimReceiptNumberDate      string `json:"claim_receipt_number_date"`
		ClaimReceiptNumberFile      string `json:"claim_receipt_number_file"`
		CreatedBy                   string `json:"created_by"`
		FileType                    string `json:"file_type"`
		ClaimRegisterNumber         string `json:"claim_register_number"`
		StatusInsurancePartner      int    `json:"status_insurance_partner"`
		ReasonInsurancePartner      string `json:"reason_insurance_partner"`
		StatusFilesInsurancePartner int    `json:"status_file_insurance_partner"`
		ReasonFilesInsurancePartner string `json:"reason_file_insurance_partner"`
	} `json:"documents"`
	Statuses []struct {
		ID           int    `json:"id"`
		ClaimsNumber string `json:"claims_number"`
		DepartmentID int    `json:"department_id"`
		Physical     int    `json:"physical"`
		Files        int    `json:"file"`
		Department   string `json:"department"`
	} `json:"statuses"`
	Shares []struct {
		Number                 string  `json:"number"`
		IsLeader               int     `json:"is_leader"`
		InsurancePartnerCode   string  `json:"insurance_partner_code"`
		InsurancePartner       string  `json:"insurance_partner"`
		InsurancePartnerBranch string  `json:"insurance_partner_branch"`
		SharePercentage        float64 `json:"share_percentage"`
	} `json:"shares"`
}

func ToClaimGetResponse(data *entity.Claim, documents []entity.ClaimDocuments) ClaimGetResponse {
	var docs []struct {
		ID                          int    `json:"id"`
		ClaimNumber                 string `json:"claim_number"`
		DocumentID                  int    `json:"document_id"`
		Description                 string `json:"description"`
		Files                       string `json:"file"`
		ReceiptNumber               string `json:"receipt_number"`
		ReceiptNumberFiles          string `json:"receipt_number_file"`
		RegisterNumber              string `json:"register_number"`
		StatusFiles                 int    `json:"status_file"`
		StatusReceive               int    `json:"status_receive"`
		ReceiptNumberDate           string `json:"receipt_number_date"`
		StatusPhysicalDocument      int    `json:"status_physical_document"`
		StatusApprove               int    `json:"status_approve"`
		XTimeStamp                  string `json:"x_time_stamp"`
		ReasonFile                  string `json:"reason_file"`
		ReasonPhysicalDocument      string `json:"reason_physical_document"`
		ReasonApprove               string `json:"reason_approve"`
		Document                    string `json:"document"`
		ClaimReceiptNumber          string `json:"claim_receipt_number"`
		ClaimReceiptNumberDate      string `json:"claim_receipt_number_date"`
		ClaimReceiptNumberFile      string `json:"claim_receipt_number_file"`
		CreatedBy                   string `json:"created_by"`
		FileType                    string `json:"file_type"`
		ClaimRegisterNumber         string `json:"claim_register_number"`
		StatusInsurancePartner      int    `json:"status_insurance_partner"`
		ReasonInsurancePartner      string `json:"reason_insurance_partner"`
		StatusFilesInsurancePartner int    `json:"status_file_insurance_partner"`
		ReasonFilesInsurancePartner string `json:"reason_file_insurance_partner"`
	}
	for _, v := range documents {
		docs = append(docs, struct {
			ID                          int    `json:"id"`
			ClaimNumber                 string `json:"claim_number"`
			DocumentID                  int    `json:"document_id"`
			Description                 string `json:"description"`
			Files                       string `json:"file"`
			ReceiptNumber               string `json:"receipt_number"`
			ReceiptNumberFiles          string `json:"receipt_number_file"`
			RegisterNumber              string `json:"register_number"`
			StatusFiles                 int    `json:"status_file"`
			StatusReceive               int    `json:"status_receive"`
			ReceiptNumberDate           string `json:"receipt_number_date"`
			StatusPhysicalDocument      int    `json:"status_physical_document"`
			StatusApprove               int    `json:"status_approve"`
			XTimeStamp                  string `json:"x_time_stamp"`
			ReasonFile                  string `json:"reason_file"`
			ReasonPhysicalDocument      string `json:"reason_physical_document"`
			ReasonApprove               string `json:"reason_approve"`
			Document                    string `json:"document"`
			ClaimReceiptNumber          string `json:"claim_receipt_number"`
			ClaimReceiptNumberDate      string `json:"claim_receipt_number_date"`
			ClaimReceiptNumberFile      string `json:"claim_receipt_number_file"`
			CreatedBy                   string `json:"created_by"`
			FileType                    string `json:"file_type"`
			ClaimRegisterNumber         string `json:"claim_register_number"`
			StatusInsurancePartner      int    `json:"status_insurance_partner"`
			ReasonInsurancePartner      string `json:"reason_insurance_partner"`
			StatusFilesInsurancePartner int    `json:"status_file_insurance_partner"`
			ReasonFilesInsurancePartner string `json:"reason_file_insurance_partner"`
		}{
			ID:                          v.ID,
			ClaimNumber:                 v.ClaimNumber,
			DocumentID:                  v.DocumentID,
			Description:                 v.Description,
			Document:                    v.Document,
			Files:                       v.Files,
			ReceiptNumber:               v.ReceiptNumber,
			StatusFiles:                 v.StatusFiles,
			StatusReceive:               v.StatusReceive,
			StatusPhysicalDocument:      v.StatusPhysicalDocument,
			StatusApprove:               v.StatusApprove,
			RegisterNumber:              v.RegisterNumber,
			ReceiptNumberFiles:          v.ReceiptNumberFiles,
			XTimeStamp:                  v.XTimeStamp,
			ClaimReceiptNumber:          v.ClaimReceiptNumber,
			ClaimReceiptNumberFile:      v.ClaimReceiptNumberFile,
			ClaimReceiptNumberDate:      v.ClaimReceiptNumberDate,
			ReceiptNumberDate:           v.ReceiptNumberDate,
			ReasonFile:                  v.ReasonFile,
			ReasonPhysicalDocument:      v.ReasonPhysicalDocument,
			ReasonApprove:               v.ReasonApprove,
			CreatedBy:                   v.CreatedBy,
			FileType:                    v.FilesType,
			ClaimRegisterNumber:         v.ClaimsRegisterNumber,
			StatusInsurancePartner:      v.StatusInsurancePartner,
			ReasonInsurancePartner:      v.ReasonInsurancePartner,
			StatusFilesInsurancePartner: v.StatusFileInsurancePartner,
			ReasonFilesInsurancePartner: v.ReasonFileInsurancePartner,
		})
	}

	return ClaimGetResponse{
		ID:                     data.ID,
		Number:                 data.Number,
		TransactionDate:        data.TransactionDate,
		ReferenceNumber:        data.ReferenceNumber,
		XFrom:                  data.XFrom,
		InsurancePartnerCode:   data.InsurancePartnerCode,
		InsurancePartnerEmail:  data.InsurancePartnersEmail,
		InsuranceTypeCode:      data.InsuranceTypeCode,
		InsuranceProductCode:   data.InsuranceProductCode,
		ClaimLetterNumber:      data.ClaimLetterNumber,
		ClaimLetterDate:        data.ClaimLetterDate,
		DateOfIncident:         data.DateOfIncident,
		Remarks:                json.RawMessage(data.Remarks),
		Description:            data.Description,
		Amount:                 data.Amount,
		StatusInsurancePartner: data.StatusInsurancePartner,
		UPBID:                  data.UPBID,
		UPB:                    data.UPB,
		UPB_City:               data.UPBCity,
		StatusSendEMail:        data.StatusSendEmail,
		StatusBlast:            data.StatusBlast,
		Approve:                data.Approve,
		Approve2:               data.Approve2,
		CreatedBy:              data.CreatedBy,
		CreatedAt:              data.CreatedAt,
		UpdatedBy:              data.UpdatedBy,
		UpdatedAt:              data.UpdatedAt,
		ApprovedBy:             data.ApprovedBy,
		ApprovedAt:             data.ApprovedAt,
		Approved2By:            data.Approved2By,
		Approved2At:            data.Approved2At,
		// Premium:                     data.P,
		InsurancePartner:   data.InsurancePartners,
		InsuranceType:      data.InsuranceType,
		Customer:           data.Customer,
		Branch:             data.Branch,
		Insured:            data.Insured,
		InsuredAddress:     data.InsuredAddress,
		InsuredDateOfBirth: data.InsuredDateOfBirth,
		InsuredPhoneNumber: data.InsuredPhoneNumber,
		NoDebitur:          data.NoDebitur,
		StartDate:          data.StartDate,
		EndDate:            data.EndDate,
		JUP:                data.JUP,
		NoPolicy:           data.NoPolicy,
		// LeaderConsortium:            data.LeaderConsortium,
		PaidAmount: data.PaidAmount,
		PaidDate:   data.PaidDate,
		Submitter:  data.Submitter,
		// StatusLack:                  data.StatusLack,
		// StatusInsuranceReceipt:      data.StatusInsuranceReceipt,
		StatusArrears:    data.StatusArrears,
		Arrears:          data.Arrears,
		StatusPenalty:    data.StatusPenalty,
		Penalty:          data.Penalty,
		StatusEscalation: data.StatusEscalation,
		Escalation:       data.Escalation,
		// ReferenceDate:               data.ReferenceDate,
		AmountRevision:       data.AmountRevision,
		DateInsurancePartner: data.DateOfIncident,
		FileInsurancePartner: data.FileInsurancePartner,
		// StatusClaimSubmission: data.StatusClaimSubmissionLetter,
		StatusPaid:       data.StatusPaid,
		ClaimApprove:     data.ClaimApprove,
		ClaimApprove2:    data.ClaimApprove2,
		ClaimApprovedBy:  data.ClaimApprovedBy,
		ClaimApprovedAt:  data.ClaimApprovedAt,
		ClaimApproved2By: data.ClaimApproved2By,
		ClaimApproved2At: data.ClaimApproved2At,
		DeletedBy:        data.DeletedBy,
		DeletedAt:        data.DeletedAt,
		DeleteStatus:     data.DeleteStatus,
		InsuranceProduct: data.InsuranceProducts,
		LeaderConsorsium: data.LeaderConsorsium,
		// TotalRecord:                 data.TotalRecord,
		// NumberLackOfDocumentLetter:  data.NumberLackOfDocumentLetter,
		// XTimeStamp:                  data.XTimeStamp,
		// InsurancePartnerAddress:     data.InsurancePartnerAddress,
		UPBCode: data.UPBCode,
		// IsLeader:                    data.IsLeader,
		// SharePercentage:             data.SharePercentage,
		// ShareAmount:                 data.ShareAmount,
		BankID:        data.BankID,
		BankName:      data.BankName,
		AccountNumber: data.AccountNumber,
		AccountName:   data.AccountName,
		// InsuredEmail:                data.InsuredEmail,
		SubmitterEmail: data.SubmitterEmail,
		SubmitterPhone: data.SubmitterPhone,
		StatusDocument: data.StatusDocument,
		// InsuredIDCard:               data.InsuredIDCard,
		// DocumentPolicy:              data.DocumentPolicy,
		LetterOfDischargeFile:       data.LetterOfDischargeFile,
		LetterOfDischargeDate:       data.LetterOfDischargeDate,
		LetterOfDischargeStatus:     data.LetterOfDischargeStatus,
		LetterOfDischargeSigned:     data.LetterOfDischargeSigned,
		LetterOfDischargeSignedDate: data.LetterOfDischargeSignedDate,
		// Note:                        data.Note,
		Documents: docs,
		// Statuses:                    statuses,
		// Shares:                      Shares,
	}
}

func ToClaimGetResponses(datas []entity.Claim, documents []entity.ClaimDocuments) []ClaimGetResponse {
	var result []ClaimGetResponse
	for _, data := range datas {
		var docs []entity.ClaimDocuments
		for _, doc := range documents {
			if doc.ClaimNumber == data.Number {
				docs = append(docs, doc)
			}
		}
		result = append(result, ToClaimGetResponse(&data, docs))
	}
	return result
}
