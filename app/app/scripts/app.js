'use strict';

/*General Configuration Area*/
var server = 'http://localhost:3000';

var app = angular.module('miranalApp', ['ngRoute', 'gridster', 'chart', 'miraServices', 'offlineControllers']);

app.factory('ApiResponseInterceptor', ['$q', '$location','$window',
function($q, $location, $window) {
	return {
		request : function(config) {
			if ($window.sessionStorage.token && $window.sessionStorage.token !== '') {
				config.headers["X-Api-Token"] = $window.sessionStorage.token;	
			}
			return config;
		},
		responseError : function(rejection) {
//			if (rejection.status === 401) {
//				$location.path('#/login');
//			}
			return $q.reject(rejection);
		}
	};
}]);

app.config(['$routeProvider', '$httpProvider',
function($routeProvider, $httpProvider) {

	$httpProvider.defaults.useXDomain = true;
	delete $httpProvider.defaults.headers.common['X-Requested-With'];
	$httpProvider.defaults.headers.common['X-Requested-With'] = "XMLHttpRequest";
	$httpProvider.interceptors.push('ApiResponseInterceptor');
	//Http Intercpetor to check auth failures for xhr requests

	$routeProvider.when('/', {
		auth : false,
		templateUrl : 'views/welcome.html',
		controller : 'WelcomeCtrl'
	}).when('/pricing', {
		auth : false,
		templateUrl : 'views/pricing.html',
		controller : 'WelcomeCtrl'
	}).when('/login', {
		auth : false,
		templateUrl : 'views/login.html',
		controller : 'LoginCtrl'
	}).when('/signup', {
		auth : false,
		templateUrl : 'views/signup.html',
		controller : 'SignUpCtrl'
	}).when('/contact', {
		auth : false,
		templateUrl : 'views/contactus.html',
		controller : 'WelcomeCtrl'
	}).when('/signout', {
		auth : true,
		templateUrl : 'views/dashboards.html',
		controller : 'SignOutCtrl'
	}).when('/dashboard/:dashboardid/chart/', {
		auth : false,
		templateUrl : 'views/newchart.html',
		controller : 'NewChartCtrl'
	}).when('/chart/:id/:type/', {
		auth : false,
		templateUrl : 'views/chart.html',
		controller : 'EditChartCtrl'
	}).when('/dashboards/:id', {
		auth : false,
		templateUrl : 'views/dashboards.html',
		controller : 'DashboardCtrl'
	}).otherwise({
		redirectTo : '/'
	});
}]).run(function($rootScope, $location, $window) {

	// register listener to watch route changes
	$rootScope.$on('$routeChangeStart', function(event, next) {
		var authRequired = next.auth;
		if (authRequired) {
			// no logged user, we should be going to #login
			if (next.templateUrl === 'views/login.html') {
				// already going to #login, no redirect needed
			} else {
				if (!$window.sessionStorage.token || $window.sessionStorage.token === '') {
					// not going to #login, we should redirect now
					console.log('no session token; redirect to login page');
					$location.path('/login');
				}
			}
		}
	});
});



app.directive('integer', function(){
    return {
        require: 'ngModel',
        link: function(scope, ele, attr, ctrl){
            ctrl.$parsers.unshift(function(viewValue){
				if (viewValue === '' || viewValue === null || typeof viewValue === 'undefined') {
					return null;
				}
                return parseInt(viewValue, 10);
            });
        }
    };
});