
mainApp.controller('accessruleListController', ['$scope','$rootScope', 'apiConstant','$http',"$location","$state",
function ($scope,$rootScope, apiConstant, $http, $location,$state ) {

      $scope.init = function(){
          $http({
              method: 'GET',
              url: apiConstant+'/accessrules',
          }).then(function successCallback(response) {
              var totalItem = response.data.Total;
              var perpage = response.data.PerPage;
              $scope.bigTotalItems = totalItem;
              $scope.bigCurrentPage = 0;
              $scope.currentPage = 1;
              $scope.per_page = perpage;
              $scope.changeData(response);
          }, function errorCallback(response) {
              console.log(response)
          });
      }
      // initial data
      $scope.init();
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
                $scope.per_page = perpage;
                $scope.changeData(response);
            }, function(rejection) {
                console.log(rejection);
            });
      }

      $scope.pageChanged = function() {
          $http({
              method: 'GET',
              url: apiConstant+'/administrators?per_page='+$scope.per_page+"&current_page="+$scope.currentPage,
              headers: {
                  'Content-type': 'application/json;charset=utf-8'
              }
          })
          .then(function(response) {
              $scope.changeData(response);
          }, function(rejection) {
              console.log(rejection);
          });
          console.log('Page changed to: ' + $scope.currentPage);
      };

      // set data for html
      $scope.changeData= function(data){
        $scope.list = data.data.Data;
      }

}]);
