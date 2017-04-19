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


mainApp.config(function($stateProvider,$urlRouterProvider) {
  $urlRouterProvider.otherwise('/home');
  $stateProvider.state('home', {
           url: '/home',
           controller: 'homeController',
           templateUrl: '/templates/admin/list.html',
       })
       .state('home.add', {
           url: '/add-admin',
           controller: 'addNewController',
           templateUrl: '/templates/admin/add.html',
       })
       .state('about', {
         name: 'about',
         url: '/about',
         template: '<h3>Its the UI-Router hello world app!</h3>'
       }
   );

});
