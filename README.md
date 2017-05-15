
# Main Idea
 A vertical puzzle sidescroller(upscroller?)
- player has to keep moving upward to continue in the game
- will only be able to go up by solving puzzles
- as the screen is always moving, the player has to do it within a certain timeframe

I'm imagining gameplay would be as such-
1. You start at the very bottom.
2. You can continue upwards by using certain objects(ex. climbing up a ladder). Getting to these objects may require you to solve a puzzle.
3. I suppose the score would be the distance you acheived. 
4. The speed  of scrolling and puzzle difficulty increase as you get higher. 

How this would be done technically-
1. Read level from text file
2. Set each character to be a certain object
3. Allow each object to have a specific purpose.

# Backup Idea

Something like Guitar Hero, or more specifically like Piano Tiles(google it if you don't know)
- piano sounds probably will be easiest
- measure the times between the notes, when the notes are, if the player presses a key within the right timeframe, the game continues
- add some gui that tells the player when to press

Three ways I can do this, ordered by simple first:
1. Just play a piano song in the background, time when the player has to press a button
2. Make the actual next note play on button press(would need a much more rigorous asset library for this).
3. Make it somehow read a .midi file.

I'm thinking the controls will just be the middle row of keys. ASDF HJKL for 8 fingers. All piano keys just compressed into that.

I'll just either make assets or find them somewhere online. I won't really need anything difficult.

Problem with this idea- it doesn't have an actual player/collisions.

I can make collisions in the gui, but I could also do without.
