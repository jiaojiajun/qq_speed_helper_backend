package model

type CarBase struct {
	Name           string
	PicUrl         string
	RaceFeature    string
	PropFeature    string
	RecordNum      int32
	LevelRecordNum int32
	Level          int32
}

type CasrBaseJson struct {
	CarId      int32
	CarBaseStr string
}
