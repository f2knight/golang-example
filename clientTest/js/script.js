/*
	Script test RESTful API
	Author : Nguyen Duc Nghia (Tommy)
	Email : tommy@likipe.se

*/
$(function(){
	var mainUrl = "http://localhost:3001/"
	$("#name").val("");
	$("#email").val("");

	function unpdateList(data){
		$("#list-users").html("");
		$.each(data, function(index) {
			var strHTML = "<tr><td class='user-id'>" + data[index].id + "</td><td class='user-name'>" + data[index].name + "</td><td class='user-email'>" + data[index].email + "</td>";
	      	strHTML += "<td><button class='btn btn-info btn-edit-user'><span class='update'><i class='fa fa-pencil-square-o'></i></span> <span class='state'>Edit</span></button> | ";
	        strHTML += "<button class='btn btn-danger btn-delete-user'><span class='action'><i class='fa fa-trash-o'></i></span>Delete</button></td></tr>";
            $("#list-users").append(strHTML);
        });
	}

	function getList($el){
		var $spinEl = $el.children(".action");
		$spinEl.html("<i class='fa fa-refresh fa-spin'></i>");
		$.ajax({
	        url: mainUrl + "users",
	        type: "GET",
	        success: function(data){
	            unpdateList(data);
	            $spinEl.html("<i class='fa'></i>");
	        },
	        error:function(err){
	        	$spinEl.html("<i class='fa'></i>");
	        }
	    });
	}
	getList($("#list-user"));
	//Test Get method
	$("#list-user").click(function(){
		getList($(this));
	});
	
	//Test Post method
	$("#create-user").click(function(){
		var sName = $("#name").val();
		var sEmail = $("#email").val();
		if(sName.trim() != "" && sEmail.trim() != ""){
			$(".action").html("<i class='fa fa-refresh fa-spin'></i>");
			$.ajax({
		        url: mainUrl + "users",
		        type: "POST",
		        data: {	"name":sName,
		    			"email":sEmail},
		        success: function(data){
		            $("#name").val("");
		            $("#email").val("");
		            $(".action").html("<i class='fa fa-plus'>");
		            getList($("#list-user"));
		        },
		        error:function(err){
		        	console.log(err);
		        	$(".action").html("<i class='fa fa-plus'>");
		        }
		    });
		}
	});

	/*Test for Method Delete */
	$(".btn-delete-user").live('click', function(){
		if(confirm("Do you want to delete this user ?")){
			var user_id = $(this).closest("tr").find(".user-id").html(),
		 	$tr = $(this).closest("tr");
			$.ajax({
			        url: mainUrl + "users/" + user_id,
			        type: "POST",
			        data : {"action" : "delete"},
			        success: function(data){
			            if(data.success){
			            	$tr.remove();
			            }
			        },
			        error:function(err){
			        	console.log(err);
			        }
			});
		}
	});


	//Test put method
	$(".btn-edit-user").live('click', function(){
		$tr = $(this).closest("tr");		
		var $state = $(this).find(".state");
		if(String($state.html()) == "Edit"){
			editHtml($tr, true);
			$state.html("Save");
		}
		else{
			if(updateUser($tr))
				$state.html("Edit");
		}
	});

	function updateUser($tr){
		var iUserId = $tr.find(".user-id").html(),
			userName = $tr.find(".user-name").children("input").val(),
			email = $tr.find(".user-email").children("input").val();
			$tr.find(".update").html("<i class='fa fa-refresh fa-spin'></i>");

			if(userName.trim() != "" && email.trim() != ""){
				$.ajax({
			        url: mainUrl + "users/" + iUserId,
			        type: "POST",
			        data: {	"name":userName,
			    			"email":email,
			    			"id": iUserId,
			    			"action": "update"
			    	},
			        success: function(data){
			            $tr.find(".update").html("<i class='fa fa-pencil-square-o'>");
			            if (data.success) {
			            	editHtml($tr, false);
			            }
			            else{
			            	alert("Can not update this user!");
			            }
			            
			        },
			        error:function(err){
			        	console.log(err);
			        	$(".update").html("<i class='fa fa-pencil-square-o'>");
			        }
			    });
			    return true;
			}
			else{
				$tr.find(".update").html("<i class='fa fa-pencil-square-o'>");
				return false;
			}

	}

	function editHtml($tr, edit){
		var iUserId = $tr.find(".user-id"),
			$userName = $tr.find(".user-name"),
			$email = $tr.find(".user-email");
		if (edit == true) {
			$userName.html("<input type='text' name='username' value='"+ $userName.html() + "' />");
			$email.html("<input type='text' name='email' value='"+ $email.html() + "' />");
		}
		else {
			$userName.html($userName.children("input").val());
			$email.html($email.children("input").val());
		}
		
	}
});