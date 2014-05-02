'use strict';

angular.module('dialog', []).directive('dialog', ['$timeout',
function($timeout) {
	return {
		restrict : 'E',
		transclude : true,
		scope : {
			'chartid' : '@'
		},
		templateUrl : 'views/inc/dialog.html',
		link : function($scope, $elm, $attr) {
			$scope.dialogIsHidden = false;
			$scope.buttonIsHidden = false;
			$elm.attr('style', $elm.parent().attr('style'));

			$scope.close = function() {
				$scope.buttonIsHidden = true;
				$scope.dialogIsHidden = true;
				console.log('call api remove chart from dashboard');
			};
		}
	};
}]);
