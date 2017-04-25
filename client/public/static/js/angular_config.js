/*
This is a javascript file for config angular
============
Author:  Hapt
Created: april 2017
*/
var mainApp = angular.module("mainApp", ['ui.router','angular-loading-bar','angular-jwt'],function($interpolateProvider,cfpLoadingBarProvider){
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

//jwt config
mainApp.config(function Config($httpProvider, jwtOptionsProvider) {
    jwtOptionsProvider.config({
      unauthenticatedRedirector: ['$state', function($state) {
        $state.go('login');
      }]
    });
});

mainApp.run(function($http,$rootScope) {
  if(localStorage.getItem('userInfor')){
      var retrievedObject = JSON.parse(localStorage.getItem('userInfor'));
      console.log((retrievedObject));
      $http.defaults.headers.common.Authorization = 'Bearer ' + retrievedObject.token;
  }
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
           url: '/logout',
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
