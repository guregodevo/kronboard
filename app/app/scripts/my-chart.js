'use strict';

angular.module('chart', []).directive('mychart', ['$location','$interval', '$timeout', '$window', '$rootScope','ChartApi',
function($location, $interval, $timeout, $window, $rootScope, ChartApi) {
	return {
		restrict : 'A',
		scope : {
			dashboardid : '@',
			chartid : '@',
			readonly: '@',
			type : '@',
			//close : '&onClose',
			onReady : '&',
			select : '&'
		},
		templateUrl : 'views/inc/my-chart.html',
		link : function($scope, $elm, $attr) {
			var timeoutId;
			//console.log('read only ' + $scope.readonly);
			//console.log('link '+  angular.toJson($elm.parent()));

			resize();
			
			function resize() {
				//adjust dimension
				var width = $elm.parent().css('width');
				var absWidth = width.substring(0, width.length - 2);
				//remove px chars
				var height = $elm.parent().css('height');
				var absHeight = height.substring(0, height.length - 2);
				//remove px chars
				$elm.find("div").css({
					width : absWidth - 10 + 'px',
					height : absHeight -10 + 'px'
				});
			}

			var chart;
 
			$scope.refresh = function() {
				updateTime();
			};

			$scope.close = function() {
		      $timeout(function () {
		       console.log('chart id' + $scope.chartid );
			   console.log('dashboard id ' + $scope.dashboardid );
			   ChartApi.deletechart($scope.dashboardid, $scope.chartid).then(function(data) {//success
					console.log('Deleted chart');
					chart.clear();
					$elm.hide();
					$elm.find("div").hide();
					$elm.parent().hide();	
				}, function(status) {
					//failed
					$scope.msg = 'Invalid chart delete';
					console.log('Get charts service failed');
				});
		      }, 200);
			};

			$scope.edit = function() {
				$location.path("/chart/"+$scope.chartid+'/'+ $scope.type+'/');
			};
			
			function create_date(day) {
				  var d = new Date();
				  return new Date(d.getFullYear(), d.getMonth(), day + 1);
			};
			
			function create_exponential_points() {
				  var i, _results;
				  _results = [];
				  for (i = 0; i <= 25; i++) {
				    _results.push([create_date(i), i * 4.]);
				  }
				  return _results;
			};

			function create_squared_points() {
				  var i, _results;
				  _results = [];
				  for (i = 0; i <= 25; i++) {
				    _results.push([create_date(i), i * (i - 1)]);
				  }
				  return _results;
			};
			
			function create_random_points() {
				  var i, _results;
				  _results = [];
				  for (i = 0; i <= 25; i++) {
				    _results.push([create_date(i), Math.random() * i]);
				  }
				  return _results;
			};		
	
				
			function updateTime() {
				if (chart) {
					chart.clear();
				}
				var e = $elm.find("div")[0];
				switch ($attr.type) {
					case "circle":
						chart = new Charts.CircleProgress(e, 'Sales', Math.random() * 100, {
							font_color : "#fff",
							fill_color : "#222",
							label_color : "#333",
							text_shadow : "rgba(0,0,0,.4)",
							stroke_color : "#6a329e"
						});
						break;
					case "index":
						chart = Charts.IndexChart(e);
						chart.add("Retail", 18316, 65);
						chart.add("Engineering/Technical", 28977, 282);
						chart.add("Education", 20839, 106);
						chart.add_guide_line("Average", 100, 1);
						chart.add_guide_line("Above Average", 200, 0.25);
						chart.add_guide_line("High", 300, 0.25);
						break;
					case "line":
						chart = new Charts.LineChart(e);
						chart.add_line({
							data : create_random_points()
						});
						break;
					case "bullet":
						chart = new Charts.BulletChart(e);
						chart.add("foo", 50, 30, 100);
						chart.add("doo", 70, 30, 100);
						chart.add("moo", 20, 30, 100);
						break;
					case "bar":
						chart = new Charts.BarChart(e, {
							x_label_color : "#333333",
							bar_width : 45,
							rounding : 10,
						});
						chart.add({
							label : "orange",
							value : 300
						});
						chart.add({
							label : "doo",
							value : 50
						});
						chart.add({
							label : "dii",
							value : 300
						});
						chart.add({
							label : "daa",
							value : 800
						});
						chart.add({
							label : "daa",
							value : 50
						});
						break;
				}
				chart.draw();
			}

			// start the UI update process; save the timeoutId for canceling
			timeoutId = $interval(function() {
				updateTime();
			}, 100 * 1000);

			// Watches, to refresh the chart when its data, title or dimensions change
			$scope.$watch('chart', function() {
				console.log('watch() mychart-js');
				resize();
				updateTime();
			}, true);

			$rootScope.$on('$destroy', function() {
				console.log('destroy chart');
				$interval.cancel(timeoutId);
				chart.clear();
			});

			// Redraw the chart if the window is resized
			$rootScope.$on('resizeMsg', function(e) {
				resize();
				updateTime();
			});

		}
	};
}]).run(['$rootScope', '$window',
function($rootScope, $window) {
	angular.element($window).bind('resize', function() {
		$rootScope.$emit('resizeMsg');
	});
	angular.element($window).bind('hashchange', function() {
		console.log('back button ');
		$rootScope.$emit('resizeMsg');
	});
}]);
