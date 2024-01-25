Kurby

1/22/24:
I am now comfortable using xml and getting nested elements.
Next time I will recursively traverse the structure.kurby document and put all data into a tree
The tree will contain Component elements. Right now, all elements are ComponentProp because that's how the xml marshals them. It's the lowest abstraction they can go
The tree will be constructed by taking these ComponentProp elements, checking which XML name it has, and running it through a factory to 
convert it to its respective Component. 

1/23/24:
  The DOM tree has been built. Now I need to experiment with rendering these components.
  What I will probably do tomorrow is brainstorm how parent/child rendering will work.
  I'm thinking each component will have to have a way to handle children. 
