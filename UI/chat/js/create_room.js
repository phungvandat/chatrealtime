//create room
$('#btnCreateRoom').click(function(){
    $nameRoomCreate = $("#txtNameRoomCreate").val()
    if($nameRoomCreate == ""){
        alert("room name is required")
        return
    }
    $roomCreate = {"room" : {"user_name" : $user, "room_name" : $nameRoomCreate}}
    $.ajax({
        url: "http://localhost:3000/rooms/create",
        type: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify($roomCreate),
        async: false,
        success: function (response) {
            displayRoom($nameRoomCreate);
        },
        error: function (xhr, ajaxOptions, thrownError) {
            err = JSON.parse(xhr.responseText)
            alert(err.error);
        }
    });
});
//create room