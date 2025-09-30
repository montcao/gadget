# gadget (beta)
Programmatically inspect layers of a container image. This can be used as a cli tool or as a package in your application. 

The beta release operates on the presumption that go is installed on the system. To install the package, run: 

``go install github.com/montcao/gadget``

Then inspect an image with: 

``gadget inspect <image-name>``

An example output is shown at the bottom.

Prebuilt binaries - TODO


#### Dev example: 

To do development, clone the repo and run this command:
``go run . inspect <imagename>``

#### Output:

<img width="636" height="533" alt="image" src="https://github.com/user-attachments/assets/2ac2e7dc-f7ff-43e2-bf4f-aeff9cab4232" />
