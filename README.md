![Go Report Card](https://goreportcard.com/badge/github.com/maheshreddy7797/docker-localdir-volume-plugin)


lugins-helpers/volume
    ```
    
# myexampledriver : docker-volume-plugin-example(2017)
  
  - This plugin was built with the reference @ (https://github.com/fntlnz/docker-volume-plugin-example) overcoming the issues       of the repository.
  
  - This plugin can mount the volume at the mountpoint provided in the path, if the mount path doesn't exists the plugin will       creates the path and tries to mounth that mount point.
  
  ## prerequisites :
  - As this is the basic part of the code some prerequests are mandatory follow the below lines 
  ```
    mkdir -p /tmp/exampledriver/myvolumename
     ```
  - We need to import/get some of the packages which will in building the code 
   ```
    go get github.com/Sirupsen/logrus 
    go get github.com/docker/go-p
  
## Install driver
  ```
      git clone https://github.com/maheshreddy7797/docker-localdir-volume-plugin.git
      cd docker-localdir-volume-plugin
      go build .
      ./docker-localdir-volume-plugin
  ```
#### Open new Terminal window
  ```Shell
      docker run -it -v myvolumename:/data --volume-driver=myexampledriver alpine sh
  ```
  > /#
  
#### Check if volume is created
     
  ```Shell
      docker volume ls
  ```
``` 
    DRIVER                 VOLUME NAME
    local                  database
    myexampledriver        myvolumename
```
