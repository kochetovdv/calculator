package model

type ID struct {
	PrevID int
	ID     int
	NextID int
}

// NewID func is used for generation ID the first time
func NewID(id int) *ID {
	i := new(ID)
	i.PrevID = id // start value for ID
	i.ID = i.PrevID + 1
	return i
}

// NewIDChain func is used for generation ID and changing the previous one
func NewIDChain(id *ID) *ID {
	i := new(ID)
	i.PrevID = id.ID
	i.ID = i.PrevID + 1
	id.NextID = i.ID
	return i
}
