let posts = document.getElementById("post-feed");
let onlineUsers = document.getElementById("onlineusers");

let postButton = document.getElementById("new-post-btn");

let users = [];
let onlineUsersFromGo = [];

let loggedInUser = "";
let homepageUsername = document.getElementById("active-username");
let crests = document.getElementsByClassName("crest-colors");

let userDetails;
let imageDiv;
let img;

let ws;

let modal = document.getElementsByClassName("modal");
let chatModal = document.getElementById("my-chat-modal");
let createPostModal = document.getElementById("create-post-modal");
let displayPostModal = document.getElementById("display-post-modal");

let submitPostButton = document.querySelector("#submit-post-button");
let postTitle = document.querySelector("#post-title");
let postContent = document.querySelector("#post-content");
let registerBtn = document.getElementById("register-btn");
let signUpForm = document.getElementById("signup-form");

let commentContainer = document.getElementById("comment-container");
let commentArrow = document.getElementById("comment-arrow");
let commentTextArea = document.getElementById("comment-input");

// Get the button that opens the modal
let btn = document.getElementById("myBtn");

postButton.addEventListener("click", function () {
    createPostModal.style.display = "block";
});

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    for (let i = 0; i < modal.length; i++) {
        // console.log("modal -> ", modal[i]);
        // console.log("evt -> ", event.target);
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
let chatRecipient = document.getElementById("chat-recipient");
let sender = true;

sendArrow.addEventListener("click", function () {
    let chatData = {};

    if (chatTextArea.value != "") {
        chatData.chatsender = loggedInUser;
        chatData.chatrecipient = chatRecipient.innerText;
        chatData.message = chatTextArea.value;
        chatData.type = "chatMessage";

        ws.send(JSON.stringify(chatData));
    }

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

    console.log("Checking values ----", chatData);

    chatContainer.appendChild(newChatBubble);

    chatBody.scrollTo(0, chatBody.scrollHeight);
});

const teamCrests = [
    "/css/img/newcastle.png",
    "/css/img/chelsea.png",
    "/css/img/man-u.png",
    "/css/img/man-city.png",
    "/css/img/liverpool.png",
    "/css/img/spurs.png",
];

const categorySelection = document.getElementById("category-selection");

for (let i = 0; i < teamCrests.length; i++) {
    let img = document.createElement("img");
    img.style.backgroundColor = "white";
    img.alt = "none";
    img.id = teamCrests[i].slice(
        teamCrests[i].lastIndexOf("/") + 1,
        teamCrests[i].length - 4
    );
    img.classList = "crest-colors";
    img.src = teamCrests[i];
    categorySelection.append(img);
}

let crestcolors = document.getElementsByClassName("crest-colors");

const colorSwitch = {
    newcastle: `linear-gradient(
      to right,
      #040108,
      #040108 50%,
      #f0f0f0 50%,
      #f0f0f0
    )`,
    spurs: "lightgrey",
    "man-u": "red",
    chelsea: "blue",
    liverpool: "red",
    "man-city": "skyblue",
};

for (let i = 0; i < crestcolors.length; i++) {
    crestcolors[i].addEventListener("mouseup", (e) => {
        if (e.target.alt == "none") {
            e.target.style.background = colorSwitch[e.target.id];
            e.target.alt = colorSwitch[e.target.id];
        } else {
            e.target.style.background = "white";
            e.target.alt = "none";
        }
    });
}

let commentData = {};
let clickedPostID;

let signupSwitch = document.getElementById("sign-up-button");
let loginBox = document.querySelector(".login-box");
let registerBox = document.querySelector(".register-box");
let loginReturn = document.querySelector("#login-return");
let loginButton = document.getElementById("login-button");
let forumBody = document.getElementById("forumbody");
let loginModal = document.querySelector(".login-modal");
let loginForm = document.getElementById("login-form");
let loginError = document.getElementById("login-error");

signupSwitch.addEventListener("click", (e) => {
    loginBox.style.display = "none";
    registerBox.style.display = "block";
});

loginReturn.addEventListener("click", (e) => {
    loginBox.style.display = "block";
    registerBox.style.display = "none";
});

const loginValidation = (data) => {
    if (data.successfulLogin) {
        loggedInUser = data.successfulusername;
        homepageUsername.innerText = loggedInUser;
        loginModal.style.display = "none";
        forumBody.style.display = "block";
        ws = new WebSocket("ws://localhost:8080/ws");
        ws.onopen = () => {
            for (let i = 0; i < data.dbposts.length; i++) {
                DisplayPosts(data.dbposts[i]);
            }

            //populateUsers(data.allUsers);

            console.log("connection established");
        };

        ws.onmessage = (e) => {
            let data = JSON.parse(e.data);
            // console.log("data when a post is clicked", data);
            if (data.tipo === "post") {
                DisplayPosts(data);
            }

            if (data.tipo === "onlineUsers") {
                onlineUsersFromGo = data.onlineUsers;

                let usersToDisplay = [];

                for (const name of data.allUsers) {
                    if (name != loggedInUser) {
                        usersToDisplay.push(name);
                    }
                }
                populateUsers(usersToDisplay);

                console.log("first OUFG", onlineUsersFromGo);
            }

            // if (data.tipo === "commentsfrompost") {
            //   console.log(data);
            // }
        };
    } else {
        loginError.style.display = "block";
    }
    // ws.onclose = () => {
    //   // window.location.reload();
    //   console.log("ONCLOSE")
    //   objData = {};
    //   objData["type"] = "Logout";
    //   objData["logoutUsername"] = loggedInUser;
    //   console.log(objData["logoutUsername"]);
    //   ws.send(JSON.stringify(objData));
    // };
};

loginButton.addEventListener("click", (e) => {
    let loginData = new FormData(loginForm);
    let loginFormToGo = Object.fromEntries(loginData);
    loginFormToGo["type"] = "login";
    console.log("---", loginFormToGo);

    // fetch, send login data to backend server

    fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(loginFormToGo),
    })
        .then((resp) => resp.json())
        .then(function (data) {
            if (data.tipo === "loginValidation") loginValidation(data);
        });
});

