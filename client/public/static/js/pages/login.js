mainApp.controller('loginController', ['$scope', 'apiConstant','$http',"$window","$state","$rootScope","$localStorage",
function ($scope, apiConstant,$http, $window ,$state,$rootScope,$localStorage) {

    $scope.submitForm = function(isValid) {
      $scope.isValid = isValid;
          var parameter = ($scope.data);
          if(isValid){
              $http.post(
                 apiConstant+'/login',
                 parameter
             ).then(function(response, status, headers, config) {
                console.log("succ");
                console.log(response);
                if(response.data.token){
                    var user = {
                      "token": response.data.token,
                      "username": response.data.username,
                    }
                    $localStorage.userInfor = user;
                    $http.defaults.headers.common.Authorization = 'Bearer ' + response.data.token;
                    $state.go("home");
                    window.location.reload();
                }else{
                  $scope.error='Username or password is incorrect';
                }
              });
          }
    };

}]).controller('logoutController', ['$scope', 'apiConstant','$http',"$window","$state","logoutService",
    function ($scope, apiConstant,$http, $window ,$state,logoutService) {
      logoutService.logout();
        // localStorage.clear();
        // $http.defaults.headers.common.Authorization = '';
        // $state.go("login")
    }
]);
