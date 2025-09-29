# ESEP Grade Calculator

This library is a simple utility to calculate your grades. It takes in assignment, exam, and essay grades, then uses a weighting formula in order to calculate the total grade. (Exams, essays, and assignments are all worth different parts of the grade, thus should be weighted differently.) The calculator then uses that total grade to, based on the letter grading curve, determine what letter grade (A-F) someone should receive. Grades are added to a NewGradeCalculator object, and computations are done with the GetFinalGrade() function.

The grading scale is as follows
- A: >= 90
- B: 80-90
- C: 70-80
- D: 60-70
- F: <70

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```
