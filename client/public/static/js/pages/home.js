mainApp.controller('homeController', ['$scope','$rootScope', 'configConstant','$http',"$location","$state",
function ($scope,$rootScope, configConstant, $http, $location,$state ) {
  console.log(localStorage);
      // show nav or not
      $rootScope.showNav = $location.path() != "/login" ? true: false;
      $scope.init = function(){
        $http({
          method: 'GET',
          url: configConstant.routerApi+'/administrators',
        }).then(function successCallback(response) {
            $scope.list = response.data;
            }, function errorCallback(response) {
            console.log(response)
          });
      }
      $scope.deleteRemind= function(element){
        var result = confirm("Want to delete?");
          if (result) {
            $http({
                method: 'POST',
                url: configConstant.routerApi+'/del-administrators/'+element.AdministratorId,
                headers: {
                    'Content-type': 'application/json;charset=utf-8'
                }
            })
            .then(function(response) {
                console.log(response);
                $scope.init();
            }, function(rejection) {
                console.log(rejection);
            });
          }
      }
      // initial data
      $scope.init();
}]);
