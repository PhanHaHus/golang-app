mainApp.controller('addNewController', ['$scope', 'configConstant','$http',"$window","$state","$stateParams",
function ($scope, configConstant,$http, $window, $state,$stateParams) {
    $scope.data = {
        accepting_host_id:null,
        description:"",
        email:"",
        enabled:false,
        name:"",
        password:"",
        permission:""
    };
     console.log($stateParams);
     //case edit
     if($stateParams.id){
       $scope.init = function(){
          $http({
             method: 'GET',
             url: configConstant.routerApi+'/administrators/'+$stateParams.id,
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


      $scope.submitForm = function(isValid) {
          if(isValid){
            var apiUrl = configConstant.routerApi+'/administrators'; //api add
            if($stateParams.id){
                apiUrl = configConstant.routerApi+'/edit-administrators/' + $stateParams.id;//api edit if exist id;
            }
            var dataPost = {
                accepting_host_id: parseInt($scope.data.accepting_host_id)==0?1:parseInt($scope.data.accepting_host_id),
                description: $scope.data.description,
                email: $scope.data.email,
                enabled: $scope.data.enabled?1:0,
                name: $scope.data.name,
                password: $scope.data.password,
                created_by_id: 1,
                permission: $scope.data.permission
            };

            $http({
                 method: 'POST',
                 url: apiUrl,
                 data: dataPost
            }).then(function(data){
                  console.log("succ");
                  console.log(data);
                  $state.go("home");
            },function(){
                console.log("err");
            });
          }
      };

}]);
