# openEBSDriver

## Install driver
  
   ```
      git clone https://github.com/maheshreddy7797/ebsDriver.git
      cd ebsDriver
      go build .
      ./ebsDriver
  ```
### Open new Terminal window`
  ```
      docker run -it -v openebsvol1:/data --volume-driver=openEBSDriver alpine sh
  ```
