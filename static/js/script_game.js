var infiniteQuestions = true;
var questionCount = 10;
var questions;
var currentQuestionIndex = -1;

var players;
var currentPlayerIndex = 0;

function initPlayers() {
    console.log("Initializing players..");
    players = new Array();

    player1 = new Object();
    player1.name = "Player 1";
    player1.correct = 0;
    player1.wrong = 0;

    player2 = new Object();
    player2.name = "Player 2";
    player2.correct = 0;
    player2.wrong = 0;
    
    players.push(player1);
    players.push(player2);
}

function initQuestions() {
    console.log("Initializing questions..");
    resetPlayerPercentage();
    questions = new Array();
    
    //addSampleQuestions();
    getNextQuestion();
}

function addSampleQuestions() {
    console.log("Adding sample questions..");

    addQuestion("What's the capital city of Germany?", ["Berlin", "Potsdam", "Köln", "Kleinmachnow"],
        0);

    addQuestion("How many times did Germany win the soccer world cup?", ["2", "3", "6", "7"],
        1);

    console.log(questions);

    showNextQuestion();
}

function addQuestion(phrase, answers, rightAnswer) {
    var question = new Object();
    question.phrase = phrase;
    question.answers = answers;
    question.rightAnswer = rightAnswer;
    questions.push(question);
}

function getNextQuestion() {
    console.log("Requesting next question..");
    getRequest("/question", function(response) {
        console.log(response);
        var question = JSON.parse(response);
        questions.push(question);
        if (questions.length == 1) {
            // This is the first question loaded
            showNextQuestion();
        }
    })
}

function showNextQuestion() {
    console.log("Showing next question");
    resetVisualizations();

    currentQuestionIndex += 1;

    if (infiniteQuestions) {
        getNextQuestion();
        showQuestion(questions[currentQuestionIndex]);
        questionCount = currentQuestionIndex + 1;
    } else {
        if (questions.length < questionCount) {
            // Preload next question
            getNextQuestion();
        }
        if (currentQuestionIndex > questions.length - 1) {
            var message = questionCount + " of " + questionCount + " questions answered, ";
            if (players[0].correct > players[1].correct) {
                message += "Player 1 ";
            } else if (players[0].correct < players[1].correct) {
                message += "Player 2 ";
            } else {
                "No one ";
            }
            message += "won the game!";
            alert(message);
        } else {
            // Show next question
            showQuestion(questions[currentQuestionIndex]);
        }
    }    
}

function showQuestion(question) {
    div_question_phrase.innerHTML = question.phrase;
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
    var defaultAnswerColor = "#50c0e9";
    document.getElementById("answer1Container").children[0].style.backgroundColor = defaultAnswerColor;
    document.getElementById("answer2Container").children[0].style.backgroundColor = defaultAnswerColor;
    document.getElementById("answer3Container").children[0].style.backgroundColor = defaultAnswerColor;
    document.getElementById("answer4Container").children[0].style.backgroundColor = defaultAnswerColor;
    setLedColor("black");
}

function visualizeRightAnswer(useLed) {
    var div = document.getElementById("answer" + (questions[currentQuestionIndex].rightAnswer + 1) + "Container").children[0];
    div.style.backgroundColor = "#a8d324";
    if (useLed) {
        setLedColor("green");
    }
}

function visualizeWrongAnswer(value) {
    var div = document.getElementById("answer" + value + "Container").children[0];
    div.style.backgroundColor = "#ff5f5f";
    setLedColor("red");
}

function processInput(value) {
    if (value > 0 && value < 5) {
        console.log("Processing input: " + value);
        if (isRightAnswer(value)) {
            console.log("Right answer");
            players[currentPlayerIndex].correct += 1;
            visualizeRightAnswer(true);
            window.setTimeout("showNextQuestion()", 2000);
        } else {
            console.log("Wrong answer");
            players[currentPlayerIndex].wrong += 1;
            players[1 - currentPlayerIndex].correct += 1;
            visualizeWrongAnswer(value);
            window.setTimeout("visualizeRightAnswer(false)", 1000);
            window.setTimeout("showNextQuestion()", 4000);
        }
        updateScore();
    } else {
        console.log("Invalid input: " + value);
    }
}

function updateScore() {
    var player1Percentage = (players[0].correct * 100) / questionCount;
    setPlayer1Percentage(player1Percentage);

    var player2Percentage = (players[1].correct * 100) / questionCount;
    setPlayer2Percentage(player2Percentage);
}

function setActivePlayer(value) {
    currentPlayerIndex = value;
    document.getElementById("timer").innerHTML = "Player " + (value + 1);
}