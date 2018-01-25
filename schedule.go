package jpushclient


import (
	"encoding/json"
)

type Single struct {
	Time string `json:"time"`
}


type Periodical struct {
	Start     string   `json:"start"`
	End       string   `json:"end"`     
	Time      string   `json:"time"`
	Time_unit string   `json:"time_unit"`
	Frequency int	   `json:"frequency"`
	Point     []string `json:"point,omitempty"`   
	
}

type Trigger struct {
	Single     *Single     `json:"single,omitempty"`
	Periodical *Periodical `json:"periodical,omitempty"`
}

type Schedule struct {
	Cid     string   `json:"cid,omitempty"`
	Name    string   `json:"name"`
	Enabled bool     `json:"enabled"`
	Trigger *Trigger `json:"trigger"`
	Push    *PayLoad `json:"push"`
}

func (this *Schedule)SetCid( cid string ) {
	this.Cid = cid
}

func (this *Schedule) SetName( name string ) {
	this.Name = name
}

func (this *Schedule) SetEnabled( enabled bool ) {
	this.Enabled = enabled
}


func (this *Schedule) SetPayload( payload *PayLoad ) {
	this.Push = payload
}


func (this *Schedule) SetSingleSchedule ( time string ) {
	single := &Single{ Time:time }

	if this.Trigger == nil {
		trigger := &Trigger{ Single:single }
		this.Trigger = trigger
	} else {
		this.Trigger.Single = single
	}

}

func (this *Schedule) SetPeriodicalSchedule ( periodical *Periodical) {

	if this.Trigger == nil {
		trigger := &Trigger{Periodical:periodical}
		this.Trigger = trigger
	} else {
		this.Trigger.Periodical = periodical
	}	
}


func (this *Schedule) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
