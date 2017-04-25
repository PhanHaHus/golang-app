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
                if(response.data.token){
                    var user = {
                      "token": response.data.token,
                      "username": response.data.username,
                    }
                    localStorage.setItem("userInfor", JSON.stringify(user));
                    $state.go("home");
                    location.reload();
                }else{
                  alert("login fail")
                }
              });
          }
    };

}]).controller('logoutController', ['$scope', 'configConstant','$http',"$window","$state",
    function ($scope, configConstant,$http, $window ,$state) {
        localStorage.clear();
        $state.go("login")
        location.reload();
    }
]);