// dummy post info being sent to server
let objData = {};
submitPostButton.addEventListener("click", function (e) {
    e.preventDefault();
    console.log("getSelTeams test -> ", getSelectedTeams());
    objData["title"] = postTitle.value;
    objData["postcontent"] = postContent.value;
    objData["type"] = "post";
    // objData["posttime"] = new Date().toISOString().slice(0, 10);
    objData["username"] = loggedInUser;
    objData["categories"] = getSelectedTeams();
    createPostModal.style.display = "none";
    postTitle.value = "";
    postContent.value = "";

    // message sent to server
    ws.send(JSON.stringify(objData));

    for (let i = 0; i < crests.length; i++) {
        crests[i].alt = "none";
        crests[i].style.background = "white";
    }
});

let successfulRegistrationMessage = document.getElementById(
    "registered-login-success"
);

let registrationErrors = document.querySelectorAll(".registration-errors");

let usernameError = document.getElementById("username-error");
let ageError = document.getElementById("age-error");
let firstnameError = document.getElementById("firstname-error");
let lastnameError = document.getElementById("lastname-error");
let emailError = document.getElementById("email-error");
let passwordError = document.getElementById("password-error");

const registrationValidation = (data) => {
    console.log("check data -> ", data);
    if (data.successfulRegistration) {
        console.log("CHECKING LOOP");
        registerBox.style.display = "none";
        loginBox.style.display = "block";
        signupSwitch.style.display = "none";
        successfulRegistrationMessage.style.display = "block";
    } else {
        // switch case for errors, validated from back end.
        // console.log(data.usernameLength);

        if (data.usernameLength || data.usernameSpace) {
            usernameError.innerText = "";
            usernameError.innerText = "min 5 characters, no spaces";
            usernameError.style.display = "block";
        }

        if (data.usernameDuplicate) {
            usernameError.innerText = "";
            usernameError.innerText = "Username exists";
            usernameError.style.display = "block";
        }

        if (data.emailDuplicate) {
            emailError.innerText = "";
            emailError.innerText = "Email exists";
            emailError.style.display = "block";
        }

        if (data.emailInvalid) {
            emailError.innerText = "";
            emailError.innerText = " Valid email required";
            emailError.style.display = "block";
        }

        if (data.passwordLength) {
            passwordError.style.display = "block";
        }
        if (data.ageEmpty) {
            console.log("age empty");

            ageError.style.display = "block";
        }

        if (data.firstnameEmpty) {
            console.log("first name empty");
            firstnameError.style.display = "block";
        }

        if (data.lastnameEmpty) {
            console.log("last name empty");

            lastnameError.style.display = "block";
        }
    }
};

registerBtn.addEventListener("click", function (e) {
    // e.preventDefault();
    registrationErrors.forEach(function (el) {
        el.style.display = "none";
        // el.innerText = "";
    });

    let signupData = new FormData(signUpForm);
    let signUpFormToGo = Object.fromEntries(signupData);
    signUpFormToGo.type = "signup";

    fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(signUpFormToGo),
    })
        .then((resp) => resp.json())
        .then(function (data) {
            if (data.tipo === "formValidation") {
                console.log("Check form val is being called");
                registrationValidation(data);
            }
        });
});

