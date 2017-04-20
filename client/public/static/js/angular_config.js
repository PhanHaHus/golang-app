/*
This is a javascript file for config angular
============
Author:  Hapt
Created: april 2017
*/
var mainApp = angular.module("mainApp", ['ui.router','angular-loading-bar'],function($interpolateProvider,cfpLoadingBarProvider){
      cfpLoadingBarProvider.includeSpinner = true;
      $interpolateProvider.startSymbol('[[');
      $interpolateProvider.endSymbol(']]');
}).constant('configConstant', {
      routerApi: "http://localhost:8081/api"
});


//Authentication in each route and controller
// mainApp.factory('Auth', function(){
//   var user;
//   return{
//       setUser : function(aUser){
//           user = aUser;
//       },
//       isLoggedIn : function(){
//           return(user)? user : false;
//       }
//   }
// });
// mainApp.run(['$rootScope', '$location', 'Auth', function ($rootScope, $location, Auth) {
//     $rootScope.$on('$routeChangeStart', function (event) {
//         if (!Auth.isLoggedIn()) {
//             console.log('DENY');
//             event.preventDefault();
//             $location.path('/login');
//         }
//         else {
//             console.log('ALLOW');
//             $location.path('/home');
//         }
//     });
// }]);

// directive navbar
mainApp.directive('navBar', function() {
  return {
    templateUrl: '/templates/_nav.html',
  };
});

// router ui
mainApp.config(function($stateProvider,$urlRouterProvider) {
  $urlRouterProvider.otherwise('/page-not-found');
  $stateProvider.state('home', {
           url: '/home',
           controller: 'homeController',
           templateUrl: '/templates/admin/list.html',
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
       .state('about', {
         name: 'about',
         url: '/about',
         template: '<h3>Its the UI-Router hello world app!</h3>'
       }
   );

});
