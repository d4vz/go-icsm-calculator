package icms

const DEFAULT_ALIQUOTA = 18.0

var aliquotasICMS = map[string]map[string]float64{
	"SP": {
		"RJ": 12.0,
		"MG": 12.0,
		"PR": 12.0,
		"SC": 7.0,
	},
	"MG": {
		"SP": 12.0,
	},
	"RJ": {
		"SP": 12.0,
	},
}

func getAliquotaICMS(ufOrigin, ufDestination string) float64 {
	if destinos, ok := aliquotasICMS[ufOrigin]; ok {
		if aliquota, ok := destinos[ufDestination]; ok {
			return aliquota
		}
	}
	return DEFAULT_ALIQUOTA
}
