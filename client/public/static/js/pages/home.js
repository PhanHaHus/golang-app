mainApp.controller('homeController', ['$scope', 'configConstant','$http',
function ($scope, configConstant,$http) {
      $scope.init = function(){
        $http({
          method: 'GET',
          url: configConstant.routerApi+'/reminder',
        }).then(function successCallback(response) {
          $scope.list = response.data;
            }, function errorCallback(response) {
            console.log(response)
          });
      }
      $scope.deleteRemind= function(element){
        var result = confirm("Want to delete?");
          if (result) {
            $http.delete(configConstant.routerApi+'/reminder/'+element.Id,{}).then(function(response){
                console.log(response);
            }).then(function(response){
                console.log(response);
            })

            // $http({
            //     method: 'DELETE',
            //     url: configConstant.routerApi+'/reminder/'+element.Id,
            //     headers: {
            //         'Content-type': 'application/json;charset=utf-8'
            //     }
            // })
            // .then(function(response) {
            //     console.log(response);
            // }, function(rejection) {
            //     console.log(rejection);
            // });
          }
      }


      $scope.init();
}]);
