var content;

var div_question_phraze;
var div_question_answer_1;
var div_question_answer_2;
var div_question_answer_3;
var div_question_answer_4;

function load() {
    content = document.getElementById("content");
    div_question_phrase = document.getElementById("question");
    div_question_answer_1 = document.getElementById("answer1Container").children[0].children[1];
    div_question_answer_2 = document.getElementById("answer2Container").children[0].children[1];
    div_question_answer_3 = document.getElementById("answer3Container").children[0].children[1];
    div_question_answer_4 = document.getElementById("answer4Container").children[0].children[1];

    resizeContent();

    initPlayers();
    initQuestions();
}

function resetPlayerPercentage() {
    setPlayer1Percentage(0);
    setPlayer2Percentage(0);
}

function setPlayer1Percentage(value) {
    var score = document.getElementById("score1");
    var valueTotal = (document.getElementById("scoreContainer").offsetWidth / 2);
    var valueAbsolute = Math.max(15, (valueTotal * value) / 100);
    score.style.left = (valueTotal - valueAbsolute) + "px";
}

function setPlayer2Percentage(value) {
    var score = document.getElementById("score2");
    var valueTotal = (document.getElementById("scoreContainer").offsetWidth / 2);
    var valueAbsolute = Math.max(15, (valueTotal * value) / 100);
    score.style.right = (valueTotal - valueAbsolute) + "px";
}

function getRequest(url, callback) {
    var xmlhttp = null;
    if (window.XMLHttpRequest) {
        xmlhttp = new XMLHttpRequest();
    } else if (window.ActiveXObject) {
        xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
    }

    xmlhttp.open("GET", url, true);
    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState != 4) {
            // pending
        }
        if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
            // sent
            callback(xmlhttp.responseText);
        }
    }
    xmlhttp.send(null);
}

function resizeContent() {
    var screen_size = getScreenSize();
    content.style.width = screen_size[0] + "px";
    content.style.height = screen_size[1] + "px";
}

function getScreenSize() {
    var myWidth = 0,
        myHeight = 0;
    if (typeof(window.innerWidth) == 'number') {
        //Non-IE
        myWidth = window.innerWidth;
        myHeight = window.innerHeight;
    } else if (document.documentElement && (document.documentElement.clientWidth || document.documentElement.clientHeight)) {
        //IE 6+ in 'standards compliant mode'
        myWidth = document.documentElement.clientWidth;
        myHeight = document.documentElement.clientHeight;
    } else if (document.body && (document.body.clientWidth || document.body.clientHeight)) {
        //IE 4 compatible
        myWidth = document.body.clientWidth;
        myHeight = document.body.clientHeight;
    }
    return [myWidth, myHeight];
}

window.onresize = function(event) {
    resizeContent();
};
