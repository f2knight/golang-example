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
		console.log(data);
	}

	function getList($el){
		var $spinEl = $el.children(".action");
		$spinEl.html("<i class='fa fa-refresh fa-spin'></i>");
		$.ajax({
	        url: mainUrl + "users",
	        type: "GET",
	        success: function(data){
	            unpdateList(data.content);
	            $spinEl.html("<i class='fa fa-refresh'></i>");
	        },
	        error:function(err){
	        	$spinEl.html("<i class='fa fa-refresh'></i>");
	        }
	    });
	}
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