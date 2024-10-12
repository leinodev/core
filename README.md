# Leino Core

## Project structure
* ./config contains configuration files
* ./internal consists of:
  * ./application contains app builder (new, configure, run)
  * ./core contains the essential logic of the application
    * ./context contains all contexts entities used for communication between layers
    * ./layers contains layers specification and implementation
* ./pkg contains independent modules that is used by our application and might be used by other apps as well