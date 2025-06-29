
#!/bin/bash
# Record 5 seconds from the working capture device
sox -t alsa plughw:0,7 -c 1 -r 16000 -b 16 input.wav trim 0 5

