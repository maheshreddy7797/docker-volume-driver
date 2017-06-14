![Go Report Card](https://goreportcard.com/badge/github.com/maheshreddy7797/docker-localdir-volume-plugin)



    
# myexampledriver : docker-volume-plugin-example(2017)
  
  - This plugin was built with the reference @ (https://github.com/fntlnz/docker-volume-plugin-example) overcoming the issues       of the repository.
  
  - This plugin can mount the volume at the mountpoint provided in the path, if the mount path doesn't exists the plugin will       creates the path and tries to mounth that mount point.
  
  ## presetup :
  - Create a mount path of the volume that needs to be mounted.
    For our convention MountPoint : /tmp/exampledriver
                       MountPath  : /tmp/exampledriver/myvolumename
  ```
    mkdir -p /tmp/exampledriver/myvolumename
     ```
  - Get the required go packages 
   ```
    go get github.com/Sirupsen/logrus 
    go get github.com/docker/go-plugins-helpers/volume
    ```
  
## 1. Install driver
  ```
      git clone https://github.com/maheshreddy7797/docker-localdir-volume-plugin.git
      cd docker-localdir-volume-plugin
      go build .
      ./docker-localdir-volume-plugin
  ```
## 2. Open new Terminal window
  ```Shell
      docker run -it -v myvolumename:/data --volume-driver=myexampledriver alpine sh
  ```
  > /#
  
## 3. Check if volume is created
     
  ```Shell
      docker volume ls
  ```
``` 
    DRIVER                 VOLUME NAME
    local                  database
    myexampledriver        myvolumename
```
