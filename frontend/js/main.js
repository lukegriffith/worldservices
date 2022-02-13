var creatures = []
var paused = true
var fps = 32
var cycleNum

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

function updateCycleStatus(){
    div = document.getElementById('cycle')
    div.innerHTML = cycleNum
}

function cycle() {
    cycleNum++
    updateCycleStatus()
    fetch('cycle')
    refreshBoard()
}

function reset() {
    cycleNum = 0
    updateCycleStatus()
    size = document.getElementById('size')
    population = document.getElementById('pop')
    fetch('reset?worldsize=' + size.value + '&pop=' + population.value)
    refreshBoard()
}

function clearBoard() {
    background(220, 75);
}

function refreshBoard() {
    function updateBoard(data) {
        creatures = data
    }
    fetch('board')
    .then(response => response.json())
    .then(data => updateBoard(data));
}


function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function drawCreature(x, y) {
    setTimeout(10)
    ellipse(x, y, 8, 8)
}


function setup() {
    createCanvas(800, 800);
    frameRate(fps);
}

function draw() {           
    if (!paused){
        cycle()
    }                                                    
    background(220, 75);
    creatures.forEach(c => drawCreature(c.X, c.Y))
}