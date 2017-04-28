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
              $scope.changeData(response);
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
      $scope.changePerpage= function(perpage){
            $http({
                method: 'GET',
                url: apiConstant+'/administrators?per_page='+perpage,
                headers: {
                    'Content-type': 'application/json;charset=utf-8'
                }
            })
            .then(function(response) {
                $scope.changeData(response);
            }, function(rejection) {
                console.log(rejection);
            });
      }
      // set data for html
      $scope.changeData= function(data){
        $scope.list = data.data.Data;
        $scope.currentPage = data.CurrentPage;
        $scope.total = data.Total;
        $scope.per_page = data.PerPage;
      }

      // initial data
      $scope.init();
}]);
