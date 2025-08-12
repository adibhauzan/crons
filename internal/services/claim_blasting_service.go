package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"regexp"
	"strings"
	"time"

	"github.com/adibhauzan/crons/helper"
	"github.com/adibhauzan/crons/internal/broker"
	"github.com/adibhauzan/crons/internal/domain/dto"
	"github.com/adibhauzan/crons/internal/repositories"
	"github.com/adibhauzan/crons/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AttachmentInfo struct {
	URL      string `json:"url"`
	Filename string `json:"filename,omitempty"`
}

type EmailTask struct {
	Email       string           `json:"email"`
	Subject     string           `json:"subject"`
	Body        string           `json:"body"`
	CC          []string         `json:"cc,omitempty"`
	Attachments []AttachmentInfo `json:"attachments,omitempty"`
}

type ClaimBlastingService interface {
	GetClaim(ctx context.Context) ([]dto.ClaimGetResponse, error)
	BlastingEmailClaim(ctx context.Context) error
}

type claimBastingService struct {
	applicationSettingRepo repositories.ApplicationSettingRepository
	claimBlastingRepo      repositories.ClaimBlastingRepository
	documentRepo           repositories.DocumentsRepository
	storageService         StorageService
	broker                 broker.RabbitMQProducerInterface
	db                     *gorm.DB
	logger                 *logrus.Logger
}

func NewClaimBastingService(
	applicationSettingRepo repositories.ApplicationSettingRepository,
	claimBlastingRepo repositories.ClaimBlastingRepository,
	documentRepo repositories.DocumentsRepository,
	storageService StorageService,
	broker broker.RabbitMQProducerInterface,
	db *gorm.DB,
	logger *logrus.Logger,
) *claimBastingService {
	return &claimBastingService{
		applicationSettingRepo: applicationSettingRepo,
		claimBlastingRepo:      claimBlastingRepo,
		documentRepo:           documentRepo,
		storageService:         storageService,
		broker:                 broker,
		db:                     db,
		logger:                 logger,
	}
}

