package grades

//  init函数在包被使用的时候首先运行
func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Carter",
			Grades: []Grade{
				{Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85},
				{Title: "final exam",
					Type:  GradeExam,
					Score: 94},
				{Title: "quiz 2",
					Type:  GradeQuiz,
					Score: 82},
			},
		},
		{
			ID:        2,
			FirstName: "roberto",
			LastName:  "Baggio",
			Grades: []Grade{
				{Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 100},
				{Title: "final exam",
					Type:  GradeExam,
					Score: 100},
				{Title: "quiz 2",
					Type:  GradeQuiz,
					Score: 81},
			},
		}, {
			ID:        3,
			FirstName: "emma",
			LastName:  "stone",
			Grades: []Grade{
				{Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 67},
				{Title: "final exam",
					Type:  GradeExam,
					Score: 0},
				{Title: "quiz 2",
					Type:  GradeQuiz,
					Score: 75},
			},
		}, {
			ID:        4,
			FirstName: "rachel",
			LastName:  "mcadams",
			Grades: []Grade{
				{Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 98},
				{Title: "final exam",
					Type:  GradeExam,
					Score: 99},
				{Title: "quiz 2",
					Type:  GradeQuiz,
					Score: 94},
			},
		}, {
			ID:        5,
			FirstName: "kelly",
			LastName:  "clarkson",
			Grades: []Grade{
				{Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 95},
				{Title: "final exam",
					Type:  GradeExam,
					Score: 100},
				{Title: "quiz 2",
					Type:  GradeQuiz,
					Score: 97},
			},
		},
	}
}
