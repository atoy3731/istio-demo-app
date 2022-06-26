var cross = new LeaderLine(
    document.getElementById('app1'),
    document.getElementById('app2'),
    { dash: {animation: true}, color: 'gold', endPlug: 'arrow1', middleLabel: LeaderLine.captionLabel({outlineColor: 'gold', color: 'white', text: 'Waiting', fontSize: '1.8em', fontFamily: '"Gill Sans", sans-serif'})}
);

var intra = new LeaderLine(
    LeaderLine.pointAnchor(document.getElementById('app1'), {x: 0, y: 50}),
    LeaderLine.pointAnchor(document.getElementById('app1'), {x: 0, y: 190}),
    { startSocketGravity: [-100, 0], endSocketGravity: [-100, 0 ], dash: {animation: true}, color: 'gold', endPlug: 'arrow1', middleLabel: LeaderLine.captionLabel('Waiting', { outlineColor: 'gold', color: 'white', fontFamily: "Gill Sans, sans-serif", offset: [20, 20], fontSize: '1.8em'})}
);

var internet = new LeaderLine(
    document.getElementById('app1'),
    document.getElementById('app3'),
    { startSocket: 'top', dash: {animation: true}, color: 'gold', endPlug: 'arrow1', middleLabel: LeaderLine.captionLabel({outlineColor: 'gold', color: 'white', text: 'Waiting', fontSize: '1.8em', fontFamily: '"Gill Sans", sans-serif'})}
);

function updateLineColor(obj, color, label) {
    obj.setOptions({
        color: color,
        middleLabel: LeaderLine.captionLabel({outlineColor: color, color: 'white', text: label, fontSize: '1.8em', fontFamily: '"Gill Sans", sans-serif'})
      });
}

const interval = setInterval(function() {
    // Check intra
    fetch('api/intra')
        .then(response => response.json())
        .then((data) => {
            if (data.Status == "success") {
                updateLineColor(intra, 'green', 'Success')
            } else {
                updateLineColor(intra, 'red', 'Failure')
            }
        });

    // Check internet
    fetch('api/internet')
        .then(response => response.json())
        .then((data) => {
            if (data.Status == "success") {
                updateLineColor(internet, 'green', 'Success')
            } else {
                updateLineColor(internet, 'red', 'Failure')
            }
        });
    

    // Check cross
    fetch('api/cross')
        .then(response => response.json())
        .then((data) => {
            if (data.Status == "success") {
                updateLineColor(cross, 'green', 'Success')
            } else {
                updateLineColor(cross, 'red', 'Failure')
            }
        });
}, 5000);
