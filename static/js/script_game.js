var questions;
var currentQuestionIndex = -1;

function initQuestions() {
	console.log("Initializing questions..");
	resetPlayerPercentage();
	questions = new Array();
	addSampleQuestions();
}

function addSampleQuestions() {
	console.log("Adding sample questions..");

	addQuestion("What's the capital city of Germany?",
		["Berlin", "Potsdam", "Köln", "Kleinmachnow"],
		0);

	addQuestion("How many times did Germany win the soccer world cup?",
		["3", "4", "6", "7"],
		1);

	console.log(questions);

	showNextQuestion();
}

function addQuestion(phraze, answers, rightAnswer) {
	var question = new Object();
	question.phraze = phraze;
	question.answers = answers;
	question.rightAnswer = rightAnswer;
	questions.push(question);
}

function showNextQuestion() {
	currentQuestionIndex += 1;
	if (currentQuestionIndex > questions.length - 1) {
		alert("All questions answered");
	} else {
		showQuestion(questions[currentQuestionIndex]);
	}
}

function showQuestion(question) {
	div_question_phraze.innerHTML = question.phraze;
	div_question_answer_1.innerHTML = question.answers[0];
	div_question_answer_2.innerHTML = question.answers[1];
	div_question_answer_3.innerHTML = question.answers[2];
	div_question_answer_4.innerHTML = question.answers[3];
}

function processInput(value) {
	if (value > 0 && value < 5) {
		console.log("Processing input: " + value);
		
	} else {
		console.log("Invalid input: " + value);
	}
}