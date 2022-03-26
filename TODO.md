- [x] debugging why network output is always the same. 
    - [x] need to create tests for Grid.GetObjectSenseData to validate its working as expected. 
- [x] ADD AN OSCILATION (Implemented, needs testing)
- ~~[ ] Add ID's to creatures to enable recall from array so debug index can be set correctly when clicked.~~
- [X] Add a different speed in stats and move further distance with it. 
- [X] Pre-Compute a certain amount of steps per simulation, and have the server export each frame of each cycle.
    - [ ] Save this as a bundle.
    - [X] This moves from having the javascript client call cylce, and having the server get overwhealemd with the requests.
    - [X] Client will be consuming pre-computed simulations, that are pre-made in advance. 
- [X] Evolutionary experiments 
    - [X] Ability to select what ones can breed and spread behaviours. 
        - [X] Spike: can I acces / change weights of the network.
            Yes you can. See cmd/crossoerverTest
- [ ] Rename package to creatureServices
- [X] Rate limit cycle calls, calling server sync every 250ms
- [ ] World size is based on the number on the form, not the size of the actual world
    - wYou can make a smaller world bigger by creating it, then changing the world size form and resetting. bit shit.


- [X] Split the world and simulation out
    - [X] Have a world registry where new configurations can be added.
    - [X] Have the simulation be replayable
- [ ] Expand breeding
    - New system to be made due to breakage from splitting worlds
    - Enable breeding the fittest of two or more worlds. 
- [ ] Fitness parameters
    - Create a system to create survival zones, instead of debug to choose the fittest.