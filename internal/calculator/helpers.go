package calculator

import (
	"encoding/json"
	"fmt"
	"io"
)

func (m *Module) CalculateGrade(split float32) {
	switch m.Type {
	case ModuleType1:
		m.grade = m.Grades.Exam
	case ModuleType2:
		m.grade = (m.Grades.Exam * split) + (m.Grades.Td * (1 - split))
	case ModuleType3:
		m.grade = (m.Grades.Exam * split) + (m.Grades.Td * ((1 - split) / 2)) + (m.Grades.Tp * ((1 - split) / 2))
	}
}

func (m *Module) Grade(split float32) float32 {
	if m.grade == 0 {
		m.CalculateGrade(split)
	}
	return m.grade
}

func (m *Modules) Grade() float32 {
	if m.grade == 0 && m.credit == 0 {
		m.CalculateGrade()
	}

	return m.grade
}

func (m *Modules) Credit() int8 {
	if m.grade == 0 && m.credit == 0 {
		m.CalculateGrade()
	}

	return m.credit
}

func (m *Modules) CalculateGrade() {
	var totalGrade float32
	var totalCoff int8
	m.credit = 0

	for _, mod := range m.Modules {
		totalGrade += mod.Grade(m.Split) * float32(mod.Coff)
		totalCoff += mod.Coff
		if mod.Grade(m.Split) >= 10 {
			m.credit += mod.Credit
		}
	}

	m.grade = totalGrade / float32(totalCoff)
}

func FromJSON(r io.Reader) (*Modules, error) {
	data := new(Modules)

	err := json.NewDecoder(r).Decode(data)
	if err != nil {
		return nil, fmt.Errorf("decoding modules: %w", err)
	}

	return data, nil
}
