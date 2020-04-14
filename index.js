var canvas = document.getElementById('tutorial');
var submit = document.getElementById('submit');
var initial = document.getElementById('initial');
var ctx = canvas.getContext("2d");
var Height = canvas.height;
var Width = canvas.width;
var BlockWidth = Width / 5;
var BlockHeight = Height / 5;
var result = "1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n6, 7, 11, 12, 13\n14, 15, 16, 17, 18\n19, 20, 21, 0, 0\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n6, 7, 11, 12, 13\n14, 15, 16, 17, 18\n19, 20, 0, 21, 0\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n6, 7, 11, 12, 13\n14, 15, 16, 17, 18\n19, 20, 0, 0, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n6, 7, 11, 12, 13\n14, 15, 16, 17, 18\n19, 0, 20, 0, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n6, 7, 11, 12, 13\n14, 0, 16, 17, 18\n19, 15, 20, 0, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n6, 7, 11, 12, 13\n0, 14, 16, 17, 18\n19, 15, 20, 0, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n6, 7, 11, 12, 13\n0, 14, 16, 0, 18\n19, 15, 20, 17, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n6, 7, 11, 12, 13\n0, 14, 0, 16, 18\n19, 15, 20, 17, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n0, 7, 11, 12, 13\n6, 14, 0, 16, 18\n19, 15, 20, 17, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n0, 7, 11, 12, 13\n6, 0, 14, 16, 18\n19, 15, 20, 17, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n0, 7, 11, 12, 13\n6, 15, 14, 16, 18\n19, 0, 20, 17, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n0, 7, 11, 12, 13\n6, 15, 14, 16, 18\n0, 19, 20, 17, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n0, 7, 11, 12, 13\n0, 15, 14, 16, 18\n6, 19, 20, 17, 21\n\n1, 2, 3, 4, 5\n7, 7, 8, 9, 10\n0, 7, 11, 12, 13\n15, 0, 14, 16, 18\n6, 19, 20, 17, 21\n\n1, 2, 3, 4, 5\n0, 0, 8, 9, 10\n7, 7, 11, 12, 13\n15, 7, 14, 16, 18\n6, 19, 20, 17, 21\n\n0, 2, 3, 4, 5\n1, 0, 8, 9, 10\n7, 7, 11, 12, 13\n15, 7, 14, 16, 18\n6, 19, 20, 17, 21\n\n2, 0, 3, 4, 5\n1, 0, 8, 9, 10\n7, 7, 11, 12, 13\n15, 7, 14, 16, 18\n6, 19, 20, 17, 21\n\n2, 0, 3, 4, 5\n1, 8, 0, 9, 10\n7, 7, 11, 12, 13\n15, 7, 14, 16, 18\n6, 19, 20, 17, 21\n\n2, 3, 0, 4, 5\n1, 8, 0, 9, 10\n7, 7, 11, 12, 13\n15, 7, 14, 16, 18\n6, 19, 20, 17, 21\n\n2, 3, 0, 4, 5\n1, 8, 11, 9, 10\n7, 7, 0, 12, 13\n15, 7, 14, 16, 18\n6, 19, 20, 17, 21\n\n2, 3, 11, 4, 5\n1, 8, 0, 9, 10\n7, 7, 0, 12, 13\n15, 7, 14, 16, 18\n6, 19, 20, 17, 21\n\n2, 3, 11, 4, 5\n1, 8, 0, 9, 10\n7, 7, 14, 12, 13\n15, 7, 0, 16, 18\n6, 19, 20, 17, 21\n\n2, 3, 11, 4, 5\n1, 8, 14, 9, 10\n7, 7, 0, 12, 13\n15, 7, 0, 16, 18\n6, 19, 20, 17, 21\n\n2, 3, 11, 4, 5\n1, 8, 14, 9, 10\n0, 7, 7, 12, 13\n15, 0, 7, 16, 18\n6, 19, 20, 17, 21\n\n2, 3, 11, 4, 5\n1, 8, 14, 9, 10\n0, 7, 7, 12, 13\n0, 15, 7, 16, 18\n6, 19, 20, 17, 21\n\n2, 3, 11, 4, 5\n0, 8, 14, 9, 10\n1, 7, 7, 12, 13\n0, 15, 7, 16, 18\n6, 19, 20, 17, 21\n\n2, 3, 11, 4, 5\n0, 8, 14, 9, 10\n0, 7, 7, 12, 13\n1, 15, 7, 16, 18\n6, 19, 20, 17, 21\n";
var states = result.split("\n\n");
states.reverse();
var numbers = states[1].replace(/\n/g, ", ").split(", ").map(function (s) { return parseInt(s); });
drawState(0);
submit.addEventListener('click', function (ev) {
    var initialState = initial.value;
    var oReq = new XMLHttpRequest();
    oReq.addEventListener("load", function (ev) {
        result = oReq.responseText;
        var states = result.split("\n\n");
        states.reverse();
        numbers = states[1].replace(/\n/g, ", ").split(", ").map(function (s) { return parseInt(s); });
        drawState(0);
    });
    oReq.open("GET", "ai/exp1?initial=" + initialState);
    oReq.send();
});
function reqListener() {
}
function drawState(i) {
    var numbers = states[i].replace(/\n/g, ", ").split(", ").map(function (s) { return parseInt(s); });
    if (canvas.getContext) {
        var ctx = canvas.getContext('2d');
        ctx.clearRect(0, 0, Width, Height);
        numbers.forEach(function (v, i) {
            if (v != 0) {
                drawBlock(i % 5, Math.floor(i / 5), v);
            }
        });
    }
    if (i < states.length - 1) {
        setTimeout(drawState, 500, i + 1);
    }
}
function drawBlock(x, y, n) {
    if (n == 0) {
        ctx.fillStyle;
    }
    ctx.fillStyle = "rgb(\n        " + (n * 15) % 255 + ",\n        " + (n * 25) % 255 + ",\n        " + (n * 25) % 255 + ")";
    ctx.fillRect(x * BlockWidth, y * BlockHeight, BlockWidth, BlockHeight);
    ctx.font = 'bold 48px serif';
    ctx.fillStyle = 'white';
    ctx.fillText(n.toString(), (x + 0.4) * BlockWidth, (y + 0.6) * BlockHeight, BlockWidth);
}
