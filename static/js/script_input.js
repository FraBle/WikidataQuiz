// Save Gamestate
// 0 = buzzer
// 1 = answer selection
var gameState = 0;
var currentSelectedNumber = -1;

// Get the canvas DOM element
var canvas = document.getElementById('leapCanvas');

// Making sure we have the proper aspect ratio for our canvas
canvas.width = canvas.clientWidth;
canvas.height = canvas.clientHeight;

// Create the context we will use for drawing
var c = canvas.getContext('2d');

// Save the canvas width and canvas height as easily accessible variables
var width = canvas.width;
var height = canvas.height;

// Creating a global Frame variable that we can access throughout the program
var frame;

// Global keyTap and screenTap arrays
var keyTaps = [];
var KEYTAP_LIFETIME = .5;
var KEYTAP_START_SIZE = 15;

var screenTaps = [];
var SCREENTAP_LIFETIME = 1;
var SCREENTAP_START_SIZE = 30;

/*
The leapToScene function takes a position in leap space
and converts it to the space in the canvas.

It does this by using the interaction box, in order to
make sure that every part of the canvas is accesible
in the interaction area of the leap
*/

function leapToScene(leapPos) {

    // Gets the interaction box of the current frame
    var iBox = frame.interactionBox;

    // Gets the left border and top border of the box
    // In order to convert the position to the proper
    // location for the canvas
    var left = iBox.center[0] - iBox.size[0] / 2;
    var top = iBox.center[1] + iBox.size[1] / 2;

    // Takes our leap coordinates, and changes them so
    // that the origin is in the top left corner
    var x = leapPos[0] - left;
    var y = leapPos[1] - top;

    // Divides the position by the size of the box
    // so that x and y values will range from 0 to 1
    // as they lay within the interaction box
    x /= iBox.size[0];
    y /= iBox.size[1];

    // Uses the height and width of the canvas to scale
    // the x and y coordinates in a way that they
    // take up the entire canvas
    x *= width;
    y *= height;

    // Returns the values, making sure to negate the sign
    // of the y coordinate, because the y basis in canvas
    // points down instead of up
    return [x, -y];
}

function calculatePlayer(pos) {
    if (pos[0] <= width/2) {
        setActivePlayer(0);
    } else {
        setActivePlayer(1);
    }
}

function onKeyTap(gesture) {
    var pos = leapToScene(gesture.position);
    var time = frame.timestamp;
    keyTaps.push([pos[0], pos[1], time]);
    if (gameState === 0){
        calculatePlayer(pos);
        gameState = 1;
        window.setTimeout("processInput(currentSelectedNumber);gameState=0", 5000);
    }
}

function updateKeyTaps() {
    for (var i = 0; i < keyTaps.length; i++) {
        var keyTap = keyTaps[i];
        var age = frame.timestamp - keyTaps[i][2];
        age /= 1000000;
        if (age >= KEYTAP_LIFETIME) {
            keyTaps.splice(i, 1);
        }
    }
}

function drawKeyTaps() {
    for (var i = 0; i < keyTaps.length; i++) {
        var keyTap = keyTaps[i];
        var x = keyTap[0];
        var y = keyTap[1];
        var age = frame.timestamp - keyTap[2];
        age /= 1000000;
        var completion = age / KEYTAP_LIFETIME;
        var timeLeft = 1 - completion;

        // Static Ring
        c.strokeStyle = "#FF2300";
        c.lineWidth = 3;
        c.beginPath();
        c.arc(x, y, KEYTAP_START_SIZE, 0, Math.PI * 2);
        c.closePath();
        c.stroke();

        var opacity = timeLeft;
        var radius = KEYTAP_START_SIZE * timeLeft;
        if (radius < 0) {
            radius = 0;
        }

        c.fillStyle = "rgba( 256 , 33 , 0 , " + opacity + ")";

        // Creating the path for the finger circle
        c.beginPath();
        c.arc(x, y, radius, 0, Math.PI * 2);
        c.closePath();
        c.fill();

    }

}

function onScreenTap(gesture) {
    var pos = leapToScene(gesture.position);
    var time = frame.timestamp;
    screenTaps.push([pos[0], pos[1], time]);
}

