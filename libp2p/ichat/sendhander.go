package main

import (
	"fmt"
	"log"
	"net/http"
)

const AddForm = `
<form method="POST" action="/send">
message: <input type="text" name="message"><br/><br/><br/>
<input type="submit" value="Send Message">
</form>
`


func Send(w http.ResponseWriter, r *http.Request)  {

	message := r.FormValue("message")
	//target := r.FormValue("target")

	if message == ""{
		w.Header().Set("Content-Type", "text/html")

		divSelfHtml := fmt.Sprintf("<div><span>mySelf : %s</span><div><br/><br/><br/>", cr.self)

		var divPeerListHtml string
		// 对等节点列表
		for e := range cr.ListPeers() {
			var p string
			p = "<span> peerID: " + cr.ListPeers()[e].String()  + "</span><br/>"

			divPeerListHtml = divPeerListHtml + p
		}

		fmt.Fprint(w, divSelfHtml + divPeerListHtml + AddForm)
		return
	}

	//log.Printf("Send target:  %s", target)
	log.Printf("Send message:  %s", message)

	// 将消息发布出去
	cr.Publish(message)

}
