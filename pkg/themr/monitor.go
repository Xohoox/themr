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
	X int	`yaml:x`
	Y int	`yaml:y`
}

type MonitorMode struct {
	Width int	`yaml:width`
	Height int	`yaml:height`
}

func (position Position) String() string {
	return strconv.Itoa(position.X) + "x" + strconv.Itoa(position.Y)
}

func (mode MonitorMode) String() string {
	return strconv.Itoa(mode.Width) + "x" + strconv.Itoa(mode.Height)
}

type Monitor struct {
	Output string		`yaml:output`
	Primary bool		`yaml:primary`
	Enabled bool		`yaml:enabled`
	// Refresh bool
	Rotation Rotation	`yaml:rotation`
	Position Position	`yaml:position`
	Mode MonitorMode	`yaml:mode`
}
