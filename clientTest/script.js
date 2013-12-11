$(function(){
	var mainUrl = "http://localhost:3001/"

	//Test Get method
	$("#list-user").click(function(){
		$.ajax({
	        url: mainUrl + "users",
	        type: "GET",
	        success: function(data){
	            console.log(data);
	        },
	        error:function(err){
	        	console.log(err);
	        }
	    });
	});
	
	//test Post method
	$("#create-user").click(function(){
		var sName = "Thuy";
		var sEmail = "abc@gmail.com";
		$.ajax({
	        url: mainUrl + "users",
	        type: "post",
	        data: {"name":sName,
	    			"email":sEmail},
	        success: function(data){
	            console.log(data);
	        },
	        error:function(err){
	        	console.log(err);
	        }
	    });
	});
		
});