![Go Report Card](https://goreportcard.com/badge/github.com/maheshreddy7797/docker-localdir-volume-plugin)



    
# myexampledriver : docker-volume-plugin-example(2017)
  
   - This plugin can mount the volume at the mountpoint provided in the path, if the mount path doesn't exists the plugin will       creates the path and mounts it.
  
## presetup :
  - Create a mount path of the volume that needs to be mounted.
  ```
     For our convention MountPoint : /tmp/exampledriver
                        MountPath  : /tmp/exampledriver/demovol
   ```
   ```
     mkdir -p /tmp/exampledriver
   ```
  - Get the required go packages 
   ```
     go get github.com/Sirupsen/logrus 
     go get github.com/docker/go-plugins-helpers/volume
   ```
  
## 1. Install driver
   ```
      Make
   ```
## 2. Check docker volume plugins list
   ```
      docker plugin ls
   ```
   ID                  NAME                                         DESCRIPTION                         ENABLED
   12fb9b1c43c8        maheshreddy7797/myexampledriver:latest       Example volume plugin               true
        
## 3. Create volume
  
  ```Shell
      docker volume create -d maheshreddy7797/myexampledriver --name=demovol
  ```
  > demovol
  
## 4. Check if volume is created
     
  ```
    docker volume ls
  ```
  ``` 
    DRIVER                               VOLUME NAME
    local                                database
    maheshreddy7797/myexampledriver      demovol
  ```

## 5. Run the container with mounted volume
  ```
    docker run -it -v demovol:/data alpine sh
  ```
