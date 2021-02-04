package formula

import (
	"testing"
	"time"
)

func TestFormula_calculatePeriods(t *testing.T) {
	tests := []struct {
		name               string
		fields             Formula
		wantFirstDuration  time.Duration
		wantSecondDuration time.Duration
	}{
		{
			name: "Run with TRUE",
			fields: Formula{
				Entry1: "08h00m",
				Exit1:  "12h00m",
				Entry2: "13h00m",
				Exit2:  "18h30m",
			},
			wantFirstDuration:  time.Duration(4) * time.Hour,
			wantSecondDuration: time.Duration(5)*time.Hour + time.Duration(30)*time.Minute,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formula{
				Entry1: tt.fields.Entry1,
				Exit1:  tt.fields.Exit1,
				Entry2: tt.fields.Entry2,
				Exit2:  tt.fields.Exit2,
			}
			firstPeriod, secondPeriod := f.calculatePeriods()

			if firstPeriod != tt.wantFirstDuration {
				t.Errorf("Run() = %v, want %v", firstPeriod, tt.wantFirstDuration)
			}

			if secondPeriod != tt.wantSecondDuration {
				t.Errorf("Run() = %v, want %v", secondPeriod, tt.wantSecondDuration)
			}
		})
	}
}
