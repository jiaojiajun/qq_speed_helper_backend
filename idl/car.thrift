namespace go helper.car

struct Car {
    1: i32 Id 
    2: CarBase CarBase
    3: CarExtra CarExtra

}

struct CarBase {
    1: string Name 
    2: string PicUrl
    3: string RaceFeature
    4: string PropFeature
    5: i32 RecordNum
    6: i32 LevelRecordNum
    7: Level Level
    8: string BurnTime
}

struct CarExtra {
    1: Attribute Attribute
    2: string MatchedMaps
    3: list<string> Comments
}

struct Attribute {
    8: double MaxCommSpeed
    9: double SteeringSpeed
    10: double SpraySpeed
    11: double NitrogenSpeed
    12: double CWWSpeed
    13: double WCWWSpeed
    14: double SprayDuration
    15: double NitrogenDuration
    16: double TurnCircleTime
    17: double DriftCircleTime
    18: double SprayPower
    19: double NitrogenPower
    20: double AccTime
    21: double BasicPower
    22: double NitrogenRatio
}



enum Level {
    S = 0
    T = 1
    A = 2
    B = 3
    C = 4
    D = 5
}


// struct Attribute {
//     1: double MaxCommSpeed
//     2: double SteeringSpeed
//     3: double SpraySpeed
//     4: double NitrogenSpeed
//     5: double CWWSpeed
//     6: double WCWWSpeed
//     7: double SprayDuration
//     8: double NitrogenDuration
//     9: double TurnCircleTime
//     10: double DriftCircleTime
//     11: double SprayPower
//     12: double NitrogenPower
//     13: double AccTime
//     14: double BasicPower
//     15: double NitrogenRatio
// }

enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}

struct GetAllCarsRequest {

}

struct GetAllCarsResp {
    1: Code Code
    2: list<Car> Cars  
}

service CarService {
    GetAllCarsResp GetAllCars(1: GetAllCarsRequest request) (api.get = "v1/car/get_all")
}


