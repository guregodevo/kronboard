'use strict';

var offlineControllers = angular.module('offlineControllers', ['gridster']);

offlineControllers.controller('WelcomeCtrl', function($scope) {
	$scope.awesomeThings = ['HTML5 Boilerplate', 'AngularJS', 'Karma'];
});

offlineControllers.controller('SignUpCtrl', ['$scope', '$http', '$window', 'AuthApi', '$location',
function($scope, $http, $window, AuthApi, $location) {
	$scope.msg = '';
	$scope.submit = function() {
		AuthApi.signup($scope.email, $scope.password, $scope.companyName).then(function(data) {//success
			console.log('Authentication successfull: store token: ' + data.token);
			$window.sessionStorage.token = data.token;
			console.log('go to dashboards page');
			$location.path('/dashboards');
		}, function(status) {//failed
			delete $window.sessionStorage.token;
			// Handle login errors here
			$scope.msg = 'Invalid email, company name or password';
			console.log('Authentication failed');
		});
	};
}]);

offlineControllers.controller('LoginCtrl', ['$scope', '$http', '$window', 'AuthApi', '$location',
function($scope, $http, $window, AuthApi, $location) {
	$scope.msg = '';
	$scope.submit = function() {
		AuthApi.auth($scope.email, $scope.password).then(function(data) {//success
			console.log('Authentication successfull: store token: ' + data.token);
			$window.sessionStorage.token = data.token;
			console.log('go to dashboards page');
			$location.path('/dashboards');
		}, function(status) {//failed
			delete $window.sessionStorage.token;
			// Handle login errors here
			$scope.msg = 'Invalid user or password';
			console.log('Authentication failed');
		});

	};
}]);

offlineControllers.controller('SignOutCtrl', ['$scope', '$http', '$window', 'AuthApi', '$location',
function($scope, $http, $window, AuthApi, $location) {
	$scope.msg = '';
	console.log('Logout successfull: store token: ' + $window.sessionStorage.token);
	delete $window.sessionStorage.token;
	$location.path('/login');
}]);

offlineControllers.controller('NewChartCtrl', ['$scope', '$routeParams', '$http', '$window', 'ChartApi', '$location',
function($scope, $routeParams, $http, $window, ChartApi, $location) {
	console.log("New chart "+ $routeParams.dashboardid);
	$scope.dashboardid = $routeParams.dashboardid;
	$scope.readonly = true;

	$scope.submit = function() {
		
	ChartApi.addchart($scope.dashboardid, $scope.type, $scope.title).then(function(data) {//success
			console.log('Chart added' + data.id);
			console.log('go to dashboards page');
			$location.path('/dashboards');
		}, function(status) {//failed
			console.log('Chart could not be added');
		});
	};
}]);

offlineControllers.controller('EditChartCtrl', ['$scope', '$routeParams', '$http', '$window', 'ChartApi', '$location',
function($scope, $routeParams, $http, $window, ChartApi, $location) {
	console.log("editing " + $routeParams.id);
	$scope.id = $routeParams.id;
	$scope.type = $routeParams.type;
	$scope.readonly = true;
	// these map directly to gridsterItem options
	ChartApi.getchart($scope.id).then(function(data) {//success
			$scope.chart = data;
			console.log('Got chart' + data);
		}, function(status) {
			$scope.msg = 'Invalid chart';
			console.log('Get chart service failed');
	});
}]);

offlineControllers.controller('DashboardCtrl', ['$scope', '$http', '$timeout', 'ChartApi',
function($scope, $http, $timeout, ChartApi) {
	//$scope.readonly = false;

	$scope.gridsterOpts = {
		margins : [5, 5],
		draggable : {
			enabled : true
		},
		resizable : {
			enabled : true
		}
	};

	// these map directly to gridsterItem options
	ChartApi.getcharts().then(function(data) {//success
			$scope.charts = data;
			console.log('Got chart:' + data);
		}, function(status) {
			//failed
			$scope.msg = 'Invalid charts';
			console.log('Get charts service failed');
	});


	$scope.hideDialog = function () {
      $timeout(function () {
       console.log('hide dialog');
      }, 200);
    };
	
	//TODO Add remove item
	/*	$scope.$watch('charts', function(newItems, oldItems) {
	 // one of the items changed
	 console.log('change detected');
	 for (var i = 0, j = newItems.length; i < j; i++) {
	 if (!angular.equals(newItems[i], oldItems[i])) {
	 if (newItems[i].sizeX != oldItems[i].sizeX || newItems[i].sizeY != oldItems[i].sizeY) {
	 console.log('changed size ' + i + newItems[i].type);
	 }
	 if (newItems[i].row != oldItems[i].row || newItems[i].col != oldItems[i].col) {
	 console.log('changed position ' + i + newItems[i].type);
	 }
	 }
	 };
	 }, true);*/

}]);
