package main

import (
	"os"
	"strings"
	"time"

	"github.com/minskylab/covid19-entity"
	neo "github.com/minskylab/neocortex"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := extractConfigFromEnv()
	if err != nil {
		panic(errors.Cause(err))
	}

	brain, err := loadCognitive(config)
	if err != nil {
		panic(errors.Cause(err))
	}

	channels, err := loadChannels(config)
	if err != nil {
		panic(errors.Cause(err))
	}
	fb := channels[0]

	engine, err := neo.Default(nil, brain, channels...)
	if err != nil {
		panic(errors.Cause(err))
	}

	accountID := os.Getenv("TWILIO_ACCOUNT_ID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	emitter, err := covid19.NewEmitter(accountID, authToken)
	if err != nil {
		panic(errors.Cause(err))
	}

	title := neo.IfDialogNodeTitleIs("Episodio 3")
	engine.Resolve(fb, title, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		c.SetContextVariable("name", c.Person.Name)
		dni := c.GetStringContextVariable("dni", "00000000")
		phone := c.GetStringContextVariable("phone", "+51957821858")

		phone = strings.TrimSpace(phone)
		if !strings.HasPrefix(phone, "+51") {
			phone = "+51" + phone
		}

		log.WithFields(log.Fields{
			"dni":   dni,
			"phone": phone,
		}).Info("generating sms alert")

		response(c, out)

		timer, err := emitter.LogMeasureBySMS(phone, c.Person.Name, 15*time.Minute)
		if err != nil {
			return errors.Wrap(err, "error at emmit message on episodio 3 resolve")
		}

		go func(to, name, dni string, timer *time.Timer) {
			<-timer.C
			if err := emitter.SendRemember(to, name, dni); err != nil {
				panic(err)
			}
		}(phone, c.Person.Name, dni, timer)

		return nil
	})

	engine.ResolveAny(fb, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		c.SetContextVariable("name", c.Person.Name)
		return response(c, out)
	})

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
