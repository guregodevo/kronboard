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
			$location.path('/dashboards/1');
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
			$location.path('/dashboards/1');
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
			if ($scope.type == "") {
				$scope.type = "circle";
			}
			console.log('Chart added' + data.id);
			console.log('go to dashboards page');
			$location.path('/dashboards/'+$scope.dashboardid);
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

offlineControllers.controller('DashboardCtrl', ['$scope', '$routeParams', '$location', '$http', '$timeout', 'ChartApi',
function($scope, $routeParams, $location, $http, $timeout, ChartApi) {
	$scope.dragging = true;
	$scope.changeditems= [];
	$scope.dashboardid = $routeParams.id
	$scope.addchart = function($event){ 
		$location.path('/dashboard/'+$scope.dashboardid+'/chart/');
    }

	$scope.gridsterOpts = {
		margins : [5, 5],
		draggable : {
			enabled : true,
			start: function(event, ui, $elm) { 
				$scope.dragging = false;
			},
			stop: function(event, ui, $elm) { 
				$scope.dragging = true;
				var e = $elm.find("div")[0];
				if (e) {
					var chartid = e.getAttribute('chartid');
					var dashboardid = e.getAttribute('dashboardid');
					if (chartid && dashboardid) {
						console.log('changed item:' + $scope.changeditems.length);
						var item;
					    for (var i = 0; i < $scope.changeditems.length; i++) {
							if ($scope.changeditems[i].id == chartid) {
								item = $scope.changeditems[i];
							}
					    }
						$scope.changeditems = [];
					}
					if (item) {
						//console.log("changed item: ");
						//console.log(item);
					    ChartApi.putchart(dashboardid, chartid, item.sizeX, item.sizeY, item.row, item.col).then(function(data) {//success
							console.log('Put chart');
							}, function(status) {
								console.log('Put chart service failed');
						});
					}
				}
				 
			} 
		},
		resizable : {
			enabled : true
		}
	};

	$scope.$watch('charts', function(oldItems, newItems) {
	   if (newItems) {
	   	   var len = newItems.length;
		   for (var i = 0; i < len; i++) {
	   			if (newItems && oldItems && $scope.dragging != true ) {
					if (newItems[i].sizeX != oldItems[i].sizeX || newItems[i].sizeY != oldItems[i].sizeY) {
						$scope.changeditems.push(oldItems[i]);
					}
					if (newItems[i].col != oldItems[i].col || newItems[i].row != oldItems[i].row) {
						$scope.changeditems.push(oldItems[i]);
					}
				}
		    }
	   }
	}, true);

	// these map directly to gridsterItem options
	ChartApi.getcharts().then(function(data) {//success
			$scope.charts = data;

			console.log('Got chart:' + data);
		}, function(status) {
			//failed
			$scope.msg = 'Invalid charts';
			console.log('Get charts service failed');
	});
	

}]);
