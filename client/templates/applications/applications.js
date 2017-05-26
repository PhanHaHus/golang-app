mainApp.controller('applicationsListController', ['$scope', '$rootScope', 'apiConstant', '$http', "$location", "$state", "RESOURCES",
    function ($scope, $rootScope, apiConstant, $http, $location, $state, RESOURCES) {
        $scope.itemPerPage = RESOURCES.itemPerPage;
        $scope.permissionList = RESOURCES.permissionList;
        var ip_search, id_search, searchAll, type_search, name_search,enabled_search,port_search,hostname_search; //init value for filter
        $scope.getFilterVal = function () {
            id_search = ($scope.search && $scope.search.id_search != undefined) ? $scope.search.id_search : '';
            searchAll = ($scope.search && $scope.search.searchAll != undefined) ? $scope.search.searchAll : '';
            type_search = ($scope.search && $scope.search.type_search != undefined) ? $scope.search.type_search : '';
            name_search = ($scope.search && $scope.search.name_search != undefined) ? $scope.search.name_search : '';
            ip_search = ($scope.search && $scope.search.ip_search != undefined) ? $scope.search.ip_search : '';
            accepting_host_name = ($scope.search && $scope.search.accepting_host_name != undefined) ? $scope.search.accepting_host_name : '';
            port_search = ($scope.search && $scope.search.port_search != undefined) ? $scope.search.port_search : '';
            hostname_search = ($scope.search && $scope.search.hostname_search != undefined) ? $scope.search.hostname_search : '';

            enabled_search = ($scope.search && $scope.search.enabled_search != undefined) ? $scope.search.enabled_search : '';
            enabled_search ?( enabled_search = 1):( enabled_search = 0);//if check enabled = 1 else =0
        }

        $scope.init = function () {
            $http({
                method: 'GET',
                url: apiConstant + '/applications',
            }).then(function successCallback(response) {
                var totalItem = response.data.Total;
                var perpage = response.data.PerPage;
                $scope.bigTotalItems = totalItem;
                $scope.bigCurrentPage = 0;
                $scope.currentPage = 1;
                $scope.per_page = perpage;
                $scope.changeData(response);
            }, function errorCallback(response) {
                console.log(response)
            });
        }
        // initial data
        $scope.init();
        $scope.deleteRemind = function (element) {
            var result = confirm("Want to delete?");
            if (result) {
                $http({
                        method: 'POST',
                        url: apiConstant + '/del-applications/' + element.AdministratorId,
                        headers: {
                            'Content-type': 'application/json;charset=utf-8'
                        }
                    })
                    .then(function (response) {
                        $scope.init();
                    }, function (rejection) {
                        console.log(rejection);
                    });
            }
        }
        $scope.changePerpage = function (perpage) {
            $scope.getFilterVal();
            $http({
                    method: 'GET',
                    url: apiConstant + '/applications?per_page=' + perpage + "&query=" + searchAll + '&id_search=' + id_search + '&email_search=' + type_search + '&type_search=' + name_search + '&ip_search=' + ip_search + '&accepting_host_name=' + accepting_host_name+ '&port_search=' + port_search+ '&ip_search=' + ip_search+'&hostname_search=' + hostname_search,
                    headers: {
                        'Content-type': 'application/json;charset=utf-8'
                    }
                })
                .then(function (response) {
                    $scope.per_page = perpage;
                    $scope.changeData(response);
                }, function (rejection) {
                    console.log(rejection);
                });
        }
        //for search box
        $scope.onSelectSearchBox = function ($item, $model, $label) {
            $scope.getFilterVal();
            return $http({
                method: 'GET',
                cache: true,
                url: apiConstant + '/applications?query=' + searchAll + '&id_search=' + id_search + '&email_search=' + email_search + '&name_search=' + name_search + '&permission=' + permission + '&accepting_host_name=' + accepting_host_name,
            }).then(function successCallback(response) {
                var totalItem = response.data.Total;
                var perpage = response.data.PerPage;
                $scope.bigTotalItems = totalItem;
                $scope.per_page = perpage;
                $scope.changeData(response);
                return $scope.list
            }, function errorCallback(response) {
                console.log("err");
                console.log(response)
            });
        };

        //for search below box( filter )
        $scope.onSearchBelowBox = function ($item, $model, $label) {
            //get filter val params
            $scope.getFilterVal();
            return $http({
                method: 'GET',
                url: apiConstant + '/applications?query=' + searchAll + '&id_search=' + id_search + '&email_search=' + email_search + '&name_search=' + name_search + '&permission=' + permission + '&accepting_host_name=' + accepting_host_name,
            }).then(function successCallback(response) {
                var totalItem = response.data.Total;
                var perpage = response.data.PerPage;
                $scope.bigTotalItems = totalItem;
                $scope.per_page = perpage;
                $scope.changeData(response);
                return $scope.list;
            }, function errorCallback(response) {
                console.log("err");
                console.log(response)
            });
        };

        $scope.pageChanged = function () {
            //get filter val params before call API
            $scope.getFilterVal();
            $http({
                    method: 'GET',
                    url: apiConstant + '/applications?per_page=' + $scope.per_page + "&current_page=" + $scope.currentPage + "&query=" + searchAll + '&id_search=' + id_search + '&email_search=' + email_search + '&name_search=' + name_search + '&permission=' + permission + '&accepting_host_name=' + accepting_host_name,
                    headers: {
                        'Content-type': 'application/json;charset=utf-8'
                    }
                })
                .then(function (response) {
                    $scope.changeData(response);
                }, function (rejection) {
                    console.log(rejection);
                });
            console.log('Page changed to: ' + $scope.currentPage);
        };

        // set data for html
        $scope.changeData = function (data) {
            $scope.list = data.data.Data;
        }
    }
]);


