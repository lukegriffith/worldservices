var creatures = []
var paused = true
var fps = 10
var cycleNum

var selectedWorld = null;

function createRadioElement( name, value, checked ) {
    var radioInput;

    radioInput = document.createElement('input');
    radioInput.setAttribute('type', 'radio');
    radioInput.setAttribute('name', name);
    radioInput.setAttribute('value', value)
    if ( checked ) {
        radioInput.setAttribute('checked', 'checked');
    }
    radioInput.onclick = x => selectedWorld = document.querySelector('input[name="worlds"]:checked').value
    return radioInput;
}

function createLabelElement( name ) {
    var label;

    label = document.createElement('label')
    label.setAttribute('for', name)
    label.innerText = name
    return label
}

// UI Elements
function getWorlds() {
    fetch("world")    
        .then(response => response.json())
        .then(data => {
            worldList = document.getElementById("worldList")
            worldList.innerHTML = ""
            data.forEach(w => {
                worldList.appendChild(createRadioElement('worlds', w, false))
                worldList.appendChild(createLabelElement(w))
            })
        })
}


function newWorld() {
    size = document.getElementById('size')
    population = document.getElementById('pop')
    worldName = document.getElementById('name')
    fetch("world?size="+ size.value +"&pop="+ population.value +"&world="+ worldName.value, {
        method: 'POST'
    })
}




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
        realCycle()
    }       
    setTimeout(cycle, 31)
}

function realCycle() {
    cycleNum++
    updateCycleStatus()
    refreshBoard()
}

function next() {
    realCycle()
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
    return fetch('board?world='+selectedWorld+'&cycle='+cycleNum)
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
    .then(r => setTimeout(refreshBoard, 250))
    .then(r => setTimeout(renderGrid, 500))

}



// p5 JS 



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
    
}


function setup() {
    size = document.getElementById('size').value
    cnv = createCanvas(size, size);
    cnv.mouseClicked(canvasClickHandler)
    frameRate(fps);
}


function draw() {           
    if (!paused){                                
        renderGrid()
    }
}

function renderGrid() {
    background(220, 75);
    creatures.forEach(function callback(c, i, a) { 
        drawCreature(c.X, c.Y, i, a)      
    });
}

cycle()