let posts = document.getElementById("postfeed")




for (let i = 0; i<10; i++){
    let postDivs = document.createElement("div")
    let postTitle = document.createElement("div")
    postTitle.id = i
    postTitle.className = "post-title-class"
    let postContent = document.createElement("div")
    postContent.id = i
postContent.className = "post-content-class"

    let postFooter = document.createElement("div")
    postFooter.id = i
    postFooter.className = "post-footer-class"
    postDivs.className = "post-class "
postDivs.id = i
postTitle.innerText = `This is post number ${i}\n`
postContent.innerText = " This is a post bla blablalala\n___________________________________________________"
postFooter.innerText = `Created by abmutungi,   Date: ${new Date().toDateString()}, Comments: ${i+13}`
postDivs.appendChild(postTitle)
postDivs.appendChild(postContent)
postDivs.appendChild(postFooter)

posts.appendChild(postDivs)
}