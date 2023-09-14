budget = document.getElementById("budget");
reputation = document.getElementById("reputation");
state = document.getElementById("etatEcole");
inventory = document.getElementById("player-inventory");

function sellItem(div, id) {
    fetch("/sell", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            id: id
        })
    })
    .then(response => response.text())
    .then(data => {
        data = JSON.parse(data);
        if (data.success) {
            console.log(data);
            alert("Vente effectuée !");
            div.parentElement.remove();
            budget.innerHTML = "Budget : " + data.budget + " €";
            reputation.innerHTML = "Reputation : " + data.reputation + " %";
            state.innerHTML = "Etat de l'école : " + data.etatEcole + " %";
        } else {
            alert(data.info);
        }
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

function buyItem(div, id) {
    fetch("/buy", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            id: id
        })
    })
    .then(response => response.text())
    .then(data => {
        data = JSON.parse(data);
        if (data.success) {
            console.log(data);
            alert("Vente effectuée !");
            div.innerHTML = "Vendre";
            div.setAttribute("onclick", "sellItem(this, " + id + ")");
            inventory.appendChild(div.parentElement);
            budget.innerHTML = "Budget : " + data.budget + " €";
            reputation.innerHTML = "Reputation : " + data.reputation + " %";
            state.innerHTML = "Etat de l'école : " + data.etatEcole + " %";
        } else {
            alert(data.info);
        }
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}
