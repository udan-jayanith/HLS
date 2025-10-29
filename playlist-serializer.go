package HLS

//Implement PlaylistSerializer struct
//Add Write method to it.
//Add AddLine method to it.
//Add AddLineTo.
//Read
//RemoveQuotes

func (pt *PlaylistToken) Serialize() []byte {
	switch pt.Type {
	case Comment, Tag:
		return []byte("#" + pt.Value)
	}
	return []byte(pt.Value)
}

func NewPlaylistToken(lineType LineType, value string) PlaylistToken {
	return PlaylistToken{
		Type:  lineType,
		Value: value,
	}
}
