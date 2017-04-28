mainApp.controller('homeController', ['$scope','$rootScope', 'apiConstant','$http',"$location","$state",
function ($scope,$rootScope, apiConstant, $http, $location,$state ) {
      // show nav or not
      $rootScope.showNav = $location.path() != "/login" ? true: false;
      $scope.init = function(){
        $http({
          method: 'GET',
          url: apiConstant+'/administrators',
        }).then(function successCallback(response) {
          console.log(response);
            $scope.list = response.data.Data;
            }, function errorCallback(response) {
            console.log(response)
          });
      }
      $scope.deleteRemind= function(element){
        var result = confirm("Want to delete?");
          if (result) {
            $http({
                method: 'POST',
                url: apiConstant+'/del-administrators/'+element.AdministratorId,
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
