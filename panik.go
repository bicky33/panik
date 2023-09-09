package panik

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type NIK struct {
	NIK string `json:"nik"`
}

type ExtractDataNIK struct {
	RegionData   regionData `json:"region_data"`
	Gender       string     `json:"gender"`
	BirthDate    string     `json:"birth_date"`
	SerialNumber string     `json:"serial_number"`
}

type regionData struct {
	ProvinceID   string `json:"provinceId"`
	ProvinceName string `json:"provinceName"`
	RegencyID    string `json:"regencyId"`
	RegencyName  string `json:"regencyName"`
	DistrictID   string `json:"districtId"`
	DistrictName string `json:"districtName"`
}

func (e *NIK) getRegionId() string {
	return e.NIK[0:6]
}

func (e *NIK) getBirthDate() string {
	return e.NIK[6:12]
}

func (e *NIK) getSerialNumber() string {
	return e.NIK[12:16]
}

func (e *NIK) IsValid() (bool, error) {
	numeric := regexp.MustCompile(`^\d+$`).MatchString(e.NIK)
	if !numeric {
		return false, nil
	}

	if len(e.NIK) != 16 {
		return false, nil
	}

	if !e.isNIKPartsValid() {
		return false, nil
	}

	return true, nil
}

func (e *NIK) isNIKPartsValid() bool {
	region, birthDate, serialNumber := e.getRegionId(), e.getBirthDate(), e.getSerialNumber()
	if _, err := os.Stat("data/" + region + ".json"); os.IsNotExist(err) {
		return false
	}

	date, err := strconv.Atoi(birthDate[0:2])
	if err != nil {
		return false
	}

	month, err := strconv.Atoi(birthDate[2:4])
	if err != nil {
		return false
	}

	if date > 71 || serialNumber == "0000" || date == 0 || month == 0 || month > 12 {
		return false
	}

	return true
}

func (e *NIK) extractRegionData(regionId string) (regionData, error) {
	var regionData regionData
	data, err := os.ReadFile("data/" + regionId + ".json")
	if err != nil {
		return regionData, err
	}

	err = json.Unmarshal(data, &regionData)
	return regionData, err
}

func (e *NIK) Data() (ExtractDataNIK, error) {
	var dataNIK ExtractDataNIK
	isValid, err := e.IsValid()
	if !isValid && err != nil {
		return dataNIK, err
	}
	regionData, err := e.extractRegionData(e.getRegionId())
	if err != nil {
		return dataNIK, err
	}

	birthdate := e.getBirthDate()
	date, err := strconv.Atoi(birthdate[0:2])
	if err != nil {
		return dataNIK, err
	}

	year, err := strconv.Atoi(birthdate[4:6])
	if err != nil {
		return dataNIK, err
	}

	thisYear := time.Now().Year() - 2000

	if year > thisYear {
		year = year + 1900
	}

	if year <= thisYear {
		year = year + 2000
	}

	if date > 31 {
		dataNIK.Gender = "Female"
		date = date - 40
		dataNIK.BirthDate = fmt.Sprintf("%d-%s-%d", date, birthdate[2:4], year)
	}

	if date <= 31 {
		dataNIK.Gender = "Male"
		dataNIK.BirthDate = fmt.Sprintf("%d-%s-%d", date, birthdate[2:4], year)
	}

	dataNIK.RegionData = regionData
	dataNIK.SerialNumber = e.getSerialNumber()
	return dataNIK, nil
}
