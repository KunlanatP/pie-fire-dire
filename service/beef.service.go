package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kunlanat/pie-fire-dire/model"
	"github.com/kunlanat/pie-fire-dire/utils"
)

type BeefService interface {
	GetBeef(ctx context.Context) (*model.BeefModel, error)
}

type beefServiceImpl struct{}

func BeefServiceImpl() BeefService {
	return &beefServiceImpl{}
}

func (s *beefServiceImpl) GetBeef(ctx context.Context) (*model.BeefModel, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return utils.FilterWorld(string(responseData)), nil
}
