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
	console.log("Showing next question");
	resetVisualizations();

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

function isRightAnswer(value) {
	if (questions[currentQuestionIndex].rightAnswer != value - 1) {
		return false;
	} else {
		return true;
	}
}

function resetVisualizations() {
	var defaultAnswerColor = "#dd630d";
	document.getElementById("answer1Container").children[0].style.backgroundColor = defaultAnswerColor;
	document.getElementById("answer2Container").children[0].style.backgroundColor = defaultAnswerColor;
	document.getElementById("answer3Container").children[0].style.backgroundColor = defaultAnswerColor;
	document.getElementById("answer4Container").children[0].style.backgroundColor = defaultAnswerColor;
}

function visualizeRightAnswer() {
	var div = document.getElementById("answer" + (questions[currentQuestionIndex].rightAnswer + 1) + "Container").children[0];
	div.style.backgroundColor = "#390";
}

function visualizeWrongAnswer(value) {
	var div = document.getElementById("answer" + value + "Container").children[0];
	div.style.backgroundColor = "#af0039";
}

function processInput(value) {
	if (value > 0 && value < 5) {
		console.log("Processing input: " + value);
		if (isRightAnswer(value)) {
			console.log("Right answer");
			visualizeRightAnswer();
			window.setTimeout("showNextQuestion()", 2000);
		} else {
			console.log("Wrong answer");
			visualizeWrongAnswer(value);
		}
	} else {
		console.log("Invalid input: " + value);
	}
}