// logout
let logoutButton = document.getElementById("log-out-button");

logoutButton.onclick = () => {
    objData = {};
    objData["type"] = "logout";
    objData["logoutUsername"] = loggedInUser;
    console.log(objData["logoutUsername"]);
    ws.send(JSON.stringify(objData));
    // ws.close()
    ws.onmessage = (e) => {
        // let commentData = JSON.parse(e.data);
        let logoutData = JSON.parse(e.data);
        if (logoutData.logoutClicked) {
            ws.close();
            window.location.reload();
        }
    };
};

const DisplayPosts = (data) => {
    let postDivs = document.createElement("div");
    let postTitle = document.createElement("div");

    postTitle.className = "post-title-class";
    // let postContent = document.createElement("div");

    //postContent.className = "post-content-class";

    let postFooter = document.createElement("div");
    postFooter.className = "post-footer-class";
    postDivs.className = "post-class ";
    // this will eventually hold the id given by go from the database (data.id)
    postDivs.id = data.postid;
    postTitle.innerText = data.title;
    // postContent.innerText = data.postcontent;
    postTitle.style.borderBottom = "0.2vh solid black";
    postFooter.innerText = `Created by ${data.username},   Date: ${
        data.posttime
    }, Comments: ${1 + 13} ${data.categories}`;
    postDivs.appendChild(postTitle);
    //postDivs.appendChild(postContent);
    postDivs.appendChild(postFooter);
    // postTitle.addEventListener("click", function (e) {
    //   console.log("Checking if listener working");

    // });
    posts.appendChild(postDivs);

    let getCommentsForPosts = {};

    // When a post is clicked on
    postDivs.addEventListener("click", (e) => {
        clickedPostID = postDivs.id;
        getCommentsForPosts["clickedPostID"] = clickedPostID;
        getCommentsForPosts["type"] = "getcommentsfrompost";
        console.log("comments for posts object", getCommentsForPosts);
        ws.send(JSON.stringify(getCommentsForPosts));

        let displayPostTitle = document.querySelector(".display-post-title");
        let displayPostContent = document.querySelector(
            ".display-post-content"
        );
        let postUsername = document.querySelector(".post-username");
        let postDate = document.querySelector(".post-date");
        displayPostTitle.innerText = data.title;
        displayPostContent.innerText = data.postcontent;
        postUsername.innerText = data.username;
        postDate.innerText = data.posttime;
        displayPostModal.style.display = "block";

        // Auto scroll to last comment
        // displayPostBody.scrollTo(0, displayPostBody.scrollHeight);
        commentContainer.innerHTML = "";

        // data coming through for comments
        ws.onmessage = (e) => {
            let commentData = JSON.parse(e.data);

            if (commentData.tipo == "allComments") {
                console.log("ALL COMMENTS DISPLAYED");
                for (let i = 0; i < commentData.comments.length; i++) {
                    let commentDiv = document.createElement("div");
                    commentDiv.style.marginBottom = "1vh";
                    commentDiv.id = `comment${commentData.comments[i].commentId}`;
                    commentDiv.innerText = `${commentData.comments[i].commentcontent} \n ${commentData.comments[i].user}, ${commentData.comments[i].commenttime}`;
                    commentContainer.appendChild(commentDiv);
                }
            }

            if (commentData.tipo == "lastComment") {
                console.log("This is the last coment", commentData);
                let commentDiv = document.createElement("div");
                commentDiv.style.marginBottom = "1vh";
                commentDiv.id = `comment${commentData.commentid}`;
                commentDiv.innerText = `${commentData.commentcontent} \n ${commentData.user}, ${commentData.commenttime}`;
                commentContainer.appendChild(commentDiv);
            }
            // console.log("is this comment data?", commentData);
        };

        console.log("commentdata---", commentData);
        console.log(commentData.type);

        // commentData.forEach((comment) => {

        // });

        // comment data in the modal needs to be cleared as we are using one modal.
    });
};

const getSelectedTeams = () => {
    let crestList = "";
    for (let i = 0; i < crests.length; i++) {
        if (crests[i].alt !== "none") {
            crestList += `${crests[i].id},`;
        }
    }
    return crestList;
};

