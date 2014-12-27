###
manager.coffee
Copyright (C) 2014 ender xu <xuender@gmail.com>

Distributed under terms of the MIT license.
###
app.config(['$routeProvider', ($routeProvider)->
  $routeProvider.
    when('/count',
      templateUrl: 'partials/count.html'
      controller: 'CountCtrl'
    ).when('/users',
      templateUrl: 'partials/users.html'
      controller: 'UsersCtrl'
    ).when('/session',
      templateUrl: 'partials/session.html'
      controller: 'SessionCtrl'
    ).otherwise({
      redirectTo: '/count'
    })
])

