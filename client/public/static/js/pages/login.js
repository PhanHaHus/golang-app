mainApp.controller('loginController', ['$scope', 'apiConstant','$http',"$window","$state","$rootScope","$localStorage","toaster",
function ($scope, apiConstant,$http, $window ,$state,$rootScope,$localStorage,toaster) {

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
                }else{
                  toaster.pop('error', "ERROR!", "Username or password is incorrect!");
                }
              });
          }else{
            toaster.pop('error', "ERROR!", "Enter valid infor!");
          }
    };

}]).controller('logoutController', ['$scope', 'apiConstant','$http',"$window","$state","logoutService",
    function ($scope, apiConstant,$http, $window ,$state,logoutService) {
      logoutService.logout();
    }
]);
