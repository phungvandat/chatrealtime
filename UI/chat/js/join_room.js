		//join room
		$('#btnJoinRoom').click(function(){
			$nameRoomJoin = $("#txtNameRoomJoin").val()
			if($nameRoomJoin == ""){
				alert("room name is required")
				return
			}
			//check room exist
			if (listRoom.get("room"+$nameRoomJoin) != null && listRoom.get("room"+$nameRoomJoin)=="room"+$nameRoomJoin){
				alert("Room name exist")
				return
			}

			$roomCreate = {"room" : {"user_name" : $user, "room_name" : $nameRoomJoin}}
			$.ajax({
				url: "http://localhost:3000/rooms/join",
				type: "POST",
				contentType: "application/json",
				dataType: "json",
				data: JSON.stringify($roomCreate),
				async: false,
				success: function (response) {
					displayRoom($nameRoomJoin);
				},
				error: function (xhr, ajaxOptions, thrownError) {
					err = JSON.parse(xhr.responseText)
					alert(err.error);
				}
			});
		});

		function displayRoom(nameRoom){

			var liNode = document.createElement("li");
			liNode.setAttribute("class", "active");
			liNode.setAttribute("id", "room"+nameRoom)
			liNode.setAttribute("onclick","ChosseRoom(this.id)")
			var divChatNode = document.createElement("div");
			divChatNode.setAttribute("class","chatList");
			var divImgNode = document.createElement("div");
			divImgNode.setAttribute("class","img");
			var imgAvatarNode = document.createElement("img");
			imgAvatarNode.setAttribute("src", "/x/xxx.jpg")
			var divDescNode = document.createElement("div");
			divDescNode.setAttribute("class","desc");
			var bNode = document.createElement("b");
			var messagebNode= document.createTextNode(nameRoom);
			bNode.appendChild(messagebNode);
			divDescNode.appendChild(bNode);
			divImgNode.appendChild(imgAvatarNode);
			divChatNode.appendChild(divImgNode);
			divChatNode.appendChild(divDescNode);
			liNode.appendChild(divChatNode);
			document.getElementById("ulRoom").appendChild(liNode);
	
			document.getElementById(idRoomActive).removeAttribute("class");
			idRoomActive = "room"+nameRoom;
			listRoom.set("room"+nameRoom,"room"+nameRoom)
		}
		function ChosseRoom(roomID){
			document.getElementById(roomID).setAttribute("class","active");
			document.getElementById(idRoomActive).removeAttribute("class");
			idRoomActive= roomID;
		}
	//join room