function updateScreenTaps() {
    for (var i = 0; i < screenTaps.length; i++) {
        var screenTap = screenTaps[i];
        var age = frame.timestamp - screenTaps[i][2];
        age /= 1000000;
        if (age >= SCREENTAP_LIFETIME) {
            screenTaps.splice(i, 1);
        }
    }
}

function drawScreenTaps() {
    for (var i = 0; i < screenTaps.length; i++) {
        var screenTap = screenTaps[i];
        var x = screenTap[0];
        var y = screenTap[1];
        var age = frame.timestamp - screenTap[2];
        age /= 1000000;
        var completion = age / SCREENTAP_LIFETIME;
        var timeLeft = 1 - completion;

        // Drawing the static ring
        c.strokeStyle = "#FFB300";
        c.lineWidth = 3;

        // Save the canvas context, so that we can restore it
        // and have it un affected
        c.save();

        // Translate the contex and rotate around the
        // center of the  square
        c.translate(x, y);

        //Starting x and y ( compared to the pivot point )
        var left = -SCREENTAP_START_SIZE / 2;
        var top = -SCREENTAP_START_SIZE / 2;
        var width = SCREENTAP_START_SIZE;
        var height = SCREENTAP_START_SIZE;

        // Draw the rectangle
        c.strokeRect(left, top, width, height);

        // Restore the context, so we don't draw everything rotated
        c.restore();

        // Drawing the non-static part
        var size = SCREENTAP_START_SIZE * timeLeft;
        var opacity = timeLeft;
        var rotation = timeLeft * Math.PI;

        c.fillStyle = "rgba( 255 , 179 , 0 , " + opacity + ")";

        c.save();

        c.translate(x, y);
        c.rotate(rotation);

        var left = -size / 2;
        var top = -size / 2;
        var width = size;
        var height = size;

        c.fillRect(left, top, width, height);

        c.restore();
    }
}

function drawHand() {
    // Loop through all hands that the frame sees
    for (var i = 0; i < frame.hands.length; i++) {
        var hand = frame.hands[i];

        // Get hands position, so that it can be passed through for drawing the connections
        var handPos = leapToScene(hand.palmPosition);

        // Loop through all the fingers
        for (var j = 0; j < hand.fingers.length; j++) {
            var finger = hand.fingers[j];

            // Get the fingers position in Canvas
            var fingerPos = leapToScene(finger.tipPosition);

            // ##First## Draw the Connection
            c.strokeStyle = "#FFA040";
            c.lineWidth = 3;
            c.beginPath();
            c.moveTo(handPos[0], handPos[1]);
            c.lineTo(fingerPos[0], fingerPos[1]);
            c.closePath();
            c.stroke();

            // ##Second## Draw the Finger
            c.strokeStyle = "#39AECF";
            c.lineWidth = 5;
            c.beginPath();
            c.arc(fingerPos[0], fingerPos[1], 20, 0, Math.PI * 2);
            c.closePath();
            c.stroke();
        }
        // ##Third## draw the hand
        c.fillStyle = "#FF5A40";
        c.beginPath();
        c.arc(handPos[0], handPos[1], 40, 0, Math.PI * 2);
        c.closePath();
        c.fill();
    }
}

function setCurrentSelectedNumber(){
    if (gameState === 1 && frame.hands.length > 0){
        currentSelectedNumber = frame.hands[0].fingers.length;
    } else {
        currentSelectedNumber = -1;
    }
    setSelectedNumber(currentSelectedNumber);
}

// Creates our Leap Controller
var controller = new Leap.Controller({
    enableGestures: true
});

// Tells the controller what to do every time it sees a frame
controller.on('frame', function(data) {
    // Assigning the data to the global frame object
    frame = data;

    // Clearing the drawing from the previous frame
    c.clearRect(0, 0, width, height);

    drawHand();
    setCurrentSelectedNumber();
    for (var i = 0; i < frame.gestures.length; i++) {
        var gesture = frame.gestures[i];
        var type = gesture.type;

        switch (type) {
            case "screenTap":
                onScreenTap(gesture);
                break;
            case "keyTap":
                onKeyTap(gesture);
                break;
        }
    }

    updateKeyTaps();
    drawKeyTaps();

    updateScreenTaps();
    drawScreenTaps();
});

// Get frames rolling by connecting the controller
controller.connect();
