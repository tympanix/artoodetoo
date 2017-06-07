package drive

type driveStyle struct{}

func (driveStyle) Color() uint {
	return 0xdb3236
}

func (driveStyle) Icon() string {
	return "fa-google"
}
