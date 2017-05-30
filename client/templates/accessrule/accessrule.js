mainApp.controller('accessruleListController', ['$scope', '$rootScope', 'apiConstant', '$http', "$location", "$state", "RESOURCES",
    function ($scope, $rootScope, apiConstant, $http, $location, $state, RESOURCES) {
        $scope.itemPerPage = RESOURCES.itemPerPage;
        $scope.accessRuleType = RESOURCES.actionType;
        var application_search, id_search, searchAll, email_search, description_search, enabled_search, user_search, device_search, group_search, byuser_search; //init value for filter
        $scope.getFilterVal = function () {
            //get  value for filter
            searchAll = ($scope.asyncSelected != undefined) ? $scope.asyncSelected : '';
            id_search = ($scope.search && $scope.search.id_search != undefined) ? $scope.search.id_search : '';
            application_search = ($scope.search && $scope.search.application_search != undefined) ? $scope.search.application_search : '';
            description_search = ($scope.search && $scope.search.description_search != undefined) ? $scope.search.description_search : '';
            group_search = ($scope.search && $scope.search.group_search != undefined) ? $scope.search.group_search : '';
            user_search = ($scope.search && $scope.search.user_search != undefined) ? $scope.search.user_search : '';
            device_search = ($scope.search && $scope.search.device_search != undefined) ? $scope.search.device_search : '';
            byuser_search = ($scope.search && $scope.search.byuser_search != undefined) ? $scope.search.byuser_search : '';
            access_rule_type = ($scope.search && $scope.search.access_rule_type != undefined) ? $scope.search.access_rule_type : '';

            enabled_search = ($scope.search && $scope.search.enabled_search != undefined) ? $scope.search.enabled_search : '';
            enabled_search ? (enabled_search = 1) : (enabled_search = 0); //if check enabled = 1 else =0
        }

        $scope.init = function () {
            $http({
                method: 'GET',
                url: apiConstant + '/accessrules',
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
                        url: apiConstant + '/del-accessrules/' + element.AccessRuleId,
                        headers: {
                            'Content-type': 'application/json;charset=utf-8'
                        }
                    })
                    .then(function (response) {
                        console.log(response.data.status);
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
                    url: apiConstant + '/accessrules?per_page=' + perpage + '&query=' + searchAll + '&id_search=' + id_search + '&application_search=' + application_search + '&description_search=' + description_search + '&group_search=' + group_search + '&user_search=' + user_search + '&device_search=' + device_search + '&byuser_search=' + byuser_search + '&access_rule_type=' + access_rule_type,
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
        $scope.onSelect = function ($item, $model, $label) {
            $scope.getFilterVal();
            return $http({
                method: 'GET',
                url: apiConstant + '/accessrules?query=' + searchAll,
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
        //for search below box( filter )
        $scope.onSearchBelowBox = function ($item, $model, $label) {
            //get filter val params
            $scope.getFilterVal();
            return $http({
                method: 'GET',
                url: apiConstant + '/accessrules?query=' + searchAll + '&id_search=' + id_search + '&application_search=' + application_search + '&description_search=' + description_search + '&group_search=' + group_search + '&user_search=' + user_search + '&device_search=' + device_search + '&byuser_search=' + byuser_search + '&access_rule_type=' + access_rule_type,
            }).then(function successCallback(response) {
                var totalItem = response.data.Total;
                var perpage = response.data.PerPage;
                $scope.bigTotalItems = totalItem;
                $scope.per_page = perpage;
                $scope.changeData(response);
                return $scope.list;
            }, function errorCallback(response) {
                console.log("err");
                console.log(response);
            });
        };

        $scope.pageChanged = function () {
            //get filter val params before call API
            $scope.getFilterVal();
            $http({
                    method: 'GET',
                    url: apiConstant + '/accessrules?per_page=' + $scope.per_page + "&current_page=" + $scope.currentPage + '&query=' + searchAll + '&id_search=' + id_search + '&application_search=' + application_search + '&description_search=' + description_search + '&group_search=' + group_search + '&user_search=' + user_search + '&device_search=' + device_search + '&byuser_search=' + byuser_search + '&access_rule_type=' + access_rule_type,
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


mainApp.controller('accessruleDetailController', ['$scope', 'apiConstant', '$http', "$window", "$state", "$stateParams", "$localStorage",
    function ($scope, apiConstant, $http, $window, $state, $stateParams, $localStorage) {
        $scope.data = {
            access_rule_id: "",
            application_id: "",
            user_id: "",
            device_id: "",
            group_id: "",
            description: "",
            access_rule_type: "",
            enabled: false,
            created_by_id: "",
            created_time: "",
            updated_time: ""
        };

        if ($stateParams.id) {
            $scope.init = function () {
                $http({
                    method: 'GET',
                    url: apiConstant + '/accessrules/' + $stateParams.id,
                }).then(function successCallback(response) {
                    console.log(response.data);
                    $scope.data = {
                        access_rule_id: response.data.AccessRuleId,
                        description: response.data.description,
                        application_id: response.data.application_id,
                        application_name: response.data.Application.name,
                        enabled: response.data.enabled,
                        created_by_id: response.data.created_by_id,
                        created_by_name: response.data.CreatedByUser.name,
                        user_id: response.data.user_id,
                        user_name: response.data.User.name,
                        device_id: response.data.device_id,
                        device_name: response.data.Device.name,
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


//add and edit accessrules
mainApp.controller('accessRuleController', ['$scope', 'apiConstant', '$http', "$window", "$state", "$stateParams", "toaster", "$localStorage", "RESOURCES",
    function ($scope, apiConstant, $http, $window, $state, $stateParams, toaster, $localStoragem, RESOURCES) {
        $scope.data = {
            access_rule_id: "",
            application_id: "",
            user_id: "",
            device_id: "",
            group_id: "",
            description: "",
            access_rule_type: "",
            enabled: false,
            created_by_id: "",
            created_time: "",
            updated_time: ""
        };
        $scope.accessRuleType = RESOURCES.actionType;
        //case edit
        if ($stateParams.id) {
            $scope.init = function () {
                $http({
                    method: 'GET',
                    url: apiConstant + '/accessrules/' + $stateParams.id,
                }).then(function successCallback(response) {
                    console.log(response.data)
                    $scope.data = {
                        application_name: response.data.Application.name,
                        user_name: response.data.User.name,
                        group_name: response.data.Group.name,
                        device_name: response.data.Device.name,

                        application_id: response.data.application_id,
                        description: response.data.description,
                        enabled: (response.data.enabled == 0 ? false : true),
                        access_rule_type: response.data.access_rule_type,
                        created_by_id: response.data.created_by_id,
                        device_id: response.data.device_id,
                        user_id: response.data.user_id,
                        group_id: response.data.group_id
                    };
                    //$scope.vm.applicationsRes= response.data.Application.name;
                }, function errorCallback(response) {
                    console.log("err");
                    console.log(response)
                });
            }
            $scope.init();
        }
        //initial data in search box
        $scope.searchResApp = [];
        $scope.searchResUser = [];
        $scope.searchResGroup = [];
        $scope.searchResDevice = [];
        //data send to server, get from search box
        $scope.vm = {
            applicationsRes: null,
            userRes: null,
            groupRes: null,
            deviceRes: null,
        };
        //when typing in search box
        $scope.searchAcl = function (value, table) {
            if (value && table) {
                $http({
                    method: 'GET',
                    url: apiConstant + '/search-acl?query=' + value + '&table=' + table,
                }).then(function successCallback(response) {
                    if (table == "applications") {
                        $scope.searchResApp = response.data;
                    }
                    if (table == "user") {
                        $scope.searchResUser = response.data;
                    }
                    if (table == "group") {
                        $scope.searchResGroup = response.data;
                    }
                    if (table == "device") {
                        $scope.searchResDevice = response.data;
                    }
                }, function errorCallback(response) {
                    console.log("err");
                    console.log(response)
                });
            }
        };
        //set selected name
        $scope.selectedFunction = function (item, table) {
            console.log(item);
            if (table == "applications") {
                $scope.data.application_name = item ? item.name : '';
            }
            if (table == "user") {
                $scope.data.user_name = item ? item.name : '';
            }
            if (table == "group") {
                $scope.data.group_name = item ? item.name : '';
            }
            if (table == "device") {
                $scope.data.device_name = item ? item.name : '';
            }
        };


        $scope.submitForm = function (isValid) {
            var apiUrl = apiConstant + '/accessrules'; //api add
            if ($stateParams.id) {
                apiUrl = apiConstant + '/edit-accessrules/' + $stateParams.id; //api edit if exist id;
            }
            var dataPost = {
                application_id:($scope.vm.applicationsRes && parseInt($scope.vm.applicationsRes.ApplicationId)) ? parseInt($scope.vm.applicationsRes.ApplicationId) : $scope.data.application_id, //if is edit application_id = $scope.data.application_id
                user_id: $scope.vm.userRes ? parseInt($scope.vm.userRes.UserId) : $scope.data.user_id,
                device_id: $scope.vm.deviceRes ? parseInt($scope.vm.deviceRes.DeviceId) : $scope.data.device_id,
                group_id: $scope.vm.groupRes ? parseInt($scope.vm.groupRes.GroupId) : $scope.data.group_id,
                description: $scope.data.description,
                access_rule_type: $scope.data.access_rule_type,
                enabled: $scope.data.enabled ? 1 : 0,
                created_by_id: 1,
            };
            if(!dataPost.user_id || !dataPost.device_id || !dataPost.group_id|| !dataPost.description|| !dataPost.access_rule_type){
                toaster.pop('error', "ERROR!", "Enter enough infor!");
                return false;
            }

            $http({
                method: 'POST',
                url: apiUrl,
                data: dataPost
            }).then(function (data) {
                console.log("succ");
                console.log(data);
                $state.go("accessrule");
            }, function () {
                console.log("err");
            });
            
        };

    }
]);