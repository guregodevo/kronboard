var module = angular.module('miraServices', []);

//var preparyQueryParam = function(queryJson) {
//  return angular.isObject(queryJson)&&!angular.equals(queryJson,{}) ? {q:JSON.stringify(queryJson)} : {};
//};

/*Entity API:Get all by Type, Get One by Id,Create One*/
module.factory('AuthApi', ['$q', '$http',
function($q, $http) {

	//Service authenticate a user by username and password
	function auth(username, password) {
		console.log("HTTP authentication : " + username);
		var deferred = $q.defer();
		//?username='+user.username+'&password='+user.password
		$http.get(server + '/authenticate?username=' + username + '&password=' + password, {}).success(function(data, status) {
			deferred.resolve(data);
		}).error(function(data, status) {
			deferred.reject(status);
		});
		return deferred.promise;
	}//factory.loadEntityById

	//Service create a new account
	function signup(username, password, company) {
		console.log("HTTP SignUp : " + username);
		var deferred = $q.defer();
		var body = {
			'username' : username,
			'password' : password,
			'company' : company
		};
		$http.post(server + '/authenticate', body).success(function(data, status) {
			deferred.resolve(data);
		}).error(function(data, status) {
			deferred.reject(status);
			console.log("Signup failed HTTP status : " + status);
		});
		return deferred.promise;
	}

	return {
		auth : auth,
		signup : signup
	};
	
}]);

module.factory('ChartApi', ['$q', '$http',
function($q, $http) {

	//Add a chart to  a collection of charts
	function addchart(dashboardid, type, title) {
		var deferred = $q.defer();
		var chart = {
			"Description" : title,
			"Type" : "circle"
		};	
		$http.post(server + '/dashboards/'+ dashboardid +'/chart/', chart).success(function(data, status) {
			deferred.resolve(data);
		}).error(function(data, status) {
			deferred.reject(status);
		});
		return deferred.promise;
	}

	//Put a chart 
	function putchart(dashboardid, id, sizeX, sizeY, row, col) {
		var deferred = $q.defer();
		var chart = {
			"id"	: id,
			"sizeX" : sizeX,
			"sizeY" : sizeY,
			"row"   : row,
			"col"   : col
		};		
		$http.put(server + '/dashboards/'+ dashboardid +'/chart/'+id, chart).success(function(data, status) {
			deferred.resolve(data);
		}).error(function(data, status) {
			deferred.reject(status);
		});
		return deferred.promise;
	}

	//Delete a chart from  a collection of charts
	function deletechart(dashboardid, id) {
		var deferred = $q.defer();
		$http.delete(server + '/dashboards/'+ dashboardid +'/chart/'+id).success(function(data, status) {
			deferred.resolve(data);
		}).error(function(data, status) {
			deferred.reject(status);
		});
		return deferred.promise;
	}

	//Get a collection of charts
	function getchart(id) {
		var deferred = $q.defer();
		$http.get(server + '/charts/'+id, {}).success(function(data, status) {
			deferred.resolve(data);
		}).error(function(data, status) {
			deferred.reject(status);
		});
		return deferred.promise;
	}

	//Get a collection of charts
	function getcharts() {
		var deferred = $q.defer();
		$http.get(server + '/dashboards/1', {}).success(function(data, status) {
			deferred.resolve(data);
		}).error(function(data, status) {
			deferred.reject(status);
		});
		return deferred.promise;
	}

	return {
		getcharts : getcharts,
		getchart : getchart,
		addchart : addchart,
		deletechart	: deletechart,
		putchart : putchart,
	};	
}]);

