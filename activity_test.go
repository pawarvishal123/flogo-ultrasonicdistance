package ultrasonicdistance

import (
	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/TIBCOSoftware/flogo-lib/flow/test"
	"testing"
)

func TestRegistered(t *testing.T) {
	act := activity.Get("ultrasonicdistance")

	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	md := activity.NewMetadata(jsonMetadata)
	act := &UltrasonicDistanceActivity{metadata: md}

	tc := test.NewTestActivityContext(md)
	//setup attrs
	tc.SetInput("triggerPin", 23)
	tc.SetInput("echoPin", 24)

	act.Eval(tc)

	//check result attr
	val := tc.GetOutput("distance")
	log.Debugf("Distance %f", val)
}
