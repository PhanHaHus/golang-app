/*
Author:  Hapt
Created: april 2017
*/
var mainApp = angular.module("mainApp", ['ui.router','angular-loading-bar','angular-jwt', 'ngStorage','toaster','ui.bootstrap'],function($interpolateProvider,cfpLoadingBarProvider){
      cfpLoadingBarProvider.includeSpinner = true;
      $interpolateProvider.startSymbol('[[');
      $interpolateProvider.endSymbol(']]');
}).constant('apiConstant','http://localhost:8081/api');

// directive navbar
mainApp.directive('navBar', function() {
  return {
    templateUrl: '/templates/_nav.html',
  };
});
mainApp.factory('logoutService', function($http,$location) {
      return {
          logout: function() {
            localStorage.clear();
            $http.defaults.headers.common.Authorization = '';
            $location.path('/login');
          }
      };
});
mainApp.run(function($http,$rootScope, $location, $localStorage,jwtHelper,logoutService,toaster) {
  if ($localStorage.userInfor) {
      $http.defaults.headers.common.Authorization = 'Bearer ' + $localStorage.userInfor.token;
  }

  // redirect to login page if not logged in and trying to access a restricted page or expired token
  $rootScope.$on('$locationChangeStart', function (event, next, current) {
      var publicPages = ['/login'];
      var restrictedPage = publicPages.indexOf($location.path()) === -1;// url !=login
      if (restrictedPage && !$localStorage.userInfor) {
          $location.path('/login');
      }
      if ($localStorage.userInfor) {
        //check expired, redirect to login
        var isTokenExpired = jwtHelper.isTokenExpired($localStorage.userInfor.token);
        if(isTokenExpired){
          logoutService.logout();
          console.log("logout");
          toaster.pop('error', "Ops!", "Your session has expired!");
        }
      }

  });
});
