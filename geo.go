package google_trend_go

// Geo used for hot trend
type Geo int

const (
	GeoUnknown Geo = iota
	GeoHK
	GeoTW
	GeoCA
	GeoRU
	GeoDE
	GeoFR
	GeoNL
	GeoBR
	GeoGB
	GeoID
	GeoIN
	GeoMX
	GeoKR
	GeoTR
	GeoPH
	GeoES
	GeoIT
	GeoVN
	GeoEG
	GeoPL
	GeoCO
	GeoLI
	GeoTH
	GeoMY
	GeoUA
	GeoSA
	GeoKE
	GeoCL
	GeoRO
	GeoJP
	GeoZA
	GeoBE
	GeoSE
	GeoCZ
	GeoAT
	GeoHU
	GeoCH
	GeoPT
	GeoGR
	GeoDK
	GeoSG
	GeoFI
	GeoNO
	GeoNG
	GeoNZ
	GeoAR
	GeoIE
	GeoUS
	GeoUK
	GeoIL
	GeoAU
	GeoPE
)

type geoPtMapType map[Geo]string

var geoPtMap geoPtMapType = map[Geo]string{
	GeoHK: "p10",
	GeoTW: "p12",
	GeoCA: "p13",
	GeoRU: "p14",
	GeoDE: "p15",
	GeoFR: "p16",
	GeoNL: "p17",
	GeoBR: "p18",
	GeoID: "p19",
	GeoMX: "p21",
	GeoKR: "p23",
	GeoTR: "p24",
	GeoPH: "p25",
	GeoES: "p26",
	GeoIT: "p27",
	GeoVN: "p28",
	GeoEG: "p29",
	GeoIN: "p3",
	GeoPL: "p31",
	GeoCO: "p32",
	GeoLI: "p33",
	GeoTH: "p33",
	GeoMY: "p34",
	GeoUA: "p35",
	GeoSA: "p36",
	GeoKE: "p37",
	GeoCL: "p38",
	GeoRO: "p39",
	GeoJP: "p4",
	GeoZA: "p40",
	GeoBE: "p41",
	GeoSE: "p42",
	GeoCZ: "p43",
	GeoAT: "p44",
	GeoHU: "p45",
	GeoCH: "p46",
	GeoPT: "p47",
	GeoGR: "p48",
	GeoDK: "p49",
	GeoSG: "p5",
	GeoFI: "p50",
	GeoNO: "p51",
	GeoNG: "p52",
	GeoNZ: "p53",
	GeoAR: "p30",
	GeoIE: "p54",
	GeoUS: "p1",
	GeoUK: "p9",
	GeoGB: "p9",
	GeoIL: "p6",
	GeoAU: "p8",
	GeoPE: "",
}

func (t geoPtMapType) getOrDefualt(geo Geo) string {
	if v, ok := t[geo]; ok && len(v) != 0 {
		return v
	}
	return "p1" // default = us
}
