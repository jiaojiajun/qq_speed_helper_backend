package mysql

import (
	"encoding/json"
	"errors"

	model "github.com/jiaojiajun/qq_speed_helper_backend/biz/model/domain_model/car"
	"github.com/jiaojiajun/qq_speed_helper_backend/biz/model/helper/car"
)

type CarDAO struct {
}

var carDao *CarDAO = initCarDAO()

func NewCarDAO() *CarDAO {
	return &CarDAO{}
}

func initCarDAO() *CarDAO {
	return NewCarDAO()
}

func GetSingletonCarDAO() *CarDAO {
	return carDao
}

func (p *CarDAO) GetAllCars() ([]*car.Car, error) {

	cars := []*car.Car{}
	db := DB.Model(car.Car{})
	_ = db.Find(&cars)
	if len(cars) == 0 {
		return nil, errors.New("查询db出错")
	}
	return cars, nil

}

func (p *CarDAO) GetCarBaseById(carId int32) (*car.CarBase, error) {
	db := DB.Model(model.CasrBaseJson{})

	var carBaseJson *model.CasrBaseJson
	_ = db.First(carBaseJson, "CarId = ?", carId)
	if carBaseJson == nil {
		return nil, errors.New("找不到 car id对应的car base")
	}

	carBase, err := p.CastJson2CarBase(carBaseJson.CarBaseStr)
	if err != nil {
		return nil, err
	}
	return carBase, nil

}

func (p *CarDAO) CastJson2CarBase(carBaseStr string) (*car.CarBase, error) {
	if carBaseStr == "" {
		return nil, errors.New("carbase string 为空，json转化失败")
	}
	var carBase *car.CarBase = &car.CarBase{}

	// 传给 unmarshal的第二个参数必须为非空指针
	err := json.Unmarshal([]byte(carBaseStr), carBase)
	if err != nil {
		return nil, errors.New("json unmashal 失败")
	}
	return carBase, nil
}

func (p *CarDAO) CastCarBase2Json(carBase *car.CarBase) (string, error) {
	bytes, err := json.Marshal(*carBase)
	if err != nil {
		return "", errors.New("转化carbaseweijson失败")
	}
	return string(bytes), nil
}
