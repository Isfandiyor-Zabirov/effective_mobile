package cars

import (
	"effective_mobile_tech_task/logger"
	"effective_mobile_tech_task/utils/env"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"sync"
	"time"
)

func GetAndSaveCars(request *AddCarsRequest) error {
	var (
		client        = &http.Client{Timeout: time.Second * 5}
		getCarInfoUrl = env.GetSettings().ExternalApi.GetCarInfoUrl + "/info"
		mu            sync.Mutex
		wg            sync.WaitGroup
		allResp       []ApiResponse
	)

	fmt.Println("checkpoint 2")

	for _, regNum := range request.RegNums {
		wg.Add(1)
		go func(num string) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			apiResp, err := sendRequestToExternalApi(getCarInfoUrl, num, client)
			if err != nil {
				return
			}
			allResp = append(allResp, apiResp)
		}(regNum)
	}

	wg.Wait()

	return saveCars(allResp)
}

func sendRequestToExternalApi(url, regNum string, client *http.Client) (apiResp ApiResponse, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error.Println("sendRequestToExternalApi func create new request error:", err.Error())
		return
	}

	apiResp.RegNum = regNum
	query := req.URL.Query()
	query.Add("regNum", regNum)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		logger.Error.Printf("sendRequestToExternalApi func regNum %s client execute request error: %s \n", regNum, err.Error())
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Printf("sendRequestToExternalApi func regNum %s read resp.Body error: %s \n", regNum, err.Error())
		return
	}

	fmt.Println("respBody:", string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &apiResp)
	if err != nil {
		logger.Error.Printf("sendRequestToExternalApi func regNum %s json unmarshal error: %s \n", regNum, err.Error())
		return
	}

	return
}

func UpdateCar(car *Car) error {
	return updateCar(car)
}

func DeleteCar(carID int) error {
	return deleteCarByID(carID)
}

func GetCars(filter GetCarsFilter) (resp GetCarsResponse, err error) {
	resp.Page = 1
	resp.Pages = 1
	resp.Cars = []Response{}

	cars, totalRows, page, pageLimit, err := getCars(filter)
	if err != nil {
		return resp, err
	}

	resp.Page = page
	resp.Pages = int(math.Ceil(float64(totalRows) / float64(pageLimit)))
	resp.TotalQuantity = int(totalRows)
	resp.Cars = cars

	return resp, nil
}
