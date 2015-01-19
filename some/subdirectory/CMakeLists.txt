
GO_GET(go_redis github.com/hoisie/redis)
ADD_GO_INSTALLABLE_PROGRAM(redis_lister # executable name
                           redis_lister.go # `package main` source file
                           go_redis) # everything else is a dependency
