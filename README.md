# Paidpodcasts
The best way to get paid for your podcast

I was thinking about what it would look like if I put [Nostrovia](https://nostrovia.org) onto [Nostrocket](https://nostrocket.org).

Here's some initial thoughts.

Previous episodes are free, but to listen to the **latest** episode, you pay an invoice. 

The "relay" then encrypts the file to the user's pubkey and sends it to the client (probably a fork of some existing podcast player that is patched to enable LN payments and decryption). In a normal podcast player that doesn't support LN payments and decryption, the latest episide is a preview.

## How to mitigate against freeloading
The "relay" will insert an audio overlay at a random point in the audio stream. 

This overlay is barely audible to the human hear but still passes through MP3 encoding and is your pubkey encoded to audo tones.

Anyone can run the file through a script to find the overlay and print out the pubkey (GNU radio would do it easily, but wouldn't be difficult to build a web frontend and host it). 

So if someone shares the file publicly, they will easily get found out and the owner of the pubkey will then, out of pure shame, commit Seppuku in order to restore honour to their family name.


