package cache

import ()

// Get() - Creates, initializes and returns the mysql instance based on the given config
func Get(conf Config) (ret CacheInterface, err error) {
	if conf.Platform == REDIS {
		ret = new(redisAdapter)
		err = ret.Init(&conf)
	} else {
		err = getErrObj(ERR_NO_PLATFORM, conf.Platform+" is not supported")
	}
	// return
	return ret, err
}