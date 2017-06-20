![Go Report Card](https://goreportcard.com/badge/github.com/maheshreddy7797/docker-localdir-volume-plugin)



    
# myexampledriver : docker-volume-plugin-example(2017)
  
   - This plugin was built and developed by following the documentations of https://docs.docker.com/engine/extend/config/#config-field-descriptions and https://blog.codeship.com/extend-docker-via-plugin/
   
   - This plugin can mount the volume at the mountpoint provided in the path, if the mount path doesn't exists the plugin will       creates the path and mounts it.
  
## 1. Install driver
   ```
        - git clone https://github.com/maheshreddy7797/myexampledriver.git
        - cd myexampledriver
        - Make
   ```
   >                                     OR
   ``` 
        - mkdir -p /tmp/exampledriver
        - docker plugin install maheshreddy7797/myexampledriver
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
