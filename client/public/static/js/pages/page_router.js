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
       .state('accessrule', {
         url: '/accessrule',
         controller: 'accessruleListController',
         template: '/templates/accessrule/list.html'
       }
   );

});
