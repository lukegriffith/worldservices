var creatures = []
var paused = true
var fps = 32
var cycleNum
var debugCreature 

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

function updateDebugCreature(value) {
    debugCreature = value
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
function updateConsole(creature){
    console_el = document.getElementById("debug_console")
    p = document.createElement("p")
    p.innerHTML = 'X: ' + creature.X + ' Y: ' + 
        creature.Y + ' Last Control Seq: ' + 
        creature.LastControlSequence
    console_el.innerHTML = ""
    console_el.appendChild(p)
}

function getCreatures(x,y){
    function updateDebug(creature_data) {
        creature_data.forEach(c => console.log(c))
        console.log("inner click")
    }
    console.log(x, y)
    fetch('creatures?X=' + x + '&Y=' + y)
    .then(response => response.json())
    .then(data => updateDebug(data));
}


function drawCreature(x, y, i, a) {
    setTimeout(10)
    if (i == debugCreature) {
        updateConsole(a[i])
        let c = color('red')
        fill(c)
    } 
    ellipse(x, y, 8, 8)
    let c = color('white')
    fill(c)
}

function canvasClickHandler() {
    console.log("click")
    getCreatures(int(mouseX), int(mouseY))
    draw()
}

function setup() {
    cnv = createCanvas(800, 800);
    cnv.mouseClicked(canvasClickHandler)
    frameRate(fps);
}

function next() {
    cycle()
}

function draw() {           
    if (!paused){
        cycle()
    }                                                    
    background(220, 75);
    creatures.forEach(function callback(c, i, a) { 
        drawCreature(c.X, c.Y, i, a)      
    });
    setTimeout(100)
}