/*
Author:  Hapt
Created: april 2017
*/
var mainApp = angular.module("mainApp", ['ui.router','angular-loading-bar','angular-jwt', 'ngStorage'],function($interpolateProvider,cfpLoadingBarProvider){
      cfpLoadingBarProvider.includeSpinner = true;
      $interpolateProvider.startSymbol('[[');
      $interpolateProvider.endSymbol(']]');
}).constant('configConstant', {
      routerApi: "http://localhost:8081/api"
});

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
mainApp.run(function($http,$rootScope, $location, $localStorage,jwtHelper,logoutService) {
  if ($localStorage.userInfor) {
      $http.defaults.headers.common.Authorization = 'Bearer ' + $localStorage.userInfor.token;
  }

  // redirect to login page if not logged in and trying to access a restricted page or expired token
  $rootScope.$on('$locationChangeStart', function (event, next, current) {
      var publicPages = ['/login'];
      var restrictedPage = publicPages.indexOf($location.path()) === -1;
      if (restrictedPage && !$localStorage.userInfor) {
          $location.path('/login');
      }
      if ($localStorage.userInfor) {
        //check expired, redirect to login
        var isTokenExpired = jwtHelper.isTokenExpired($localStorage.userInfor.token);
        if(isTokenExpired){
          logoutService.logout();
          alert('Your session has expired!');
        }else{
          //if isset session
          $location.path('/home');
        }
      }

  });
});

// router ui config
mainApp.config(function($stateProvider,$urlRouterProvider) {
  $urlRouterProvider.otherwise('/page-not-found');
  $stateProvider.state('home', {
           url: '/home',
           controller: 'homeController',
           templateUrl: '/templates/admin/list.html',
       })
       .state('dashboard', {
            url: '/',
            controller: 'dashboardController',
            templateUrl: '/templates/dashboard/dashboard.html',
        })
       .state('addadmin', {
           url: '/add-admin',
           controller: 'addNewController',
           templateUrl: '/templates/admin/form.html',
       })
       .state('editadmin', {
           url: '/edit-admin/:id',
           params: {
              id: null
            },
           controller: 'addNewController',
           templateUrl: '/templates/admin/form.html',
       })
       .state('detailadmin', {
           url: '/detail-admin/:id',
           params: {
              id: null
            },
           controller: 'detailController',
           templateUrl: '/templates/admin/form_detail.html',
       })
       .state('page-not-found', {
         url: '/page-not-found',
           templateUrl: '/templates/error/404.html',
           data: {
               displayName: false
           }
       })
       .state('login', {
           url: '/login',
           controller: 'loginController',
           templateUrl: '/templates/login.html',
       })
       .state('logout', {
          //  url: '/logout',
           controller: 'logoutController',
           templateUrl: "/templates/login.html",
       })
       .state('about', {
         name: 'about',
         url: '/about',
         template: '<h3>Its the UI-Router hello world app!</h3>'
       }
   );

});
