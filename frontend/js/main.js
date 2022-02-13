port = 8080
url = 'http://localhost:' + port + '/'

var creatures = []
var paused = true


function pause() {
    button = document.getElementById('pause')
    if (paused) {
        button.value = "pause"
        paused = false
    }
    else {
        button.value = "start"
        paused = true
    }
}

function cycle() {
    fetch(url + 'cycle')
    refreshBoard()
}

function reset() {
    fetch(url + 'reset')
    refreshBoard()
}

function refreshBoard() {

    function updateBoard(data) {
        creatures = data
        console.log(creatures)
    }

    fetch(url + 'board')
    .then(response => response.json())
    .then(data => updateBoard(data));
}


function generateWorld() {
    fetch(url + 'worldsize')
    .then(response => response.json())
    .then(data => createDivs(data));
}

function drawCreature(x, y) {
    ellipse(x, y, 8, 8)
}


function setup() {
    createCanvas(800, 800);
    frameRate(3);
}

function draw() {           
    if (!paused){
        cycle()
    }                                                    

    background(220, 75);
    creatures.forEach(c => drawCreature(c.X, c.Y))
}