let posts = document.getElementById("post-feed");
let onlineUsers = document.getElementById("onlineusers");

let postButton = document.getElementById("new-post-btn");

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
    displayPostModal.style.display = "block";
  });
});

let modal = document.getElementsByClassName("modal");
let chatModal = document.getElementById("my-chat-modal");
let createPostModal = document.getElementById("create-post-modal");
let displayPostModal = document.getElementById("display-post-modal");

postButton.addEventListener("click", function () {
  createPostModal.style.display = "block";
});

let userRg = document.querySelectorAll(".registered-user");
let chatRecipient = document.getElementById("chat-recipient");

// Get the button that opens the modal
let btn = document.getElementById("myBtn");

// Get the <span> element that closes the modal
let span = document.getElementsByClassName("close");

// When the user clicks the button, open the modal

for (let i = 0; i < userRg.length; i++) {
  userRg[i].onclick = function () {
    chatRecipient.innerText = userRg[i].id;

    chatModal.style.display = "block";
  };
}

// When the user clicks on <span> (x), close the modal
for (let i = 0; i < span.length; i++) {
  span[i].onclick = function () {
    modal[i].style.display = "none";
  };
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
  for (let i = 0; i < modal.length; i++) {
    if (event.target == modal[i]) {
      modal[i].style.display = "none";
    }
  }
};

let sendArrow = document.getElementById("chat-arrow");
let chatTextArea = document.getElementById("chat-input");
let chatContainer = document.getElementById("chat-container");
let chatBody = document.getElementById("chat-box-body");
let displayPostBody = document.getElementById("display-post-body");
let sender = true;

sendArrow.addEventListener("click", function () {
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

  chatBody.scrollTo(0, chatBody.scrollHeight);
});

let commentContainer = document.getElementById("comment-container");
let commentArrow = document.getElementById("comment-arrow");
let commentTextArea = document.getElementById("comment-input");

commentArrow.addEventListener("click", function () {
  let i = 0;
  let comment = document.createElement("div");
  let commentDetails = document.createElement("div");
  commentDetails.innerText = `Created by: McTom Date: ${
    new Date().toISOString().split("T")[0]
  } ${new Date().toISOString().split("T")[1].substring(0, 5)}`;
  comment.style.marginBottom = "1vh";
  comment.id = `comment-${i}`;
  commentDetails.id = `comment-detail-${i}`;
  comment.innerText = `${commentTextArea.value}`;
  commentTextArea.value = "";
  comment.appendChild(commentDetails);
  commentContainer.appendChild(comment);
  displayPostBody.scrollTo(0, displayPostBody.scrollHeight);
});
