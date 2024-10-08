console.log("Hola");

var imgArray = [
    "static/img/hand_rock.png",
    "static/img/hand_paper.png",
    "static/img/hand_scissors.png"
]

function choose(x){
    fetch("/play?c="+x).then(response => response.json())
    .then(data => {
        console.log(data);
        document.getElementById("computer_choice").innerHTML = data.computer_choice;
        document.getElementById("round_result").innerHTML = data.round_result;            
        document.getElementById("computer_score").innerHTML = data.computer_score;
        document.getElementById("player_score").innerHTML = data.player_score;
        document.getElementById("round_message").innerHTML = data.message;

        if(x ==0){
            document.getElementById("player_choice").innerHTML = "El jugador eligió Piedra";
        }else if (x ==1){
            document.getElementById("player_choice").innerHTML = "La jugador eligió Papel";
        }else if (x ==2){
            document.getElementById("player_choice").innerHTML = "El jugador eligió Tijera";
        }

        var imgElement = document.getElementById("img_computer");
        imgElement.src = imgArray[data.computer_choice_int]
    })
}
