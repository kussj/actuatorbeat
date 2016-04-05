package beater

import (
        "net/http"
        "io/ioutil"
        "encoding/json"
        "github.com/elastic/beats/libbeat/logp"
)

func (ab *Actuatorbeat) GetMetricsActuator(u string) (map[string]float64, error) {
        metrics := make(map[string]float64)

        resp, err := http.Get(u)
        defer resp.Body.Close()

        if err != nil {
                logp.Err("An error occured while executing HTTP request: %v", err)
                return metrics, err
        }

        // read json http response
	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

        if err != nil {
                logp.Err("An error occured while reading HTTP response: %v", err)
                return metrics, err
        }

	err = json.Unmarshal([]byte(jsonDataFromHttp), &metrics)

        if err != nil {
                logp.Err("An error occured while unmarshaling metrics actuator data: %v", err)
                return metrics, err
        }
        return metrics, nil
}
