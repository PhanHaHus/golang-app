mainApp.controller('homeController', ['$scope','$rootScope', 'apiConstant','$http',"$location","$state",
function ($scope,$rootScope, apiConstant, $http, $location,$state ) {
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
      $scope.setPage = function(pageNo) {
        $scope.currentPage = pageNo;
      };

      // set data for html
      $scope.changeData= function(data){
        var totalItem = data.data.Total;
        $scope.list = data.data.Data;
        $scope.currentPage = data.data.CurrentPage;
        $scope.per_page = data.data.PerPage;

        $scope.bigTotalItems = totalItem;
        $scope.bigCurrentPage = 1;
      }


}]);


mainApp.controller('accessruleListController', ['$scope', 'apiConstant','$http',"$window","$state","$stateParams",
function ($scope, apiConstant,$http, $window, $state,$stateParams) {
  $scope.init = function(){
      $http({
          method: 'GET',
          url: apiConstant+'/accessrules',
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
    $scope.totalItems = data.Total;
    $scope.per_page = data.PerPage;
  }
  console.log("aaa");
  $scope.setPage = function (pageNo) {
     $scope.currentPage = pageNo;
   };
   $scope.pageChanged = function() {
    $log.log('Page changed to: ' + $scope.currentPage);
  };
  $scope.maxSize = 1;
  $scope.bigTotalItems = 175;
  $scope.bigCurrentPage = 1;
  // initial data
  $scope.init();

}]);
