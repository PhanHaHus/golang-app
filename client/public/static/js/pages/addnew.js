mainApp.controller('addNewController', ['$scope', 'configConstant','$http',"$window",
function ($scope, configConstant,$http, $window) {
    $scope.data = {
      accepting_host_id:null,
      description:"",
      email:"",
      enabled:false,
      name:"",
      password:"",
      permission:""
    };
    $scope.submitForm = function(isValid) {
          var parameter = ($scope.data);
          console.log(parameter);
            $http.post(
               configConstant.routerApi+'/administrators',
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
              $window.location.href = '/';
            }).then(function(data, status, headers, config) {
              console.log("err");
            });
    };
}]);
