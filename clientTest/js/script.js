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
			var strHTML = "<tr><td>" + data[index].id + "</td><td>" + data[index].name + "</td><td>" + data[index].email + "</td>";
	      	strHTML += "<td><button class='btn btn-info'><i class='fa fa-pencil-square-o'></i> Edit</button>";
	        strHTML += "<button class='btn btn-danger'><span class='action'><i class='fa fa-trash-o'></i></span> Delete</button></td></tr>";
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
	            $spinEl.html("<i class='fa fa-refresh'></i>");
	        },
	        error:function(err){
	        	$spinEl.html("<i class='fa fa-refresh'></i>");
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
});