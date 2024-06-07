namespace go helper.car

struct Car {
    1: CarBase CarBase
    2: CarExtra CarExtra
}

struct CarBase {
    1: string Name 
    2: string PicUrl
    3: Feature Feature
    4: i32 RecordNum
}

struct CarExtra {
    1: Feature Feature
    2: list<string> Attributes
    3: list<string> MatchedMaps
    4: string Comments
}

struct Feature {
    1: string RaceFeature //竞速特性
    2: string PropFeature //道具特性
}


