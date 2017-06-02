package dropbox

// Token is provided by the user and used to access dropbox
type Token string

type dropboxStyle struct{}

func (dropboxStyle) Color() uint {
	return 0x007ee5
}

func (dropboxStyle) Icon() string {
	return "fa-dropbox"
}
