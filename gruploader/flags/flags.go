package flags

import (
	"github.com/pkg6/gruploader-action/gruploader/logger"
	"math/rand"
	"time"
)

type Flags struct {
	Cloud     string `validate:"required"`
	File      string `validate:"required"`
	Repo      string `validate:"required"`
	Tag       string `validate:"required"`
	Overwrite bool   `validate:"required"`
	Retry     uint   `validate:"required"`
}

func (f *Flags) DoRetryUpload(fn func(args *Flags) error) error {
	retry := f.Retry
	minDuration := 3
	maxDuration := 15
	for {
		retry--
		if err := fn(f); err != nil {
			if retry == 0 {
				return err
			} else {
				randomDuration := time.Duration(minDuration + rand.Intn(maxDuration-minDuration))
				retryDuration := time.Second * randomDuration
				logger.WarningF("Upload asset error, will retry in %s: %v\n", retryDuration.String(), err)
				time.Sleep(retryDuration) // retry after 3-15 seconds
				continue
			}
		}
		break
	}
	return nil
}
