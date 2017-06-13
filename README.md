# openEBSDriver

## Install driver
  
   ```
      git clone https://github.com/maheshreddy7797/ebsDriver.git
      cd ebsDriver
      go build .
      ./ebsDriver
  ```
#### Open new Terminal window
  ```Shell
      docker run -it -v openebsvol1:/data --volume-driver=openEBSDriver alpine sh
  ```
  > /#
  
#### Check if volume is created
     
  ```Shell
      docker volume ls
  ```
> local                database
> openEBSDriver        openebsvol1
