package utils

import (
	"regexp"

	"github.com/kunlanat/pie-fire-dire/model"
)

func FilterWorld(data string) *model.BeefModel {
	re := regexp.MustCompile("[^A-Za-z-]")

	counter := make(map[string]int)

	results := re.Split(data, -1)
	for _, v := range results {
		if v == "" {
			continue
		}
		if counter[v] != 0 {
			counter[v] = counter[v] + 1
		} else {
			counter[v] = 1
		}

	}

	return &model.BeefModel{Beef: counter}
}
