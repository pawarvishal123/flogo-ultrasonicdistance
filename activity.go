package ultrasonicdistance

import (
	"errors"
	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/op/go-logging"
	"github.com/stianeikeland/go-rpio"
	"time"
)

// log is the default package logger
var log = logging.MustGetLogger("activity-tibco-rest")

const (
	triggerPin = "triggerPin"
	echoPin    = "echoPin"
	distance   = "distance"
)

// UltrasonicDistanceActivity type 
type UltrasonicDistanceActivity struct {
	metadata *activity.Metadata
}

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(&UltrasonicDistanceActivity{metadata: md})
}

// Metadata implements activity.Activity.Metadata
func (a *UltrasonicDistanceActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *UltrasonicDistanceActivity) Eval(context activity.Context) (done bool, err error) {

	log.Debug("Running ultrasonic distance activity.")

	//get pinNumber
	triggerPinNumber := context.GetInput(triggerPin).(int)
	echoPinNumber := context.GetInput(echoPin).(int)

	log.Debugf("Trigger Pin: %d  Echo Pin: %d", triggerPinNumber, echoPinNumber)

	//Open pin
	openErr := rpio.Open()
	if openErr != nil {
		log.Errorf("Open RPIO error: %+v", openErr.Error())
		return true, errors.New("Open RPIO error: " + openErr.Error())
	}

	trig := rpio.Pin(triggerPinNumber)
	echo := rpio.Pin(echoPinNumber)
	trig.Output()
	echo.Input()

	trig.Low()
	time.Sleep(time.Second * 2)

	trig.High()
	time.Sleep(time.Microsecond * 10)
	trig.Low()

	startTime := time.Now()
	endTime := time.Now()
	
	log.Debugf("Start Time: %t  End Time: %t", startTime, endTime)
	
	for int(echo.Read()) == 0 {
		startTime = time.Now()
	}

	for int(echo.Read()) == 1 {
		endTime = time.Now()
	}

	log.Debugf("Start Time: %t  End Time: %t", startTime, endTime)
	
	duration := endTime.Sub(startTime)
	distanceVal := duration * 17150
	//distanceVal := round(distanceVal, 2)
	log.Debugf("Distance(cm): %f", distanceVal)

	context.SetOutput(distance, distanceVal)

	return true, nil
}
