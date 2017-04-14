mainApp.controller('addNewController', ['$scope', 'configConstant','$http',"$window",
function ($scope, configConstant,$http, $window) {
    $scope.data = {
        message: ""
    };
    $scope.submitForm = function(isValid) {
      if($scope.data.message){
          var parameter = ($scope.data);
          $http.post(
             configConstant.routerApi+'/reminder',
             parameter
         ).then(function(data, status, headers, config) {
            console.log(data);
            console.log(status);
            // $window.location.href = '/';
          }).then(function(data, status, headers, config) {
            console.log(data);
            console.log(status);
          });
      }
    };
}]);