func (service *claimBastingService) GetClaim(ctx context.Context) ([]dto.ClaimGetResponse, error) {
	data, err := service.claimBlastingRepo.GetAllClaimToBlasting(ctx)
	if err != nil {
		return nil, err
	}

	var claims []dto.ClaimGetResponse
	for _, claim := range data {
		claimDocuments, err := service.claimBlastingRepo.ClaimGetDocuments(ctx, claim.Number)
		if err != nil {
			return nil, err
		}

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
		for _, v := range claimDocuments {
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
		claims = append(claims, dto.ClaimGetResponse{
			ID:                     claim.ID,
			Number:                 claim.Number,
			TransactionDate:        claim.TransactionDate,
			ReferenceNumber:        claim.ReferenceNumber,
			XFrom:                  claim.XFrom,
			InsurancePartnerCode:   claim.InsurancePartnerCode,
			InsurancePartnerEmail:  claim.InsurancePartnersEmail,
			InsuranceTypeCode:      claim.InsuranceTypeCode,
			InsuranceProductCode:   claim.InsuranceProductCode,
			ClaimLetterNumber:      claim.ClaimLetterNumber,
			ClaimLetterDate:        claim.ClaimLetterDate,
			DateOfIncident:         claim.DateOfIncident,
			Remarks:                json.RawMessage(claim.Remarks),
			Description:            claim.Description,
			Amount:                 claim.Amount,
			StatusInsurancePartner: claim.StatusInsurancePartner,
			UPBID:                  claim.UPBID,
			UPB:                    claim.UPB,
			StatusSendEMail:        claim.StatusSendEmail,
			StatusBlast:            claim.StatusBlast,
			Approve:                claim.Approve,
			Approve2:               claim.Approve2,
			CreatedBy:              claim.CreatedBy,
			CreatedAt:              claim.CreatedAt,
			UpdatedBy:              claim.UpdatedBy,
			UpdatedAt:              claim.UpdatedAt,
			ApprovedBy:             claim.ApprovedBy,
			ApprovedAt:             claim.ApprovedAt,
			Approved2By:            claim.Approved2By,
			Approved2At:            claim.Approved2At,
			// Premium:                     claim.P,
			InsurancePartner:   claim.InsurancePartners,
			InsuranceType:      claim.InsuranceType,
			Customer:           claim.Customer,
			Branch:             claim.Branch,
			Insured:            claim.Insured,
			InsuredAddress:     claim.InsuredAddress,
			InsuredDateOfBirth: claim.InsuredDateOfBirth,
			InsuredPhoneNumber: claim.InsuredPhoneNumber,
			NoDebitur:          claim.NoDebitur,
			StartDate:          claim.StartDate,
			EndDate:            claim.EndDate,
			JUP:                claim.JUP,
			NoPolicy:           claim.NoPolicy,
			// LeaderConsortium:            claim.LeaderConsortium,
			PaidAmount: claim.PaidAmount,
			PaidDate:   claim.PaidDate,
			Submitter:  claim.Submitter,
			// StatusLack:                  claim.StatusLack,
			// StatusInsuranceReceipt:      claim.StatusInsuranceReceipt,
			StatusArrears:    claim.StatusArrears,
			Arrears:          claim.Arrears,
			StatusPenalty:    claim.StatusPenalty,
			Penalty:          claim.Penalty,
			StatusEscalation: claim.StatusEscalation,
			Escalation:       claim.Escalation,
			// ReferenceDate:               claim.ReferenceDate,
			AmountRevision:       claim.AmountRevision,
			DateInsurancePartner: claim.DateOfIncident,
			FileInsurancePartner: claim.FileInsurancePartner,
			// StatusClaimSubmission: claim.StatusClaimSubmissionLetter,
			StatusPaid:       claim.StatusPaid,
			ClaimApprove:     claim.ClaimApprove,
			ClaimApprove2:    claim.ClaimApprove2,
			ClaimApprovedBy:  claim.ClaimApprovedBy,
			ClaimApprovedAt:  claim.ClaimApprovedAt,
			ClaimApproved2By: claim.ClaimApproved2By,
			ClaimApproved2At: claim.ClaimApproved2At,
			DeletedBy:        claim.DeletedBy,
			DeletedAt:        claim.DeletedAt,
			DeleteStatus:     claim.DeleteStatus,
			InsuranceProduct: claim.InsuranceProducts,
			LeaderConsorsium: claim.LeaderConsorsium,
			// TotalRecord:                 claim.TotalRecord,
			// NumberLackOfDocumentLetter:  claim.NumberLackOfDocumentLetter,
			// XTimeStamp:                  claim.XTimeStamp,
			// InsurancePartnerAddress:     claim.InsurancePartnerAddress,
			UPBCode: claim.UPBCode,
			// IsLeader:                    claim.IsLeader,
			// SharePercentage:             claim.SharePercentage,
			// ShareAmount:                 claim.ShareAmount,
			BankID:        claim.BankID,
			BankName:      claim.BankName,
			AccountNumber: claim.AccountNumber,
			AccountName:   claim.AccountName,
			// InsuredEmail:                claim.InsuredEmail,
			SubmitterEmail: claim.SubmitterEmail,
			SubmitterPhone: claim.SubmitterPhone,
			StatusDocument: claim.StatusDocument,
			// InsuredIDCard:               claim.InsuredIDCard,
			// DocumentPolicy:              claim.DocumentPolicy,
			LetterOfDischargeFile:       claim.LetterOfDischargeFile,
			LetterOfDischargeDate:       claim.LetterOfDischargeDate,
			LetterOfDischargeStatus:     claim.LetterOfDischargeStatus,
			LetterOfDischargeSigned:     claim.LetterOfDischargeSigned,
			LetterOfDischargeSignedDate: claim.LetterOfDischargeSignedDate,
			Documents:                   docs,
		})
	}

	return claims, nil
}

func (service *claimBastingService) BlastingEmailClaim(ctx context.Context) error {
	err := service.db.Transaction(func(tx *gorm.DB) error {

		claim, err := service.GetClaim(ctx)
		if err != nil {
			return err
		}

		for _, claim := range claim {
			appSetting, err := service.applicationSettingRepo.Get(ctx)
			if err != nil {
				return err
			}

			layout := time.RFC3339
			dob, err := time.Parse(layout, claim.InsuredDateOfBirth)
			if err != nil {
				service.logger.Error("Error parsing date:", err)
				return err
			}

			dob2, err := time.Parse(layout, claim.TransactionDate)
			if err != nil {
				service.logger.Error("Error parsing date:", err)
				return err
			}

			dob3, err := time.Parse(layout, claim.StartDate)
			if err != nil {
				service.logger.Error("Error parsing date:", err)
				return err
			}

			dob4, err := time.Parse(layout, claim.DateOfIncident)
			if err != nil {
				service.logger.Error("Error parsing date:", err)
				return err
			}

			dob5, err := time.Parse(layout, claim.StartDate)
			if err != nil {
				service.logger.Error("Error parsing date:", err)
				return err
			}
			dob6, err := time.Parse(layout, claim.EndDate)
			if err != nil {
				service.logger.Error("Error parsing date:", err)
				return err
			}

			tglLahirDebitur := helper.FormatTanggalIndonesia(dob, false)
			tglTransaksi := helper.FormatTanggalIndonesia(dob2, false)
			tglRealisasi := helper.FormatTanggalIndonesia(dob3, false)
			tglKejadian := helper.FormatTanggalIndonesia(dob4, false)
			hariIni := time.Now()
			hariIniBenar := helper.FormatTanggalIndonesia(hariIni, false)
			hariIniBenarDenganKoma := helper.FormatTanggalIndonesia(hariIni, true)
			tanggalIniBenar := helper.FormatTanggalIndonesia(hariIni, false)

			startDate := helper.FormatTanggalKeIndonesiaV2(dob5)
			endDate := helper.FormatTanggalKeIndonesiaV2(dob6)

			periodePertanggungan := fmt.Sprintf("%s s/d %s", startDate, endDate)

			ser1, err := service.storageService.GetFileBase64(appSetting.Sertificate1)
			if err != nil {
				return errors.New("failed to get certificate 1: " + err.Error())
			}

			ser2, err := service.storageService.GetFileBase64(appSetting.Sertificate2)
			if err != nil {
				return errors.New("failed to get certificate 2: " + err.Error())
			}
			ser3, err := service.storageService.GetFileBase64(appSetting.Sertificate3)
			if err != nil {
				return errors.New("failed to get certificate 3: " + err.Error())
			}
			ser4, err := service.storageService.GetFileBase64(appSetting.Sertificate4)
			if err != nil {
				return errors.New("failed to get certificate 4: " + err.Error())
			}
			ser5, err := service.storageService.GetFileBase64(appSetting.Sertificate5)
			if err != nil {
				return errors.New("failed to get certificate 5: " + err.Error())
			}

			var docs []struct {
				ID          int
				Name        string
				Description string
				File        string
				Link        string
			}

			for _, doc := range claim.Documents {
				docum, err := service.documentRepo.GetByID(ctx, doc.DocumentID)
				if err != nil {
					return err
				}

				docs = append(docs, struct {
					ID          int
					Name        string
					Description string
					File        string
					Link        string
				}{
					ID:          doc.DocumentID,
					Name:        docum.Name,
					Description: doc.Description,
					File:        doc.Files,
					Link:        doc.Files,
				})
			}

			var wilayahPrefixRegex = regexp.MustCompile(`(?i)^(kota|kabupaten)\s+`)
			upbCity := func(nama string) string {
				return strings.TrimSpace(wilayahPrefixRegex.ReplaceAllString(nama, ""))
			}(claim.UPB_City)

			htmlData := map[string]interface{}{
				"NamaPelapor":         claim.Submitter,
				"NamaDebitur":         claim.Insured,
				"TanggalLahirDebitur": tglLahirDebitur,
				"NomorPeserta":        claim.ReferenceNumber,
				"TanggalRealisasi":    tglRealisasi,
				"TanggalMusibah":      tglKejadian,
				"AlamatDebitur":       claim.InsuredAddress,
				"NomorTeleponDebitur": claim.SubmitterPhone,
				"ProdukAsuransi":      claim.InsuranceProduct,
				"HariIni":             hariIniBenar,
				"TanggalIni":          tanggalIniBenar,
				"KodeUPB":             claim.UPBCode,
				"UPB":                 upbCity,
				"CompanyName":         appSetting.CompanyName,
				"CompanyPrintName":    appSetting.CompanyPrintName,
				"CompanyPhone":        appSetting.Phone1,
				"CompanyEmail":        appSetting.Email,
				"TanggalTransaksi":    tglTransaksi,
				"Documents":           docs,
				"Certificate1":        template.HTMLAttr(ser1),
				"Certificate2":        template.HTMLAttr(ser2),
				"Certificate3":        template.HTMLAttr(ser3),
				"Certificate4":        template.HTMLAttr(ser4),
				"Certificate5":        template.HTMLAttr(ser5),
			}

			pdfBuffer, err := helper.GeneratePDFWITHHTML("templates/pages/claim_pelapor_attachment.html", htmlData)
			if err != nil {
				return err
			}

			uploadPath, err := service.storageService.UploadRawFile(dto.UploadRawFile{
				FolderName:  "pengajuan-claims",
				FileName:    "Laporan Pengajuan Claim" + ".pdf",
				ContentType: "application/pdf",
				Content:     pdfBuffer,
			})

			service.logger.Info("uploadPath: ", uploadPath)
			if err != nil {
				service.logger.Error("Error uploading file to storage:", err)
				return err
			}

			tmpl, err := template.ParseFiles("templates/pages/claim_pelapor_email.html")
			if err != nil {
				return fmt.Errorf("failed to parse email template: %w", err)
			}

			var emailTasks []EmailTask

			bodyToPelapor := map[string]string{
				"NamaPelapor":    claim.Submitter,
				"TanggalMusibah": tglKejadian,
				"ProdukAsuransi": claim.InsuranceProduct,
				"NamaDebitur":    claim.Insured,
			}

			var tplBuffer1 bytes.Buffer
			if err := tmpl.Execute(&tplBuffer1, bodyToPelapor); err != nil {
				return err
			}

			emailTasks = append(emailTasks, EmailTask{
				Email:   "adibhauzan1@gmail.com",
				Subject: "Laporan Klaim " + claim.InsuranceProduct,
				Body:    tplBuffer1.String(),
				Attachments: []AttachmentInfo{{
					URL:      uploadPath,
					Filename: "Laporan Pengajuan Claim" + ".pdf",
				},
				},
			})

			type dataTertanggungs struct {
				NamaTertanggung      string
				NoPolicies           string
				TanggalMusibah       string
				TanggalLapor         string
				PeriodePertanggungan string
			}

			var dataTertanggung []dataTertanggungs

			dataTertanggung = append(dataTertanggung, dataTertanggungs{
				NamaTertanggung:      claim.Insured,
				NoPolicies:           claim.ReferenceNumber,
				TanggalMusibah:       tglKejadian,
				TanggalLapor:         tglTransaksi,
				PeriodePertanggungan: periodePertanggungan,
			})

			bodyToAsuransi := map[string]any{
				"HariIni":          hariIniBenarDenganKoma,
				"Asuransi":         claim.InsurancePartner,
				"InsuranceProduct": claim.InsuranceProduct,
				"Data":             dataTertanggung,
			}

			tmpl2, err := template.New("claim_to_asuransi.html").
				Funcs(template.FuncMap{
					"add": func(a, b int) int { return a + b },
				}).
				ParseFiles("templates/pages/claim_to_asuransi.html")
			if err != nil {
				return fmt.Errorf("failed to parse email template: %w", err)
			}
			var tplBuffer2 bytes.Buffer
			if err := tmpl2.Execute(&tplBuffer2, bodyToAsuransi); err != nil {
				return err
			}

			emailTasks = append(emailTasks, EmailTask{
				Email:   "adibhauzan1@gmail.com",
				Subject: "Laporan Klaim " + claim.InsuranceProduct,
				Body:    tplBuffer2.String(),
			})

			if err := utils.PublishTasks(service.broker, "email_queue", emailTasks); err != nil {
				return err
			}

			// err = service.claimBlastingRepo.UpdateStatusBlasting(ctx, claim.Number)
			// if err != nil {
			// 	service.logger.Errorf("Error updating status blasting: %v", err)
			// 	return err
			// }
		}

		return nil
	})

	if err != nil {
		service.logger.Errorf("insert claim receipt err:%v", err)
		return err
	}
	return nil
}
