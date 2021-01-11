package main

type config struct {
	Info             string       `json:"info"`
	Abbrevation      abbrevation  `json:"abbrevation"`
	Addional         addional     `json:"addional"`
	Ages             bool         `json:"ages"`
	Augment          bool         `json:"augment"`
	Currency         currency     `json:"currency"`
	Dash             dash         `json:"dash"`
	Digital          digitales    `json:"digital"`
	DigitalPlace     digitalPlace `json:"digital_place"`
	Ellipsis         bool         `json:"ellipsis"`
	French           french       `json:"french"`
	Initials         initials     `json:"initials"`
	LastWord         bool         `json:"last_word"`
	Measurements     measurements `json:"measurements"`
	Monosyllabic     monosyllabic `json:"monosyllabic"`
	NoBreak          bool         `json:"nobreak"`
	Particle         particle     `json:"particle"`
	Phone            bool         `json:"phone"`
	PercentHairSpace bool         `json:"percent_hair_space"`
	Quotes           quotes       `json:"quotes"`
	Removes          removes      `json:"removes"`
	Slashes          slashes      `json:"slashes"`
	Square           bool         `json:"square"`
	WebNoBreak       bool         `json:"web_nobreak"`
	Rules            []rule       `json:"rules"`
	// ChkOut           bool         `json:"chk_out"`
}

type enabled struct {
	Enable bool `json:"enable"`
}

type spaced struct {
	Space uint `json:"space"`
}

type listed struct {
	List []string `json:"list"`
}

type regexped struct {
	Regexp bool `json:"regexp"`
}

type abbrevation struct {
	enabled
	listed
	spaced
}

type addional struct {
	enabled
	Currencies          []string `json:"currencies"`
	NoBreakHyphenAfter  []string `json:"no_break_hyphen_after"`
	NoBreakHyphenBefore []string `json:"no_break_hyphen_before"`
	ForMinusSign        []string `json:"minus_signes"`
}

type slashes struct {
	enabled
	HairSpace bool `json:"hair_space"`
}

type removes struct {
	DoubleParagraphs bool `json:"double_paragraphs"`
	DoubleSpaces     bool `json:"double_spaces"`
	TabFirst         bool `json:"tab_first"`
	TabLast          bool `json:"tab_last"`
}

type quotes struct {
	enabled
	QuoteType uint `json:"type"`
}

type monosyllabic struct {
	enabled
	listed
	regexped
	CaseSens bool `json:"casesens"`
}

type measurements struct {
	Abbrevation bool `json:"abbrevation"`
	listed
	spaced
}

type initials struct {
	enabled
	Before  uint `json:"before"`
	Between uint `json:"between"`
}

type digitalPlace struct {
	enabled
	spaced
	Exists bool `json:"exists"`
}

type dash struct {
	enabled
	BeginSpace         uint `json:"begin_space"`
	SpaceAfterPunct    bool `json:"space_after_punct"`
	Em                 bool `json:"em"`
	En                 bool `json:"en"`
	SpaceAfter         uint `json:"space_after"`
	SpaceAfterNoBreak  bool `json:"space_after_no_break"`
	SpaceAfterWidth    uint `json:"space_after_width"`
	SpaceBefore        uint `json:"space_before"`
	SpaceBeforeNoBreak bool `json:"space_before_no_break"`
	SpaceBeforeWidth   uint `json:"space_before_width"`
}

type french struct {
	enabled
	spaced
}

type particle struct {
	enabled
	listed
	regexped
}

type currency struct {
	enabled
	spaced
}

type digitales struct {
	Arabes    bool `json:"arabes"`
	Latin     bool `json:"latin"`
	Dash      uint `json:"dash"`
	DashSpace uint `json:"dash_space"`
	Minus     bool `json:"minus"`
}

type rule struct {
	enabled
	Name         string `json:"name"`
	Find         string `json:"find"`
	FindPref     string `json:"find_pref"`
	Change       string `json:"change"`
	ChangePref   string `json:"change_pref"`
	Footnote     bool   `json:"footnote"`
	MasterPage   bool   `json:"masters"`
	HiddenLayers bool   `json:"hiddens"`
}
