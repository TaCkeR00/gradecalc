package main

import (
	"fmt"

	"simple-api/calculator"
)

func ReadGradesFromUser(m *calculator.Modules) error {
	for i := 0; i < len(m.Modules); i++ {
		var err error
		fmt.Println(m.Modules[i].Name)
		switch m.Modules[i].Type {
		case calculator.ModuleType3:
			m.Modules[i].Grades.Tp, err = scanGrade("TP", m.Modules[i].Name)
			if err != nil {
				return err
			}
			fallthrough
		case calculator.ModuleType2:
			m.Modules[i].Grades.Td, err = scanGrade("TD", m.Modules[i].Name)
			if err != nil {
				return err
			}
			fallthrough
		case calculator.ModuleType1:
			m.Modules[i].Grades.Exam, err = scanGrade("Exam", m.Modules[i].Name)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type %d for %s", m.Modules[i].Type, m.Modules[i].Name)
		}
	}
	return nil
}

func scanGrade(label, moduleName string) (float32, error) {
	fmt.Printf("  %s: ", label)
	var val float32
	if _, err := fmt.Scan(&val); err != nil {
		return 0, fmt.Errorf("reading %s for %s: %w", label, moduleName, err)
	}
	if val < 0 || val > 20 {
		return 0, fmt.Errorf("%s grade for %s must be between 0 and 20 (got %v)", label, moduleName, val)
	}
	return val, nil
}

func PrintResult(m *calculator.Modules) {
	if m.Grade() == 0 && m.Credit() == 0 {
		m.CalculateGrade()
	}
	fmt.Printf("grade: %f, credit: %d \n", m.Grade(), m.Credit())
}
