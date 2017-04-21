mainApp.controller('detailController', ['$scope', 'configConstant','$http',"$window","$state","$stateParams",
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

}]);
