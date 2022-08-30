const ws = new WebSocket("ws://localhost:8080/ws");

let btn = document.getElementById("btn");
let txtBox = document.getElementById("input");
let output = document.getElementById("msgOut");
let postButton = document.getElementById("create-post-button");
let submitPostButton = document.getElementById("submit-post-button");
let submitCommentButton = document.getElementById("submit-comment-button");
let postModal = document.getElementById("post-modal");
let postContent = document.getElementById("post-content-text");
let postTitle = document.getElementById("post-title-text");
let commentContent = document.getElementById("comment-content-text");
let postDisplay = document.getElementById("post-feed");
let content = document.getElementById("content");
let msgArr = [];
let objData = {};
let commentData = {};
let sendingArr = [];
const loginSubmitButton = document.getElementById("login-submit-button");
const signUpModal = document.getElementById("signupModal");
const signUpButton = document.getElementById("signup-submit-button");
const signUpForm = document.getElementById("signup-form")

// client side websocket
ws.onopen = () => {
  console.log("Connection to server established...");
};

submitPostButton.addEventListener("click", function (e) {
  e.preventDefault();
  console.log("submit---", postButton);
  postButton.style.display = "block";
  postModal.style.display = "none";

  objData["title"] = postTitle.value;
  objData["postcontent"] = postContent.value;
  objData["type"] = "post";
  objData["posttime"] = new Date().toISOString().slice(0, 10);

  postTitle.value = "";
  postContent.value = "";

  // message sent to server
  ws.send(JSON.stringify(objData));
});

// message received from server side
ws.onmessage = (e) => {
  let data = JSON.parse(e.data);
  console.log(data.tipo);
  console.log(data);

  if (data.tipo === "post") {
    msgArr.push(data.title + "\n" + data.postcontent + "\n" + data.user + " " + data.posttime);
    content.textContent += "\n" + msgArr[msgArr.length - 1];
  }

  if (data.tipo === "comment") {
    alert(`COmment received${data.commentcontent}`);
  }
  console.log("Received this message from server....", data);
};

submitCommentButton.addEventListener("click", function (e) {
  e.preventDefault();
  postButton.style.display = "block";
  postModal.style.display = "none";

  commentData["commentcontent"] = commentContent.value;
  commentData["type"] = "comment";
  commentData["commenttime"] = new Date().toISOString().slice(0, 10);

  commentContent.value = "";
  ws.send(JSON.stringify(commentData));
});

postButton.addEventListener("click", function (e) {
  e.preventDefault();

  postButton.style.display = "none";
  postModal.style.display = "block";
});

document.addEventListener("click", (event) => {
  // if  (!event.target.closest("post-modal") && event.target != postButton )
  if (event.target != postModal && event.target != postButton) {
    postModal.style.display = "none";
    postButton.style.display = "block";
  }
});

signUpButton.addEventListener("click", (e) => {
  var data = new FormData(signUpForm);

  const formProps = Object.fromEntries(data);

  console.log(formProps);
});
