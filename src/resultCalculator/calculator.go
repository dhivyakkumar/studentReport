package resultCalculator

import (
	"sort"
	"studentReports/src/model"
)

func CalculateRank(studList []model.Student) []model.Student {
	for _, stud := range studList {
		stud.Total = stud.Subject1 + stud.Subject2
		stud.Avg = float32(stud.Total / 2)
	}

	studList = sortList(studList)

	for i := 0; i < len(studList); i++ {
		studList[i].Rank = i + 1
	}

	return studList
}

func sortList(studList []model.Student) []model.Student {
	sort.SliceStable(studList, func(i, j int) bool {
		return studList[i].Avg > studList[j].Avg
	})

	return studList
}
