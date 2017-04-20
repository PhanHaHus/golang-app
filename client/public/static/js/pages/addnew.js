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
             }, function errorCallback(response) {
               console.log("err");
               console.log(response)
         });
       }
       $scope.init()
   }


    $scope.submitForm = function(isValid) {
          if(isValid){
            $http.post(
                 configConstant.routerApi+'/administrators' +($stateParams.id?('/'+$stateParams.id):''),
                 {
                   accepting_host_id: parseInt($scope.data.accepting_host_id),
                   description: $scope.data.description,
                   email: $scope.data.email,
                   enabled: $scope.data.enabled?1:0,
                   name: $scope.data.name,
                   password: $scope.data.password,
                   created_by_id: 1,
                   permission: $scope.data.permission
                 }
              ).then(function(data, status, headers, config) {
                console.log("succ");
                console.log(data);
                $state.go("home");
              }).then(function(data, status, headers, config) {
                console.log("err");
              });
          }
    };
}]);
