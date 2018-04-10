package ultrasonicdistance

var jsonMetadata = `{
  "name": "ultrasonicdistance",
  "version": "0.0.1",
  "description": "activity to read distance using ultrasonic sensor on raspberry pi",
  "author": "Vishal Pawar <vpawar@tibco.com>",
  "inputs":[
    {
      "name": "triggerPin",
      "type": "integer",
	  "required": true
    },
	{
      "name": "echoPin",
      "type": "integer",
	  "required": true
    }
  ],
  "outputs": [
    {
      "name": "distance",
      "type": "float"
    }
  ]
}`
