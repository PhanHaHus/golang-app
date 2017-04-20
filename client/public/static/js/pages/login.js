mainApp.controller('loginController', ['$scope', 'configConstant','$http',"$window","$state",
function ($scope, configConstant,$http, $window ,$state) {

    $scope.submitForm = function(isValid) {
      $scope.isValid = isValid;
          var parameter = ($scope.data);
          if(isValid){
              $http.post(
                 configConstant.routerApi+'/login',
                 parameter
             ).then(function(response, status, headers, config) {
                console.log("succ");
                console.log(response);
                if(response.data.status=="true"){
                    localStorage.setItem("Token", response.data.token);
                    $state.go("home");
                }
              });
          }

    };

}]);
