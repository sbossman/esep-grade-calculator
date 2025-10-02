package esepunittests


type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades:       make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{ Name:  name, Grade: grade, Type:  gradeType })
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignments := make([]Grade, 0)
	exams := make([]Grade, 0)
	essays := make([]Grade, 0)
	
	for i := 0; i < len(gc.grades); i++{
		gradeType := gc.grades[i].Type
		switch gradeType{
		case Assignment:
			assignments = append(assignments, gc.grades[i]);
		case Exam:
			exams = append(exams, gc.grades[i]);
		case Essay:
			essays = append(essays, gc.grades[i]);
		}

	}
	
	assignment_average := computeAverage(assignments)
	exam_average := computeAverage(exams)
	essay_average := computeAverage(exams)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(gradeList []Grade) int {
	sum := 0

	for _, grade := range gradeList {
		sum += grade.Grade
	}


	return sum / len(gradeList)
}
