mainApp.controller('applicationsListController', ['$scope', '$rootScope', 'apiConstant', '$http', "$location", "$state", "RESOURCES","toaster",
    function ($scope, $rootScope, apiConstant, $http, $location, $state, RESOURCES,toaster) {
        $scope.itemPerPage = RESOURCES.itemPerPage;
        $scope.applicationType = RESOURCES.applicationType;
        $scope.status = RESOURCES.status;
        var ip_search, id_search, searchAll, type_search, name_search, enabled_search, port_search, hostname_search; //init value for filter
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
            enabled_search ? (enabled_search = 1) : (enabled_search = 0); //if check enabled = 1 else =0
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
                        url: apiConstant + '/del-applications/' + element.ApplicationId,
                        headers: {
                            'Content-type': 'application/json;charset=utf-8'
                        }
                    })
                    .then(function (response) {
                        $scope.init();
                    }, function (rejection) {
                        console.log(rejection);
                        if(rejection.data.Message){
                            toaster.pop('error', "ERROR!", rejection.data.Message);
                        }
                    });
            }
        }
        $scope.changePerpage = function (perpage) {
            $scope.getFilterVal();
            $http({
                    method: 'GET',
                    url: apiConstant + '/applications?per_page=' + perpage + "&query=" + searchAll + '&id_search=' + id_search + '&name_search=' + name_search  + '&type_search=' + type_search + '&ip_search=' + ip_search + '&accepting_host_name=' + accepting_host_name + '&port_search=' + port_search + '&ip_search=' + ip_search + '&hostname_search=' + hostname_search,
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
                url: apiConstant + '/applications?query=' + searchAll ,
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
                url: apiConstant + '/applications?query=' + searchAll + '&id_search=' + id_search + '&name_search=' + name_search  + '&type_search=' + type_search + '&ip_search=' + ip_search + '&accepting_host_name=' + accepting_host_name + '&port_search=' + port_search + '&ip_search=' + ip_search + '&hostname_search=' + hostname_search,
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
                    url: apiConstant + '/applications?per_page=' + $scope.per_page + "&current_page=" + $scope.currentPage + "&query=" + searchAll + '&id_search=' + id_search + '&name_search=' + name_search  + '&type_search=' + type_search + '&ip_search=' + ip_search + '&accepting_host_name=' + accepting_host_name + '&port_search=' + port_search + '&ip_search=' + ip_search + '&hostname_search=' + hostname_search,
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


//add and edit application
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
        $scope.applicationType = RESOURCES.applicationType;
        //case edit
        if ($stateParams.id) {
            $scope.init = function () {
                $http({
                    method: 'GET',
                    url: apiConstant + '/applications/' + $stateParams.id,
                }).then(function successCallback(response) {
                    $scope.data = {
                        accepting_host_id: response.data.accepting_host_id,
                        accepting_host_name: response.data.AcceptingHost.name,
                        description: response.data.description,
                        application_type: response.data.application_type,
                        email: response.data.email,
                        enabled: (response.data.enabled == 0 ? false : true),
                        name: response.data.name,
                        ip: response.data.ip,
                        host_name: response.data.host_name,
                        port: response.data.port
                    };
                }, function errorCallback(response) {
                    console.log("err");
                    console.log(response)
                });
            }
            $scope.init();
        }
        //initial data in search box
        $scope.searchResAcceptinghost = [];
        // search for select option
        $scope.searchForApp = function (value, table) {
            if (value && table) {
                $http({
                    method: 'GET',
                    url: apiConstant + '/search-in-app?query=' + value + '&table=' + table,
                }).then(function successCallback(response) {
                    if (table == "accepting_hosts") {
                        $scope.searchResAcceptinghost = response.data;
                    }
                    
                }, function errorCallback(response) {
                    console.log("err");
                    console.log(response)
                });
            }
        };
        //set selected name
        $scope.selectedFunction = function (item, table) {
            if (table == "accepting_hosts") {
                $scope.data.accepting_hosts_name = item ? item.name : '';
            }
            
        };
         //data send to server, get from search box
        $scope.vm = {
            searchResAcceptinghost: null
        };

        $scope.submitForm = function (isValid) {
            var searchResAcceptinghost = ($scope.vm && $scope.vm.searchResAcceptinghost) ? $scope.vm.searchResAcceptinghost: "";//get from search box
            var accepting_host_id = $scope.data.accepting_host_id?$scope.data.accepting_host_id:'';//if edit, get from api
            //case search and change accepting host 
            if(searchResAcceptinghost){
                accepting_host_id = searchResAcceptinghost.AcceptingHostId;
            }

            if (accepting_host_id &&  $scope.data.name) {
                var apiUrl = apiConstant + '/applications'; //api add
                if ($stateParams.id) {
                    apiUrl = apiConstant + '/edit-applications/' + $stateParams.id; //api edit if exist id;
                }
                
                var dataPost = {
                    name: $scope.data.name,
                    accepting_host_id: accepting_host_id,
                    description: $scope.data.description,
                    application_type: $scope.data.application_type,
                    ip: $scope.data.ip,
                    port: $scope.data.port,
                    enabled: $scope.data.enabled ? 1 : 0,
                    host_name: $scope.data.host_name,
                    is_valid_user_required: 1,
                    is_valid_device_required: 1,
                    created_by_id: 1,
                };
                $http({
                    method: 'POST',
                    url: apiUrl,
                    data: dataPost
                }).then(function (data) {
                    console.log("succ");
                    console.log(data);
                    $state.go("applications");
                }, function (err) {
                    console.log(err);
                    toaster.pop('error', "ERROR!", err.data.message);
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
                        application_type: response.data.application_type,
                        enabled: response.data.enabled,
                        name: response.data.name,
                        host_name: response.data.host_name,
                        port: response.data.port,
                        created_time: response.data.created_time,
                        ip: response.data.ip
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