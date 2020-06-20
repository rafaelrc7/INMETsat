# inmet-doppler
Simple shell script that downloads the latest satellite images from the Instituto Nacional de Metereologia (at http://www.inmet.gov.br/satelites/) and process them into a GIF.

# Configuration
Configuration is simple. First, to specify the scope of the aquired images, modify the "region" and "regionFull" variables near the top of the script. The default is "br" and "BRASIL", which get the full country satellite images.
You can also choose what kind of informations is retrieved, modifying the "infoType" variable, the default is "TN" that gets "Topo das Nuvens".
And the variable "gifImageSize" controls how many images before the last one are retrieved, the default is 30.

# Install
Installation is pretty simple, just move the script to a folder inside your PATH so that it can be easily accessed. The script uses the /tmp directory to download and process the images. 

# Requirements
The script only has two requirements:
  - MPV  (Used to reproduce the gif on a loop)
  - imagemagick  (Used to make the downloaded images into a GIF)
  
# Thanks for your time
I would also like to add that this is my first public script, and I'm not that experienced, if you have any critics/suggestions feel free to contact me.
