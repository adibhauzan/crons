package helper

import (
	"fmt"
	"strings"
	"time"
)

func FormatTanggalKeIndonesia(tanggal string) (string, error) {
	parsedDate, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return "", fmt.Errorf("error parsing date: %w", err)
	}

	formattedDate := parsedDate.Format("01 JANUARY 2006")

	indonesianMonths := map[string]string{
		"JANUARY":   "JANUARI",
		"FEBRUARY":  "FEBRUARI",
		"MARCH":     "MARET",
		"APRIL":     "APRIL",
		"MAY":       "MEI",
		"JUNE":      "JUNI",
		"JULY":      "JULI",
		"AUGUST":    "AGUSTUS",
		"SEPTEMBER": "SEPTEMBER",
		"OCTOBER":   "OKTOBER",
		"NOVEMBER":  "NOVEMBER",
		"DECEMBER":  "DESEMBER",
	}

	for eng, ind := range indonesianMonths {
		formattedDate = strings.ReplaceAll(formattedDate, eng, ind)
	}

	return formattedDate, nil
}

func FormatTanggalKeIndonesiaV2(t time.Time) string {
	formattedDate := t.Format("02 January 2006")

	indonesianMonths := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	for eng, ind := range indonesianMonths {
		formattedDate = strings.ReplaceAll(formattedDate, eng, ind)
	}

	return formattedDate
}

var hari = []string{
	"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu",
}

var bulan = []string{
	"Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

// FormatTanggalIndonesia mengubah time.Time ke format "Senin 01 Januari 2025"
func FormatTanggalIndonesia(t time.Time, withKoma bool) string {
	koma := ","
	if !withKoma {
		koma = ""
	}

	return fmt.Sprintf("%s%s %d %s %d",
		hari[t.Weekday()],
		koma,
		t.Day(),
		bulan[int(t.Month())-1],
		t.Year(),
	)
}
