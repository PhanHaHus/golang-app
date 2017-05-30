// router ui config
mainApp.config(function($stateProvider,$urlRouterProvider) {
  $urlRouterProvider.otherwise('/home');
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
      // acceptingHost
      .state('acceptinghost', {
           url: '/accepting-host',
           controller: 'acceptingHostListController',
           templateUrl: '/templates/acceptinghost/list.html',
       })
       .state('addacceptinghost', {
           url: '/add-accepting-host',
           controller: 'acceptingHostController',
           templateUrl: '/templates/acceptinghost/form.html',
       })
       .state('editacceptinghost', {
           url: '/edit-accepting-host/:id',
           params: {
              id: null
            },
           controller: 'acceptingHostController',
           templateUrl: '/templates/acceptinghost/form.html',
       })
       .state('detailacceptinghost', {
           url: '/detail-accepting-host/:id',
           params: {
              id: null
            },
           controller: 'detailAcceptingHostController',
           templateUrl: '/templates/acceptinghost/form_detail.html',
       })
      //  end acceptingHost
      // initial Host
      .state('initiatinghosts', {
           url: '/initiating-host',
           controller: 'initiatingHostListController',
           templateUrl: '/templates/initiatinghosts/list.html',
       })
       .state('addinitiatinghosts', {
           url: '/add-initiating-host',
           controller: 'initiatingHostController',
           templateUrl: '/templates/initiatinghosts/form.html',
       })
       .state('editinitiatinghosts', {
           url: '/edit-initiating-host/:id',
           params: {
              id: null
            },
           controller: 'initiatingHostController',
           templateUrl: '/templates/initiatinghosts/form.html',
       })
       .state('detailinitiatinghosts', {
           url: '/detail-initiating-host/:id',
           params: {
              id: null
            },
           controller: 'detailInitiatingHostController',
           templateUrl: '/templates/initiatinghosts/form_detail.html',
       })
      //  end initialHost
      // applications
      .state('applications', {
           url: '/applications',
           controller: 'applicationsListController',
           templateUrl: '/templates/applications/list.html',
       })
       .state('addapplications', {
           url: '/add-applications',
           controller: 'applicationsController',
           templateUrl: '/templates/applications/form.html',
       })
       .state('editapplications', {
           url: '/edit-applications-host/:id',
           params: {
              id: null
            },
           controller: 'applicationsController',
           templateUrl: '/templates/applications/form.html',
       })
       .state('detailapplications', {
           url: '/detail-applications/:id',
           params: {
              id: null
            },
           controller: 'detailApplicationsController',
           templateUrl: '/templates/applications/form_detail.html',
       })
      //  end applications
      // applications
      .state('certificate', {
           url: '/certificate',
           controller: 'certificateListController',
           templateUrl: '/templates/certificate/list.html',
       })
       .state('addcertificate', {
           url: '/add-certificate',
           controller: 'certificateController',
           templateUrl: '/templates/certificate/form.html',
       })
       .state('editcertificate', {
           url: '/edit-certificate-host/:id',
           params: {
              id: null
            },
           controller: 'certificateController',
           templateUrl: '/templates/certificate/form.html',
       })
       .state('detailcertificate', {
           url: '/detail-certificate/:id',
           params: {
              id: null
            },
           controller: 'detailCertificateController',
           templateUrl: '/templates/certificate/form_detail.html',
       })
      //  end certificate
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
            templateUrl: '/templates/accessrule/form.html',
        })
        // end accessRules
        .state('dashboard', {
             url: '/',
             //controller: 'dashboardController',
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
