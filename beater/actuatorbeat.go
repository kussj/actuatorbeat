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

type Actuatorbeat struct {
	beatConfig *config.Config
	done       chan struct{}
	period     time.Duration
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

	ticker := time.NewTicker(bt.period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
		}
		b.Events.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (bt *Actuatorbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (bt *Actuatorbeat) Stop() {
	close(bt.done)
}
