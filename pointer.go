package wayland

type Button int

const (
	ButtonMisc Button = 0x100
	Button0    Button = 0x100
	Button1    Button = 0x101
	Button2    Button = 0x102
	Button3    Button = 0x103
	Button4    Button = 0x104
	Button5    Button = 0x105
	Button6    Button = 0x106
	Button7    Button = 0x107
	Button8    Button = 0x108
	Button9    Button = 0x109

	ButtonMouse   Button = 0x110
	ButtonLeft    Button = 0x110
	ButtonRight   Button = 0x111
	ButtonMiddle  Button = 0x112
	ButtonSide    Button = 0x113
	ButtonExtra   Button = 0x114
	ButtonForward Button = 0x115
	ButtonBack    Button = 0x116
	ButtonTask    Button = 0x117

	ButtonJoystick Button = 0x120
	ButtonTrigger  Button = 0x120
	ButtonThumb    Button = 0x121
	ButtonThumb2   Button = 0x122
	ButtonTop      Button = 0x123
	ButtonTop2     Button = 0x124
	ButtonPinkie   Button = 0x125
	ButtonBase     Button = 0x126
	ButtonBase2    Button = 0x127
	ButtonBase3    Button = 0x128
	ButtonBase4    Button = 0x129
	ButtonBase5    Button = 0x12a
	ButtonBase6    Button = 0x12b
	ButtonDead     Button = 0x12f

	ButtonGamepad Button = 0x130
	ButtonSouth   Button = 0x130
	ButtonA       Button = ButtonSouth
	ButtonEast    Button = 0x131
	ButtonB       Button = ButtonEast
	ButtonC       Button = 0x132
	ButtonNorth   Button = 0x133
	ButtonX       Button = ButtonNorth
	ButtonWest    Button = 0x134
	ButtonY       Button = ButtonWest
	ButtonZ       Button = 0x135
	ButtonTl      Button = 0x136
	ButtonTr      Button = 0x137
	ButtonTl2     Button = 0x138
	ButtonTr2     Button = 0x139
	ButtonSelect  Button = 0x13a
	ButtonStart   Button = 0x13b
	ButtonMode    Button = 0x13c
	ButtonThumbl  Button = 0x13d
	ButtonThumbr  Button = 0x13e

	ButtonDigi          Button = 0x140
	ButtonToolPen       Button = 0x140
	ButtonToolRubber    Button = 0x141
	ButtonToolBrush     Button = 0x142
	ButtonToolPencil    Button = 0x143
	ButtonToolAirbrush  Button = 0x144
	ButtonToolFinger    Button = 0x145
	ButtonToolMouse     Button = 0x146
	ButtonToolLens      Button = 0x147
	ButtonToolQuinttap  Button = 0x148 /* Five fingers on trackpad */
	ButtonStylus3       Button = 0x149
	ButtonTouch         Button = 0x14a
	ButtonStylus        Button = 0x14b
	ButtonStylus2       Button = 0x14c
	ButtonToolDoubletap Button = 0x14d
	ButtonToolTripletap Button = 0x14e
	ButtonToolQuadtap   Button = 0x14f /* Four fingers on trackpad */

	ButtonWheel    Button = 0x150
	ButtonGearDown Button = 0x150
	ButtonGearUp   Button = 0x151
)
