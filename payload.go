package gopresence

type Handshake struct {
	V        int    `json:"v"`
	ClientID string `json:"client_id"`
}

type SetActivityArg struct {
	Pid      int      `json:"pid"`
	Activity Activity `json:"activity"`
}

type Activity struct {
	Timestamps    *Timestamps `json:"timestamps,omitempty"`
	ApplicationID int         `json:"application_id,omitempty"`
	Details       string      `json:"details,omitempty"`
	State         string      `json:"state,omitempty"`
	Emoji         *Emoji      `json:"emoji,omitempty"`
	Party         *Party      `json:"party,omitempty"`
	Assets        *Assets     `json:"assets,omitempty"`
	Secrets       *Secrets    `json:"secrets,omitempty"`
	Instance      bool        `json:"instance,omitempty"`
	Flags         int         `json:"flags,omitempty"`
	Buttons       *[2]Button  `json:"buttons,omitempty"`
}

type Timestamps struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
}

type Emoji struct {
	Name     string `json:"name,omitempty"`
	ID       int    `json:"id,omitempty"`
	Animated bool   `json:"animated,omitempty"`
}

type Party struct {
	ID   string `json:"id,omitempty"`
	Size [2]int `json:"size,omitempty"`
}

type Assets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type Secrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

type Button struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}

type Payload struct {
	Cmd   string         `json:"cmd"`
	Nonce string         `json:"nonce"`
	Args  SetActivityArg `json:"args"`
	Evt   string         `json:"evt,omitempty"`
	Data  string         `json:"data,omitempty"`
}
