package calculator_test

import (
	"strings"
	"testing"

	"github.com/TaCkeR00/gradecalc/calculator"
)

func TestInternalLogic(t *testing.T) {
	jsonInput := `
	{
		"split": 0.6,
  	"modules": [
    	{
      	"name": "analysis",
      	"coff": 4,
      	"credit": 6,
      	"type": 2,
				"grades": {
					"exam": 4.0,
					"td": 10.0
				}
			},
			{
				"name": "algorithm",
				"coff": 4,
				"credit": 6,
				"type": 3,
				"grades": {
					"exam": 14.75,
					"td": 16.75,
					"tp": 16.5
				}
			}
		]
	}
	`

	modules, err := calculator.FromJSON(strings.NewReader(jsonInput))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if modules.Split != 0.6 {
		t.Fatalf("got %v want 0.6", modules.Split)
	}

	if len(modules.Modules) != 2 {
		t.Fatalf("expected 2 modules, got %d", len(modules.Modules))
	}

	if modules.Modules[0].Name != "analysis" {
		t.Errorf("expected 'analysis', got %s", modules.Modules[0].Name)
	}

	modules.CalculateGrade()

	expectedGrade := float32(10.95)
	expectedCredit := int8(6)

	if modules.Grade() != expectedGrade {
		t.Errorf("Grade() = %v, want %v", modules.Grade(), expectedGrade)
	}
	if modules.Credit() != expectedCredit {
		t.Errorf("Credit() = %v, want %v", modules.Credit(), expectedCredit)
	}
}
