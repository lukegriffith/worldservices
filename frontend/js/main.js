var creatures = []
var paused = true
var fps = 10
var cycleNum


function pause() {
    button = document.getElementById('pause')
    if (paused) {
        button.value = "pause"
        paused = false
        loop()
    }
    else {
        button.value = "start"
        paused = true
        noLoop()
    }
}

function updateCycleStatus(){
    div = document.getElementById('cycle')
    div.innerHTML = cycleNum
}




function cycle() {
    if (!paused){
        cycleNum++
        updateCycleStatus()
        fetch('cycle')
        refreshBoard()
    }       

}

function next() {
    updateCycleStatus()
    fetch('cycle')
    refreshBoard()
    background(220, 75);
    creatures.forEach(function callback(c, i, a) { 
        drawCreature(c.X, c.Y, i, a)      
    });
}


function resetCanvas() {
    // remove canvas
    cycleNum = 0
    world_section = document.getElementById("world")
    world_section.innerHTML = ""
    if (! paused) {
        pause()
    }
    setup()
}

function reset() {
    updateCycleStatus()
    size = document.getElementById('size')
    population = document.getElementById('pop')
    fetch('reset?worldsize=' + size.value + '&pop=' + population.value)
    resetCanvas()
}

function breed() {
    function updateWorld(data) {

    }
    fetch('breed')
    .then(response => resetCanvas())
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
    }
    console.log(x, y)
    fetch('creatures?X=' + x + '&Y=' + y)
    .then(response => response.json())
    .then(data => updateDebug(data))
    .then(r => setTimeout(next(), 100))

}


function drawCreature(x, y, i, a) {
    setTimeout(10)

    if (a[i].Debug) {
        let c = color('red')
        fill(c)
    }
    ellipse(x, y, 8, 8)
    let c = color('white')
    fill(c)
}

function canvasClickHandler() {
    getCreatures(int(mouseX), int(mouseY))
    draw()
}

function setup() {
    size = document.getElementById('size').value
    cnv = createCanvas(size, size);
    cnv.mouseClicked(canvasClickHandler)
    frameRate(fps);
}


function draw() {           
                                             
    background(220, 75);
    creatures.forEach(function callback(c, i, a) { 
        drawCreature(c.X, c.Y, i, a)      
    });
}

// server sync every 250ms.
setInterval(cycle, 250)