# **A golang template which base gin, designed by yy**

## Installing and Getting started

1. Clone the repository.

       git clone https://github.com/guaidashu/go_proc_cache.git

## Usage

      key := "custom key"
      reward, _ := proc_cache.ProcCache.Get(key, func() (interface{}, error) {
         var data type(any type)
         // do something
      
         return data, nil
      }, expired(过期时间，例：time.Second*5))

## FAQ

Contact to me with email "1023767856@qq.com" or "song42960@gmail.com"

## Running Tests

Add files to /test and run it.

## Finally Thanks 