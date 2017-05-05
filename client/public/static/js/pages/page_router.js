// router ui config
mainApp.config(function($stateProvider,$urlRouterProvider) {
  $urlRouterProvider.otherwise('/');
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
       .state('detailadmin', {
           url: '/detail-admin/:id',
           params: {
              id: null
            },
           controller: 'detailController',
           templateUrl: '/templates/admin/form_detail.html',
       })
      //  end admin
        // accessrule
        .state('accessrule', {
             url: '/accessrule',
             controller: 'accessruleListController',
             templateUrl: '/templates/accessrule/list.html',
        })
        .state('detailaccessrule', {
             url: '/accessrule/:id',
             params: {
                id: null
              },
             controller: 'accessruleDetailController',
             templateUrl: '/templates/accessrule/form_detail.html',
        })
        .state('addaccessrule', {
             url: '/add-accessrule',
             controller: 'accessRuleController',
             templateUrl: '/templates/accessrule/form.html',
        })
        .state('editaccessrule', {
            url: '/edit-accessrule/:id',
            params: {
               id: null
             },
            controller: 'accessRuleController',
            templateUrl: '/templates/admin/form.html',
        })
        // end accessRules
        .state('dashboard', {
             url: '/',
             controller: 'dashboardController',
             templateUrl: '/templates/dashboard/dashboard.html',
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
       });

});
