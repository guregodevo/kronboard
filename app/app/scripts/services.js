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

	//Get a collection of charts
	function getchart(id) {
		var deferred = $q.defer();
		$http.get(server + '/charts?id='+id, {}).success(function(data, status) {
			deferred.resolve(data);
		}).error(function(data, status) {
			deferred.reject(status);
		});
		return deferred.promise;
	}

	//Get a collection of charts
	function getcharts() {
		var deferred = $q.defer();
		$http.get(server + '/dashboards?id=1', {}).success(function(data, status) {
			deferred.resolve(data);
		}).error(function(data, status) {
			deferred.reject(status);
		});
		return deferred.promise;
	}

	return {
		getcharts : getcharts,
		getchart : getchart
	};	
}]);

