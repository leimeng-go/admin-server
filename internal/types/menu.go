package types

type AuthItem struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthMark string `json:"auth_mark"`
}

type Meta struct {
	Title             string     `json:"title"`
	Icon              string     `json:"icon,omitempty"`
	KeepAlive         bool       `json:"keepAlive"`
	ShowTextBadge     string     `json:"showTextBadge,omitempty"`
	ShowBadge         bool       `json:"showBadge,omitempty"`
	Link              string     `json:"link,omitempty"`
	IsIframe          bool       `json:"isIframe,omitempty"`
	IsHide            bool       `json:"isHide,omitempty"`
	IsHideTab         bool       `json:"isHideTab,omitempty"`
	IsInMainContainer bool       `json:"isInMainContainer,omitempty"`
	AuthList          []AuthItem `json:"authList,omitempty"`
}

type Route struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	Component string   `json:"component"`
	Meta      Meta     `json:"meta"`
	Children  []*Route `json:"children,omitempty"`
}
