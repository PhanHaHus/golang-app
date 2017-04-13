/*
This is a javascript file for config angular
============
Author:  Hapt
Created: april 2017
*/
var mainApp = angular.module("mainApp", ['ui.bootstrap'],function($interpolateProvider){
      $interpolateProvider.startSymbol('[[');
      $interpolateProvider.endSymbol(']]');
}).constant('configConstant', {
      routerApi: "http://localhost:8081/api"
});
