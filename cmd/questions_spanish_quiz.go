package cmd

var (
	FirstQuestion = Question{
		question: "What does 'arma' mean in Spanish?",
		options: []QuestionOptions{
			{
				option: "Arm",
				value:  false,
			},
			{
				option: "Weapon",
				value:  true,
			},
			{
				option: "Army",
				value:  false,
			},
		},
	}

	SecondQuestion = Question{
		question: "How do you say 'dog' in Spanish",
		options: []QuestionOptions{
			{
				option: "Perro",
				value:  true,
			},
			{
				option: "Dogo",
				value:  false,
			},
			{
				option: "Gato",
				value:  false,
			},
		},
	}

	ThridQuestion = Question{
		question: "How do you say 'embarrassed' in Spanish",
		options: []QuestionOptions{
			{
				option: "Embarazada",
				value:  false,
			},
			{
				option: "avergonzada",
				value:  true,
			},
			{
				option: "Embarrada",
				value:  false,
			},
		},
	}

	FourthQuestion = Question{
		question: "What does 'Disculparse' mean in Spanish?",
		options: []QuestionOptions{
			{
				option: "Disappointed",
				value:  false,
			},
			{
				option: "Peach",
				value:  false,
			},
			{
				option: "Apologize",
				value:  true,
			},
		},
	}

	FifthQuestion = Question{
		question: "If someone asks you in Spanish, '¿En dónde está la farmacia?' you answer...",
		options: []QuestionOptions{
			{
				option: "En la esquina.",
				value:  true,
			},
			{
				option: "Ven a la biblioteca.",
				value:  false,
			},
			{
				option: "Para eso estamos.",
				value:  false,
			},
		},
	}
)

// Here the questions and their options are added to a list:

var Questions = QuestionsList{
	questions: []Question{
		FirstQuestion,
		SecondQuestion,
		ThridQuestion,
		FourthQuestion,
		FifthQuestion,
	},
}