//add and edit admin
mainApp.controller('applicationsController', ['$scope', 'apiConstant', '$http', "$window", "$state", "$stateParams", "toaster", "RESOURCES",
    function ($scope, apiConstant, $http, $window, $state, $stateParams, toaster, RESOURCES) {
        $scope.data = {
            accepting_host_id: null,
            description: "",
            email: "",
            enabled: false,
            name: "",
            password: "",
            permission: ""
        };
        $scope.permissionList = RESOURCES.permissionList;
        //case edit
        if ($stateParams.id) {
            $scope.init = function () {
                $http({
                    method: 'GET',
                    url: apiConstant + '/applications/' + $stateParams.id,
                }).then(function successCallback(response) {
                    $scope.data = {
                        accepting_host_id: response.data.accepting_host_id,
                        description: response.data.description,
                        email: response.data.email,
                        enabled: (response.data.enabled == 0 ? false : true),
                        name: response.data.name,
                        password: response.data.password,
                        permission: response.data.permission
                    };
                }, function errorCallback(response) {
                    console.log("err");
                    console.log(response)
                });
            }
            $scope.init();
        }

        $scope.submitForm = function (isValid) {
            console.log($scope.data);
            if (isValid) {
                var apiUrl = apiConstant + '/applications'; //api add
                if ($stateParams.id) {
                    apiUrl = apiConstant + '/edit-applications/' + $stateParams.id; //api edit if exist id;
                }
                var accepting_host_id = parseInt($scope.data.accepting_host_id) == 0 ? 1 : parseInt($scope.data.accepting_host_id);
                var dataPost = {
                    accepting_host_id: accepting_host_id,
                    description: $scope.data.description,
                    email: $scope.data.email,
                    enabled: $scope.data.enabled ? 1 : 0,
                    name: $scope.data.name,
                    password: $scope.data.password,
                    created_by_id: 1,
                    permission: $scope.data.permission
                };

                $http({
                    method: 'POST',
                    url: apiUrl,
                    data: dataPost
                }).then(function (data) {
                    console.log("succ");
                    console.log(data);
                    $state.go("home");
                }, function () {
                    console.log("err");
                });
            } else {
                toaster.pop('error', "ERROR!", "Enter valid infor!");
            }
        };

    }
]);

mainApp.controller('detailApplicationsController', ['$scope', 'apiConstant', '$http', "$window", "$state", "$stateParams",
    function ($scope, apiConstant, $http, $window, $state, $stateParams) {
        $scope.data = {
            accepting_host_id: null,
            description: "",
            email: "",
            enabled: false,
            name: "",
            password: "",
            permission: ""
        };

        if ($stateParams.id) {
            $scope.init = function () {
                $http({
                    method: 'GET',
                    url: apiConstant + '/applications/' + $stateParams.id,
                }).then(function successCallback(response) {
                    console.log(response);
                    $scope.data = {
                        accepting_host_id: response.data.accepting_host_id,
                        accepting_host_name: response.data.AcceptingHost.name,
                        description: response.data.description,
                        email: response.data.email,
                        enabled: response.data.enabled,
                        name: response.data.name,
                        password: response.data.password,
                        permission: response.data.permission
                    };
                }, function errorCallback(response) {
                    console.log("err");
                    console.log(response)
                });
            }
            $scope.init()
        }

    }
]);