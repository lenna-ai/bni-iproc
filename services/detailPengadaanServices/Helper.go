package detailpengadaanservices

func splitStringWhere(whereKeySplit []string, whereValueSplit []string, whereSymbolSplit []string) string {
	var dataWhereResultMap = make(map[string][]interface{})

	for wks := 0; wks < len(whereKeySplit); wks++ {
		for wss := 0; wss < len(whereSymbolSplit); wss++ {
			dataWhereResultMap[whereKeySplit[wss]] = append(dataWhereResultMap[whereKeySplit[wss]], whereSymbolSplit[wss])
		}
		for wvs := 0; wvs < len(whereValueSplit); wvs++ {
			dataWhereResultMap[whereKeySplit[wvs]] = append(dataWhereResultMap[whereKeySplit[wvs]], whereValueSplit[wvs])
		}
		break
	}
	tempWhereClauses := ""
	var countDataWhereResultMap = 0
	for key, values := range dataWhereResultMap {
		tempWhereClauses += key
		for _, value := range values {
			tempWhereClauses += " " + value.(string)
		}
		countDataWhereResultMap++
		if countDataWhereResultMap < len(dataWhereResultMap) {
			tempWhereClauses += " AND "
		}
	}

	return tempWhereClauses
}
