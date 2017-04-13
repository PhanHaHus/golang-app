mainApp.controller('homeController', ['$scope', 'configConstant','$http',
function ($scope, configConstant,$http) {
      $scope.init = function(){
        $http({
          method: 'GET',
          url: configConstant.routerApi+'/reminder',
        }).then(function successCallback(response) {
          $scope.list = response.data;
          console.log($scope.list);
            }, function errorCallback(response) {
            console.log(response)
          });
      }
      $scope.test =23;
      $scope.init();
}]);
