const ws = new WebSocket("ws://localhost:8080/ws")

let btn = document.getElementById("btn");
let txtBox = document.getElementById("input");
let output = document.getElementById("msgOut");
let postButton = document.getElementById("create-post-button");
let submitPostButton = document.getElementById("submit-post-button");
let postModal = document.getElementById("post-modal");
let postContent = document.getElementById("post-content-text");
let postTitle = document.getElementById("post-title-text");
let postDisplay = document.getElementById("post-feed");
let content = document.getElementById("content")