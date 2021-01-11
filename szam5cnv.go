package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	appName = "sZam5 config converter"
	appVer  = "0.0.1"
)

// Set - ...
type Set struct {
	XMLName                xml.Name `xml:"Sets"`
	Ages                   bool     `xml:"setAges"`
	ArabDigital            bool     `xml:"setArabDigital"`
	Chast                  bool     `xml:"setChast"`
	ChastList              []string `xml:"setChastList"`
	ChastRegEx             bool     `xml:"setChastRegEx"`
	Currency               bool     `xml:"setCurrency"`
	CurrencySpace          uint     `xml:"setCurrencySpace"`
	Dash                   bool     `xml:"setDash"`
	DashBeginSpace         uint     `xml:"setDashBeginSpace"`
	DashSpaceAfterPunct    bool     `xml:"setDashSpaceAfterPunct"`
	DashEm                 bool     `xml:"setDashEm"`
	DashEn                 bool     `xml:"setDashEn"`
	DashSpaceAfter         uint     `xml:"setDashSpaceAfter"`
	DashSpaceAfterNoBreak  bool     `xml:"setDashSpaceAfterNoBreak"`
	DashSpaceAfterWidth    uint     `xml:"setDashSpaceAfterWidth"`
	DashSpaceBefore        uint     `xml:"setDashSpaceBefore"`
	DashSpaceBeforeNoBreak bool     `xml:"setDashSpaceBeforeNoBreak"`
	DashSpaceBeforeWidth   uint     `xml:"setDashSpaceBeforeWidth"`
	DigitalDash            uint     `xml:"setDigitalDash"`
	DigitalDashSpace       uint     `xml:"setDigitalDashSpace"`
	DigitalMinus           bool     `xml:"setDigitalMinus"`
	FrenchEnable           bool     `xml:"setFrenchEnable"`
	FrenchSpace            uint     `xml:"setFrenchSpace"`
	ExistRazriad           bool     `xml:"setExistRazriad"`
	Info                   string   `xml:"setInfo"`
	Initials               bool     `xml:"setInitials"`
	InitialsBefore         uint     `xml:"setInitialsBefore"`
	InitialsBetween        uint     `xml:"setInitialsBetween"`
	LastWord               bool     `xml:"setLastWord"`
	LatinDigital           bool     `xml:"setLatinDigital"`
	MesSocr                bool     `xml:"setMesSocr"`
	MesSocrList            []string `xml:"setMesSocrList"`
	MesSocrSpace           uint     `xml:"setMesSocrSpace"`
	Ellipsis               bool     `xml:"setMultidots"`
	Augment                bool     `xml:"setNar"`
	NoBreak                bool     `xml:"setNoBreak"`
	Monosyllabic           bool     `xml:"setOneWord"`
	MonosyllabicCaseSens   bool     `xml:"setOneWordCaseSens"`
	MonosyllabicList       []string `xml:"setOneWordList"`
	MonosyllabicRegEx      bool     `xml:"setOneWordRegEx"`
	Phone                  bool     `xml:"setPhone"`
	PercentHairSpace       bool     `xml:"setProcentHairSpace"`
	Square                 bool     `xml:"setQuadrat"`
	Quotes                 bool     `xml:"setQuotes"`
	QuotesType             uint     `xml:"setQuotesType"`
	Razriad                bool     `xml:"setRazriad"`
	RazriadSpace           uint     `xml:"setRazriadSpace"`
	RemoveDoublePara       bool     `xml:"setRemoveDoublePara"`
	RemoveDoubleSpaces     bool     `xml:"setRemoveDoubleSpaces"`
	RemoveFirstTab         bool     `xml:"setRemoveFirstTab"`
	RemoveLastTab          bool     `xml:"setRemoveLastTab"`
	Slashes                bool     `xml:"setSlashes"`
	SlashesHairSpace       bool     `xml:"setSlashesHairSpace"`
	Socr                   bool     `xml:"setSocr"`
	SocrList               []string `xml:"setSocrList"`
	SocrSpace              uint     `xml:"setSocrSpace"`
	WebNoBreak             bool     `xml:"setWebNoBreak"`
	ChkOut                 bool     `xml:"setChkOut"`
	UseAddional            bool     `xml:"setUseAddional"`
	AdCurrencies           []string `xml:"setAdCurrencies"`
	AdYears                []string `xml:"setAdYears"`
	AdNoBreakHyppenBefore  []string `xml:"setAdNoBreakHyppenBefore"`
	AdNoBreakHyppenAfter   []string `xml:"setAdNoBreakHyppenAfter"`
	AdForMinusSigns        []string `xml:"setAdForMinusSigns"`
	Rule                   []Rule
}

// Rule - ...
type Rule struct {
	XMLName      xml.Name `xml:"Rule"`
	Name         string   `xml:"name,attr"`
	Enable       bool     `xml:"enable,attr"`
	Find         string
	FindPref     string
	Change       string
	ChangePref   string
	Footnote     bool
	MasterPage   bool
	HiddenLayers bool
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(appName, appVer)
		return
	}

	inFile := os.Args[1]
	if f, err := os.Stat(inFile); err != nil || !strings.HasSuffix(f.Name(), ".xn") {
		fmt.Println(appName, appVer)
		log.Println("Файл", inFile, "не существует или не является файлом .XN")
		return
	}

	var set Set
	bytes, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Println(err)
		return
	}

	if err = xml.Unmarshal(bytes, &set); err != nil {
		log.Println(err)
		return
	}

	err = buildConfig(inFile, set)
	if err != nil {
		log.Println(err)
		return
	}
}

func buildConfig(f string, set Set) error {
	var cfg config
	cfg.Info = set.Info

	cfg.Abbrevation.Enable = set.Socr
	cfg.Abbrevation.List = set.SocrList
	cfg.Abbrevation.Space = set.SocrSpace

	cfg.Addional.Enable = set.UseAddional
	cfg.Addional.Currencies = set.AdCurrencies
	cfg.Addional.ForMinusSign = set.AdForMinusSigns
	cfg.Addional.NoBreakHyphenAfter = set.AdNoBreakHyppenAfter
	cfg.Addional.NoBreakHyphenBefore = set.AdNoBreakHyppenBefore

	cfg.Ages = set.Ages

	cfg.Augment = set.Augment

	cfg.Currency.Enable = set.Currency
	cfg.Currency.Space = set.CurrencySpace

	cfg.Dash.Enable = set.Dash
	cfg.Dash.BeginSpace = set.DashBeginSpace
	cfg.Dash.SpaceAfter = set.DashSpaceAfter
	cfg.Dash.SpaceAfterNoBreak = set.DashSpaceAfterNoBreak
	cfg.Dash.SpaceAfterPunct = set.DashSpaceAfterPunct
	cfg.Dash.SpaceAfterWidth = set.DashSpaceAfterWidth
	cfg.Dash.SpaceBefore = set.DashSpaceBefore
	cfg.Dash.SpaceBeforeNoBreak = set.DashSpaceBeforeNoBreak
	cfg.Dash.SpaceBeforeWidth = set.DashSpaceBeforeWidth

	cfg.Digital.Arabes = set.ArabDigital
	cfg.Digital.Latin = set.LatinDigital
	cfg.Digital.Dash = set.DigitalDash
	cfg.Digital.DashSpace = set.DigitalDashSpace
	cfg.Digital.Minus = set.DigitalMinus

	// Запись JSON
	result, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(strings.TrimSuffix(f, ".xn")+".json", result, os.ModePerm)
}
