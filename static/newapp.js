// import { fileURLToPath } from "url"
// import { dirname } from "path"

let posts = document.getElementById("post-feed");
let onlineUsers = document.getElementById("onlineusers");

let users = ["tb38r", "abmutungi", "eternal17", "million"];

for (let i = 0; i < 10; i++) {
    let postDivs = document.createElement("div");
    let postTitle = document.createElement("div");
    postTitle.id = i;
    postTitle.className = "post-title-class";
    let postContent = document.createElement("div");
    postContent.id = i;
    postContent.className = "post-content-class";

    let postFooter = document.createElement("div");
    postFooter.id = i;
    postFooter.className = "post-footer-class";
    postDivs.className = "post-class ";
    postDivs.id = i;
    postTitle.innerText = `This is post number ${i}\n`;
    postContent.innerText =
        " This is a post bla blablalala\n___________________________________________________";
    postFooter.innerText = `Created by abmutungi,   Date: ${new Date().toDateString()}, Comments: ${
        i + 13
    }`;
    postDivs.appendChild(postTitle);
    postDivs.appendChild(postContent);
    postDivs.appendChild(postFooter);

    posts.appendChild(postDivs);
}

let userDetails;
let imageDiv;
let img;

for (let i = 0; i < 4; i++) {
    userDetails = document.createElement("div");
    let username = document.createElement("div");
    imageDiv = document.createElement("div");
    img = document.createElement("img");
    let onlineIcon = document.createElement("div");

    onlineIcon.className = "online-icon-class";

    img.src = "/css/img/newcastle.png";
    img.style.width = "2vw";
    imageDiv.appendChild(onlineIcon);
    userDetails.id = `${users[i]}`;

    //   userDetails.setAttribute("type", "button");

    userDetails.setAttribute("data-bs-target", "#chatModal");
    userDetails.setAttribute("data-bs-toggle", "modal");

    userDetails.className = "registered-user";
    username.innerText = `${users[i]}`;
    imageDiv.append(img);
    userDetails.appendChild(username);
    userDetails.appendChild(imageDiv);
    onlineUsers.appendChild(userDetails);
}

let postTitlesClick = document.getElementsByClassName("post-title-class");
Array.from(postTitlesClick).forEach(function (postTitle) {
    postTitle.addEventListener("click", function (e) {
        alert(postTitle.innerText);
    });
});

var modal = document.querySelector(".modal");

var userRg = document.querySelectorAll(".registered-user");
let chatRecipient = document.getElementById("chat-recipient");

// Get the button that opens the modal
var btn = document.getElementById("myBtn");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

// When the user clicks the button, open the modal

for (let i = 0; i < userRg.length; i++) {
    userRg[i].onclick = function (e) {
        chatRecipient.innerText = userRg[i].id;

        console.log("Users clicked");
        modal.style.display = "block";
    };
}

// userRg.onclick = function() {
//   console.log("Users clicked");
//   modal.style.display = "block";
// }

// When the user clicks on <span> (x), close the modal
span.onclick = function () {
    modal.style.display = "none";
};

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    if (event.target == modal) {
        modal.style.display = "none";
    }
};

let sendArrow = document.getElementById("chat-arrow");
let chatTextArea = document.getElementById("chat-input");
let chatContainer = document.getElementById("chat-container");
let sender = true;

sendArrow.addEventListener("click", function () {
    console.log("arrow clicked");
    let newChatBubble = document.createElement("div");

    newChatBubble.innerText = chatTextArea.value;
    chatTextArea.value = "";
    if (sender) {
        newChatBubble.id = "chat-message-sender";
        sender = false;
    } else {
        newChatBubble.id = "chat-message-recipient";
        sender = true;
    }

    chatContainer.appendChild(newChatBubble);
});
