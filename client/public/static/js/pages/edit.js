mainApp.controller('editController', ['$scope', 'configConstant','$http','$window', '$location',
function ($scope, configConstant , $http, $window, $location) {
    $scope.init = function(){
      console.log($location.absUrl().split("/"));
        $scope.id = $location.absUrl().split("/")[4];
        $http({
            method: 'GET',
            url: configConstant.routerApi+'/administrators/'+$scope.id,
        }).then(function successCallback(response) {
              console.log(response);
              $scope.data = response.data;
          }, function errorCallback(response) {
              console.log(response)
        });
    }
    // initial data
    $scope.init();

    $scope.submitForm = function(isValid) {
      if($scope.data.message){
          var parameter = ($scope.data);
          $http.post(
             configConstant.routerApi+'/edit-administrators/'+ $scope.id,
             parameter
         ).then(function(data, status, headers, config) {
            console.log("succ");
            console.log(data);
            $window.location.href = '/';
          }).then(function(data, status, headers, config) {
            console.log("err");
          });
      }
    };

}]);
