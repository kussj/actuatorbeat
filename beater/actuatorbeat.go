package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/kussj/actuatorbeat/config"
)

const selector = "actuatorbeat"
const selectorDetail = "json"

type Actuatorbeat struct {
	beatConfig *config.Config
	done       chan struct{}
	period     time.Duration
	urls       []string
}

// Creates beater
func New() *Actuatorbeat {
	return &Actuatorbeat{
		done: make(chan struct{}),
	}
}

/// *** Beater interface methods ***///

func (bt *Actuatorbeat) Config(b *beat.Beat) error {

	// Load beater beatConfig
	err := cfgfile.Read(&bt.beatConfig, "")
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	// Sets default period if not specified
//	if bt.beatConfig.Actuatorbeat.Period != nil {
//		bt.period = time.Duration(*bt.beatConfig.Actuatorbeat.Period) * time.Second	
//	} else {
//		bt.period = 10 * time.Second
//	}

	// define default URL if none provided
	var urlConfig []string
	if bt.beatConfig.Actuatorbeat.URLs != nil {
		urlConfig = bt.beatConfig.Actuatorbeat.URLs
	} else {
		urlConfig = []string{"http://localhost:8080/metrics"}
	}

	bt.urls = make([]string, len(urlConfig))
	for i := 0; i < len(urlConfig); i++ {
		u := urlConfig[i]
		bt.urls[i] = u
	}

	logp.Debug(selector, "Init actuatorbeat")
	logp.Debug(selector, "Period %v\n", bt.period)
	logp.Debug(selector, "Watching: %v", bt.urls)

	return nil
}

func (bt *Actuatorbeat) Setup(b *beat.Beat) error {

	// Setting default period if not set
	if bt.beatConfig.Actuatorbeat.Period == "" {
		bt.beatConfig.Actuatorbeat.Period = "1s"
	}

	var err error
	bt.period, err = time.ParseDuration(bt.beatConfig.Actuatorbeat.Period)
	if err != nil {
		return err
	}

	return nil
}

func (bt *Actuatorbeat) Run(b *beat.Beat) error {
	logp.Info("actuatorbeat is running! Hit CTRL-C to stop it.")

	for _, u := range bt.urls {
		go func(u string) {
			ticker := time.NewTicker(bt.period)
			defer ticker.Stop()

			for {
				select {
				case <-bt.done:
					goto GotoFinish
				case <-ticker.C:
				}

				metrics, err := bt.GetMetricsActuator(u)
				if err != nil {
					logp.Err("Error reading metrics endpoint: %v", err)
				} else {
					logp.Debug(selectorDetail, "Metrics: %+v", metrics)
					event := common.MapStr{
						"@timestamp": common.Time(time.Now()),
						"type":       b.Name,
						"metrics":    metrics,
					}
					b.Events.PublishEvent(event)
					logp.Info("Event sent")
				}
			}
		GotoFinish:
		}(u)
	}
	<-bt.done
	return nil
}

func (bt *Actuatorbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (bt *Actuatorbeat) Stop() {
	close(bt.done)
}
