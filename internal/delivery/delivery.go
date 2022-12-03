package delivery

import "gorm.io/gorm"

//Packet is the physical object that we want to send
type Packet struct {
	gorm.Model
	SenderID    uint
	From        string
	Destination string
	Weight      float32
}
