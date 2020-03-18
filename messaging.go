package covid19

import (
	"fmt"
	"time"

	"github.com/sfreiberg/gotwilio"
	log "github.com/sirupsen/logrus"
)

type Emitter struct {
	twilio *gotwilio.Twilio
}

func NewEmitter(accountID, authToken string) (*Emitter, error) {
	twilio := gotwilio.NewTwilioClient(accountID, authToken)
	return &Emitter{twilio: twilio}, nil
}

func (emitter *Emitter) SendRemember(to string, name string, dni string) error {
	from := "+16122840701"

	message := fmt.Sprintf("Hola %s, es hora de una nueva evaluación; por favor comunicate con el chat 'Peru COVID-19' de messenger o entra aquí http://m.me/101423784830598", name)
	log.Infof("message to send: %s", message)

	res, exp, err := emitter.twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		return err
	}

	if exp != nil {
		return exp
	}

	log.Infof("total price: %s", res.Price)
	return nil
}

func (emitter *Emitter) LogMeasureBySMS(to string, name string, nextMeasurementDuration time.Duration) (*time.Timer, error) {
	from := "+16122840701"

	next := time.Now().Add(5 * time.Minute).Add(-5 * time.Hour)
	nextTimer := time.NewTimer(nextMeasurementDuration)

	nextString := next.Format("January 02, 2006 a las 15:04")
	nextMeasure := fmt.Sprintf("Tu siguiente medición, esta programada para %s, te estaremos avisando.", nextString)

	message := fmt.Sprintf("Hola %s, somos de ChatCOVID vamos estar pendientes de usted para reconocer cómo van evolucionando sus síntomas, por eso le estamos programando un seguimiento personalizado. %s", name, nextMeasure)

	log.Infof("message to send: %s", message)

	res, exp, err := emitter.twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		return nil, err
	}

	if exp != nil {
		return nil, exp
	}

	log.Infof("total price: %s", res.Price)
	return nextTimer, nil
}
