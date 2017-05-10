mainApp.controller('acceptingHostListController', ['$scope','$rootScope', 'apiConstant','$http',"$location","$state",
function ($scope,$rootScope, apiConstant, $http, $location,$state ) {
      $scope.itemPerPage = [
        {value: '10', displayName: '10 items'},
        {value: '20', displayName: '20 items'},
        {value: '50', displayName: '50 items'},
        {value: '100', displayName: '100 items'}
     ];

      $scope.init = function(){
          $http({
              method: 'GET',
              url: apiConstant+'/acceptinghosts',
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
                url: apiConstant+'/del-acceptinghosts/'+element.AdministratorId,
                headers: {
                    'Content-type': 'application/json;charset=utf-8'
                }
            })
            .then(function(response) {
                $scope.init();
            }, function(rejection) {
                console.log(rejection);
            });
          }
      }
      $scope.changePerpage= function(perpage){
            $http({
                method: 'GET',
                url: apiConstant+'/acceptinghosts?per_page='+perpage+"&query="+($scope.asyncSelected?$scope.asyncSelected:''),//if not search, query = ''
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
      //for search box
        $scope.onSelect = function ($item, $model, $label) {
            return $http({
                method: 'GET',
                url: apiConstant+'/acceptinghosts?query='+($scope.asyncSelected?$scope.asyncSelected:''),
                }).then(function successCallback(response) {
                    var totalItem = response.data.Total;
                    var perpage = response.data.PerPage;
                    $scope.bigTotalItems = totalItem;
                    $scope.per_page = perpage;
                    $scope.changeData(response);
                    return $scope.list
                }, function errorCallback(response) {
                    console.log("err");
                    console.log(response)
                });
        };
        

      $scope.pageChanged = function() {
          $http({
              method: 'GET',
              url: apiConstant+'/acceptinghosts?per_page='+$scope.per_page+"&current_page="+$scope.currentPage+"&query="+($scope.asyncSelected?$scope.asyncSelected:''),//if not search, query = ''
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


//add and edit admin
mainApp.controller('acceptingHostController', ['$scope', 'apiConstant','$http',"$window","$state","$stateParams","toaster",
function ($scope, apiConstant,$http, $window, $state,$stateParams,toaster) {
      $scope.data = {
          description:"",
          enabled:false,
          name:"",
          created_by_id:"",
          permission:""
      };
     //case edit
     if($stateParams.id){
       $scope.init = function(){
          $http({
             method: 'GET',
             url: apiConstant+'/acceptinghosts/'+$stateParams.id,
           }).then(function successCallback(response) {
                $scope.data = {
                    description:response.data.description,
                    enabled:(response.data.enabled==0?false:true),
                    name:response.data.name,
                };
           }, function errorCallback(response) {
                 console.log("err");
                 console.log(response)
           });
         }
         $scope.init();
     }

      $scope.submitForm = function(isValid) {
        console.log($scope.data);
          if(isValid){
            var apiUrl = apiConstant+'/acceptinghosts'; //api add
            if($stateParams.id){
                apiUrl = apiConstant+'/edit-acceptinghosts/' + $stateParams.id;//api edit if exist id;
            }
            var dataPost = {
                description: $scope.data.description,
                enabled: $scope.data.enabled?1:0,
                name: $scope.data.name,
                created_by_id: 1,
            };

            $http({
                 method: 'POST',
                 url: apiUrl,
                 data: dataPost
            }).then(function(data){
                  console.log("succ");
                  console.log(data);
                  $state.go("acceptinghost");
            },function(){
                console.log("err");
            });
          }else{
            toaster.pop('error', "ERROR!", "Enter valid infor!");
          }
      };
}]);

mainApp.controller('detailAcceptingHostController', ['$scope', 'apiConstant','$http',"$window","$state","$stateParams",
function ($scope, apiConstant,$http, $window, $state,$stateParams) {
    $scope.data = {
        accepting_host_id:null,
        description:"",
        email:"",
        enabled:false,
        name:"",
        password:"",
        permission:""
    };

     if($stateParams.id){
       $scope.init = function(){
          $http({
             method: 'GET',
             url: apiConstant+'/acceptinghosts/'+$stateParams.id,
           }).then(function successCallback(response) {
                console.log(response);
                $scope.data = {
                    accepting_host_id:response.data.accepting_host_id,
                    description:response.data.description,
                    email:response.data.email,
                    enabled:response.data.enabled,
                    name:response.data.name,
                    password:response.data.password,
                    permission:response.data.permission
                };
           }, function errorCallback(response) {
                 console.log("err");
                 console.log(response)
           });
         }
         $scope.init()
     }

}]);
