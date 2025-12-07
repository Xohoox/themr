package themr

import "strconv"

type Rotation string

const (
	RotationNormal Rotation = "normal"
	RotationLeft   Rotation = "left"
	RotationRight  Rotation = "right"
	RotationInvert Rotation = "inverted"
)

type Position struct {
	X int
	Y int
}

type MonitorMode struct {
	Width int
	Height int
}

func (position Position) String() string {
	return strconv.Itoa(position.X) + "x" + strconv.Itoa(position.Y)
}

func (mode MonitorMode) String() string {
	return strconv.Itoa(mode.Width) + "x" + strconv.Itoa(mode.Height)
}

type Monitor struct {
	Output string
	Primary bool
	Enabled bool
	// Refresh bool
	Rotation Rotation
	Position Position
	Mode MonitorMode
}
