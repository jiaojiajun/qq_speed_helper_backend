package mysql_test

import (
	"testing"

	"github.com/jiaojiajun/qq_speed_helper_backend/biz/dal/mysql"
	"github.com/jiaojiajun/qq_speed_helper_backend/biz/model/helper/car"
	"github.com/stretchr/testify/assert"
)

func TestCarDao_GetAllCarBases_happay_case_1(t *testing.T) {
	mysql.Init()
	dao := mysql.GetSingletonCarDAO()
	var assert = assert.New(t)
	assert.NotNil(dao)
	cars, err := dao.GetAllCars()
	assert.NoError(err)
	assert.NotNil(cars)

	//assert
	assert.Equal(len(cars), 1000)

}

func TestInit_happyCase(t *testing.T) {
	mysql.Init()
}

func TestCastCarBase2Json(t *testing.T) {
	var assert = assert.New(t)
	// var ctx =
	var carBase = &car.CarBase{
		Name:           "冰魄",
		PicUrl:         "pic_url_test",
		RaceFeature:    "集满氮气给小喷",
		PropFeature:    "无",
		RecordNum:      1,
		LevelRecordNum: 2,
		Level:          car.Level_A,
		BurnTime:       "2020-02-07",
	}
	var dao = mysql.GetSingletonCarDAO()

	//act
	carBaseStr, err := dao.CastCarBase2Json(carBase)
	assert.NoError(err)

	//assert
	t.Log(carBaseStr)

}

func TestCastJson2CarBase_happy_case(t *testing.T) {
	var assert = assert.New(t)
	// var ctx =
	var carBase = &car.CarBase{
		Name:           "冰魄",
		PicUrl:         "pic_url_test",
		RaceFeature:    "集满氮气给小喷",
		PropFeature:    "无",
		RecordNum:      1,
		LevelRecordNum: 2,
		Level:          car.Level_A,
		BurnTime:       "2020-02-07",
	}
	var dao = mysql.GetSingletonCarDAO()

	//act
	carBaseStr, err := dao.CastCarBase2Json(carBase)
	if err != nil {
		t.Log(err)
	}

	carBase1, err := dao.CastJson2CarBase(carBaseStr)
	assert.NoError(err)

	assert.NotNil(carBase1)
	assert.Equal(carBase1.Name, "冰魄")
}
