package appiumgen

const (
	SpecActionGo             = "go"
	SpecActionNone           = "none"
	SpecActionDone           = "done"
	SpecActionSend           = "send"
	SpecActionNext           = "next"
	SpecActionJoin           = "join"
	SpecActionRoute          = "route"
	SpecActionSearch         = "search"
	SpecActionNewline        = "newline"
	SpecActionPrevious       = "previous"
	SpecActionContinueAction = "continueAction"
	SpecActionEmergencyCall  = "emergencyCall"
	SpecActionUnspecified    = "unspecified"
)

const (
	SpecMatcherNothing = "nothing"
	SpecMatcherOne     = "one"
)

const (
	SpecTypeVerify = "verify"
	SpecTypeTap    = "tap"
	SpecTypeEnter  = "enter"
)

type Screen struct {
	Name  string
	Specs []Spec
}

type Spec struct {
	Name   string
	Config *SpecConfig
}

type SpecConfig struct {
	Key     *string `json:"key"`
	Text    *string `json:"text"`
	Type    string  `json:"type"`
	Action  *string `json:"action"`
	Delay   *int    `json:"delay"`
	Matcher *string `json:"matcher"`
}
