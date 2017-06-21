[![Go Report Card](https://goreportcard.com/badge/github.com/maheshreddy7797/docker-volume-driver)](https://goreportcard.com/report/github.com/maheshreddy7797/docker-volume-driver)

# Docker-Volume-Driver (2017)
  
   - This driver was built and developed by following the documentations of https://docs.docker.com/engine/extend/config/#config-field-descriptions and https://blog.codeship.com/extend-docker-via-plugin/
   
   - This driver can mount the volume at the mountpoint provided in the path(should be an absolute host path only), if the mount volume doesn't exists, a new volume with the specified volume-name is created at the path and mounts it.
  
## 1. Install driver
   ```
        - git clone https://github.com/maheshreddy7797/docker-volume-driver.git
        - cd docker-volume-driver
        - Make
   ```
   >                                     OR
   ``` 
        - mkdir -p /tmp/exampledriver /etc/myexampledriver
        - docker plugin install maheshreddy7797/docker-volume-driver
   ```
## 2. Check docker volume plugins list
   ```
      docker plugin ls
   ```
      ID                  NAME                                              DESCRIPTION                         ENABLED
      12fb9b1c43c8        maheshreddy7797/docker-volume-driver:latest       Example volume plugin               true
        
## 3. Create volume
  
  ```Shell
      docker volume create -d maheshreddy7797/docker-volume-driver -o user=admin -o password=open -o size=1G --name=demovol
  ```
  > demovol
  
## 4. Check if volume is created
     
  ```
    docker volume ls
  ```
  ``` 
    DRIVER                                    VOLUME NAME
    local                                     database
    maheshreddy7797/docker-volume-driver      demovol
  ```

## 5. Run the container with mounted volume
  ```
    docker run -it -v demovol:/data alpine sh
  ```
