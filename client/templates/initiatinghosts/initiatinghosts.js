mainApp.controller('initiatingHostListController', ['$scope', '$rootScope', 'apiConstant', '$http', "$location", "$state", "RESOURCES",
    function ($scope, $rootScope, apiConstant, $http, $location, $state, RESOURCES) {
        $scope.itemPerPage = RESOURCES.itemPerPage;
        var name_search, id_search, searchAll, description_search, enabled_search; //init value for filter
        $scope.getFilterVal = function () {
            //get  value for filter
            searchAll = ($scope.asyncSelected != undefined) ? $scope.asyncSelected : '';
            id_search = ($scope.search && $scope.search.id_search != undefined) ? $scope.search.id_search : '';
            name_search = ($scope.search && $scope.search.name_search != undefined) ? $scope.search.name_search : '';
            description_search = ($scope.search && $scope.search.description_search != undefined) ? $scope.search.description_search : '';

            enabled_search = ($scope.search && $scope.search.enabled_search != undefined) ? $scope.search.enabled_search : '';
            enabled_search ? (enabled_search = 1) : (enabled_search = 0); //if check enabled = 1 else =0
        }

        $scope.init = function () {
            $http({
                method: 'GET',
                url: apiConstant + '/initiatinghosts',
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
                        url: apiConstant + '/del-initiatinghosts/' + element.InitiatingId,
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
            //get value filter
            $scope.getFilterVal();
            $http({
                    method: 'GET',
                    url: apiConstant + '/initiatinghosts?per_page=' + perpage + "&query=" + searchAll + '&id_search=' + id_search + '&name_search=' + name_search + '&description_search=' + description_search, //if not search, query = ''
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
            //get value filter
            $scope.getFilterVal();
            return $http({
                method: 'GET',
                url: apiConstant + '/initiatinghosts?query=' + searchAll,
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
                url: apiConstant + '/initiatinghosts?query=' + searchAll + '&id_search=' + id_search + '&name_search=' + name_search + '&description_search=' + description_search,
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
            //get value filter
            $scope.getFilterVal();
            $http({
                    method: 'GET',
                    url: apiConstant + '/initiatinghosts?per_page=' + $scope.per_page + "&current_page=" + $scope.currentPage + "&query=" + searchAll + '&id_search=' + id_search + '&name_search=' + name_search + '&description_search=' + description_search, //if not search, query = ''
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


//add and edit ACH
mainApp.controller('initiatingHostController', ['$scope', 'apiConstant', '$http', "$window", "$state", "$stateParams", "toaster",
    function ($scope, apiConstant, $http, $window, $state, $stateParams, toaster) {
        $scope.data = {
            description: "",
            enabled: false,
            name: "",
            created_by_id: "",
            permission: ""
        };
        //case edit
        if ($stateParams.id) {
            $scope.init = function () {
                $http({
                    method: 'GET',
                    url: apiConstant + '/initiatinghosts/' + $stateParams.id,
                }).then(function successCallback(response) {
                    $scope.data = {
                        description: response.data.description,
                        enabled: (response.data.enabled == 0 ? false : true),
                        name: response.data.name,
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
                var apiUrl = apiConstant + '/initiatinghosts'; //api add
                if ($stateParams.id) {
                    apiUrl = apiConstant + '/edit-initiatinghosts/' + $stateParams.id; //api edit if exist id;
                }
                var dataPost = {
                    description: $scope.data.description,
                    enabled: $scope.data.enabled ? 1 : 0,
                    name: $scope.data.name,
                    created_by_id: 1,
                };

                $http({
                    method: 'POST',
                    url: apiUrl,
                    data: dataPost
                }).then(function (data) {
                    console.log("succ");
                    console.log(data);
                    $state.go("initiatinghosts");
                }, function () {
                    console.log("err");
                });
            } else {
                toaster.pop('error', "ERROR!", "Enter valid infor!");
            }
        };
    }
]);

mainApp.controller('detailInitiatingHostController', ['$scope', 'apiConstant', '$http', "$window", "$state", "$stateParams",
    function ($scope, apiConstant, $http, $window, $state, $stateParams) {
        $scope.data = {
            initiating_host_id: null,
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
                    url: apiConstant + '/initiatinghosts/' + $stateParams.id,
                }).then(function successCallback(response) {
                    console.log(response);
                    $scope.data = {
                        initiating_host_id: response.data.initiating_host_id,
                        description: response.data.description,
                        enabled: response.data.enabled,
                        name: response.data.name,
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