commentArrow.addEventListener("click", function () {
    commentData["commentcontent"] = commentTextArea.value;
    commentData["user"] = loggedInUser;
    commentData["postid"] = clickedPostID;
    commentData["type"] = "comment";

    ws.send(JSON.stringify(commentData));
    commentTextArea.value = "";
    //
    // ws.onmessage = (e) => {
    //     let lastComment = JSON.parse(e.data);
    //     console.log("This is the last coment", lastComment);
    //     if (lastComment.tipo === "lastcomment") {
    //         let commentDiv = document.createElement("div");
    //         commentDiv.style.marginBottom = "1vh";
    //         commentDiv.id = `comment${lastComment.commentId}`;
    //         commentDiv.innerText = `${lastComment.commentcontent} \n ${lastComment.user}, ${lastComment.commenttime}`;
    //         commentContainer.appendChild(commentDiv);
    //     }
    //     // console.log("is this comment data?", commentData);
    // };
});

// const colouredCategories = (data) => {
//   console.log("cat before split -> ", data);

//   let crestArr = data.split(",");
//   console.log("cat after split --> ", crestArr);

//   let masterCrestList = document.createElement("div");
//   for (let i = 0; i < crestArr.length; i++) {
//     //let crestList = document.createElement("p");
//     let clubName = document.createTextNode(`${crestArr[i]}`);

//     // crestList.insertAdjacentText("afterend", `#${crestArr[i]}`);
//     //clubName.style.color = `${colorSwitch[crestArr[i]]}`;
//     console.log("crest list before append --> ", clubName);
//     masterCrestList.append(clubName);
//   }
//   return masterCrestList;
// };

// const getSelectedTeams = () => {
//   let masterCrestList = document.createElement("p");
//   for (let i = 0; i < crests.length; i++) {
//     let crestList = document.createElement("p");
//     if (crests[i].alt !== "none") {
//       crestList.innerText += `#${crests[i].id},`;
//       crestList.style.color = `${crests[i].alt}`;
//       masterCrestList.append(crestList[i]);
//     }
//   }
//   return masterCrestList;
// };

const populateUsers = (users) => {
    onlineUsers.innerHTML = "";

    console.log("online from Go", onlineUsersFromGo);
    for (let i = 0; i < users.length; i++) {
        userDetails = document.createElement("div");
        let username = document.createElement("div");
        imageDiv = document.createElement("div");
        img = document.createElement("img");
        let onlineIcon = document.createElement("div");

        img.src = "/css/img/newcastle.png";
        img.style.width = "2vw";
        imageDiv.appendChild(onlineIcon);
        userDetails.id = `${users[i]}`;

        userDetails.className = "registered-user";

        if (onlineUsersFromGo.includes(users[i])) {
            onlineIcon.className = "online-icon-class";
        } else {
            onlineIcon.className = "offline-icon-class";
        }
        username.innerText = `${users[i]}`;
        imageDiv.append(img);
        userDetails.appendChild(username);
        userDetails.appendChild(imageDiv);
        onlineUsers.appendChild(userDetails);
    }
    let userRg = document.getElementsByClassName("registered-user");

    // Get the <span> element that closes the modal
    let span = document.getElementsByClassName("close");

    // When the user clicks the button, open the modal

    for (let i = 0; i < userRg.length; i++) {
        userRg[i].onclick = function () {
            chatContainer.innerHTML = "";
            chatRecipient.innerText = userRg[i].id;
            let requestChatData = {};
            requestChatData["chatsender"] = loggedInUser;
            requestChatData["chatrecipient"] = chatRecipient.innerText;
            requestChatData["type"] = "requestChatHistory";
            ws.send(JSON.stringify(requestChatData));

            ws.onmessage = (e) => {
                let data = JSON.parse(e.data);
                console.log("on message check");
                console.log(data.tipo);
                if (data.tipo == "messagehistoryfromgo") {
                    console.log("MESSAGE HISTORY FROM GO----------", data);
                    for (let i = 0; i < data.chathistory.length; i++) {
                        let newChatBubble = document.createElement("div");
                        newChatBubble.innerText = data.chathistory[i].message;
                        if (data.chathistory[i].chatsender == loggedInUser) {
                            newChatBubble.id = "chat-message-sender";
                        } else {
                            newChatBubble.id = "chat-message-recipient";
                        }
                        chatContainer.appendChild(newChatBubble);
                        chatBody.scrollTo(0, chatBody.scrollHeight);
                    }
                }
            };
            console.log("Users clicked");
            chatModal.style.display = "block";
        };
    }

    // When the user clicks on <span> (x), close the modal
    for (let i = 0; i < span.length; i++) {
        span[i].onclick = function () {
            modal[i].style.display = "none";
        };
    }
};
